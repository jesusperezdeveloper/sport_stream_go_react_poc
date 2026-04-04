import { Building2, Radio, Calendar, Eye } from 'lucide-react';
import { StatsCard } from '../../ui/StatsCard';
import { useDashboard } from '../../../hooks/useDashboard';

function formatNumber(n: number): string {
  if (n >= 1_000_000) return `${(n / 1_000_000).toFixed(1)}M`;
  if (n >= 1_000) return `${(n / 1_000).toFixed(1)}K`;
  return n.toLocaleString();
}

export function DashboardSummaryCards() {
  const { data, isLoading } = useDashboard();

  if (isLoading || !data) {
    return (
      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
        {Array.from({ length: 4 }).map((_, i) => (
          <div
            key={i}
            className="h-32 animate-pulse rounded-2xl bg-surface-container-lowest shadow-sm"
          />
        ))}
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
      <StatsCard
        icon={Building2}
        iconColor="#006d35"
        borderColor="#006d35"
        label="Total Clubs"
        value={formatNumber(data.totalClubs)}
        extraLabel="+12% WoW"
      />
      <StatsCard
        icon={Radio}
        iconColor="#ba1a1a"
        borderColor="#ba1a1a"
        label="Active Streams"
        value={formatNumber(data.liveStreams)}
        extraLabel="Concurrent"
        isLive
      />
      <StatsCard
        icon={Calendar}
        iconColor="#c77c00"
        borderColor="#f59e0b"
        label="Upcoming Events"
        value={formatNumber(data.upcomingEvents)}
      />
      <StatsCard
        icon={Eye}
        iconColor="#00618f"
        borderColor="#00618f"
        label="Total Network Views"
        value={formatNumber(data.totalViews)}
      />
    </div>
  );
}
