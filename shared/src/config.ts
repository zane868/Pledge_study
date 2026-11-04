import deploymentJson from '../deployment.json';

export interface Deployment {
  chainId: number;
  chainName: string;
  rpcUrl: string;
  explorer: string;
  apiBase: string;
  pool: `0x${string}`;
  oracle: `0x${string}`;
  multisig: `0x${string}`;
  stable: `0x${string}`;
  weth: `0x${string}`;
  router: `0x${string}`;
  factory: `0x${string}`;
  sp: `0x${string}`;
  jp: `0x${string}`;
  operator: `0x${string}`;
  admins: `0x${string}`[];
  threshold: number;
  deploymentStartBlock: number;
  stableDecimals: number;
  voucherDecimals: number;
}

export const DEPLOYMENT = deploymentJson as Deployment;

/** 原生 ETH 在池中以零地址标识 */
export const NATIVE_ETH = '0x0000000000000000000000000000000000000000' as const;

export function isNative(token: string): boolean {
  return token.toLowerCase() === NATIVE_ETH;
}

export function explorerTx(hash: string): string {
  return `${DEPLOYMENT.explorer}/tx/${hash}`;
}

export function explorerAddress(address: string): string {
  return `${DEPLOYMENT.explorer}/address/${address}`;
}

export function shortAddress(address?: string, size = 4): string {
  if (!address) return '';
  return `${address.slice(0, 2 + size)}…${address.slice(-size)}`;
}