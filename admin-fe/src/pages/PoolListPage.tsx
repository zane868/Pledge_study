import { Link } from 'react-router-dom';
import { useReadContract, useWriteContract, useWaitForTransactionReceipt, useBlock } from 'wagmi';
import { DEPLOYMENT, POOL_ABI, shortAddress } from '@shared/index';
import { formatUnits } from 'viem';
import { useState, useEffect } from 'react';
import { useAccount } from 'wagmi';
import { useMultisigApproval } from '@/hooks/useMultisigApproval';

const POOL = DEPLOYMENT.pool as `0x${string}`;

export function PoolListPage() {
  const { data: poolCount } = useReadContract({
    address: POOL,
    abi: POOL_ABI,
    functionName: 'poolLength',
  });

  const totalPools = poolCount ? Number(poolCount) : 0;

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h2 className="font-semibold text-white">池管理</h2>
          <p className="mt-1 text-sm text-slate-400">当前共 {totalPools} 个池</p>
        </div>
        <div className="flex gap-3">
          <Link to="/pools/create" className="btn-primary">
            创建新池
          </Link>
        </div>
      </div>

      <div className="space-y-4">
        {totalPools === 0 ? (
          <div className="card text-center text-slate-400">
            <p>暂无池，点击上方「创建新池」开始</p>
          </div>
        ) : (
          Array.from({ length: totalPools }, (_, i) => <PoolCard key={i} pid={i} />)
        )}
      </div>
    </div>
  );
}

function PoolCard({ pid }: { pid: number }) {
  const [open, setOpen] = useState(false);
  const [forceOpen, setForceOpen] = useState(false);
  const [forceMode, setForceMode] = useState<'finish' | 'liquidate'>('finish');

  const { data: baseInfo } = useReadContract({
    address: POOL,
    abi: POOL_ABI,
    functionName: 'poolBaseInfo',
    args: [BigInt(pid)],
  });

  const { data: poolState } = useReadContract({
    address: POOL,
    abi: POOL_ABI,
    functionName: 'getPoolState',
    args: [BigInt(pid)],
  });

  if (!baseInfo) {
    return (
      <div className="card">
        <p className="text-sm text-slate-400">加载池 #{pid + 1} 信息...</p>
      </div>
    );
  }

  const [
    _settleTime,
    _endTime,
    interestRate,
    maxSupply,
    lendSupply,
    borrowSupply,
    _martgageRate,
    _lendToken,
    _borrowToken,
    state,
  ] = baseInfo as any[];

  const stateLabel =
    state === 0
      ? 'MATCH 募集中'
      : state === 1
        ? 'EXECUTION 执行中'
        : state === 2
          ? 'FINISH 已结束'
          : state === 3
            ? 'LIQUIDATION 已清算'
            : state === 4
              ? 'UNDONE 已撤销'
              : '未知';

  const displayId = pid + 1;
  const interestRatePercent = ((Number(interestRate) * 100) / 1e8).toFixed(2);

  return (
    <div className="card">
      <div className="flex items-start justify-between">
        <div>
          <div className="flex items-center gap-3">
            <h3 className="text-lg font-semibold text-white">池 #{displayId}</h3>
            <span
              className={`rounded px-2 py-0.5 text-xs font-medium ${
                state === 0
                  ? 'bg-accent-500/20 text-accent-400'
                  : state === 1
                    ? 'bg-blue-500/20 text-blue-400'
                    : 'bg-slate-700 text-slate-400'
              }`}
            >
              {stateLabel}
            </span>
          </div>

          <div className="mt-4 grid grid-cols-2 gap-x-6 gap-y-3 md:grid-cols-4">
            <div>
              <div className="text-xs text-slate-500">年利率</div>
              <div className="mt-0.5 text-sm text-slate-200">{interestRatePercent}%</div>
            </div>
            <div>
              <div className="text-xs text-slate-500">出借上限</div>
              <div className="mt-0.5 text-sm text-slate-200">{formatUnits(maxSupply, 6)} USDT</div>
            </div>
            <div>
              <div className="text-xs text-slate-500">已入金</div>
              <div className="mt-0.5 text-sm text-slate-200">{formatUnits(lendSupply, 6)} USDT</div>
            </div>
            <div>
              <div className="text-xs text-slate-500">已抵押</div>
              <div className="mt-0.5 text-sm text-slate-200">{formatUnits(borrowSupply, 18)} ETH</div>
            </div>
          </div>
        </div>

        <div className="flex flex-wrap gap-2">
          {state === 0 && (
            <button className="btn-ghost text-sm" onClick={() => setOpen(!open)}>
              {open ? '收起' : '完成募集'}
            </button>
          )}
          {state === 1 && (
            <button className="btn-ghost text-sm" onClick={() => setOpen(!open)}>
              {open ? '收起' : '结束池子'}
            </button>
          )}
          {/* 强制操作：跳过时间校验，仅限应急使用 */}
          {state === 0 && (
            <button
              className="btn-ghost border-amber-500/40 text-sm text-amber-400 hover:bg-amber-500/10"
              onClick={() => setForceOpen(!forceOpen)}
            >
              {forceOpen ? '收起强制募集完成' : '强制募集完成'}
            </button>
          )}
          {state === 1 && (
            <button
              className="btn-ghost border-amber-500/40 text-sm text-amber-400 hover:bg-amber-500/10"
              onClick={() => setForceOpen(!forceOpen)}
            >
              {forceOpen ? '收起强制结算' : '强制结算'}
            </button>
          )}
          {state === 1 && (
            <button
              className="btn-ghost border-amber-500/40 text-sm text-amber-400 hover:bg-amber-500/10"
              onClick={() => {
                setForceMode('liquidate');
                setForceOpen(!forceOpen);
              }}
            >
              {forceOpen && forceMode === 'liquidate' ? '收起强制清算' : '强制清算'}
            </button>
          )}
        </div>
      </div>

      {open && (
        <OperationPanel pid={pid} displayId={displayId} poolState={Number(poolState ?? state)} />
      )}

      {forceOpen && (
        <ForceOperationPanel
          pid={pid}
          displayId={displayId}
          poolState={Number(poolState ?? state)}
          initialMode={forceMode}
        />
      )}
    </div>
  );
}

