import { MoreHorizontal } from 'lucide-react';
import { useDashboard } from '../../../hooks/useDashboard';

const TYPE_CONFIG: Record<string, { color: string; label: string }> = {
  live: { color: '#3fff8b', label: 'Live Match' },
  vod: { color: '#a7d7ff', label: 'VOD Replays' },
  highlight: { color: '#fbbf24', label: 'Highlights' },
  behind_the_scenes: { color: '#c084fc', label: 'Behind the Scenes' },
};

export function StreamsByTypeChart() {
  const { data, isLoading } = useDashboard();

  const chartData = data
    ? Object.entries(data.streamsByType).map(([name, value]) => ({
        name,
        value,
        config: TYPE_CONFIG[name] ?? { color: '#9e9e9e', label: name },
      }))
    : [];

  const maxValue = chartData.length > 0 ? Math.max(...chartData.map((d) => d.value)) : 1;

  return (
    <div className="col-span-1 lg:col-span-2 bg-inverse-surface rounded-2xl p-4 md:p-8">
      <div className="flex justify-between items-center mb-4 md:mb-8">
        <h2 className="text-lg md:text-xl font-bold tracking-tight text-white">Streams by Type</h2>
        <MoreHorizontal size={20} className="text-white/40" />
      </div>

      {isLoading ? (
        <div className="space-y-6">
          {Array.from({ length: 4 }).map((_, i) => (
            <div key={i} className="space-y-2">
              <div className="h-3 w-24 animate-pulse rounded bg-white/10" />
              <div className="h-3 w-full animate-pulse rounded-full bg-white/10" />
            </div>
          ))}
        </div>
      ) : chartData.length === 0 ? (
        <p className="flex h-48 items-center justify-center text-sm text-white/40">
          No stream data available
        </p>
      ) : (
        <div className="space-y-4 md:space-y-6">
          {chartData.map(({ name, value, config }) => {
            const percentage = maxValue > 0 ? (value / maxValue) * 100 : 0;
            return (
              <div key={name} className="space-y-2">
                <div className="flex justify-between text-xs font-bold uppercase tracking-widest mb-1">
                  <span className="text-white/80">{config.label}</span>
                  <span style={{ color: config.color }}>
                    {value} {value === 1 ? 'Stream' : 'Streams'}
                  </span>
                </div>
                <div className="h-3 w-full bg-white/10 rounded-full overflow-hidden">
                  <div
                    className="h-full rounded-full transition-all duration-500"
                    style={{ width: `${Math.max(percentage, 5)}%`, backgroundColor: config.color }}
                  />
                </div>
              </div>
            );
          })}
        </div>
      )}
    </div>
  );
}
