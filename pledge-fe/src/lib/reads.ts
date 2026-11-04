import { readContract } from 'wagmi/actions';
import { config } from '@/lib/wagmi';
import { PledgePoolAbi, POOL_ADDRESS, SP_ADDRESS, JP_ADDRESS, ERC20Abi } from '@/lib/contracts';
import { PoolState } from '@shared/index';

/**
 * 链上读取层
 *
 * poolBaseInfo / poolDataInfo 是 public 数组的 getter，返回的是扁平化的值元组，
 * 不是对象。viem 的类型推导会把它们变成 readonly 元组，字段顺序必须与合约结构体一致。
 * 这里统一封装成命名对象，避免页面里到处写 base[0]、base[3]。
 */

export interface PoolBase {
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

export interface PoolData {
  settleAmountLend: bigint;
  settleAmountBorrow: bigint;
  finishAmountLend: bigint;
  finishAmountBorrow: bigint;
  liquidationAmounLend: bigint;
  liquidationAmounBorrow: bigint;
}

export interface UserRecord {
  stakeAmount: bigint;
  refundAmount: bigint;
  hasNoRefund: boolean;
  hasNoClaim: boolean;
}

export class ReadError extends Error {
  constructor(
    message: string,
    readonly cause?: unknown,
  ) {
    super(message);
    this.name = 'ReadError';
  }
}

export async function readPoolBase(poolIndex: number): Promise<PoolBase | null> {
  try {
    const r = (await readContract(config, {
      address: POOL_ADDRESS,
      abi: PledgePoolAbi,
      functionName: 'poolBaseInfo',
      args: [BigInt(poolIndex)],
    })) as readonly [
      bigint,
      bigint,
      bigint,
      bigint,
      bigint,
      bigint,
      bigint,
      `0x${string}`,
      `0x${string}`,
      number,
      `0x${string}`,
      `0x${string}`,
      bigint,
    ];
    return {
      settleTime: r[0],
      endTime: r[1],
      interestRate: r[2],
      maxSupply: r[3],
      lendSupply: r[4],
      borrowSupply: r[5],
      martgageRate: r[6],
      lendToken: r[7],
      borrowToken: r[8],
      state: Number(r[9]) as PoolState,
      spCoin: r[10],
      jpCoin: r[11],
      autoLiquidateThreshold: r[12],
    };
  } catch (err) {
    throw new ReadError('读取池基础信息失败', err);
  }
}

export async function readPoolData(poolIndex: number): Promise<PoolData | null> {
  try {
    const r = (await readContract(config, {
      address: POOL_ADDRESS,
      abi: PledgePoolAbi,
      functionName: 'poolDataInfo',
      args: [BigInt(poolIndex)],
    })) as readonly [bigint, bigint, bigint, bigint, bigint, bigint];
    return {
      settleAmountLend: r[0],
      settleAmountBorrow: r[1],
      finishAmountLend: r[2],
      finishAmountBorrow: r[3],
      liquidationAmounLend: r[4],
      liquidationAmounBorrow: r[5],
    };
  } catch {
    return null;
  }
}

export async function readUserLend(account: `0x${string}`, poolIndex: number): Promise<UserRecord | null> {
  try {
    const r = (await readContract(config, {
      address: POOL_ADDRESS,
      abi: PledgePoolAbi,
      functionName: 'userLendInfo',
      args: [account, BigInt(poolIndex)],
    })) as readonly [bigint, bigint, boolean, boolean];
    return { stakeAmount: r[0], refundAmount: r[1], hasNoRefund: r[2], hasNoClaim: r[3] };
  } catch {
    return null;
  }
}

export async function readUserBorrow(account: `0x${string}`, poolIndex: number): Promise<UserRecord | null> {
  try {
    const r = (await readContract(config, {
      address: POOL_ADDRESS,
      abi: PledgePoolAbi,
      functionName: 'userBorrowInfo',
      args: [account, BigInt(poolIndex)],
    })) as readonly [bigint, bigint, boolean, boolean];
    return { stakeAmount: r[0], refundAmount: r[1], hasNoRefund: r[2], hasNoClaim: r[3] };
  } catch {
    return null;
  }
}

export async function readBalance(token: `0x${string}`, account: `0x${string}`): Promise<bigint> {
  try {
    return (await readContract(config, {
      address: token,
      abi: ERC20Abi,
      functionName: 'balanceOf',
      args: [account],
    })) as bigint;
  } catch {
    return 0n;
  }
}

export async function readAllowance(token: `0x${string}`, owner: `0x${string}`, spender: `0x${string}`): Promise<bigint> {
  try {
    return (await readContract(config, {
      address: token,
      abi: ERC20Abi,
      functionName: 'allowance',
      args: [owner, spender],
    })) as bigint;
  } catch {
    return 0n;
  }
}

export async function readPoolLength(): Promise<number> {
  try {
    const n = (await readContract(config, {
      address: POOL_ADDRESS,
      abi: PledgePoolAbi,
      functionName: 'poolLength',
    })) as bigint;
    return Number(n);
  } catch {
    return 0;
  }
}

export async function readGlobalPaused(): Promise<boolean> {
  try {
    return (await readContract(config, {
      address: POOL_ADDRESS,
      abi: PledgePoolAbi,
      functionName: 'globalPaused',
    })) as boolean;
  } catch {
    return false;
  }
}

export async function readMinAmount(): Promise<bigint> {
  try {
    return (await readContract(config, {
      address: POOL_ADDRESS,
      abi: PledgePoolAbi,
      functionName: 'minAmount',
    })) as bigint;
  } catch {
    return 0n;
  }
}

/** 一次性读取用户在某池的全部权益数据 */
export interface PositionData {
  pool: PoolBase;
  data: PoolData | null;
  lend: UserRecord | null;
  borrow: UserRecord | null;
  spBalance: bigint;
  jpBalance: bigint;
}

export async function readPosition(account: `0x${string}`, poolIndex: number): Promise<PositionData | null> {
  const pool = await readPoolBase(poolIndex);
  if (!pool) return null;
  const [data, lend, borrow, spBalance, jpBalance] = await Promise.all([
    readPoolData(poolIndex),
    readUserLend(account, poolIndex),
    readUserBorrow(account, poolIndex),
    readBalance(pool.spCoin || SP_ADDRESS, account),
    readBalance(pool.jpCoin || JP_ADDRESS, account),
  ]);
  return { pool, data, lend, borrow, spBalance, jpBalance };
}
