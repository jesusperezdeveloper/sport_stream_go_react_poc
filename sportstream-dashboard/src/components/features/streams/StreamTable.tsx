import { useState, useMemo } from 'react';
import { StopCircle, MoreVertical, TrendingUp, Zap, Play } from 'lucide-react';
import { StatusBadge } from '../../ui/StatusBadge';
import { StreamFilters } from './StreamFilters';
import { StreamModal } from '../../ui/StreamModal';
import { useStreams } from '../../../hooks/useStreams';
import { useClubs } from '../../../hooks/useClubs';
import type { Stream, StreamStatus, StreamType } from '../../../types/stream';

function formatViews(n: number): string {
  if (n >= 1_000_000) return `${(n / 1_000_000).toFixed(1)}M`;
  if (n >= 1_000) return `${Math.round(n / 1_000)}K`;
  return n.toLocaleString();
}

function formatDuration(seconds: number): string {
  if (seconds === 0) return '--';
  const h = Math.floor(seconds / 3600);
  const m = Math.floor((seconds % 3600) / 60);
  const s = seconds % 60;
  if (h > 0) return `${h}:${m.toString().padStart(2, '0')}:${s.toString().padStart(2, '0')}`;
  return `${m}:${s.toString().padStart(2, '0')}`;
}

const TYPE_BADGE_STYLES: Record<string, string> = {
  live: 'bg-primary-container/20 text-on-primary-container',
  vod: 'bg-blue-100 text-blue-800',
  highlight: 'bg-secondary-container/30 text-on-secondary-container',
  behind_the_scenes: 'bg-tertiary-container/30 text-on-tertiary-container',
};

function StreamMobileCard({ stream, clubName, onWatch }: { stream: Stream; clubName: string; onWatch: () => void }) {
  return (
    <div className="bg-surface-container-lowest rounded-xl p-4 shadow-sm border border-outline-variant/10">
      <div className="flex gap-3">
        {/* Thumbnail */}
        <div className="w-20 h-14 rounded-lg bg-surface-container-high overflow-hidden shrink-0">
          {stream.thumbnailUrl ? (
            <img className="w-full h-full object-cover" src={stream.thumbnailUrl} alt={stream.title} />
          ) : (
            <div className="w-full h-full bg-gradient-to-br from-surface-container to-surface-container-high" />
          )}
        </div>
        <div className="flex-1 min-w-0">
          <h3 className="font-semibold text-on-surface text-sm truncate">{stream.title}</h3>
          <p className="text-xs text-on-surface-variant font-medium truncate">{clubName}</p>
        </div>
      </div>
      <div className="flex items-center justify-between mt-3 pt-3 border-t border-outline-variant/10">
        <div className="flex items-center gap-2 flex-wrap">
          <StatusBadge status={stream.status} />
          <span className={`inline-flex items-center px-2 py-0.5 rounded-full text-[10px] font-black tracking-widest uppercase ${TYPE_BADGE_STYLES[stream.type] ?? 'bg-surface-container text-on-surface-variant'}`}>
            {stream.type === 'behind_the_scenes' ? 'BTS' : stream.type}
          </span>
        </div>
        <div className="flex items-center gap-3">
          {stream.status !== 'scheduled' && (
            <span className="text-xs font-semibold text-on-surface-variant">
              {formatViews(stream.viewCount)} views
            </span>
          )}
          {stream.status !== 'scheduled' && (
            <button
              onClick={onWatch}
              className="text-primary hover:bg-primary/10 px-2 py-1 rounded-md text-xs font-bold transition-all inline-flex items-center gap-1 min-h-[44px]"
            >
              <Play size={14} /> Watch
            </button>
          )}
          {stream.status === 'live' ? (
            <button className="text-error hover:bg-error-container/30 px-2 py-1 rounded-md text-xs font-bold transition-all inline-flex items-center gap-1 min-h-[44px]">
              <StopCircle size={14} /> End
            </button>
          ) : stream.status === 'scheduled' ? (
            <button className="bg-primary text-on-primary hover:bg-on-primary-container px-3 py-1 rounded-md text-xs font-bold transition-all shadow-sm min-h-[44px]">
              Go Live
            </button>
          ) : (
            <button className="text-on-surface-variant hover:text-primary transition-colors min-w-[44px] min-h-[44px] flex items-center justify-center">
              <MoreVertical size={18} />
            </button>
          )}
        </div>
      </div>
    </div>
  );
}

