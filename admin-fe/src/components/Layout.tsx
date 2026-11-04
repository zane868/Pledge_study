import { Link, useLocation } from 'react-router-dom';
import { useAccount, useDisconnect } from 'wagmi';
import { DEPLOYMENT } from '@shared/index';

export function Layout({ children }: { children: React.ReactNode }) {
  const location = useLocation();
  const { address } = useAccount();
  const { disconnect } = useDisconnect();

  const nav = [
    { path: '/', label: '总览' },
    { path: '/pools', label: '池管理' },
    { path: '/proposals', label: '多签提案' },
  ];

  const adminLabel =
    address?.toLowerCase() === DEPLOYMENT.admins[0].toLowerCase()
      ? 'Owner A'
      : address?.toLowerCase() === DEPLOYMENT.admins[1].toLowerCase()
        ? 'Owner B'
        : 'Owner C';

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
            <div className="text-right">
              <div className="text-xs text-slate-500">{adminLabel}</div>
              <div className="font-mono text-xs text-slate-300">
                {address?.slice(0, 6)}…{address?.slice(-4)}
              </div>
            </div>
            <button className="btn-ghost text-sm" onClick={() => disconnect()}>
              断开
            </button>
          </div>
        </div>
      </header>
      <main className="mx-auto max-w-7xl px-4 py-6">{children}</main>
    </div>
  );
}
