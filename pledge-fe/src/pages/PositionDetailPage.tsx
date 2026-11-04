import { useEffect, useMemo, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { useAccount, useWaitForTransactionReceipt, useWriteContract } from 'wagmi';
import { BaseError, ContractFunctionRevertedError } from 'viem';
import {
  DEPLOYMENT,
  displayUnits,
  displayOrLessThan,
  formatRate,
  formatTimestampWithTz,
  PoolState,
  jpDenominator,
  spDenominator,
  expectedSp,
  expectedJp,
  expectedLoan,
  refundAmount,
  previewSpRedeem,
  previewJpRedeem,
} from '@shared/index';
import { PledgePoolAbi, POOL_ADDRESS } from '@/lib/contracts';
import { StateBadge } from '@/components/StateBadge';
import { useChainNow, usePoolGlobals } from '@/hooks/usePools';
import { useTxStore, classifyTxError } from '@/lib/tx';
import { readPosition, type PositionData } from '@/lib/reads';

const STABLE = DEPLOYMENT.stableDecimals;
const ETH = 18;
const SCALE = 10n ** BigInt(18 - STABLE);

type Action =
  | 'claimLend'
  | 'claimBorrow'
  | 'refundLend'
  | 'refundBorrow'
  | 'withdrawLend'
  | 'withdrawBorrow'
  | 'emergencyLendWithdrawal'
  | 'emergencyBorrowWithdrawal';

export function PositionDetailPage() {
  const { poolIndex } = useParams<{ chainId: string; poolIndex: string }>();
  const pid = Number(poolIndex ?? 0);
  const { address, chainId } = useAccount();
  const chainNow = useChainNow();
  const { paused } = usePoolGlobals();

  const [pos, setPos] = useState<PositionData | null>(null);
  const [spInput, setSpInput] = useState('');
  const [jpInput, setJpInput] = useState('');

  const { writeContract, data: hash, error: writeError, isPending } = useWriteContract();
  const { isLoading: confirming } = useWaitForTransactionReceipt({ hash });
  const tx = useTxStore();
  const [lastAction, setLastAction] = useState<Action | null>(null);

  const wrongChain = Boolean(chainId && chainId !== DEPLOYMENT.chainId);

  useEffect(() => {
    if (!address || wrongChain) {
      setPos(null);
      return;
    }
    let cancelled = false;
    const load = async () => {
      const p = await readPosition(address, pid);
      if (!cancelled) setPos(p);
    };
    void load();
    const timer = setInterval(() => void load(), 15_000);
    return () => {
      cancelled = true;
      clearInterval(timer);
    };
  }, [address, pid, wrongChain]);

  const pool = pos?.pool;
  const data = pos?.data;

  // ─── 链上固定分母（PRD 17.2.3）────────────────────────────────
  const settleAmountLend = data?.settleAmountLend ?? 0n;
  const totalSpAmount = useMemo(() => spDenominator(settleAmountLend, SCALE), [settleAmountLend]);
  const totalJpAmount = useMemo(
    () => jpDenominator(settleAmountLend, SCALE, pool?.martgageRate ?? 0n),
    [settleAmountLend, pool?.martgageRate],
  );

  // ─── 领取预览（直接用 shared 里的链上等价函数，避免与合约算法漂移）────
  const preview = useMemo(() => {
    if (!pool || !pos) return null;
    const { lend, borrow } = pos;
    const lendStake = lend?.stakeAmount ?? 0n;
    const borrowStake = borrow?.stakeAmount ?? 0n;

    // refundBorrow 的分母是 pool.borrowSupply，未成交抵押额是 settleAmountBorrow
    const borrowRefund =
      pool.borrowSupply > 0n
        ? refundAmount(borrowStake, pool.borrowSupply, data?.settleAmountBorrow ?? 0n)
        : 0n;

    return {
      spExpected: expectedSp(settleAmountLend, SCALE, lendStake, pool.lendSupply),
      jpExpected: expectedJp(settleAmountLend, SCALE, pool.martgageRate, borrowStake, pool.borrowSupply),
      loanExpected: expectedLoan(settleAmountLend, borrowStake, pool.borrowSupply),
      lendRefund: refundAmount(lendStake, pool.lendSupply, settleAmountLend),
      borrowRefund,
    };
  }, [pool, pos, settleAmountLend, data]);

  // ─── 赎回预览 ────────────────────────────────────────────────
  const spRedeem = useMemo(() => {
    if (!pool || !data) return null;
    const amt = BigInt(spInput || '0');
    const { amount, willRevert } = previewSpRedeem(
      pool.state,
      amt,
      totalSpAmount,
      data.finishAmountLend,
      data.liquidationAmounLend,
    );
    // 链上 require(redeemAmount > 0) 会回滚
    return { amount, willRevert: amt > 0n && willRevert };
  }, [spInput, pool, data, totalSpAmount]);

  const jpRedeem = useMemo(() => {
    if (!pool || !data) return null;
    const amt = BigInt(jpInput || '0');
    const { amount } = previewJpRedeem(
      pool.state,
      amt,
      totalJpAmount,
      data.finishAmountBorrow,
      data.liquidationAmounBorrow,
    );
    // 链上 JP 侧没有 require(>0)，零兑付会成功执行 —— DEC-14 需要披露
    return { amount, zeroPayout: amt > 0n && amount === 0n };
  }, [jpInput, pool, data, totalJpAmount]);

  function run(action: Action, args: unknown[] = []) {
    setLastAction(action);
    const key = `${action}-${pid}`;
    tx.set(key, { stage: 'awaiting-wallet' });
    writeContract(
      {
        address: POOL_ADDRESS,
        abi: PledgePoolAbi,
        functionName: action,
        args: args as never,
      } as never,
      {
        onSuccess: (h) => tx.set(key, { stage: 'pending', hash: h }),
        onError: (err) => tx.set(key, classifyTxError(err)),
      },
    );
  }

  if (!address) return <div className="card text-sm text-slate-400">连接钱包后查看权益。</div>;
  if (wrongChain) {
    return (
      <div className="card border-warn/20 bg-warn/5 text-sm text-warn">
        钱包网络与池所属网络不符，请先切换至 {DEPLOYMENT.chainName}。切换前不会读取该账户在池中的权益，
        也不会提交任何交易。
      </div>
    );
  }
  if (!pool) return <div className="card text-sm text-slate-400">正在读取池与权益数据…</div>;

  const state = pool.state;
  const afterSettle = chainNow > pool.settleTime;
  const afterEnd = chainNow > pool.endTime;
  const claimable = state === PoolState.EXECUTION || state === PoolState.FINISH || state === PoolState.LIQUIDATION;
  const redeemable = state === PoolState.FINISH || state === PoolState.LIQUIDATION;
  const undone = state === PoolState.UNDONE;
  const lend = pos?.lend;
  const borrow = pos?.borrow;

  return (
    <div className="space-y-5">
      <div className="flex flex-wrap items-center gap-3">
        <Link to="/portfolio" className="text-sm text-slate-400 hover:text-slate-200">
          ← 我的仓位
        </Link>
        <h1 className="text-xl font-semibold text-white">池 #{pid + 1} 权益</h1>
        <StateBadge state={state} />
      </div>

      {afterSettle && !afterEnd && claimable && (
        <div className="card border-warn/15 bg-warn/5 text-xs text-slate-400">
          当前处于执行期：没有补仓、主动还款或取回本金的入口，只能等到期结算或清算。
        </div>
      )}

      <div className="grid gap-5 lg:grid-cols-2">
        {/* 出借侧 */}
        <section className="card">
          <h2 className="text-sm font-semibold text-white">出借侧（SP）</h2>

          <div className="mt-3 space-y-1 text-sm">
            <Row label="我的出借本金" value={`${displayOrLessThan(lend?.stakeAmount ?? 0n, STABLE)} USDT`} />
            <Row label="已领取退款" value={`${displayOrLessThan(lend?.refundAmount ?? 0n, STABLE)} USDT`} />
            <Row label="SP 凭证余额" value={displayOrLessThan(pos?.spBalance ?? 0n, ETH)} />
            <Row label="预计可得 SP" value={displayOrLessThan(preview?.spExpected ?? 0n, ETH)} />
          </div>

          <ActionBox
            title="领取 SP 凭证"
            desc="匹配确定权益后领取 SP。SP 代表该池的出借权益，可用于到期赎回本金与利息。"
            available={claimable && afterSettle && (lend?.stakeAmount ?? 0n) > 0n && !lend?.hasNoClaim}
            blockedReason={
              !claimable
                ? '该池尚未进入匹配后的执行期'
                : !afterSettle
                  ? '募集截止后才能领取'
                  : (lend?.stakeAmount ?? 0n) === 0n
                    ? '没有出借本金记录'
                    : lend?.hasNoClaim
                      ? '已领取过 SP，不能重复领取'
                      : undefined
            }
            onRun={() => run('claimLend', [BigInt(pid)])}
            busy={isPending && lastAction === 'claimLend'}
          />

          <ActionBox
            title="领取未匹配退款"
            desc="未参与匹配的部分按你的本金占比退回，不会自动到账。"
            available={
              claimable &&
              afterSettle &&
              (lend?.stakeAmount ?? 0n) > 0n &&
              !lend?.hasNoRefund &&
              (preview?.lendRefund ?? 0n) > 0n
            }
            blockedReason={
              !claimable
                ? '该池尚未进入匹配后的执行期'
                : !afterSettle
                  ? '募集截止后才能退款'
                  : (lend?.stakeAmount ?? 0n) === 0n
                    ? '没有出借本金记录'
                    : lend?.hasNoRefund
                      ? '已领取过退款'
                      : (preview?.lendRefund ?? 0n) === 0n
                        ? '当前没有可退金额'
                        : undefined
            }
            onRun={() => run('refundLend', [BigInt(pid)])}
            busy={isPending && lastAction === 'refundLend'}
            extra={
              (preview?.lendRefund ?? 0n) > 0n ? (
                <p className="mt-1 text-xs text-slate-400">
                  预计退回 {displayUnits(preview!.lendRefund, STABLE, 6)} USDT
                </p>
              ) : null
            }
          />

          {redeemable && (pos?.spBalance ?? 0n) > 0n && (
            <div className="mt-4 border-t border-white/5 pt-4">
              <h3 className="text-sm font-medium text-slate-200">赎回 SP</h3>
              <p className="mt-1 text-xs text-slate-500">
                销毁 SP 换取该池分配的稳定币。
                {state === PoolState.FINISH ? '按到期结算分配额计算。' : '按清算分配额计算。'}
              </p>
              <input
                className="field tabular mt-2"
                inputMode="decimal"
                placeholder="0.0"
                value={spInput}
                onChange={(e) => setSpInput(e.target.value)}
              />
              <div className="mt-1 flex items-center justify-between text-xs">
                <button
                  className="text-accent-500 hover:underline"
                  onClick={() => setSpInput((pos?.spBalance ?? 0n).toString())}
                >
                  最大
                </button>
                <span className="tabular text-slate-400">
                  {spRedeem ? `预计赎回 ${displayUnits(spRedeem.amount, STABLE, 6)} USDT` : ''}
                </span>
              </div>
              {spRedeem?.willRevert && (
                <p className="mt-1 text-xs text-negative">
                  该数量赎回结果为 0，链上会因「Zero USDT redemption」回滚。请调小数量。
                </p>
              )}
              <button
                className="btn-primary mt-2 w-full"
                disabled={isPending || confirming || spRedeem?.willRevert || BigInt(spInput || '0') === 0n}
                onClick={() => run('withdrawLend', [BigInt(pid), BigInt(spInput || '0')])}
              >
                {isPending && lastAction === 'withdrawLend' ? '请在钱包确认…' : '赎回 SP'}
              </button>
            </div>
          )}
        </section>

        {/* 借款侧 */}
        <section className="card">
          <h2 className="text-sm font-semibold text-white">抵押借款侧（JP）</h2>

          <div className="mt-3 space-y-1 text-sm">
            <Row label="我的抵押本金" value={`${displayOrLessThan(borrow?.stakeAmount ?? 0n, ETH, 6)} ETH`} />
            <Row label="已领取退款" value={`${displayOrLessThan(borrow?.refundAmount ?? 0n, ETH, 6)} ETH`} />
            <Row label="JP 凭证余额" value={displayOrLessThan(pos?.jpBalance ?? 0n, ETH)} />
            <Row label="预计可得 JP" value={displayOrLessThan(preview?.jpExpected ?? 0n, ETH)} />
            <Row label="预计领取借款" value={`${displayOrLessThan(preview?.loanExpected ?? 0n, STABLE, 6)} USDT`} />
          </div>

          <ActionBox
            title="领取借款与 JP 凭证"
            desc="领取借款会同时支付本笔借款并铸造 JP 凭证。JP 代表该池的抵押权益，可用于到期取回抵押物。"
            available={claimable && afterSettle && (borrow?.stakeAmount ?? 0n) > 0n && !borrow?.hasNoClaim}
            blockedReason={
              !claimable
                ? '该池尚未进入匹配后的执行期'
                : !afterSettle
                  ? '募集截止后才能领取'
                  : (borrow?.stakeAmount ?? 0n) === 0n
                    ? '没有抵押本金记录'
                    : borrow?.hasNoClaim
                      ? '已领取过借款与 JP'
                      : undefined
            }
            onRun={() => run('claimBorrow', [BigInt(pid)])}
            busy={isPending && lastAction === 'claimBorrow'}
            extra={
              (preview?.loanExpected ?? 0n) > 0n ? (
                <p className="mt-1 text-xs text-slate-400">
                  预计到账 {displayUnits(preview!.loanExpected, STABLE, 6)} USDT
                </p>
              ) : null
            }
          />

          <ActionBox
            title="领取未匹配退款"
            desc="未参与匹配的抵押物按占比退回。"
            available={
              claimable &&
              afterSettle &&
              (borrow?.stakeAmount ?? 0n) > 0n &&
              !borrow?.hasNoRefund &&
              (preview?.borrowRefund ?? 0n) > 0n
            }
            blockedReason={
              !claimable
                ? '该池尚未进入匹配后的执行期'
                : !afterSettle
                  ? '募集截止后才能退款'
                  : (borrow?.stakeAmount ?? 0n) === 0n
                    ? '没有抵押本金记录'
                    : borrow?.hasNoRefund
                      ? '已领取过退款'
                      : (preview?.borrowRefund ?? 0n) === 0n
                        ? '当前没有可退金额'
                        : undefined
            }
            onRun={() => run('refundBorrow', [BigInt(pid)])}
            busy={isPending && lastAction === 'refundBorrow'}
            extra={
              (preview?.borrowRefund ?? 0n) > 0n ? (
                <p className="mt-1 text-xs text-slate-400">
                  预计退回 {displayUnits(preview!.borrowRefund, ETH, 8)} ETH
                </p>
              ) : null
            }
          />

          {redeemable && (pos?.jpBalance ?? 0n) > 0n && (
            <div className="mt-4 border-t border-white/5 pt-4">
              <h3 className="text-sm font-medium text-slate-200">赎回 JP</h3>
              <p className="mt-1 text-xs text-slate-500">
                销毁 JP 换取该池分配的抵押物。取回的是当时池内剩余的抵押物，不保证等于最初投入的数量。
              </p>
              <input
                className="field tabular mt-2"
                inputMode="decimal"
                placeholder="0.0"
                value={jpInput}
                onChange={(e) => setJpInput(e.target.value)}
              />
              <div className="mt-1 flex items-center justify-between text-xs">
                <button
                  className="text-accent-500 hover:underline"
                  onClick={() => setJpInput((pos?.jpBalance ?? 0n).toString())}
                >
                  最大
                </button>
                <span className="tabular text-slate-400">
                  {jpRedeem ? `预计赎回 ${displayUnits(jpRedeem.amount, ETH, 8)} ETH` : ''}
                </span>
              </div>
              {jpRedeem?.zeroPayout && (
                <p className="mt-1 text-xs text-warn">
                  该数量赎回结果为 0。链上 JP 侧允许零兑付并会成功执行，凭证将被销毁但不产生回收。
                </p>
              )}
              <button
                className="btn-primary mt-2 w-full"
                disabled={isPending || confirming || BigInt(jpInput || '0') === 0n}
                onClick={() => run('withdrawBorrow', [BigInt(pid), BigInt(jpInput || '0')])}
              >
                {isPending && lastAction === 'withdrawBorrow' ? '请在钱包确认…' : '赎回 JP'}
              </button>
            </div>
          )}
        </section>
      </div>

      {/* 紧急退出（仅 UNDONE 状态）*/}
      {undone && (
        <section className="card border-warn/20 bg-warn/5">
          <h2 className="text-sm font-semibold text-warn">募集未成交 · 紧急取回</h2>
          <p className="mt-1 text-xs text-slate-400">
            该池未形成匹配，你的本金原额退回。注意：紧急取回会直接标记为已退款，不能再领取退款或凭证。
          </p>
          <div className="mt-3 flex flex-wrap gap-2">
            <button
              className="btn-ghost"
              disabled={wrongChain || paused || (lend?.stakeAmount ?? 0n) === 0n || lend?.hasNoRefund}
              onClick={() => run('emergencyLendWithdrawal', [BigInt(pid)])}
            >
              取回出借本金 {displayUnits(lend?.stakeAmount ?? 0n, STABLE, 2)} USDT
            </button>
            <button
              className="btn-ghost"
              disabled={wrongChain || paused || (borrow?.stakeAmount ?? 0n) === 0n || borrow?.hasNoRefund}
              onClick={() => run('emergencyBorrowWithdrawal', [BigInt(pid)])}
            >
              取回抵押本金 {displayUnits(borrow?.stakeAmount ?? 0n, ETH, 6)} ETH
            </button>
          </div>
        </section>
      )}

      {/* 状态与错误 */}
      <div className="card space-y-2 text-xs text-slate-400">
        <p>
          募集截止 {formatTimestampWithTz(pool.settleTime)} · 到期 {formatTimestampWithTz(pool.endTime)} · 利率{' '}
          {formatRate(pool.interestRate)}
        </p>
        <p>链上时间 {chainNow > 0n ? formatTimestampWithTz(chainNow) : '读取中'}</p>
        {paused && <p className="text-negative">当前用户操作已暂停，链上会拒绝入金、领取、退款与赎回。</p>}
        <TxFeedback hash={hash} error={writeError} confirming={confirming} action={lastAction} />
      </div>
    </div>
  );
}

function Row({ label, value }: { label: string; value: string }) {
  return (
    <div className="flex items-baseline justify-between gap-4 border-b border-white/5 py-1.5 last:border-0">
      <span className="shrink-0 text-slate-400">{label}</span>
      <span className="tabular text-right text-slate-200">{value}</span>
    </div>
  );
}

function ActionBox({
  title,
  desc,
  available,
  blockedReason,
  onRun,
  busy,
  extra,
}: {
  title: string;
  desc: string;
  available: boolean;
  blockedReason?: string;
  onRun: () => void;
  busy: boolean;
  extra?: React.ReactNode;
}) {
  return (
    <div className="mt-4 border-t border-white/5 pt-4">
      <h3 className="text-sm font-medium text-slate-200">{title}</h3>
      <p className="mt-1 text-xs text-slate-500">{desc}</p>
      {extra}
      <button className="btn-primary mt-2 w-full" disabled={!available || busy} onClick={onRun}>
        {busy ? '请在钱包确认…' : title}
      </button>
      {!available && blockedReason && <p className="mt-1 text-center text-xs text-slate-500">{blockedReason}</p>}
    </div>
  );
}

function TxFeedback({
  hash,
  error,
  confirming,
  action,
}: {
  hash?: `0x${string}`;
  error: unknown;
  confirming: boolean;
  action: Action | null;
}) {
  const revert = (() => {
    if (!(error instanceof BaseError)) return undefined;
    const r = error.walk((e) => e instanceof ContractFunctionRevertedError);
    return r instanceof ContractFunctionRevertedError ? (r.reason ?? r.shortMessage) : error.shortMessage;
  })();

  if (revert) return <p className="text-negative">交易失败：{revert}</p>;
  if (confirming) return <p className="text-accent-500">交易已广播，等待确认…</p>;
  if (hash) {
    return (
      <a
        className="text-accent-500 hover:underline"
        href={`${DEPLOYMENT.explorer}/tx/${hash}`}
        target="_blank"
        rel="noreferrer"
      >
        {action ?? '交易'}已提交 · {hash.slice(0, 10)}…{hash.slice(-8)}
      </a>
    );
  }
  return null;
}
