import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { useAccount } from 'wagmi';
import { DEPLOYMENT, displayOrLessThan } from '@shared/index';
import { useMarketPools } from '@/hooks/useMarket';
import { useChainNow, usePoolGlobals } from '@/hooks/usePools';
import { readUserLend, readUserBorrow, readBalance } from '@/lib/reads';

const STABLE = DEPLOYMENT.stableDecimals;

interface PositionRow {
  poolIndex: number;
  displayPoolId: number;
  lendStake: bigint;
  borrowStake: bigint;
  lendRefunded: boolean;
  lendClaimed: boolean;
  borrowRefunded: boolean;
  borrowClaimed: boolean;
  spBalance: bigint;
  jpBalance: bigint;
}

/**
 * 我的仓位（PRD 11.4）
 *
 * 关键点（GAP-18 / UI-11）：
 *   1. 权益来源有两类 —— 原始入金记录（userLendInfo/userBorrowInfo）和 SP/JP 凭证余额。
 *      只按 stakeAmount != 0 筛选会漏掉「只收到凭证转账、没有原始入金」的账户。
 *      本页按池逐个查两者并集。
 *   2. 首版不扫描历史转账事件，因此凭证来源的持仓只在已同步的池里可见。
 */
export function PortfolioPage() {
  const { address, isConnected } = useAccount();
  const { rows } = useMarketPools();
  const chainNow = useChainNow();
  const { paused } = usePoolGlobals();
  const positions = usePositions(address, rows);

  if (!isConnected) {
    return (
      <div className="card">
        <h1 className="text-lg font-semibold text-white">我的仓位</h1>
        <p className="mt-2 text-sm text-slate-400">连接钱包后查看你参与过的池与可领取的权益。</p>
      </div>
    );
  }

  return (
    <div className="space-y-5">
      <div>
        <h1 className="text-xl font-semibold text-white">我的仓位</h1>
        <p className="mt-1 text-sm text-slate-400">
          权益分为两类：原始入金记录（可领取退款或凭证）与 SP / JP 凭证余额（可赎回）。两者独立核算。
        </p>
      </div>

      {positions.length === 0 && (
        <div className="card">
          <p className="text-sm text-slate-400">当前账户在这些池中没有原始入金记录，也没有 SP / JP 凭证余额。</p>
          <p className="mt-1 text-xs text-slate-500">
            首版仅按已同步的池列表逐个查询，不扫描历史凭证转账事件。若你通过转账收到过凭证但此处未显示，说明该池不在当前列表内。
          </p>
          <Link to="/" className="mt-3 inline-block text-sm text-accent-500 hover:underline">
            去市场看看
          </Link>
        </div>
      )}

      <div className="space-y-3">
        {positions.map((p) => (
          <Link
            key={p.poolIndex}
            to={`/portfolio/${DEPLOYMENT.chainId}/${p.poolIndex}`}
            className="card block transition-colors hover:border-white/10"
          >
            <div className="flex flex-wrap items-center justify-between gap-3">
              <div className="flex items-center gap-3">
                <span className="tabular text-base font-semibold text-white">池 #{p.displayPoolId}</span>
                <span className="text-xs text-slate-500">索引 {p.poolIndex}</span>
              </div>
              <span className="text-xs text-slate-500">查看权益与操作 →</span>
            </div>

            <div className="mt-4 grid grid-cols-2 gap-x-6 gap-y-3 md:grid-cols-4">
              <Cell label="我的出借本金" value={`${displayOrLessThan(p.lendStake, STABLE)} USDT`} />
              <Cell label="我的抵押本金" value={`${displayOrLessThan(p.borrowStake, 18, 6)} ETH`} />
              <Cell label="SP 凭证余额" value={displayOrLessThan(p.spBalance, 18)} />
              <Cell label="JP 凭证余额" value={displayOrLessThan(p.jpBalance, 18)} />
            </div>

            <div className="mt-3 flex flex-wrap gap-2 text-xs">
              {p.lendStake > 0n && !p.lendRefunded && !p.lendClaimed && (
                <span className="tag bg-warn/15 text-warn">出借待领取</span>
              )}
              {p.borrowStake > 0n && !p.borrowRefunded && !p.borrowClaimed && (
                <span className="tag bg-warn/15 text-warn">抵押待领取</span>
              )}
              {p.spBalance > 0n && <span className="tag bg-accent-500/15 text-accent-500">SP 待赎回</span>}
              {p.jpBalance > 0n && <span className="tag bg-accent-500/15 text-accent-500">JP 待赎回</span>}
              {paused && <span className="tag bg-negative/15 text-negative">操作暂停中</span>}
            </div>
          </Link>
        ))}
      </div>

      <p className="text-xs text-slate-500">
        链上时间 {chainNow > 0n ? new Date(Number(chainNow) * 1000).toISOString() : '读取中'} · 状态与可执行性以链上时间为准
      </p>
    </div>
  );
}

function Cell({ label, value }: { label: string; value: string }) {
  return (
    <div>
      <div className="text-xs text-slate-500">{label}</div>
      <div className="tabular mt-0.5 text-sm text-slate-200">{value}</div>
    </div>
  );
}

/**
 * 逐池查询原始记录与凭证余额。
 * 用 readPoolBase 拿到该池的 spCoin / jpCoin 地址，再读余额 —— 不能假设凭证地址固定。
 */
function usePositions(
  account: `0x${string}` | undefined,
  pools: { poolIndex: number; spCoin: string; jpCoin: string }[],
) {
  const [data, setData] = useState<PositionRow[]>([]);
  const key = `${account ?? ''}-${pools.map((p) => p.poolIndex).join(',')}`;

  useEffect(() => {
    let cancelled = false;

    async function load() {
      if (!account || pools.length === 0) {
        setData([]);
        return;
      }
      const results: PositionRow[] = [];

      for (const p of pools) {
        try {
          const [lend, borrow, spBalance, jpBalance] = await Promise.all([
            readUserLend(account, p.poolIndex),
            readUserBorrow(account, p.poolIndex),
            readBalance(p.spCoin as `0x${string}`, account),
            readBalance(p.jpCoin as `0x${string}`, account),
          ]);
          if (
            (lend?.stakeAmount ?? 0n) > 0n ||
            (borrow?.stakeAmount ?? 0n) > 0n ||
            spBalance > 0n ||
            jpBalance > 0n
          ) {
            results.push({
              poolIndex: p.poolIndex,
              displayPoolId: p.poolIndex + 1,
              lendStake: lend?.stakeAmount ?? 0n,
              borrowStake: borrow?.stakeAmount ?? 0n,
              lendRefunded: lend?.hasNoRefund ?? false,
              lendClaimed: lend?.hasNoClaim ?? false,
              borrowRefunded: borrow?.hasNoRefund ?? false,
              borrowClaimed: borrow?.hasNoClaim ?? false,
              spBalance,
              jpBalance,
            });
          }
        } catch {
          // 单个池读取失败不应影响其他池；该池在本次结果中缺席
        }
      }

      if (!cancelled) setData(results);
    }

    void load();
    return () => {
      cancelled = true;
    };
  }, [key]);

  return data;
}
