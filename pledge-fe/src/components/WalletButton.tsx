import { useAccount, useConnect, useDisconnect, useSwitchChain } from 'wagmi';
import { shortAddress, DEPLOYMENT } from '@shared/config';
import { EXPECTED_CHAIN_ID } from '@/lib/wagmi';

export function WalletButton() {
  const { address, isConnected, chainId } = useAccount();
  const { connect, connectors, isPending } = useConnect();
  const { disconnect } = useDisconnect();
  const { switchChain, isPending: switching } = useSwitchChain();

  const wrongChain = isConnected && chainId !== EXPECTED_CHAIN_ID;

  if (!isConnected) {
    const injected = connectors[0];
    return (
      <button className="btn-primary" disabled={!injected || isPending} onClick={() => injected && connect({ connector: injected })}>
        {isPending ? '连接中…' : '连接钱包'}
      </button>
    );
  }

  return (
    <div className="flex items-center gap-2">
      {wrongChain && (
        <button className="btn bg-warn/20 text-warn hover:bg-warn/30" disabled={switching} onClick={() => switchChain({ chainId: EXPECTED_CHAIN_ID })}>
          {switching ? '切换中…' : `切换至 ${DEPLOYMENT.chainName}`}
        </button>
      )}
      <button className="btn-ghost tabular" onClick={() => disconnect()} title={address}>
        {shortAddress(address)}
      </button>
    </div>
  );
}