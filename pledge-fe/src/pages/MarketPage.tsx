import { useMarketPools, useSyncStatus } from '@/hooks/useMarket';
import { useChainNow } from '@/hooks/usePools';
import { usePoolGlobals } from '@/hooks/usePools';
import { StateBadge, FlagBadges } from '@/components/StateBadge';
import { Link } from 'react-router-dom';
import { ApiError } from '@shared/api';
import {
  derivePoolState,
  displayUnits,
  formatRate,
  formatTimestampWithTz,
  remainingLabel,
  termDays,
  DEPLOYMENT,
} from '@shared/index';

export function MarketPage() {
  const { rows, isLoading, error } = useMarketPools();
  const { data: sync } = useSyncStatus();
  const chainNow = useChainNow();
  const { paused } = usePoolGlobals();

  // 同步时间：lastSuccess 是后端最后成功同步时刻（RFC3339 UTC）。
  // 同步周期固定 120s（见 pledge-backend/schedule/pledge_task.go），
  // 下次计划同步按 lastSuccess + 120s 估算，误差为单轮同步耗时（约几秒）。
  const SYNC_INTERVAL_S = 120;
  const syncInfo = (() => {
    const last = sync?.data?.lastSuccess;
    if (!last) return null;
    const lastUnix = Math.floor(new Date(last).getTime() / 1000);
    if (!Number.isFinite(lastUnix)) return null;
    return { lastUnix, nextUnix: lastUnix + SYNC_INTERVAL_S };
  })();

  return (
    <div className="space-y-5">
      <div className="flex flex-wrap items-end justify-between gap-3">
        <div>
          <h1 className="text-xl font-semibold text-white">借贷市场</h1>
          <p className="mt-1 text-sm text-slate-400">
            出借人提供 USDT，借款人抵押原生 ETH，募集截止后统一匹配，未匹配部分可领取退款。
          </p>
        </div>
        {sync?.data && (
          <div className="text-right text-xs leading-relaxed text-slate-500">
            <p>
              最后同步 {syncInfo ? formatTimestampWithTz(syncInfo.lastUnix) : '读取中'}
              {sync.stale ? ' · 快照可能已过期' : ''}
            </p>
            <p>
              下次同步 {syncInfo ? formatTimestampWithTz(syncInfo.nextUnix) : '—'} · 区块{' '}
              {sync.data.blockNumber}
            </p>
          </div>
        )}
      </div>

      {error && <SyncUnavailable error={error} />}

      {!error && isLoading && (
        <div className="card text-sm text-slate-400">正在读取池列表…</div>
      )}

      {!error && !isLoading && rows.length === 0 && (
        <div className="card text-sm text-slate-400">当前网络暂无借贷池。</div>
      )}

      <div className="space-y-3">
        {rows.map((row) => {
          const derived = derivePoolState({
            state: row.state as 0 | 1 | 2 | 3 | 4,
            settleTime: row.settleTime,
            endTime: row.endTime,
            chainNow,
            paused,
          });
          const progress =
            row.maxSupply > 0n ? Number((row.lendSupply * 10_000n) / row.maxSupply) / 100 : 0;

          return (
            <Link
              key={`${row.poolContract}-${row.poolIndex}`}
              to={`/pool/${DEPLOYMENT.chainId}/${row.poolIndex}`}
              className="card block transition-colors hover:border-white/10"
            >
              <div className="flex flex-wrap items-start justify-between gap-3">
                <div className="flex items-center gap-3">
                  <span className="tabular text-lg font-semibold text-white">#{row.displayPoolId}</span>
                  <StateBadge state={row.state} label={derived?.label} />
                  <FlagBadges flags={derived?.flags ?? []} />
                </div>
                <span className="text-sm text-slate-400">
                  募集截止 {formatTimestampWithTz(row.settleTime)}
                </span>
              </div>

              <div className="mt-4 grid grid-cols-2 gap-x-6 gap-y-3 md:grid-cols-5">
                <Metric label="年利率" value={formatRate(row.interestRate)} accent />
                <Metric label="约定期限" value={`${termDays(row.settleTime, row.endTime)} 天`} />
                <Metric label="募集抵押率" value={formatRate(row.martgageRate)} />
                <Metric
                  label="出借募集进度"
                  value={`${displayUnits(row.lendSupply, DEPLOYMENT.stableDecimals, 2)} / ${displayUnits(
                    row.maxSupply,
                    DEPLOYMENT.stableDecimals,
                    0,
                  )} ${row.lendSymbol}`}
                />
                <Metric
                  label="抵押募集量"
                  value={`${displayUnits(row.borrowSupply, 18, 4)} ${row.borrowSymbol}`}
                />
              </div>

              {row.state === 0 && (
                <div className="mt-4">
                  <div className="h-1.5 overflow-hidden rounded-full bg-white/5">
                    <div className="h-full bg-accent-500" style={{ width: `${Math.min(progress, 100)}%` }} />
                  </div>
                  <p className="mt-1.5 text-xs text-slate-500">
                    {progress.toFixed(1)}% · 募集进度以累计出借量 ÷ 出借上限计算，不代表已成交比例
                    {' · '}
                    {remainingLabel(row.settleTime, chainNow)} 后截止
                  </p>
                </div>
              )}
            </Link>
          );
        })}
      </div>
    </div>
  );
}

function Metric({ label, value, accent }: { label: string; value: string; accent?: boolean }) {
  return (
    <div>
      <div className="text-xs text-slate-500">{label}</div>
      <div className={`tabular mt-0.5 text-sm ${accent ? 'font-semibold text-accent-500' : 'text-slate-200'}`}>
        {value}
      </div>
    </div>
  );
}

/** 区分「同步不可用」与「确实没有池」（PRD 11.2.1 / 13.4） */
function SyncUnavailable({ error }: { error: unknown }) {
  const is503 = error instanceof ApiError && error.status === 503;
  return (
    <div className="card border-warn/20 bg-warn/5">
      <p className="text-sm text-warn">
        {is503 ? '数据同步中或快照已过期，暂无法展示市场列表。' : '市场数据加载失败。'}
      </p>
      <p className="mt-1 text-xs text-slate-400">
        {error instanceof Error ? error.message : String(error)} · 这不代表没有正在募集的池，请稍后刷新。
      </p>
    </div>
  );
}