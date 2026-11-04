/**
 * 池状态与页面派生状态
 *
 * 依据 Pledge_PRD.md 5.3.1（状态定义）与 5.3.4（页面派生状态）。
 * 链上状态只有 5 种，不得合并成 "Finished" 之类的笼统文案。
 */

export enum PoolState {
  MATCH = 0,
  EXECUTION = 1,
  FINISH = 2,
  LIQUIDATION = 3,
  UNDONE = 4,
}

export const POOL_STATE_LABEL: Record<PoolState, string> = {
  [PoolState.MATCH]: '募集',
  [PoolState.EXECUTION]: '执行',
  [PoolState.FINISH]: '已结算',
  [PoolState.LIQUIDATION]: '已清算',
  [PoolState.UNDONE]: '募集未成交',
};

export interface DeriveInput {
  state: PoolState;
  settleTime: bigint;
  endTime: bigint;
  /** 链上区块时间（秒）。不要用本地 Date.now() 推断可执行性。 */
  chainNow: bigint;
  paused: boolean;
}

export interface DerivedState {
  /** 页面主状态文案 */
  label: string;
  /** 补充标记，不覆盖主状态 */
  flags: string[];
  /** 是否允许入金（仍需合约层校验） */
  canDeposit: boolean;
  /** 是否已过募集截止 */
  pastSettle: boolean;
  /** 是否已过到期 */
  pastEnd: boolean;
}

/**
 * 按 5.3.4 组合链上状态、链上时间、暂停开关生成展示文案。
 *
 * 时间边界严格性（5.3.3）：
 *   入金要求 t < settleTime（严格）
 *   匹配要求 t > settleTime（严格）
 *   到期结算要求 t > endTime（严格）
 * 因此 t === settleTime 时既不满足入金，也不满足匹配 —— 页面提示"等待可执行时间"。
 */
export function derivePoolState(input: DeriveInput): DerivedState {
  const { state, settleTime, endTime, chainNow, paused } = input;
  const pastSettle = chainNow >= settleTime;
  const pastEnd = chainNow >= endTime;
  const flags: string[] = [];
  if (paused) flags.push('平台暂停');

  let label: string;
  let canDeposit = false;

  switch (state) {
    case PoolState.MATCH:
      if (chainNow < settleTime) {
        label = '募集中';
        canDeposit = !paused;
      } else if (chainNow === settleTime) {
        // 严格边界：入金要求 t < settleTime，匹配要求 t > settleTime
        label = '募集已截止，待匹配';
        flags.push('等待可执行时间');
      } else {
        label = '募集已截止，待匹配';
      }
      break;

    case PoolState.EXECUTION:
      label = pastEnd ? '已到期，待结算' : '借款执行中';
      break;

    case PoolState.FINISH:
      label = '到期结算完成';
      break;

    case PoolState.LIQUIDATION:
      label = '已清算';
      break;

    case PoolState.UNDONE:
      label = '募集未成交';
      break;

    default:
      label = `未知状态(${state})`;
  }

  return { label, flags, canDeposit, pastSettle, pastEnd };
}