export function StreamTable() {
  const [status, setStatus] = useState<StreamStatus | ''>('');
  const [type, setType] = useState<StreamType | ''>('');
  const [watchStream, setWatchStream] = useState<Stream | null>(null);

  const filters = {
    ...(status ? { status } : {}),
    ...(type ? { type } : {}),
  };

  const { data: streams, isLoading, error } = useStreams(
    Object.keys(filters).length > 0 ? filters : undefined,
  );
  const { data: clubs } = useClubs();

  const clubNameMap = useMemo(() => {
    const map = new Map<string, string>();
    clubs?.forEach((c) => map.set(c.id, c.name));
    return map;
  }, [clubs]);

  if (error) {
    return (
      <div className="rounded-2xl border border-error/30 bg-error-container p-6 text-center text-sm text-on-error-container">
        Failed to load streams: {error.message}
      </div>
    );
  }

  const liveCount = streams?.filter((s) => s.status === 'live').length ?? 0;
  const totalViews = streams?.reduce((sum, s) => sum + s.viewCount, 0) ?? 0;

  return (
    <div className="space-y-4 md:space-y-8">
      {/* Filters row */}
      <StreamFilters
        status={status}
        type={type}
        onStatusChange={setStatus}
        onTypeChange={setType}
      />

      {/* Premium Table */}
      {isLoading ? (
        <div className="h-96 animate-pulse rounded-2xl bg-surface-container-lowest shadow-xl" />
      ) : (
        <>
          {/* Mobile: Card list */}
          <div className="md:hidden space-y-3">
            {(streams ?? []).length === 0 ? (
              <div className="rounded-2xl bg-surface-container-lowest p-12 text-center text-on-surface-variant shadow-sm">
                No streams found
              </div>
            ) : (
              (streams ?? []).map((stream: Stream) => (
                <StreamMobileCard
                  key={stream.id}
                  stream={stream}
                  clubName={clubNameMap.get(stream.clubId) ?? stream.clubId}
                  onWatch={() => setWatchStream(stream)}
                />
              ))
            )}
          </div>

          {/* Desktop: Table */}
          <div className="hidden md:block bg-surface-container-lowest rounded-2xl shadow-xl overflow-hidden border border-outline-variant/10">
            <div className="overflow-x-auto">
              <table className="w-full text-left border-collapse">
                <thead>
                  <tr className="bg-surface-container-low border-b border-outline-variant/20">
                    <th className="py-4 px-6 text-[11px] font-bold text-outline uppercase tracking-wider">Title</th>
                    <th className="py-4 px-6 text-[11px] font-bold text-outline uppercase tracking-wider">Club</th>
                    <th className="py-4 px-6 text-[11px] font-bold text-outline uppercase tracking-wider text-center">Type</th>
                    <th className="py-4 px-6 text-[11px] font-bold text-outline uppercase tracking-wider">Status</th>
                    <th className="py-4 px-6 text-[11px] font-bold text-outline uppercase tracking-wider text-right">Views</th>
                    <th className="py-4 px-6 text-[11px] font-bold text-outline uppercase tracking-wider text-center">Duration</th>
                    <th className="py-4 px-6 text-[11px] font-bold text-outline uppercase tracking-wider text-right">Actions</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-outline-variant/10">
                  {(streams ?? []).map((stream: Stream, index: number) => (
                    <tr
                      key={stream.id}
                      className={`hover:bg-primary/5 transition-colors group ${index % 2 !== 0 ? 'bg-surface-container-low/30' : ''}`}
                    >
                      {/* Title with thumbnail */}
                      <td className="py-5 px-6">
                        <div className="flex items-center gap-3">
                          <div className="w-10 h-10 rounded-lg bg-surface-container-high flex items-center justify-center overflow-hidden shrink-0">
                            {stream.thumbnailUrl ? (
                              <img className="w-full h-full object-cover" src={stream.thumbnailUrl} alt={stream.title} />
                            ) : (
                              <div className="w-full h-full bg-gradient-to-br from-surface-container to-surface-container-high" />
                            )}
                          </div>
                          <span className="font-semibold text-on-surface text-sm">{stream.title}</span>
                        </div>
                      </td>

                      {/* Club */}
                      <td className="py-5 px-6 text-sm text-on-surface-variant font-medium">
                        {clubNameMap.get(stream.clubId) ?? stream.clubId}
                      </td>

                      {/* Type badge */}
                      <td className="py-5 px-6 text-center">
                        <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-[10px] font-black tracking-widest uppercase ${TYPE_BADGE_STYLES[stream.type] ?? 'bg-surface-container text-on-surface-variant'}`}>
                          {stream.type === 'behind_the_scenes' ? 'BTS' : stream.type}
                        </span>
                      </td>

                      {/* Status */}
                      <td className="py-5 px-6">
                        <StatusBadge status={stream.status} />
                      </td>

                      {/* Views */}
                      <td className="py-5 px-6 text-right font-mono text-sm font-semibold">
                        {stream.status === 'scheduled' ? (
                          <span className="text-outline">--</span>
                        ) : (
                          formatViews(stream.viewCount)
                        )}
                      </td>

                      {/* Duration */}
                      <td className="py-5 px-6 text-center text-on-surface-variant text-sm font-medium">
                        {formatDuration(stream.duration)}
                      </td>

                      {/* Actions */}
                      <td className="py-5 px-6 text-right">
                        <div className="flex items-center justify-end gap-2">
                          {stream.status !== 'scheduled' && (
                            <button
                              onClick={() => setWatchStream(stream)}
                              className="text-primary hover:bg-primary/10 px-3 py-1 rounded-md text-xs font-bold transition-all inline-flex items-center gap-1"
                            >
                              <Play size={14} /> Watch
                            </button>
                          )}
                          {stream.status === 'live' ? (
                            <button className="text-error hover:bg-error-container/30 px-3 py-1 rounded-md text-xs font-bold transition-all inline-flex items-center gap-1">
                              <StopCircle size={14} /> End
                            </button>
                          ) : stream.status === 'scheduled' ? (
                            <button className="bg-primary text-on-primary hover:bg-on-primary-container px-3 py-1 rounded-md text-xs font-bold transition-all shadow-sm">
                              Go Live
                            </button>
                          ) : (
                            <button className="text-on-surface-variant hover:text-primary transition-colors">
                              <MoreVertical size={18} />
                            </button>
                          )}
                        </div>
                      </td>
                    </tr>
                  ))}
                  {(streams ?? []).length === 0 && (
                    <tr>
                      <td colSpan={7} className="px-6 py-12 text-center text-on-surface-variant">
                        No streams found
                      </td>
                    </tr>
                  )}
                </tbody>
              </table>
            </div>

            {/* Pagination */}
            <div className="flex items-center justify-between px-8 py-4 bg-surface-container-low/50">
              <p className="text-xs font-semibold text-on-surface-variant">
                Showing {(streams ?? []).length} streams
              </p>
              <div className="flex items-center gap-1">
                <button className="w-8 h-8 flex items-center justify-center rounded-lg bg-primary text-on-primary text-xs font-bold">
                  1
                </button>
              </div>
            </div>
          </div>
        </>
      )}

      {/* Bottom Bento Stats */}
      <div className="grid grid-cols-2 md:grid-cols-4 gap-3 md:gap-6">
        <div className="bg-surface-container-low p-4 md:p-6 rounded-2xl flex flex-col justify-between h-28 md:h-32 border border-outline-variant/10">
          <span className="text-[10px] font-black uppercase tracking-widest text-outline">Total Active Streams</span>
          <div className="flex items-end justify-between">
            <span className="text-2xl md:text-3xl font-black text-on-surface">{(streams ?? []).length}</span>
            <span className="text-xs font-bold text-primary flex items-center gap-0.5">
              +{liveCount} <TrendingUp size={14} />
            </span>
          </div>
        </div>

        <div className="bg-surface-container-low p-4 md:p-6 rounded-2xl flex flex-col justify-between h-28 md:h-32 border border-outline-variant/10">
          <span className="text-[10px] font-black uppercase tracking-widest text-outline">Concurrent Viewers</span>
          <div className="flex items-end justify-between">
            <span className="text-2xl md:text-3xl font-black text-on-surface">
              {totalViews >= 1_000_000 ? `${(totalViews / 1_000_000).toFixed(1)}M` : totalViews >= 1_000 ? `${(totalViews / 1_000).toFixed(0)}K` : totalViews}
            </span>
            <span className="text-xs font-bold text-primary flex items-center gap-0.5">
              +18% <TrendingUp size={14} />
            </span>
          </div>
        </div>

        <div className="bg-surface-container-low p-4 md:p-6 rounded-2xl flex flex-col justify-between h-28 md:h-32 border border-outline-variant/10">
          <span className="text-[10px] font-black uppercase tracking-widest text-outline">Storage Used</span>
          <div className="flex items-end justify-between">
            <span className="text-2xl md:text-3xl font-black text-on-surface">4.2 TB</span>
            <div className="w-16 md:w-20 h-1.5 bg-surface-container-high rounded-full overflow-hidden">
              <div className="h-full bg-primary rounded-full" style={{ width: '65%' }} />
            </div>
          </div>
        </div>

        <div className="bg-sidebar-bg p-4 md:p-6 rounded-2xl flex flex-col justify-between h-28 md:h-32 relative overflow-hidden group">
          <div className="absolute inset-0 bg-primary/10 opacity-0 group-hover:opacity-100 transition-opacity" />
          <span className="text-[10px] font-black uppercase tracking-widest text-white/40 relative z-10">Bandwidth Health</span>
          <div className="flex items-center gap-2 relative z-10">
            <span className="text-2xl md:text-3xl font-black text-white">99.9</span>
            <span className="text-xs font-bold text-sidebar-accent">%</span>
          </div>
          <Zap size={80} className="absolute -right-4 -bottom-4 text-white/5 rotate-12" />
        </div>
      </div>

      {/* Watch Stream Modal */}
      {watchStream && (
        <StreamModal
          stream={watchStream}
          clubName={clubNameMap.get(watchStream.clubId) ?? watchStream.clubId}
          onClose={() => setWatchStream(null)}
        />
      )}
    </div>
  );
}
