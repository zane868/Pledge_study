import { useMemo, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import {
  useAccount,
  useReadContract,
  useWriteContract,
  useWaitForTransactionReceipt,
  useDeployContract,
} from 'wagmi';
import { DEPLOYMENT, DEBT_TOKEN_ABI, POOL_ABI, parseUnits, BASE_DECIMAL, NATIVE_ETH, shortAddress } from '@shared/index';
import { useMultisigApproval } from '@/hooks/useMultisigApproval';
import { WalletGate } from '@/components/WalletGate';
import { DEBT_TOKEN_BYTECODE } from '@/lib/debtTokenBytecode';

const POOL = DEPLOYMENT.pool as `0x${string}`;
const MULTISIG = DEPLOYMENT.multisig as `0x${string}`;
const STABLE = DEPLOYMENT.stable as `0x${string}`;
const STABLE_DECIMALS = DEPLOYMENT.stableDecimals;
// 百分比（%）→ 1e8 基数的缩放：1e8 = 100%，x% = x × 1e6
const PCT_DECIMALS = 6;

interface PoolParams {
  interestRate: string; // 年利率 %
  martgageRate: string; // 抵押率 %
  maxSupply: string; // 出借上限 USDT
  autoLiquidateThreshold: string; // 清算边界增量 %
  settleTime: string; // 募集截止 datetime-local
  endTime: string; // 到期 datetime-local
}

function defaultParams(): PoolParams {
  const now = new Date();
  const tomorrow = new Date(now.getTime() + 24 * 60 * 60 * 1000);
  const nextWeek = new Date(now.getTime() + 7 * 24 * 60 * 60 * 1000);
  const toLocal = (d: Date) => {
    const p = (n: number) => String(n).padStart(2, '0');
    return `${d.getFullYear()}-${p(d.getMonth() + 1)}-${p(d.getDate())}T${p(d.getHours())}:${p(d.getMinutes())}`;
  };
  return {
    interestRate: '20',
    martgageRate: '200',
    maxSupply: '10000',
    autoLiquidateThreshold: '20',
    settleTime: toLocal(tomorrow),
    endTime: toLocal(nextWeek),
  };
}

/**
 * 仅校验表单参数本身（时间、比例、额度），不涉及 SP/JP 是否已部署。
 * 单位口径见 Pledge_PRD.md 6.2.1：比例用 1e8 基数，金额用稳定币最小单位。
 * 返回 null 表示参数合法，否则返回错误文案。
 */
function paramError(params: PoolParams): string | null {
  try {
    const settleTime = Math.floor(new Date(params.settleTime).getTime() / 1000);
    const endTime = Math.floor(new Date(params.endTime).getTime() / 1000);
    if (!Number.isFinite(settleTime) || !Number.isFinite(endTime)) {
      return '请填写有效的募集截止与到期时间';
    }
    if (endTime <= settleTime) {
      return '到期时间必须晚于募集截止时间';
    }
    if (settleTime <= Math.floor(Date.now() / 1000)) {
      return '募集截止时间必须晚于当前时间';
    }
    const interestRate = parseUnits(params.interestRate, PCT_DECIMALS);
    const martgageRate = parseUnits(params.martgageRate, PCT_DECIMALS);
    parseUnits(params.autoLiquidateThreshold, PCT_DECIMALS);
    const maxSupply = parseUnits(params.maxSupply, STABLE_DECIMALS);
    void interestRate;
    if (martgageRate <= BASE_DECIMAL) {
      return '抵押率必须大于 100%（ETH 超额抵押）';
    }
    if (maxSupply <= 0n) {
      return '出借上限必须大于 0';
    }
    return null;
  } catch (e) {
    return (e as Error).message;
  }
}

/**
 * 在参数合法且 SP/JP 已部署的前提下，组装 createPoolInfo 的链上参数。
 * 只有走到第四步（创建池）才需要完整 args。
 */
function buildPoolArgs(
  params: PoolParams,
  spAddress: string | null,
  jpAddress: string | null,
): { args: (bigint | `0x${string}`)[] | null; error: string | null } {
  const err = paramError(params);
  if (err) return { args: null, error: err };

  if (!spAddress || !jpAddress) {
    return { args: null, error: '请先部署 SP / JP 凭证' };
  }
  if (spAddress.toLowerCase() === jpAddress.toLowerCase()) {
    return { args: null, error: 'SP 与 JP 不能是同一地址' };
  }

  const settleTime = BigInt(Math.floor(new Date(params.settleTime).getTime() / 1000));
  const endTime = BigInt(Math.floor(new Date(params.endTime).getTime() / 1000));
  const interestRate = parseUnits(params.interestRate, PCT_DECIMALS);
  const martgageRate = parseUnits(params.martgageRate, PCT_DECIMALS);
  const threshold = parseUnits(params.autoLiquidateThreshold, PCT_DECIMALS);
  const maxSupply = parseUnits(params.maxSupply, STABLE_DECIMALS);

  return {
    args: [
      settleTime,
      endTime,
      interestRate,
      maxSupply,
      martgageRate,
      STABLE,
      NATIVE_ETH,
      spAddress as `0x${string}`,
      jpAddress as `0x${string}`,
      threshold,
    ],
    error: null,
  };
}

const STEPS = ['参数', '部署凭证', '授权与配置', '创建池'] as const;
type StepIndex = 1 | 2 | 3 | 4;

export function CreatePoolPage() {
  const navigate = useNavigate();
  const [step, setStep] = useState<StepIndex>(1);
  const [params, setParams] = useState<PoolParams>(defaultParams());
  const [spAddress, setSpAddress] = useState<string | null>(null);
  const [jpAddress, setJpAddress] = useState<string | null>(null);

  const poolArgs = useMemo(() => buildPoolArgs(params, spAddress, jpAddress), [params, spAddress, jpAddress]);

  return (
    <div className="mx-auto max-w-3xl space-y-6">
      <div>
        <button className="text-sm text-slate-400 hover:text-slate-200" onClick={() => navigate(-1)}>
          ← 返回
        </button>
        <h2 className="mt-2 text-xl font-semibold text-white">创建新池</h2>
        <p className="mt-1 text-sm text-slate-400">
          每池需要独立的 SP / JP 凭证，创建分为四步：填写参数 → 部署凭证 → 多签授权与 minter → 提交创建。
        </p>
      </div>

      {/* 中断风险提示 */}
      <div className="card border-warn/30 bg-warn/10 text-sm text-warn">
        <p className="font-medium">⚠️ 请勿中途关闭或刷新本页面</p>
        <p className="mt-1 text-xs leading-relaxed text-slate-400">
          创建进度（当前步骤、已填参数、已部署的 SP/JP 地址）只保存在本页面内存中，关闭或刷新后不会恢复，
          需要从头重新填写并重新部署凭证（会产生额外的合约创建 gas）。
        </p>
        <p className="mt-1 text-xs leading-relaxed text-slate-400">
          如果已经部署过 SP / JP 但页面中断，合约本身仍存在于链上，请记录下部署时的合约地址，再回到本页核对，避免重复部署。
        </p>
      </div>

      {/* 步骤指示器 */}
      <div className="flex items-center gap-2 text-xs text-slate-400">
        {STEPS.map((label, i) => {
          const n = (i + 1) as StepIndex;
          const active = n === step;
          const done = n < step;
          return (
            <div key={label} className="flex items-center gap-2">
              <span
                className={`flex h-6 w-6 items-center justify-center rounded-full ${
                  active ? 'bg-accent-500 text-slate-900' : done ? 'bg-green-500/20 text-green-400' : 'bg-slate-700 text-slate-400'
                }`}
              >
                {done ? '✓' : n}
              </span>
              <span className={active ? 'text-white' : ''}>{label}</span>
              {n < 4 && <span className="text-slate-600">→</span>}
            </div>
          );
        })}
      </div>

      {step === 1 && (
        <ParamsStep
          params={params}
          setParams={setParams}
          onNext={() => setStep(2)}
          error={paramError(params)}
        />
      )}
      {step === 2 && (
        <DeployStep
          spAddress={spAddress}
          jpAddress={jpAddress}
          onSpDeployed={setSpAddress}
          onJpDeployed={setJpAddress}
          onBack={() => setStep(1)}
          onNext={() => setStep(3)}
        />
      )}
      {step === 3 && (
        <SetupStep
          spAddress={spAddress}
          jpAddress={jpAddress}
          onBack={() => setStep(2)}
          onNext={() => setStep(4)}
        />
      )}
      {step === 4 && (
        <CreateStep
          poolArgs={poolArgs}
          onBack={() => setStep(3)}
        />
      )}
    </div>
  );
}

// ─── Step 1：参数 ───────────────────────────────────────────────────

function ParamsStep({
  params,
  setParams,
  onNext,
  error,
}: {
  params: PoolParams;
  setParams: (p: PoolParams) => void;
  onNext: () => void;
  error: string | null;
}) {
  const set = (name: keyof PoolParams) => (e: React.ChangeEvent<HTMLInputElement>) =>
    setParams({ ...params, [name]: e.target.value });

  // 参数预览：基于表单原始值做单位换算，不依赖 SP/JP 是否部署
  const preview = (() => {
    try {
      const supply = parseUnits(params.maxSupply, STABLE_DECIMALS);
      const interest = Number(parseUnits(params.interestRate, PCT_DECIMALS)) / 1e6;
      const mortgage = Number(parseUnits(params.martgageRate, PCT_DECIMALS)) / 1e6;
      const threshold = Number(parseUnits(params.autoLiquidateThreshold, PCT_DECIMALS)) / 1e6;
      return { supply: supply.toString(), interest, mortgage, threshold };
    } catch {
      return null;
    }
  })();

  return (
    <form
      className="card space-y-5"
      onSubmit={(e) => {
        e.preventDefault();
        if (!error) onNext();
      }}
    >
      <div>
        <label className="label">借款资产地址（USDT）</label>
        <input type="text" className="field font-mono" value={STABLE} readOnly />
        <p className="mt-1 text-xs text-slate-500">借款资产固定为当前部署的 6 位 TestUSDT</p>
      </div>

      <div>
        <label className="label">抵押资产（原生 ETH）</label>
        <input type="text" className="field font-mono" value={NATIVE_ETH} readOnly />
        <p className="mt-1 text-xs text-slate-500">零地址表示原生 ETH</p>
      </div>

      <div className="grid gap-5 md:grid-cols-2">
        <div>
          <label className="label">固定年利率（%）</label>
          <input
            type="number"
            name="interestRate"
            className="field tabular"
            value={params.interestRate}
            onChange={set('interestRate')}
            step="0.01"
            min="0"
            required
          />
          <p className="mt-1 text-xs text-slate-500">如 20% 输入 20（1e8 基数 20000000）</p>
        </div>
        <div>
          <label className="label">募集抵押率（%）</label>
          <input
            type="number"
            name="martgageRate"
            className="field tabular"
            value={params.martgageRate}
            onChange={set('martgageRate')}
            step="0.01"
            min="100"
            required
          />
          <p className="mt-1 text-xs text-slate-500">ETH 超额抵押，需大于 100%</p>
        </div>
      </div>

      <div className="grid gap-5 md:grid-cols-2">
        <div>
          <label className="label">出借上限（USDT）</label>
          <input
            type="number"
            name="maxSupply"
            className="field tabular"
            value={params.maxSupply}
            onChange={set('maxSupply')}
            step="0.01"
            min="0"
            required
          />
        </div>
        <div>
          <label className="label">清算边界增量（%）</label>
          <input
            type="number"
            name="autoLiquidateThreshold"
            className="field tabular"
            value={params.autoLiquidateThreshold}
            onChange={set('autoLiquidateThreshold')}
            step="0.01"
            min="0"
            required
          />
          <p className="mt-1 text-xs text-slate-500">如 20% 输入 20（边界 = 本金 × 120%）</p>
        </div>
      </div>

      <div className="grid gap-5 md:grid-cols-2">
        <div>
          <label className="label">募集截止时间</label>
          <input
            type="datetime-local"
            name="settleTime"
            className="field"
            value={params.settleTime}
            onChange={set('settleTime')}
            required
          />
        </div>
        <div>
          <label className="label">到期时间</label>
          <input
            type="datetime-local"
            name="endTime"
            className="field"
            value={params.endTime}
            onChange={set('endTime')}
            required
          />
        </div>
      </div>

      <div className="border-t border-slate-700 pt-5">
        <h3 className="text-sm font-medium text-white">参数预览</h3>
        <div className="mt-3 space-y-1.5 rounded bg-slate-800/50 p-3 font-mono text-xs text-slate-300">
          <div>年利率：{params.interestRate}%（1e8 基数 {preview ? Math.round(preview.interest * 1e6) : '—'}）</div>
          <div>抵押率：{params.martgageRate}%</div>
          <div>出借上限：{params.maxSupply} USDT（6 位原始 {preview ? preview.supply : '—'}）</div>
          <div>清算边界增量：{params.autoLiquidateThreshold}%</div>
          <div>募集截止：{params.settleTime} → 到期：{params.endTime}</div>
        </div>
        {error && (
          <p className="mt-3 text-sm text-amber-400">⚠️ {error}</p>
        )}
      </div>

      <div className="flex gap-3">
        <button type="submit" className="btn-primary flex-1" disabled={Boolean(error)}>
          下一步：部署凭证
        </button>
      </div>
    </form>
  );
}

// ─── Step 2：部署 SP / JP ───────────────────────────────────────────

function DeployStep({
  spAddress,
  jpAddress,
  onSpDeployed,
  onJpDeployed,
  onBack,
  onNext,
}: {
  spAddress: string | null;
  jpAddress: string | null;
  onSpDeployed: (a: string) => void;
  onJpDeployed: (a: string) => void;
  onBack: () => void;
  onNext: () => void;
}) {
  const bothDeployed = Boolean(spAddress && jpAddress);
  return (
    <div className="space-y-4">
      <DeployVoucher
        name="Pledge SP"
        symbol="SP"
        deployedAddress={spAddress}
        onDeployed={onSpDeployed}
      />
      <DeployVoucher
        name="Pledge JP"
        symbol="JP"
        deployedAddress={jpAddress}
        onDeployed={onJpDeployed}
      />
      <div className="flex gap-3">
        <button type="button" className="btn-ghost" onClick={onBack}>
          上一步
        </button>
        <button
          type="button"
          className="btn-primary flex-1"
          disabled={!bothDeployed}
          onClick={onNext}
        >
          {bothDeployed ? '下一步：授权与配置' : '请先部署 SP 与 JP 两个凭证'}
        </button>
      </div>
    </div>
  );
}

function DeployVoucher({
  name,
  symbol,
  deployedAddress,
  onDeployed,
}: {
  name: string;
  symbol: string;
  deployedAddress: string | null;
  onDeployed: (a: string) => void;
}) {
  const { deployContractAsync, isPending } = useDeployContract();
  const [hash, setHash] = useState<`0x${string}` | null>(null);
  const [error, setError] = useState<string | null>(null);
  const { data: receipt } = useWaitForTransactionReceipt({ hash: hash ?? undefined });

  // 交易确认后，从 receipt 取部署地址
  if (receipt && hash && !deployedAddress && receipt.contractAddress) {
    onDeployed(receipt.contractAddress);
  }

  const deploy = async () => {
    setError(null);
    try {
      const h = await deployContractAsync({
        abi: DEBT_TOKEN_ABI,
        bytecode: DEBT_TOKEN_BYTECODE,
        args: [name, symbol, MULTISIG],
      });
      setHash(h);
    } catch (e: any) {
      setError(e.shortMessage || e.message || '部署失败');
    }
  };

  return (
    <div className="card">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="font-medium text-white">{symbol} 凭证</h3>
          <p className="mt-1 text-xs text-slate-500">
            DebtToken（name={name}，symbol={symbol}，多签 {shortAddress(MULTISIG)}）
          </p>
        </div>
        {deployedAddress ? (
          <span className="rounded bg-green-500/10 px-2 py-0.5 font-mono text-xs text-green-400">
            已部署 {shortAddress(deployedAddress)}
          </span>
        ) : (
          <button className="btn-primary text-sm" onClick={deploy} disabled={isPending || Boolean(hash)}>
            {isPending ? '部署中…' : hash ? '等待确认…' : `部署 ${symbol}`}
          </button>
        )}
      </div>
      {deployedAddress && (
        <a
          className="mt-2 block break-all font-mono text-xs text-accent-400 hover:underline"
          href={`${DEPLOYMENT.explorer}/address/${deployedAddress}`}
          target="_blank"
          rel="noreferrer"
        >
          {deployedAddress}
        </a>
      )}
      {hash && (
        <a
          className="mt-2 block text-xs text-accent-400 hover:underline"
          href={`${DEPLOYMENT.explorer}/tx/${hash}`}
          target="_blank"
          rel="noreferrer"
        >
          查看部署交易 →
        </a>
      )}
      {error && <p className="mt-2 text-sm text-red-400">❌ {error}</p>}
    </div>
  );
}

// ─── Step 3：多签授权 + addMinter ───────────────────────────────────

function SetupStep({
  spAddress,
  jpAddress,
  onBack,
  onNext,
}: {
  spAddress: string | null;
  jpAddress: string | null;
  onBack: () => void;
  onNext: () => void;
}) {
  return (
    <div className="space-y-4">
      <div className="card border-warn/20 bg-warn/5 text-xs text-slate-400">
        <p>本步骤为 SP / JP 建立多签授权，并把池合约设为凭证的 minter。</p>
        <p className="mt-1">
          申请人固定为 operator（<span className="font-mono">{shortAddress(DEPLOYMENT.operator)}</span>）。
          多签门槛 2/3：operator 创建授权申请后，需切换到另外两个管理员钱包
          （<span className="font-mono">{shortAddress(DEPLOYMENT.admins[1])}</span>、<span className="font-mono">{shortAddress(DEPLOYMENT.admins[2])}</span>）
          各自点「签名批准」，凑满 2 个签名后才能 addMinter。
        </p>
        <p className="mt-1 text-amber-300">
          签名时在钱包里切换管理员账户即可，申请与「签名批准」按钮会正常显示，无需刷新页面。
        </p>
      </div>
      <VoucherSetup voucher={spAddress} label="SP" />
      <VoucherSetup voucher={jpAddress} label="JP" />
      <div className="flex gap-3">
        <button type="button" className="btn-ghost" onClick={onBack}>
          上一步
        </button>
        <button type="button" className="btn-primary flex-1" onClick={onNext}>
          下一步：创建池
        </button>
      </div>
    </div>
  );
}

function VoucherSetup({ voucher, label }: { voucher: string | null; label: string }) {
  const { address } = useAccount();
  // 申请人固定为 operator：部署凭证与 createApplication 都由 operator 发起，
  // msghash = keccak256(operator, voucher)。签名由任意管理员用自己钱包完成，
  // 所以这里不能用「当前钱包」当申请人，否则切换管理员钱包后 msghash 就对不上、
  // 申请会"消失"，其他管理员无法签名。
  const approval = useMultisigApproval(DEPLOYMENT.operator, voucher ?? undefined);

  const [minterHash, setMinterHash] = useState<`0x${string}` | null>(null);
  const [minterError, setMinterError] = useState<string | null>(null);
  const { writeContractAsync } = useWriteContract();
  const { isLoading: minterMining } = useWaitForTransactionReceipt({ hash: minterHash ?? undefined });

  const { data: isMinter } = useReadContract({
    address: voucher as `0x${string}` | undefined,
    abi: DEBT_TOKEN_ABI,
    functionName: 'isMinter',
    args: [POOL],
    query: { enabled: !!voucher },
  });

  const addMinter = async () => {
    if (!voucher) return;
    setMinterError(null);
    try {
      const h = await writeContractAsync({
        address: voucher as `0x${string}`,
        abi: DEBT_TOKEN_ABI,
        functionName: 'addMinter',
        args: [POOL],
      });
      setMinterHash(h);
    } catch (e: any) {
      setMinterError(e.shortMessage || e.message || 'addMinter 失败');
    }
  };

  if (!voucher) {
    return <div className="card text-sm text-slate-400">等待部署 {label} 凭证…</div>;
  }

  return (
    <div className="card">
      <div className="flex items-center justify-between">
        <h3 className="font-medium text-white">{label} 授权与 minter</h3>
        <span className="font-mono text-xs text-slate-500">{shortAddress(voucher)}</span>
      </div>

      {approval.error && <p className="mt-2 text-sm text-red-400">❌ {approval.error}</p>}
      {minterError && <p className="mt-2 text-sm text-red-400">❌ {minterError}</p>}

      {/* 复制签名链接：发给其他管理员，点开即可签名，无需进入创建向导 */}
      {voucher && (
        <CopySignLink voucher={voucher} label={label} />
      )}

      {/* 授权 */}
      <div className="mt-3 rounded bg-slate-900/40 p-3">
        <p className="text-xs text-slate-400">
          授权（{approval.signCount}/{approval.required}）
          {approval.approved ? ' ✅ 已达标' : ' · 待其他管理员签名'}
        </p>
        <p className="mt-1 text-xs text-slate-500">
          申请人 <span className="font-mono">{shortAddress(DEPLOYMENT.operator)}</span> · 当前钱包{' '}
          <span className="font-mono">{shortAddress(address)}</span>
          {approval.hasSigned ? ' · 你已签名' : ' · 你未签名'}
        </p>
        {approval.hash && (
          <p className="mt-1 break-all font-mono text-xs text-slate-500">{approval.hash}</p>
        )}
        <div className="mt-2 flex flex-wrap gap-2">
          {approval.appCount === 0 ? (
            <button className="btn-primary text-sm" onClick={approval.createApplication} disabled={approval.isPending}>
              {approval.isPending ? '提交中…' : '创建授权申请'}
            </button>
          ) : (
            <>
              {!approval.hasSigned ? (
                <button className="btn-primary text-sm" onClick={approval.signApplication} disabled={approval.isPending}>
                  {approval.isPending ? '签名中…' : '签名批准'}
                </button>
              ) : (
                <span className="text-sm text-green-400">✅ 你已签名</span>
              )}
              {approval.approved && (
                <button
                  className="btn-ghost text-sm"
                  onClick={addMinter}
                  disabled={minterMining || Boolean(isMinter)}
                >
                  {minterMining ? '设置中…' : isMinter ? '✅ 池已是 minter' : '设置池为 minter'}
                </button>
              )}
            </>
          )}
        </div>
      </div>
    </div>
  );
}

// ─── 复制签名链接 ───────────────────────────────────────────────────

function CopySignLink({ voucher, label }: { voucher: string; label: string }) {
  const [copied, setCopied] = useState(false);

  const copy = async () => {
    const url = `${window.location.origin}/admin/sign?to=${voucher}&label=${encodeURIComponent(label)}`;
    try {
      await navigator.clipboard.writeText(url);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    } catch {
      // 剪贴板不可用时降级为手动复制提示
      window.prompt('复制下面的签名链接发给其他管理员：', url);
    }
  };

  return (
    <div className="mt-3 rounded bg-blue-500/10 p-3">
      <p className="text-xs text-slate-300">
        发给其他管理员签名：点「复制签名链接」，对方打开后连接钱包即可直接签名，无需进入创建向导。
      </p>
      <div className="mt-2 flex items-center gap-2">
        <button className="btn-ghost text-sm" onClick={copy}>
          {copied ? '✅ 已复制' : `复制 ${label} 签名链接`}
        </button>
      </div>
    </div>
  );
}

// ─── Step 4：创建池 ─────────────────────────────────────────────────

function CreateStep({
  poolArgs,
  onBack,
}: {
  poolArgs: { args: unknown[] | null; error: string | null };
  onBack: () => void;
}) {
  const { address } = useAccount();
  const approval = useMultisigApproval(address, DEPLOYMENT.pool);

  const [createHash, setCreateHash] = useState<`0x${string}` | null>(null);
  const [createError, setCreateError] = useState<string | null>(null);
  const { writeContractAsync } = useWriteContract();
  const { isLoading: creating, data: receipt } = useWaitForTransactionReceipt({ hash: createHash ?? undefined });

  const isOperator = address?.toLowerCase() === DEPLOYMENT.operator.toLowerCase();

  // 交易确认成功：状态为 1 表示 createPoolInfo 成功，进入成功态（防止重复点击）
  const succeeded = Boolean(createHash && receipt && receipt.status === 'success');

  const createPool = async () => {
    if (!poolArgs.args) return;
    setCreateError(null);
    try {
      const h = await writeContractAsync({
        address: POOL,
        abi: POOL_ABI,
        functionName: 'createPoolInfo',
        args: poolArgs.args as never,
      });
      setCreateHash(h);
    } catch (e: any) {
      setCreateError(e.shortMessage || e.message || '创建池失败');
    }
  };

  if (succeeded) {
    return (
      <div className="space-y-4">
        <div className="card border-green-500/30 bg-green-500/5 text-center">
          <h3 className="text-lg font-semibold text-green-400">✅ 池创建成功</h3>
          <p className="mt-2 text-sm text-slate-300">
            交易已确认，新池已写入链上并进入「募集中」状态。
          </p>
          <a
            className="mt-2 block break-all font-mono text-xs text-accent-400 hover:underline"
            href={`${DEPLOYMENT.explorer}/tx/${createHash}`}
            target="_blank"
            rel="noreferrer"
          >
            查看创建交易 →
          </a>
          <div className="mt-4 flex justify-center gap-3">
            <Link to="/pools" className="btn-primary">
              返回池管理查看新池
            </Link>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <div className="card">
        <h3 className="font-medium text-white">提交创建交易</h3>
        <p className="mt-1 text-xs text-slate-400">
          createPoolInfo 需通过多签授权（validCall 校验 keccak256(调用者, Pool)）。
          {isOperator
            ? ' 当前钱包是 operator，部署时已预授权 Pool，可直接创建。'
            : ' 当前钱包非 operator，需先完成对 Pool 的多签授权。'}
        </p>

        {poolArgs.error && <p className="mt-2 text-sm text-amber-400">⚠️ {poolArgs.error}</p>}
        {approval.error && <p className="mt-2 text-sm text-red-400">❌ {approval.error}</p>}
        {createError && <p className="mt-2 text-sm text-red-400">❌ {createError}</p>}

        {/* 非 operator 需要先授权 Pool */}
        {!isOperator && !approval.approved && (
          <div className="mt-3 rounded bg-slate-900/40 p-3">
            <p className="text-xs text-slate-400">
              授权 Pool（{approval.signCount}/{approval.required}）
            </p>
            {approval.hash && <p className="mt-1 break-all font-mono text-xs text-slate-500">{approval.hash}</p>}
            <div className="mt-2 flex flex-wrap gap-2">
              {approval.appCount === 0 ? (
                <button className="btn-primary text-sm" onClick={approval.createApplication} disabled={approval.isPending}>
                  {approval.isPending ? '提交中…' : '创建授权申请'}
                </button>
              ) : !approval.hasSigned ? (
                <button className="btn-primary text-sm" onClick={approval.signApplication} disabled={approval.isPending}>
                  {approval.isPending ? '签名中…' : '签名批准'}
                </button>
              ) : (
                <span className="text-sm text-green-400">✅ 你已签名，等待其他管理员签名</span>
              )}
            </div>
          </div>
        )}

        <div className="mt-4 flex gap-3">
          <button type="button" className="btn-ghost" onClick={onBack}>
            上一步
          </button>
          <button
            type="button"
            className="btn-primary flex-1"
            disabled={creating || succeeded || Boolean(poolArgs.error) || (!isOperator && !approval.approved)}
            onClick={createPool}
          >
            {creating ? '创建中…' : '创建池'}
          </button>
        </div>

        <div className="mt-3">
          <WalletGate action="创建池">
            <p className="text-xs text-green-400">✅ 管理员钱包已连接，可以创建池</p>
          </WalletGate>
        </div>

        {createHash && (
          <a
            className="mt-3 block text-xs text-accent-400 hover:underline"
            href={`${DEPLOYMENT.explorer}/tx/${createHash}`}
            target="_blank"
            rel="noreferrer"
          >
            查看创建交易 →
          </a>
        )}
      </div>

      <div className="card text-xs text-slate-500">
        <p>创建成功后：</p>
        <ul className="mt-2 space-y-1">
          <li>· 池立即进入 MATCH（募集）状态，链上索引 = 当前 poolLength。</li>
          <li>· 建议回到「池管理」列表核对新池参数与提交值是否一致。</li>
          <li>· 创建前确认 SP / JP 已 addMinter，否则资金进入后无法领取或赎回凭证。</li>
        </ul>
      </div>
    </div>
  );
}
