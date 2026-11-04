import { useMemo } from 'react';
import { useBlock, useReadContract, useReadContracts } from 'wagmi';
import { PledgePoolAbi, POOL_ADDRESS, STABLE_ADDRESS, SP_ADDRESS, JP_ADDRESS } from '@/lib/contracts';
import { DEPLOYMENT } from '@shared/config';

const ERC20_BALANCE_ABI = [
  { type: 'function', name: 'balanceOf', stateMutability: 'view', inputs: [{ name: 'a', type: 'address' }], outputs: [{ type: 'uint256' }] },
  { type: 'function', name: 'allowance', stateMutability: 'view', inputs: [{ name: 'o', type: 'address' }, { name: 's', type: 'address' }], outputs: [{ type: 'uint256' }] },
] as const;

import { derivePoolState, lendScale, PoolState, type DerivedState } from '@shared/index';

/**
 * 链上时间：以最新区块时间为准，不用浏览器本地时间推断可执行性（PRD 5.3.3 / 11.2.2）。
 * 每 15 秒刷新区块头。
 */
export function useChainNow(refetchMs = 15_000) {
  const { data } = useBlock({ query: { refetchInterval: refetchMs } });
  return data?.timestamp ?? 0n;
}

export interface OnChainPool {
  settleTime: bigint;
  endTime: bigint;
  interestRate: bigint;
  maxSupply: bigint;
  lendSupply: bigint;
  borrowSupply: bigint;
  martgageRate: bigint;
  lendToken: `0x${string}`;
  borrowToken: `0x${string}`;
  state: PoolState;
  spCoin: `0x${string}`;
  jpCoin: `0x${string}`;
  autoLiquidateThreshold: bigint;
}

/** 读取单个池的基础信息（链上为准，后端副本仅用于列表展示） */
export function usePoolBaseInfo(poolIndex: number) {
  return useReadContract({
    address: POOL_ADDRESS,
    abi: PledgePoolAbi,
    functionName: 'poolBaseInfo',
    args: [BigInt(poolIndex)],
    query: { refetchInterval: 15_000 },
  });
}

export function usePoolDataInfo(poolIndex: number) {
  return useReadContract({
    address: POOL_ADDRESS,
    abi: PledgePoolAbi,
    functionName: 'poolDataInfo',
    args: [BigInt(poolIndex)],
    query: { refetchInterval: 15_000 },
  });
}

export function usePoolLength() {
  return useReadContract({
    address: POOL_ADDRESS,
    abi: PledgePoolAbi,
    functionName: 'poolLength',
    query: { refetchInterval: 30_000 },
  });
}

export function usePoolGlobals() {
  const { data, ...rest } = useReadContracts({
    contracts: [
      { address: POOL_ADDRESS, abi: PledgePoolAbi, functionName: 'globalPaused' },
      { address: POOL_ADDRESS, abi: PledgePoolAbi, functionName: 'lendFee' },
      { address: POOL_ADDRESS, abi: PledgePoolAbi, functionName: 'borrowFee' },
      { address: POOL_ADDRESS, abi: PledgePoolAbi, functionName: 'minAmount' },
    ],
    query: { refetchInterval: 30_000 },
  });

  return {
    ...rest,
    paused: Boolean(data?.[0]?.result),
    lendFee: (data?.[1]?.result as bigint | undefined) ?? 0n,
    borrowFee: (data?.[2]?.result as bigint | undefined) ?? 0n,
    minAmount: (data?.[3]?.result as bigint | undefined) ?? 0n,
  };
}

/** 用户在该池的原始参与记录 */
export function useUserLendInfo(poolIndex: number, account?: `0x${string}`) {
  return useReadContract({
    address: POOL_ADDRESS,
    abi: PledgePoolAbi,
    functionName: 'userLendInfo',
    args: account ? [account, BigInt(poolIndex)] : undefined,
    query: { enabled: Boolean(account), refetchInterval: 15_000 },
  });
}

export function useUserBorrowInfo(poolIndex: number, account?: `0x${string}`) {
  return useReadContract({
    address: POOL_ADDRESS,
    abi: PledgePoolAbi,
    functionName: 'userBorrowInfo',
    args: account ? [account, BigInt(poolIndex)] : undefined,
    query: { enabled: Boolean(account), refetchInterval: 15_000 },
  });
}

/** SP/JP 余额：覆盖"纯凭证持有人"（只收到转账、没有原始入金）的情况（GAP-18 / UI-11） */
export function useVoucherBalances(account?: `0x${string}`, spAddress = SP_ADDRESS, jpAddress = JP_ADDRESS) {
  const { data, ...rest } = useReadContracts({
    contracts: [
      { address: spAddress, abi: ERC20_BALANCE_ABI, functionName: 'balanceOf', args: account ? [account] : undefined },
      { address: jpAddress, abi: ERC20_BALANCE_ABI, functionName: 'balanceOf', args: account ? [account] : undefined },
    ],
    query: { enabled: Boolean(account), refetchInterval: 15_000 },
  });
  return {
    ...rest,
    spBalance: (data?.[0]?.result as bigint | undefined) ?? 0n,
    jpBalance: (data?.[1]?.result as bigint | undefined) ?? 0n,
  };
}

/** 稳定币余额与对池合约的 allowance */
export function useStableBalance(account?: `0x${string}`, spender: `0x${string}` = POOL_ADDRESS) {
  const { data, ...rest } = useReadContracts({
    contracts: [
      { address: STABLE_ADDRESS, abi: ERC20_BALANCE_ABI, functionName: 'balanceOf', args: account ? [account] : undefined },
      { address: STABLE_ADDRESS, abi: ERC20_BALANCE_ABI, functionName: 'allowance', args: account ? [account, spender] : undefined },
    ],
    query: { enabled: Boolean(account), refetchInterval: 15_000 },
  });
  return {
    ...rest,
    balance: (data?.[0]?.result as bigint | undefined) ?? 0n,
    allowance: (data?.[1]?.result as bigint | undefined) ?? 0n,
  };
}

/** 组合：派生状态 */
export function useDerived(pool: OnChainPool | undefined, chainNow: bigint, paused: boolean): DerivedState | null {
  return useMemo(() => {
    if (!pool) return null;
    return derivePoolState({
      state: pool.state,
      settleTime: pool.settleTime,
      endTime: pool.endTime,
      chainNow,
      paused,
    });
  }, [pool, chainNow, paused]);
}

export const SCALE = lendScale(DEPLOYMENT.stableDecimals);