/**
 * 操作面板：三级流程（与自定义 multiSignature 合约一致，非 Safe）
 * 1. 创建授权申请 createApplication(pool)
 * 2. 其他管理员签名 signApplication(msghash)
 * 3. 达到阈值后，由申请人调用 pool.settle / pool.finish
 */
function OperationPanel({
  pid,
  displayId,
  poolState,
}: {
  pid: number;
  displayId: number;
  poolState: number;
}) {
  const { address } = useAccount();
  const [execHash, setExecHash] = useState<string | null>(null);
  const [execError, setExecError] = useState<string | null>(null);

  const { writeContractAsync } = useWriteContract();
  const { isLoading: execMining } = useWaitForTransactionReceipt({ hash: execHash as `0x${string}` });

  // 以当前连接地址作为申请人（msghash = keccak256(msg.sender, pool)）
  const approval = useMultisigApproval(address, DEPLOYMENT.pool);

  const action = poolState === 0 ? 'settle' : 'finish';
  const actionLabel = poolState === 0 ? '完成募集' : '结束池子';

  useEffect(() => {
    if (execHash && !execMining) {
      approval.refetch();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [execHash, execMining]);

  const execute = async () => {
    if (!address) return;
    setExecError(null);
    try {
      const h = await writeContractAsync({
        address: POOL,
        abi: POOL_ABI,
        functionName: action as 'settle' | 'finish',
        args: [BigInt(pid)],
      });
      setExecHash(h);
    } catch (e: any) {
      console.error('执行失败:', e);
      setExecError(e.shortMessage || e.message || '执行失败');
    }
  };

  return (
    <div className="mt-4 space-y-4 rounded-lg border border-accent-500/30 bg-accent-500/5 p-4">
      <div className="flex items-center justify-between">
        <h4 className="font-medium text-white">操作授权：{actionLabel}</h4>
        <span className="rounded bg-slate-800/60 px-2 py-0.5 text-xs text-slate-400">
          目标：Pool #{displayId} · {action}({pid})
        </span>
      </div>

      {approval.error && (
        <div className="rounded bg-red-500/10 p-3 text-sm text-red-400">❌ {approval.error}</div>
      )}
      {execError && (
        <div className="rounded bg-red-500/10 p-3 text-sm text-red-400">❌ {execError}</div>
      )}

      {/* 步骤 1：创建授权申请 */}
      <StepBox step={1} title="创建授权申请" done={approval.appCount > 0}>
        <p className="mb-2 text-xs text-slate-400">
          申请后由其他管理员对下方 msghash 签名。msghash = keccak256(申请人, Pool 合约)
        </p>
        <div className="mb-2 rounded bg-slate-800/50 p-2 font-mono text-xs text-slate-300 break-all">
          {approval.hash ?? '计算中…'}
        </div>
        {approval.appCount === 0 ? (
          <button
            className="btn-primary text-sm"
            onClick={approval.createApplication}
            disabled={approval.isPending}
          >
            {approval.isPending ? '提交中…' : '创建授权申请'}
          </button>
        ) : (
          <div className="text-sm text-green-400">✅ 申请已创建（申请人 {shortAddress(approval.applicant)}）</div>
        )}
      </StepBox>

      {/* 步骤 2：签名 */}
      <StepBox
        step={2}
        title={`签名（${approval.signCount}/${approval.required}）`}
        done={approval.approved}
      >
        {approval.appCount === 0 ? (
          <p className="text-xs text-slate-500">请先创建授权申请</p>
        ) : (
          <>
            <div className="mb-2 flex flex-wrap gap-1.5">
              {approval.signers.length === 0 ? (
                <span className="text-xs text-slate-500">暂无签名</span>
              ) : (
                approval.signers.map((s) => (
                  <span
                    key={s}
                    className="rounded bg-green-500/10 px-2 py-0.5 font-mono text-xs text-green-400"
                  >
                    {shortAddress(s)}
                  </span>
                ))
              )}
            </div>
            {!approval.hasSigned ? (
              <button
                className="btn-primary text-sm"
                onClick={approval.signApplication}
                disabled={approval.isPending}
              >
                {approval.isPending ? '签名中…' : '签名批准'}
              </button>
            ) : (
              <div className="flex items-center gap-2">
                <span className="text-sm text-green-400">✅ 你已签名</span>
                <button
                  className="btn-ghost text-sm"
                  onClick={approval.revokeSign}
                  disabled={approval.isPending}
                >
                  撤销签名
                </button>
              </div>
            )}
          </>
        )}
      </StepBox>

      {/* 步骤 3：执行 */}
      <StepBox step={3} title="执行操作" done={false}>
        {!approval.approved ? (
          <p className="text-xs text-slate-500">
            需要 {approval.required} 个签名达标后才可执行（当前 {approval.signCount}）
          </p>
        ) : (
          <>
            <div className="mb-2 rounded bg-green-500/10 p-2 text-xs text-green-400">
              ✅ 已获多签批准，可以执行 {action}({pid})
            </div>
            {!approval.isApplicant && (
              <p className="mb-2 text-xs text-amber-400">
                ⚠️ 该授权绑定申请人 {shortAddress(approval.applicant)}，
                需由该地址发起交易才能通过 validCall 校验。
              </p>
            )}
            <button
              className="btn-primary text-sm"
              onClick={execute}
              disabled={execMining || !approval.isApplicant}
            >
              {execMining ? '执行中…' : `执行 ${actionLabel}`}
            </button>
            {execHash && (
              <a
                className="mt-2 block text-xs text-accent-400 hover:underline"
                href={`${DEPLOYMENT.explorer}/tx/${execHash}`}
                target="_blank"
                rel="noopener noreferrer"
              >
                查看交易 →
              </a>
            )}
          </>
        )}
      </StepBox>
    </div>
  );
}

function StepBox({
  step,
  title,
  done,
  children,
}: {
  step: number;
  title: string;
  done: boolean;
  children: React.ReactNode;
}) {
  return (
    <div className="rounded bg-slate-900/40 p-3">
      <div className="mb-2 flex items-center gap-2">
        <span
          className={`flex h-5 w-5 items-center justify-center rounded-full text-xs ${
            done ? 'bg-green-500/20 text-green-400' : 'bg-slate-700 text-slate-300'
          }`}
        >
          {done ? '✓' : step}
        </span>
        <span className="text-sm font-medium text-slate-200">{title}</span>
      </div>
      {children}
    </div>
  );
}

/**
 * 强制操作面板：调用 forceSettle / forceFinish / forceLiquidate，跳过时间校验。
 *
 * 与常规 settle / finish / liquidate 的唯一区别是合约层面不校验 block.timestamp，
 * 其余状态机、计价、交换逻辑完全一致，且同样需要多签批准。
 *
 * 注意：多签授权的 msghash = keccak256(申请人, Pool) 与具体调用哪个函数无关，
 * 因此同一个授权可用于 settle / forceSettle / finish / forceFinish / forceLiquidate。
 *
 * EXECUTION 状态下有两个可选动作：
 * - 强制结算（forceFinish）：按到期时间结束池子
 * - 强制清算（forceLiquidate）：提前清算池子
 */
function ForceOperationPanel({
  pid,
  displayId,
  poolState,
  initialMode = 'finish',
}: {
  pid: number;
  displayId: number;
  poolState: number;
  initialMode?: 'finish' | 'liquidate';
}) {
  const { address } = useAccount();
  const [execHash, setExecHash] = useState<string | null>(null);
  const [execError, setExecError] = useState<string | null>(null);
  const [confirming, setConfirming] = useState(false);
  // EXECUTION 状态下可切换「强制结算 / 强制清算」
  const [mode, setMode] = useState<'finish' | 'liquidate'>(initialMode);

  const { writeContractAsync } = useWriteContract();
  const { isLoading: execMining } = useWaitForTransactionReceipt({ hash: execHash as `0x${string}` });

  const approval = useMultisigApproval(address, DEPLOYMENT.pool);

  // 池状态 0(MATCH) 对应募集阶段，1(EXECUTION) 对应执行阶段
  const isLiquidate = poolState === 1 && mode === 'liquidate';
  const action = poolState === 0 ? 'forceSettle' : isLiquidate ? 'forceLiquidate' : 'forceFinish';
  const actionLabel =
    poolState === 0 ? '强制募集完成' : isLiquidate ? '强制清算' : '强制结算';

  // 读取链上时间与池子时间，用于提示当前是否真的卡在时间校验上
  const { data: block } = useBlock({ watch: true });
  const { data: baseInfo } = useReadContract({
    address: POOL,
    abi: POOL_ABI,
    functionName: 'poolBaseInfo',
    args: [BigInt(pid)],
  });
  const chainNow = block?.timestamp ?? 0n;
  const settleTime = (baseInfo as any[])?.[0] as bigint | undefined;
  const endTime = (baseInfo as any[])?.[1] as bigint | undefined;
  // 各动作对应的时间闸门：settle/liquidate 看 settleTime，finish 看 endTime
  const timeGate = isLiquidate || poolState === 0 ? settleTime : endTime;
  const gateName = isLiquidate || poolState === 0 ? '募集截止' : '到期';
  const blockedByTime = timeGate !== undefined && chainNow > 0n && chainNow <= timeGate;

  useEffect(() => {
    if (execHash && !execMining) {
      approval.refetch();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [execHash, execMining]);

  const execute = async () => {
    if (!address) return;
    setExecError(null);
    try {
      const h = await writeContractAsync({
        address: POOL,
        abi: POOL_ABI,
        functionName: action as 'forceSettle' | 'forceFinish' | 'forceLiquidate',
        args: [BigInt(pid)],
      });
      setExecHash(h);
      setConfirming(false);
    } catch (e: any) {
      console.error('强制执行失败:', e);
      setExecError(e.shortMessage || e.message || '强制执行失败');
    }
  };

  return (
    <div className="mt-4 space-y-4 rounded-lg border border-amber-500/40 bg-amber-500/5 p-4">
      <div className="flex items-center justify-between">
        <h4 className="font-medium text-amber-300">强制操作：{actionLabel}</h4>
        <span className="rounded bg-slate-800/60 px-2 py-0.5 text-xs text-slate-400">
          目标：Pool #{displayId} · {action}({pid})
        </span>
      </div>

      <div className="rounded bg-amber-500/10 p-3 text-xs text-amber-300">
        ⚠️ 该操作<b>跳过时间校验</b>，无论是否到达募集截止/到期时间都可执行。
        状态机、计价与兑换逻辑与常规流程一致，且同样需要多签批准。请谨慎使用。
      </div>

      {/* EXECUTION 状态下有两个可选的强制动作 */}
      {poolState === 1 && (
        <div className="flex gap-2">
          <button
            className={`rounded px-3 py-1.5 text-sm ${
              !isLiquidate
                ? 'bg-amber-500/30 text-amber-200 ring-1 ring-amber-500/60'
                : 'bg-slate-800/60 text-slate-400 hover:text-slate-200'
            }`}
            onClick={() => {
              setMode('finish');
              setConfirming(false);
            }}
          >
            强制结算
          </button>
          <button
            className={`rounded px-3 py-1.5 text-sm ${
              isLiquidate
                ? 'bg-amber-500/30 text-amber-200 ring-1 ring-amber-500/60'
                : 'bg-slate-800/60 text-slate-400 hover:text-slate-200'
            }`}
            onClick={() => {
              setMode('liquidate');
              setConfirming(false);
            }}
          >
            强制清算
          </button>
        </div>
      )}

      {blockedByTime && (
        <div className="rounded bg-slate-800/40 p-3 text-xs text-slate-300">
          当前链上时间 {new Date(Number(chainNow) * 1000).toLocaleString()} 尚未超过
          {gateName}时间{' '}
          {timeGate ? new Date(Number(timeGate) * 1000).toLocaleString() : '—'}，
          常规接口会被时间校验拦截，强制接口可以绕过。
        </div>
      )}

      {approval.error && (
        <div className="rounded bg-red-500/10 p-3 text-sm text-red-400">❌ {approval.error}</div>
      )}
      {execError && (
        <div className="rounded bg-red-500/10 p-3 text-sm text-red-400">❌ {execError}</div>
      )}

      {/* 步骤 1：创建授权申请 */}
      <StepBox step={1} title="创建授权申请" done={approval.appCount > 0}>
        <p className="mb-2 text-xs text-slate-400">
          申请后由其他管理员对下方 msghash 签名。msghash = keccak256(申请人, Pool 合约)
        </p>
        <div className="mb-2 rounded bg-slate-800/50 p-2 font-mono text-xs text-slate-300 break-all">
          {approval.hash ?? '计算中…'}
        </div>
        {approval.appCount === 0 ? (
          <button
            className="btn-primary text-sm"
            onClick={approval.createApplication}
            disabled={approval.isPending}
          >
            {approval.isPending ? '提交中…' : '创建授权申请'}
          </button>
        ) : (
          <div className="text-sm text-green-400">
            ✅ 申请已创建（申请人 {shortAddress(approval.applicant)}）
          </div>
        )}
      </StepBox>

      {/* 步骤 2：签名 */}
      <StepBox step={2} title={`签名（${approval.signCount}/${approval.required}）`} done={approval.approved}>
        {approval.appCount === 0 ? (
          <p className="text-xs text-slate-500">请先创建授权申请</p>
        ) : (
          <>
            <div className="mb-2 flex flex-wrap gap-1.5">
              {approval.signers.length === 0 ? (
                <span className="text-xs text-slate-500">暂无签名</span>
              ) : (
                approval.signers.map((s) => (
                  <span
                    key={s}
                    className="rounded bg-green-500/10 px-2 py-0.5 font-mono text-xs text-green-400"
                  >
                    {shortAddress(s)}
                  </span>
                ))
              )}
            </div>
            {!approval.hasSigned ? (
              <button
                className="btn-primary text-sm"
                onClick={approval.signApplication}
                disabled={approval.isPending}
              >
                {approval.isPending ? '签名中…' : '签名批准'}
              </button>
            ) : (
              <div className="flex items-center gap-2">
                <span className="text-sm text-green-400">✅ 你已签名</span>
                <button
                  className="btn-ghost text-sm"
                  onClick={approval.revokeSign}
                  disabled={approval.isPending}
                >
                  撤销签名
                </button>
              </div>
            )}
          </>
        )}
      </StepBox>

      {/* 步骤 3：执行 */}
      <StepBox step={3} title="执行强制操作" done={false}>
        {!approval.approved ? (
          <p className="text-xs text-slate-500">
            需要 {approval.required} 个签名达标后才可执行（当前 {approval.signCount}）
          </p>
        ) : (
          <>
            <div className="mb-2 rounded bg-green-500/10 p-2 text-xs text-green-400">
              ✅ 已获多签批准，可以执行 {action}({pid})
            </div>
            {!approval.isApplicant && (
              <p className="mb-2 text-xs text-amber-400">
                ⚠️ 该授权绑定申请人 {shortAddress(approval.applicant)}，
                需由该地址发起交易才能通过 validCall 校验。
              </p>
            )}
            {!confirming ? (
              <button
                className="rounded bg-amber-500/20 px-4 py-2 text-sm font-medium text-amber-300 hover:bg-amber-500/30 disabled:opacity-50"
                onClick={() => setConfirming(true)}
                disabled={execMining || !approval.isApplicant}
              >
                {execMining ? '执行中…' : `强制执行 ${actionLabel}`}
              </button>
            ) : (
              <div className="rounded border border-amber-500/40 bg-amber-500/10 p-3">
                <p className="mb-2 text-sm text-amber-200">
                  确认对池 #{displayId} 执行 <b>{action}({pid})</b>？该操作不可撤销。
                  {isLiquidate && ' 清算会按当前抵押率折算分配给借贷双方。'}
                </p>
                <div className="flex gap-2">
                  <button
                    className="rounded bg-amber-500 px-4 py-1.5 text-sm font-medium text-slate-900 hover:bg-amber-400"
                    onClick={execute}
                    disabled={execMining}
                  >
                    {execMining ? '执行中…' : '确认执行'}
                  </button>
                  <button
                    className="btn-ghost text-sm"
                    onClick={() => setConfirming(false)}
                    disabled={execMining}
                  >
                    取消
                  </button>
                </div>
              </div>
            )}
            {execHash && (
              <a
                className="mt-2 block text-xs text-accent-400 hover:underline"
                href={`${DEPLOYMENT.explorer}/tx/${execHash}`}
                target="_blank"
                rel="noopener noreferrer"
              >
                查看交易 →
              </a>
            )}
          </>
        )}
      </StepBox>
    </div>
  );
}
