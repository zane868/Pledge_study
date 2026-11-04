import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAccount, useConnect, useDisconnect } from 'wagmi';
import { DEPLOYMENT } from '@shared/index';

const ADMINS = DEPLOYMENT.admins.map((a) => a.toLowerCase());

export function LoginPage() {
  const navigate = useNavigate();
  const { address, isConnected } = useAccount();
  const { connect, connectors, error, isPending } = useConnect();
  const { disconnect } = useDisconnect();

  useEffect(() => {
    if (isConnected && address && ADMINS.includes(address.toLowerCase())) {
      navigate('/', { replace: true });
    }
  }, [isConnected, address, navigate]);

  return (
    <div className="flex min-h-screen items-center justify-center bg-ink-900 p-4">
      <div className="card max-w-md">
        <h1 className="text-2xl font-semibold text-white">Pledge 运营后台</h1>
        <p className="mt-2 text-sm text-slate-400">连接管理员钱包以继续</p>

        <div className="mt-6 space-y-3">
          {connectors.map((connector) => (
            <button
              key={connector.uid}
              className="btn-primary w-full"
              disabled={isPending}
              onClick={() => connect({ connector })}
            >
              {isPending ? '连接中...' : `连接 ${connector.name}`}
            </button>
          ))}
        </div>

        {error && (
          <p className="mt-4 text-sm text-negative">{error.message}</p>
        )}

        {isConnected && address && !ADMINS.includes(address.toLowerCase()) && (
          <div className="mt-4 rounded border border-negative/20 bg-negative/5 p-3">
            <p className="text-sm text-negative">当前地址无权限</p>
            <p className="mt-1 break-all font-mono text-xs text-slate-400">{address}</p>
            <button className="btn-ghost mt-3" onClick={() => disconnect()}>
              断开连接
            </button>
          </div>
        )}

        <div className="mt-8 border-t border-slate-700 pt-4">
          <p className="text-xs text-slate-500">管理员地址：</p>
          <div className="mt-2 space-y-1">
            {DEPLOYMENT.admins.map((admin, idx) => (
              <p key={admin} className="font-mono text-xs text-slate-400">
                {idx === 0 ? 'Owner A' : idx === 1 ? 'Owner B' : 'Owner C'}: {admin.slice(0, 10)}…{admin.slice(-8)}
              </p>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
