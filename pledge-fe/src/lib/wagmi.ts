import { http, createConfig } from 'wagmi';
import { sepolia } from 'wagmi/chains';
import { injected } from 'wagmi/connectors';
import { DEPLOYMENT } from '@shared/config';

/**
 * 只支持部署清单里声明的网络（DEC-01 / PRD 11.1.2）：
 * 钱包网络不匹配时禁止提交，不回退到其他网络的合约地址。
 */
export const config = createConfig({
  chains: [sepolia],
  connectors: [injected()],
  transports: {
    [sepolia.id]: http(DEPLOYMENT.rpcUrl),
  },
});

export const EXPECTED_CHAIN_ID = DEPLOYMENT.chainId as (typeof sepolia)['id'];

declare module 'wagmi' {
  interface Register {
    config: typeof config;
  }
}