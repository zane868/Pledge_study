import { useMemo } from 'react';
import { useQuery } from '@tanstack/react-query';
import { api } from '@shared/index';
import { DEPLOYMENT } from '@shared/config';

export interface MarketRow {
  poolIndex: number;
  displayPoolId: number;
  poolContract: string;
  lendSymbol: string;
  borrowSymbol: string;
  lendSupply: bigint;
  borrowSupply: bigint;
  maxSupply: bigint;
  interestRate: bigint;
  martgageRate: bigint;
  liquidateThreshold: bigint;
  settleTime: bigint;
  endTime: bigint;
  state: number;
  spCoin: string;
  jpCoin: string;
  syncBlock: string;
}

function toMarketRow(item: api.PoolBaseInfoRes): MarketRow {
  const p = item.pool_data;
  return {
    poolIndex: item.poolIndex,
    displayPoolId: item.displayPoolId,
    poolContract: p.poolContract,
    lendSymbol: p.lendTokenInfo?.tokenName ?? 'USDT',
    borrowSymbol: p.borrowTokenInfo?.tokenName ?? 'ETH',
    lendSupply: BigInt(p.lendSupply || '0'),
    borrowSupply: BigInt(p.borrowSupply || '0'),
    maxSupply: BigInt(p.maxSupply || '0'),
    interestRate: BigInt(p.interestRate || '0'),
    martgageRate: BigInt(p.martgageRate || '0'),
    liquidateThreshold: BigInt(p.autoLiquidateThreshold || '0'),
    settleTime: BigInt(p.settleTime || '0'),
    endTime: BigInt(p.endTime || '0'),
    state: Number(p.state ?? 0),
    spCoin: p.spCoin,
    jpCoin: p.jpCoin,
    syncBlock: p.syncBlockNumber,
  };
}

/**
 * 市场列表数据。后端在无成功快照 / 同步失败 / 快照超 5 分钟时返回 503，
 * 这里会抛 ApiError，页面必须区分「加载失败」与「确实无池」（PRD 11.2.1 / 13.4）。
 */
export function useMarketPools() {
  const query = useQuery({
    queryKey: ['market', DEPLOYMENT.chainId],
    queryFn: () => api.fetchPoolBaseInfo(DEPLOYMENT.chainId),
    refetchInterval: 30_000,
    retry: 1,
  });

  const rows = useMemo<MarketRow[]>(() => {
    const list = query.data?.data ?? [];
    return list
      .map(toMarketRow)
      .sort((a, b) => {
        // 默认优先展示仍可参与的池，同组按募集截止由近到远，编号补充稳定次序（PRD 11.2.1）
        const aOpen = a.state === 0 ? 0 : 1;
        const bOpen = b.state === 0 ? 0 : 1;
        if (aOpen !== bOpen) return aOpen - bOpen;
        if (a.settleTime !== b.settleTime) return a.settleTime < b.settleTime ? -1 : 1;
        return a.poolIndex - b.poolIndex;
      });
  }, [query.data]);

  return { ...query, rows };
}

export function useSyncStatus() {
  return useQuery({
    queryKey: ['sync-status'],
    queryFn: () => api.fetchSyncStatus(),
    refetchInterval: 30_000,
    retry: 1,
  });
}