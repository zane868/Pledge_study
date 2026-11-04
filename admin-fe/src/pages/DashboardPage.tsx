import { Link } from 'react-router-dom';
import { DEPLOYMENT } from '@shared/index';

export function DashboardPage() {
  const contracts = [
    { label: 'PledgePool 合约', address: DEPLOYMENT.pool },
    { label: 'TestUSDT 合约', address: DEPLOYMENT.stable },
    { label: '多签合约', address: DEPLOYMENT.multisig },
    { label: '预言机合约', address: DEPLOYMENT.oracle },
    { label: 'SP 凭证', address: DEPLOYMENT.sp },
    { label: 'JP 凭证', address: DEPLOYMENT.jp },
  ];

  const quickActions = [
    { label: '创建新池', path: '/pools/create', color: 'accent' },
    { label: '查看池列表', path: '/pools', color: 'slate' },
    { label: '多签提案', path: '/proposals', color: 'slate' },
  ];

  return (
    <div className="space-y-6">
      <div>
        <h2 className="text-xl font-semibold text-white">运营后台总览</h2>
        <p className="mt-1 text-sm text-slate-400">
          Sepolia 测试网 · 当前多签阈值 2/3
        </p>
      </div>

      <div className="grid gap-5 md:grid-cols-2">
        <div className="card">
          <h3 className="font-semibold text-white">快捷操作</h3>
          <div className="mt-4 space-y-2">
            {quickActions.map((action) => (
              <Link
                key={action.path}
                to={action.path}
                className={`block rounded-lg border px-4 py-3 transition ${
                  action.color === 'accent'
                    ? 'border-accent-500/30 bg-accent-500/10 hover:bg-accent-500/20'
                    : 'border-slate-700 bg-slate-800/50 hover:bg-slate-700/50'
                }`}
              >
                <span className="text-sm font-medium text-white">{action.label}</span>
              </Link>
            ))}
          </div>
        </div>

        <div className="card">
          <h3 className="font-semibold text-white">合约地址</h3>
          <div className="mt-4 space-y-3">
            {contracts.map((contract) => (
              <div key={contract.address}>
                <div className="text-xs text-slate-500">{contract.label}</div>
                <div className="mt-0.5 flex items-center gap-2">
                  <a
                    href={`${DEPLOYMENT.explorer}/address/${contract.address}`}
                    target="_blank"
                    rel="noreferrer"
                    className="font-mono text-xs text-accent-400 hover:underline"
                  >
                    {contract.address.slice(0, 10)}…{contract.address.slice(-8)}
                  </a>
                  <button
                    className="text-xs text-slate-500 hover:text-slate-300"
                    onClick={() => navigator.clipboard.writeText(contract.address)}
                    title="复制地址"
                  >
                    复制
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>

      <div className="card border-amber-500/20 bg-amber-500/5">
        <h3 className="font-semibold text-amber-400">系统状态</h3>
        <ul className="mt-3 space-y-2 text-sm text-slate-300">
          <li>✓ 后端同步服务运行中</li>
          <li>✓ 多签合约可用（2/3 阈值）</li>
          <li>✓ TestUSDT 支持 EIP-2612 permit</li>
        </ul>
      </div>
    </div>
  );
}
