import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { useAccount, useWaitForTransactionReceipt, useWriteContract } from 'wagmi';
import { BaseError, ContractFunctionRevertedError, parseEther } from 'viem';
import {
  DEPLOYMENT,
  displayUnits,
  displayOrLessThan,
  formatRate,
  formatTimestampWithTz,
  mortgageToLtv,
  remainingLabel,
  termDays,
  PoolState,
  derivePoolState,
} from '@shared/index';
import { PledgePoolAbi, POOL_ADDRESS, STABLE_ADDRESS, ERC20Abi } from '@/lib/contracts';
import { StateBadge, FlagBadges } from '@/components/StateBadge';
import { useChainNow, usePoolGlobals } from '@/hooks/usePools';
import { useTxStore, classifyTxError } from '@/lib/tx';
import { readPoolBase, readPoolData, readBalance, readAllowance, ReadError } from '@/lib/reads';

const STABLE = DEPLOYMENT.stableDecimals;
const ETH = 18;

export function PoolDetailPage() {
  const { poolIndex } = useParams<{ chainId: string; poolIndex: string }>();
  const pid = Number(poolIndex ?? 0);
  const [tab, setTab] = useState<'lend' | 'borrow'>('lend');
  const [amount, setAmount] = useState('');

  const { address, chainId } = useAccount();
  const chainNow = useChainNow();
  const { paused } = usePoolGlobals();

  const [pool, setPool] = useState<Awaited<ReturnType<typeof readPoolBase>>>(null);
  const [poolData, setPoolData] = useState<Awaited<ReturnType<typeof readPoolData>>>(null);
  const [readError, setReadError] = useState<string | null>(null);
  const [stableBalance, setStableBalance] = useState(0n);
  const [allowance, setAllowance] = useState(0n);

  const tx = useTxStore();
  const txKey = `deposit-${tab}-${pid}`;
  const current = tx.txs[txKey];

  const { writeContract, data: hash, isPending, error: writeError } = useWriteContract();
  const { isLoading: confirming } = useWaitForTransactionReceipt({ hash });

  // 链上数据轮询（列表页已读后端快照，详情页以链上为准）
  useEffect(() => {
    let cancelled = false;
    const load = async () => {
      try {
        const [p, d] = await Promise.all([readPoolBase(pid), readPoolData(pid)]);
        if (cancelled) return;
        setPool(p);
        setPoolData(d);
        setReadError(null);
      } catch (err) {
        // 读取失败不能显示成"加载中"或空状态 —— 让用户看到真实原因并重试
        if (cancelled) return;
        const cause = err instanceof ReadError ? (err.cause as { shortMessage?: string; message?: string }) : undefined;
        setReadError(
          err instanceof ReadError
            ? `${err.message}: ${cause?.shortMessage ?? cause?.message ?? '未知原因'}`
            : String(err),
        );
      }
    };
    void load();
    const timer = setInterval(() => void load(), 15_000);
    return () => {
      cancelled = true;
      clearInterval(timer);
    };
  }, [pid]);

  useEffect(() => {
    if (!address) return;
    let cancelled = false;
    const load = async () => {
      const [b, a] = await Promise.all([
        readBalance(STABLE_ADDRESS, address),
        readAllowance(STABLE_ADDRESS, address, POOL_ADDRESS),
      ]);
      if (!cancelled) {
        setStableBalance(b);
        setAllowance(a);
      }
    };
    void load();
    const timer = setInterval(() => void load(), 15_000);
    return () => {
      cancelled = true;
      clearInterval(timer);
    };
  }, [address]);

  const derived = pool
    ? derivePoolState({
        state: pool.state,
        settleTime: pool.settleTime,
        endTime: pool.endTime,
        chainNow,
        paused,
      })
    : null;

  const wrongChain = Boolean(chainId && chainId !== DEPLOYMENT.chainId);
  const remaining = pool && pool.maxSupply > pool.lendSupply ? pool.maxSupply - pool.lendSupply : 0n;

  // ─── 输入校验（PRD 11.3.2）：拒绝负数/零/非法字符/超精度，不静默截断
  const parsed = (() => {
    if (!amount) return { value: 0n, error: undefined as string | undefined };
    const decimals = tab === 'lend' ? STABLE : ETH;
    if (!/^\d*\.?\d*$/.test(amount)) return { value: 0n, error: '只允许数字和小数点' };
    const [, frac = ''] = amount.split('.');
    if (frac.length > decimals) return { value: 0n, error: `最多 ${decimals} 位小数` };
    const value = tab === 'lend' ? parseEther(amount) / 10n ** BigInt(18 - STABLE) : parseEther(amount);
    return { value, error: undefined };
  })();

  // ─── 按钮阻塞原因（PRD 11.5.2）：显示最直接的那一条
  const blockReason = (() => {
    if (!address) return '连接钱包后继续';
    if (wrongChain) return `请切换至 ${DEPLOYMENT.chainName}`;
    if (paused) return '用户操作暂时暂停，已有权益可在此查看';
    if (!pool) return '正在读取最新池状态，请稍候';
    if (pool.state !== PoolState.MATCH) return '该池已进入执行期，不再接受入金';
    if (!derived?.canDeposit) return '募集已截止，等待匹配';
    if (parsed.error) return parsed.error;
    if (parsed.value === 0n) return '请输入金额';
    if (tab === 'lend') {
      if (parsed.value > remaining) return '当前出借额度已用尽或不足';
      if (parsed.value > stableBalance) return `${displayUnits(stableBalance, STABLE, 2)} USDT 余额不足`;
    }
    return undefined;
  })();

  const needsApproval = tab === 'lend' && parsed.value > 0n && allowance < parsed.value;

  function handleApprove() {
    tx.set(txKey, { stage: 'awaiting-wallet' });
    writeContract(
      { address: STABLE_ADDRESS, abi: ERC20Abi, functionName: 'approve', args: [POOL_ADDRESS, parsed.value] },
      {
        onSuccess: () => tx.set(txKey, { stage: 'pending' }),
        onError: (err) => tx.set(txKey, classifyTxError(err)),
      },
    );
  }

  function handleDeposit() {
    tx.set(txKey, { stage: 'awaiting-wallet' });
    if (tab === 'lend') {
      writeContract(
        { address: POOL_ADDRESS, abi: PledgePoolAbi, functionName: 'depositLend', args: [BigInt(pid), parsed.value] },
        {
          onSuccess: (h) => tx.set(txKey, { stage: 'pending', hash: h }),
          onError: (err) => tx.set(txKey, classifyTxError(err)),
        },
      );
    } else {
      writeContract(
        {
          address: POOL_ADDRESS,
          abi: PledgePoolAbi,
          functionName: 'depositBorrow',
          args: [BigInt(pid), parsed.value],
          value: parsed.value,
        },
        {
          onSuccess: (h) => tx.set(txKey, { stage: 'pending', hash: h }),
          onError: (err) => tx.set(txKey, classifyTxError(err)),
        },
      );
    }
  }

  if (readError) {
    return (
      <div className="card border-negative/20 bg-negative/5">
        <p className="text-sm text-negative">无法读取池 #{pid + 1} 的链上状态</p>
        <p className="mt-1 text-xs text-slate-400">{readError}</p>
        <p className="mt-3 text-xs text-slate-500">
          这不代表该池不存在。可先检查钱包网络是否为 {DEPLOYMENT.chainName}，或稍后刷新重试。
        </p>
        <button className="btn-ghost mt-3" onClick={() => location.reload()}>
          重新加载
        </button>
      </div>
    );
  }

  if (!pool) {
    return (
      <div className="card text-sm text-slate-400">
        正在读取池 #{pid + 1} 的链上状态…
        <p className="mt-2 text-xs text-slate-500">RPC {DEPLOYMENT.rpcUrl}</p>
      </div>
    );
  }

  return (
    <div className="grid gap-5 lg:grid-cols-[1fr_380px]">
      {/* 左：概况与风险 */}
      <div className="space-y-5">
        <div className="card">
          <div className="flex flex-wrap items-center gap-3">
            <h1 className="text-xl font-semibold text-white">池 #{pid + 1}</h1>
            <StateBadge state={pool.state} label={derived?.label} />
            <FlagBadges flags={derived?.flags ?? []} />
          </div>

          <div className="mt-5 grid grid-cols-2 gap-x-6 gap-y-4 md:grid-cols-3">
            <Field label="借款资产" value="USDT · 6 位" />
            <Field label="抵押资产" value="原生 ETH · 18 位" />
            <div>
              <div className="text-xs text-slate-500">USDT 合约</div>
              <a
                href={`${DEPLOYMENT.explorer}/address/${STABLE_ADDRESS}`}
                target="_blank"
                rel="noreferrer"
                className="mt-0.5 block break-all font-mono text-xs text-accent-500 hover:underline"
              >
                {STABLE_ADDRESS}
              </a>
            </div>
            <Field label="固定年利率" value={formatRate(pool.interestRate)} accent />
            <Field label="约定期限" value={`${termDays(pool.settleTime, pool.endTime)} 天`} />
            <Field label="募集抵押率" value={formatRate(pool.martgageRate)} />
            <Field label="同口径 LTV" value={mortgageToLtv(pool.martgageRate)} />
            <Field label="出借上限" value={`${displayUnits(pool.maxSupply, STABLE, 0)} USDT`} />
            <Field label="募集截止" value={formatTimestampWithTz(pool.settleTime)} />
            <Field label="到期时间" value={formatTimestampWithTz(pool.endTime)} />
            <Field label="清算估值边界" value={`本金的 ${formatRate(100_000_000n + pool.autoLiquidateThreshold)}`} />
            <Field label="SP 凭证" value={pool.spCoin} mono />
            <Field label="JP 凭证" value={pool.jpCoin} mono />
          </div>

          {pool.state === PoolState.MATCH && (
            <p className="mt-4 text-xs text-slate-500">
              距募集截止 {remainingLabel(pool.settleTime, chainNow)}。截止后由授权账户执行匹配，届时才确定实际成交本金。
            </p>
          )}

          {poolData && (
            <div className="mt-5 grid grid-cols-2 gap-x-6 gap-y-3 border-t border-white/5 pt-4 md:grid-cols-4 text-sm">
              <Field label="匹配出借" value={`${displayUnits(poolData.settleAmountLend, STABLE, 2)} USDT`} />
              <Field label="匹配抵押" value={`${displayUnits(poolData.settleAmountBorrow, ETH, 6)} ETH`} />
              <Field label="到期结算出借" value={`${displayUnits(poolData.finishAmountLend, STABLE, 2)} USDT`} />
              <Field label="到期结算抵押" value={`${displayUnits(poolData.finishAmountBorrow, ETH, 6)} ETH`} />
            </div>
          )}
        </div>

        <div className="card border-warn/15 bg-warn/5">
          <h2 className="text-sm font-semibold text-warn">参与前须知</h2>
          <ul className="mt-2 space-y-1.5 text-xs leading-relaxed text-slate-400">
            <li>· 入金后不会立即生成 SP 或 JP，也不会立即获得借款；募集截止后统一匹配。</li>
            <li>· 未匹配部分需主动领取退款，不会自动退回钱包。</li>
            <li>· 成交后没有补仓、主动提前还款或执行期取回本金的入口。</li>
            <li>· 提前清算按完整约定期限计息，不按实际经过时间减免；抵押不足可能造成损失和结算停滞。</li>
            <li>· 「固定年利率」指约定计息参数，不代表保本或固定回收金额。</li>
          </ul>
        </div>
      </div>

      {/* 右：参与表单 */}
      <div className="card h-fit">
        <div className="mb-4 flex rounded-lg bg-ink-900 p-1">
          {(
            [
              ['lend', '出借 USDT'],
              ['borrow', '抵押 ETH 借款'],
            ] as const
          ).map(([key, label]) => (
            <button
              key={key}
              onClick={() => {
                setTab(key);
                setAmount('');
              }}
              className={`flex-1 rounded-md px-3 py-2 text-sm transition-colors ${
                tab === key ? 'bg-ink-700 text-white' : 'text-slate-400 hover:text-slate-200'
              }`}
            >
              {label}
            </button>
          ))}
        </div>

        <label className="label">{tab === 'lend' ? '出借金额（USDT）' : '抵押金额（ETH）'}</label>
        <div className="relative">
          <input
            className="field tabular pr-16"
            inputMode="decimal"
            placeholder="0.0"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
          />
          {tab === 'lend' && stableBalance > 0n && (
            <button
              className="absolute right-2 top-1/2 -translate-y-1/2 text-xs text-accent-500 hover:underline"
              onClick={() => setAmount(displayUnits(stableBalance, STABLE, STABLE))}
            >
              最大
            </button>
          )}
        </div>

        {tab === 'lend' ? (
          <div className="mt-2 space-y-1 text-xs text-slate-500">
            <div className="flex justify-between">
              <span>钱包余额</span>
              <span className="tabular">{displayOrLessThan(stableBalance, STABLE)} USDT</span>
            </div>
            <div className="flex justify-between">
              <span>剩余出借额度</span>
              <span className="tabular">{displayUnits(remaining, STABLE, 2)} USDT</span>
            </div>
            {needsApproval && parsed.value > 0n && (
              <div className="mt-2 rounded-md bg-accent-500/10 px-2.5 py-2 text-accent-500">
                需先授权 {displayUnits(parsed.value, STABLE, 6)} USDT 给池合约
              </div>
            )}
          </div>
        ) : (
          <div className="mt-2 text-xs text-slate-500">原生 ETH 抵押通过交易 value 支付，不需要 ERC20 授权。</div>
        )}

        {/* 预估（PRD 11.3.2）：明确标注为预估，不是确定结果 */}
        {parsed.value > 0n && pool.state === PoolState.MATCH && (
          <div className="mt-4 rounded-lg bg-ink-900 p-3 text-xs text-slate-400">
            <p className="mb-1 font-medium text-slate-300">匹配预估</p>
            <p>
              {tab === 'lend'
                ? '预计成交本金取决于抵押侧募集量与价格，最终以匹配结果为准。'
                : '预计可借额度取决于出借侧募集量与价格，最终以匹配结果为准。'}
            </p>
            <p className="mt-1 text-slate-500">预估不是确定借款、也不是确定凭证余额；未匹配部分不计借贷利息。</p>
          </div>
        )}

        <div className="mt-4 space-y-2">
          {needsApproval ? (
            <button
              className="btn-primary w-full"
              disabled={Boolean(blockReason) || isPending || confirming}
              onClick={handleApprove}
            >
              {isPending ? '请在钱包确认授权…' : `授权 ${displayUnits(parsed.value, STABLE, 2)} USDT`}
            </button>
          ) : (
            <button
              className="btn-primary w-full"
              disabled={Boolean(blockReason) || isPending || confirming}
              onClick={handleDeposit}
            >
              {isPending
                ? '请在钱包确认…'
                : confirming
                  ? '交易已广播，等待确认…'
                  : tab === 'lend'
                    ? '出借稳定币'
                    : '抵押 ETH 借款'}
            </button>
          )}
          {blockReason && <p className="text-center text-xs text-slate-500">{blockReason}</p>}
          <TxFeedback txKey={txKey} stage={current?.stage} writeError={writeError} />
        </div>
      </div>
    </div>
  );
}

