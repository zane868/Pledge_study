import { useAccount, useConnect } from 'wagmi';
import { DEPLOYMENT, shortAddress } from '@shared/index';

const ADMINS = DEPLOYMENT.admins.map((a) => a.toLowerCase());

/**
 * 链上操作守卫：在具体操作前检测钱包连接与管理员身份。
 * 页面浏览不受影响，只有发交易时才需要。
 *
 * 用法：把需要钱包的按钮包在 <WalletGate> 内，或调用 useWalletGuard()。
 */
export function WalletGate({
  children,
  action,
}: {
  children: React.ReactNode;
  /** 操作名称，用于提示文案，如「创建池」「结算」 */
  action: string;
}) {
  const { address, isConnected } = useAccount();
  const { connect, connectors } = useConnect();
  const isAdmin = Boolean(address && ADMINS.includes(address.toLowerCase()));

  if (isConnected && isAdmin) return <>{children}</>;

  return (
    <div className="rounded border border-amber-500/30 bg-amber-500/5 p-3 text-xs">
      {!isConnected ? (
        <>
          <p className="text-amber-400">{action}需要连接管理员钱包</p>
          <button
            className="btn-ghost mt-2 text-xs"
            onClick={() => connectors[0] && connect({ connector: connectors[0] })}
          >
            连接钱包
          </button>
        </>
      ) : (
        <>
          <p className="text-amber-400">
            当前钱包 <span className="font-mono">{shortAddress(address)}</span> 不是管理员，无法{action}
          </p>
          <p className="mt-1 text-slate-400">请切换至管理员账户（{DEPLOYMENT.admins.map((a) => shortAddress(a)).join(' / ')}）</p>
        </>
      )}
    </div>
  );
}

/** 返回当前钱包是否可用于链上操作 */
export function useWalletReady() {
  const { address, isConnected } = useAccount();
  return {
    ready: Boolean(isConnected && address && ADMINS.includes(address.toLowerCase())),
    isConnected,
    address,
  };
}