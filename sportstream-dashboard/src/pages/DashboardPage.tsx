import { TrendingUp, TrendingDown } from 'lucide-react';
import { Header } from '../components/layout/Header';
import { DashboardSummaryCards } from '../components/features/dashboard/DashboardSummary';
import { StreamsByTypeChart } from '../components/features/dashboard/StreamsByTypeChart';
import { LiveNowPanel } from '../components/features/dashboard/LiveNowPanel';
import { UpcomingEventsPanel } from '../components/features/dashboard/UpcomingEvents';
import { useDashboard } from '../hooks/useDashboard';

function formatViews(n: number): string {
  if (n >= 1_000_000) return `${(n / 1_000_000).toFixed(1)}M`;
  if (n >= 1_000) return `${(n / 1_000).toFixed(1)}K`;
  return n.toLocaleString();
}

const RANK_COLORS: Record<number, { text: string; bg: string }> = {
  1: { text: 'text-amber-500', bg: 'bg-amber-500/10' },
  2: { text: 'text-slate-400', bg: 'bg-slate-400/10' },
  3: { text: 'text-amber-700', bg: 'bg-amber-700/10' },
};

function TopClubsByViews() {
  const { data, isLoading } = useDashboard();

  return (
    <div className="bg-surface-container-lowest rounded-2xl p-4 md:p-8 shadow-sm">
      <h2 className="text-lg md:text-xl font-bold tracking-tight mb-4 md:mb-8 text-on-surface">Top Clubs by Views</h2>

      {isLoading || !data ? (
        <div className="space-y-3">
          {Array.from({ length: 5 }).map((_, i) => (
            <div key={i} className="h-16 animate-pulse rounded-xl bg-surface-container-low" />
          ))}
        </div>
      ) : !data.topClubsByViews || data.topClubsByViews.length === 0 ? (
        <p className="py-8 text-center text-sm text-on-surface-variant">No data available</p>
      ) : (
        <div className="space-y-2">
          {data.topClubsByViews.map((club, index) => {
            const rank = index + 1;
            const color = RANK_COLORS[rank] ?? {
              text: 'text-on-surface-variant',
              bg: 'bg-surface-container',
            };
            // Simulate trend data based on rank
            const isPositive = rank !== 3;
            const trendValue = [4.2, 1.8, 0.5, 8.4, 2.1][index] ?? 1.0;

            return (
              <div
                key={club.clubId}
                className="flex items-center gap-3 md:gap-4 p-3 md:p-4 hover:bg-surface-container-low rounded-xl transition-colors group"
              >
                <div className={`w-8 h-8 md:w-10 md:h-10 flex items-center justify-center font-black ${rank <= 3 ? 'text-lg md:text-xl' : 'text-base md:text-lg'} ${color.text} ${color.bg} rounded-full`}>
                  {rank}
                </div>
                <div className="flex-1 min-w-0">
                  <h4 className="font-bold text-on-surface text-sm md:text-base">{club.clubName}</h4>
                  <p className="text-[10px] text-on-surface-variant font-bold uppercase tracking-widest hidden sm:block">
                    {rank === 1 ? 'Global Audience' : rank === 2 ? 'National Fanbase' : rank === 3 ? 'Local Heroes' : rank === 4 ? 'Elite Circle' : 'Coastal Sports'}
                  </p>
                </div>
                <div className="text-right">
                  <div className="text-xs md:text-sm font-black text-on-surface">
                    {formatViews(club.totalViews)}
                  </div>
                  <div className={`text-[10px] font-bold flex items-center justify-end gap-0.5 ${isPositive ? 'text-primary' : 'text-error'}`}>
                    {isPositive ? (
                      <><TrendingUp size={10} /> {trendValue}%</>
                    ) : (
                      <><TrendingDown size={10} /> {trendValue}%</>
                    )}
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      )}
    </div>
  );
}

export function DashboardPage() {
  return (
    <div>
      <Header title="Dashboard" />
      <div className="p-4 md:p-8 space-y-4 md:space-y-8 bg-background min-h-screen">
        {/* Bento Grid: Top Stats */}
        <DashboardSummaryCards />

        {/* Middle Section: Asymmetric Layout (2+1) */}
        <section className="grid grid-cols-1 lg:grid-cols-3 gap-4 md:gap-8">
          <StreamsByTypeChart />
          <LiveNowPanel />
        </section>

        {/* Bottom Section: 2 columns */}
        <section className="grid grid-cols-1 lg:grid-cols-2 gap-4 md:gap-8">
          <UpcomingEventsPanel />
          <TopClubsByViews />
        </section>
      </div>
    </div>
  );
}
