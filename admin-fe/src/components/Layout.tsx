import { Link, useLocation, useNavigate } from 'react-router-dom';
import { useAccount, useConnect, useDisconnect } from 'wagmi';
import { DEPLOYMENT, shortAddress } from '@shared/index';
import { getUsername, logout } from '@/lib/auth';

const ADMINS = DEPLOYMENT.admins.map((a) => a.toLowerCase());

export function Layout({ children }: { children: React.ReactNode }) {
  const location = useLocation();
  const navigate = useNavigate();
  const { address, isConnected } = useAccount();
  const { connect, connectors } = useConnect();
  const { disconnect } = useDisconnect();

  const nav = [
    { path: '/', label: '总览' },
    { path: '/pools', label: '池管理' },
    { path: '/proposals', label: '多签提案' },
  ];

  const username = getUsername();
  const isAdminWallet = Boolean(address && ADMINS.includes(address.toLowerCase()));

  const doLogout = async () => {
    await logout();
    navigate('/login', { replace: true });
  };

  return (
    <div className="min-h-screen bg-ink-900">
      <header className="border-b border-slate-700 bg-ink-800">
        <div className="mx-auto flex max-w-7xl items-center justify-between px-4 py-4">
          <div className="flex items-center gap-8">
            <h1 className="text-lg font-semibold text-white">Pledge 运营后台</h1>
            <nav className="flex gap-1">
              {nav.map((item) => (
                <Link
                  key={item.path}
                  to={item.path}
                  className={`rounded px-3 py-1.5 text-sm transition ${
                    location.pathname === item.path
                      ? 'bg-accent-500/20 text-accent-400'
                      : 'text-slate-400 hover:bg-slate-700 hover:text-slate-200'
                  }`}
                >
                  {item.label}
                </Link>
              ))}
            </nav>
          </div>

          <div className="flex items-center gap-4">
            {/* 钱包：可选，仅链上操作需要 */}
            {isConnected && address ? (
              <div className="flex items-center gap-2 text-right">
                <div>
                  <div className="text-xs text-slate-500">
                    {isAdminWallet ? '管理员钱包' : '非管理员钱包'}
                  </div>
                  <div className={`font-mono text-xs ${isAdminWallet ? 'text-slate-300' : 'text-amber-400'}`}>
                    {shortAddress(address)}
                  </div>
                </div>
                <button className="btn-ghost text-xs" onClick={() => disconnect()}>
                  断开
                </button>
              </div>
            ) : (
              <button
                className="btn-ghost text-xs"
                onClick={() => connectors[0] && connect({ connector: connectors[0] })}
              >
                连接钱包（操作时）
              </button>
            )}

            <div className="h-6 w-px bg-slate-700" />

            <div className="flex items-center gap-2">
              <span className="text-xs text-slate-400">{username ?? '管理员'}</span>
              <button className="btn-ghost text-xs" onClick={doLogout}>
                退出登录
              </button>
            </div>
          </div>
        </div>
      </header>
      <main className="mx-auto max-w-7xl px-4 py-6">{children}</main>
    </div>
  );
}