function Field({ label, value, mono, accent }: { label: string; value: string; mono?: boolean; accent?: boolean }) {
  return (
    <div>
      <div className="text-xs text-slate-500">{label}</div>
      <div
        className={`mt-0.5 text-sm ${mono ? 'break-all font-mono text-xs' : 'tabular'} ${
          accent ? 'font-semibold text-accent-500' : 'text-slate-200'
        }`}
      >
        {value}
      </div>
    </div>
  );
}

function TxFeedback({ txKey, stage, writeError }: { txKey: string; stage?: string; writeError: unknown }) {
  // selector 必须返回稳定引用：返回新对象会让 zustand 每次比较都判定"变了"，
  // 触发无限渲染循环（React error #185 / getSnapshot should be cached）。
  const hash = useTxStore((s) => s.txs[txKey]?.hash);

  const revertMessage = (() => {
    if (!(writeError instanceof BaseError)) return undefined;
    const revert = writeError.walk((e) => e instanceof ContractFunctionRevertedError);
    if (revert instanceof ContractFunctionRevertedError) {
      return revert.reason ?? revert.shortMessage;
    }
    return writeError.shortMessage;
  })();

  if (revertMessage) return <p className="text-center text-xs text-negative">{revertMessage}</p>;
  if (stage === 'rejected') {
    return <p className="text-center text-xs text-warn">本次操作未提交（用户在钱包中拒绝）。</p>;
  }
  if (stage === 'broadcast-unknown') {
    return <p className="text-center text-xs text-warn">暂未取得明确提交结果，请先核对钱包记录，避免重复发送。</p>;
  }
  if (hash) {
    return (
      <a
        className="block text-center text-xs text-accent-500 hover:underline"
        href={`${DEPLOYMENT.explorer}/tx/${hash}`}
        target="_blank"
        rel="noreferrer"
      >
        查看交易 {hash.slice(0, 10)}…{hash.slice(-8)}
      </a>
    );
  }
  return null;
}
