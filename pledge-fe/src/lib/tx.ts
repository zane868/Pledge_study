import { create } from 'zustand';

export type TxStage =
  | 'idle'
  | 'awaiting-wallet'
  | 'rejected'
  | 'broadcast-unknown'
  | 'pending'
  | 'confirmed'
  | 'failed'
  | 'syncing';

export interface PendingTx {
  key: string;
  hash?: `0x${string}`;
  stage: TxStage;
  message?: string;
  updatedAt: number;
}

interface TxStore {
  txs: Record<string, PendingTx>;
  set: (key: string, tx: Partial<PendingTx> & { stage: TxStage }) => void;
  clear: (key: string) => void;
  get: (key: string) => PendingTx | undefined;
}

/**
 * 交易状态机（PRD 11.5.1）：
 * 等待钱包确认 / 用户拒绝 / 广播结果未知 / 已广播待确认 / 确认成功 / 确认失败 / 数据同步中。
 * 页面刷新后仍保留待确认交易，不因刷新丢失归属（UI-16）。
 */
export const useTxStore = create<TxStore>((set, get) => ({
  txs: {},
  set: (key, tx) =>
    set((s) => ({
      txs: { ...s.txs, [key]: { key, stage: tx.stage, hash: tx.hash, message: tx.message, updatedAt: Date.now() } },
    })),
  clear: (key) =>
    set((s) => {
      const { [key]: _, ...rest } = s.txs;
      return { txs: rest };
    }),
  get: (key) => get().txs[key],
}));

/** 从错误对象里区分「用户拒绝」与「未知广播结果」（UI-15） */
export function classifyTxError(err: unknown): { stage: TxStage; message: string } {
  const e = err as { code?: number; name?: string; shortMessage?: string; message?: string };
  const code = e?.code;
  const msg = e?.shortMessage ?? e?.message ?? String(err);
  // viem 的 UserRejectedRequestError
  if (code === 4001 || e?.name === 'UserRejectedRequestError' || /user rejected|rejected the request/i.test(msg)) {
    return { stage: 'rejected', message: '本次操作未提交（用户在钱包中拒绝）' };
  }
  if (/timeout|timed out/i.test(msg)) {
    return { stage: 'broadcast-unknown', message: '暂未取得明确提交结果，请核对钱包记录后再重试' };
  }
  return { stage: 'failed', message: msg };
}