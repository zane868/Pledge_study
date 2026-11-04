import type { ReactNode } from 'react';
import { useAccount } from 'wagmi';
import { Link, useLocation } from 'react-router-dom';
import { WalletButton } from './WalletButton';
import { DEPLOYMENT } from '@shared/config';
import { EXPECTED_CHAIN_ID } from '@/lib/wagmi';

export function Layout({ children }: { children: ReactNode }) {
  const { isConnected, chainId } = useAccount();
  const { pathname } = useLocation();
  const wrongChain = isConnected && chainId !== EXPECTED_CHAIN_ID;

  const nav = [
    { to: '/', label: '市场' },
    { to: '/portfolio', label: '我的仓位' },
  ];

  return (
    <div className="min-h-screen">
      <header className="border-b border-white/5 bg-ink-800/60 backdrop-blur">
        <div className="mx-auto flex max-w-6xl items-center justify-between gap-6 px-5 py-3.5">
          <div className="flex items-center gap-7">
            <Link to="/" className="text-base font-semibold tracking-tight text-white">
              Pledge
            </Link>
            <nav className="flex items-center gap-1">
              {nav.map((item) => (
                <Link
                  key={item.to}
                  to={item.to}
                  className={`rounded-md px-3 py-1.5 text-sm transition-colors ${
                    pathname === item.to ? 'bg-white/5 text-white' : 'text-slate-400 hover:text-slate-200'
                  }`}
                >
                  {item.label}
                </Link>
              ))}
            </nav>
          </div>
          <div className="flex items-center gap-3">
            <span className="tag bg-white/5 text-slate-400">{DEPLOYMENT.chainName}</span>
            <WalletButton />
          </div>
        </div>
      </header>

      {wrongChain && (
        <div className="border-b border-warn/20 bg-warn/10">
          <div className="mx-auto max-w-6xl px-5 py-2 text-sm text-warn">
            钱包网络与池所属网络不符，请先切换至 {DEPLOYMENT.chainName}。切换前不会提交任何池交易。
          </div>
        </div>
      )}

      <main className="mx-auto max-w-6xl px-5 py-7">{children}</main>

      <footer className="mx-auto max-w-6xl px-5 pb-10 pt-4 text-xs text-slate-500">
        <div className="flex flex-wrap gap-x-4 gap-y-1">
          <span>池合约 {DEPLOYMENT.pool}</span>
          <span>·</span>
          <span>测试环境，使用自建 6 位测试 USDT，非发行方官方资产</span>
        </div>
      </footer>
    </div>
  );
}