import clsx from 'clsx';
import { PoolState, POOL_STATE_LABEL } from '@shared/index';

/** 池状态标签：五态分开展示，不合并成 "Finished"（GAP-18 / PRD 11.2.2） */
export function StateBadge({ state, label }: { state: number; label?: string }) {
  const tone =
    state === PoolState.MATCH
      ? 'bg-accent-500/15 text-accent-500'
      : state === PoolState.EXECUTION
        ? 'bg-sky-500/15 text-sky-400'
        : state === PoolState.FINISH
          ? 'bg-positive/15 text-positive'
          : state === PoolState.LIQUIDATION
            ? 'bg-negative/15 text-negative'
            : 'bg-slate-500/15 text-slate-400';

  return <span className={clsx('tag', tone)}>{label ?? POOL_STATE_LABEL[state as PoolState] ?? '未知'}</span>;
}

/** 补充标记，不覆盖主状态（PRD 11.2.2） */
export function FlagBadges({ flags }: { flags: string[] }) {
  if (!flags.length) return null;
  return (
    <div className="flex flex-wrap gap-1.5">
      {flags.map((f) => (
        <span key={f} className="tag bg-warn/15 text-warn">
          {f}
        </span>
      ))}
    </div>
  );
}