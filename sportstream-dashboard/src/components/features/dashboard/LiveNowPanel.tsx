import { useStreams } from '../../../hooks/useStreams';

function formatViewers(n: number): string {
  if (n >= 1_000_000) return `${(n / 1_000_000).toFixed(1)}M`;
  if (n >= 1_000) return `${Math.round(n / 1_000)}K`;
  return n.toLocaleString();
}

export function LiveNowPanel() {
  const { data: streams, isLoading } = useStreams({ status: 'live' });

  const featured = streams?.[0];
  const secondary = streams?.slice(1) ?? [];

  return (
    <div className="bg-inverse-surface rounded-2xl p-4 md:p-6 flex flex-col gap-4">
      <div className="flex items-center gap-3 mb-2">
        <div className="w-2 h-2 bg-error rounded-full kinetic-pulse" />
        <h2 className="text-lg font-bold tracking-tight text-white">LIVE NOW</h2>
      </div>

      {isLoading ? (
        <div className="space-y-3">
          <div className="aspect-video w-full animate-pulse rounded-xl bg-white/10" />
          <div className="h-14 animate-pulse rounded-xl bg-white/10" />
        </div>
      ) : !streams || streams.length === 0 ? (
        <p className="py-8 text-center text-sm text-white/40">
          No live streams right now
        </p>
      ) : (
        <>
          {/* Featured stream */}
          {featured && (
            <div className="group relative bg-white/5 rounded-xl overflow-hidden hover:bg-white/10 transition-colors">
              <div className="aspect-video w-full bg-slate-800 relative">
                {featured.thumbnailUrl ? (
                  <img alt={featured.title} className="w-full h-full object-cover opacity-80" src={featured.thumbnailUrl} />
                ) : (
                  <div className="w-full h-full bg-gradient-to-br from-slate-700 to-slate-900" />
                )}
                <div className="absolute bottom-2 left-2 bg-error text-[10px] font-black px-2 py-0.5 rounded-sm text-white">
                  {formatViewers(featured.viewCount)} VIEWERS
                </div>
              </div>
              <div className="p-3 md:p-4">
                <div className="flex justify-between items-start mb-1">
                  <h3 className="font-bold text-white leading-tight text-sm">{featured.title}</h3>
                  {featured.tags?.[0] && (
                    <span className="text-[10px] bg-white/10 px-2 py-0.5 rounded text-white/60 uppercase">
                      {featured.tags[0]}
                    </span>
                  )}
                </div>
                <p className="text-[10px] text-white/40 uppercase tracking-widest font-bold">
                  {featured.type === 'live' ? 'Live Match' : featured.type}
                </p>
              </div>
            </div>
          )}

          {/* Secondary streams - horizontal scroll on mobile */}
          <div className="flex flex-row overflow-x-auto gap-3 md:flex-col md:overflow-x-visible -mx-4 px-4 md:mx-0 md:px-0">
            {secondary.map((stream) => (
              <div
                key={stream.id}
                className="group flex gap-3 md:gap-4 p-3 bg-white/5 rounded-xl hover:bg-white/10 transition-colors shrink-0 w-[280px] md:w-auto"
              >
                <div className="w-20 h-14 md:w-24 md:h-16 shrink-0 bg-slate-800 rounded-lg overflow-hidden">
                  {stream.thumbnailUrl ? (
                    <img alt={stream.title} className="w-full h-full object-cover opacity-80" src={stream.thumbnailUrl} />
                  ) : (
                    <div className="w-full h-full bg-gradient-to-br from-slate-700 to-slate-900" />
                  )}
                </div>
                <div className="flex-1 min-w-0">
                  <div className="flex justify-between">
                    <h3 className="font-bold text-sm text-white truncate">{stream.title}</h3>
                    <span className="text-[10px] text-white/40 shrink-0 ml-2">
                      {formatViewers(stream.viewCount)}
                    </span>
                  </div>
                  {stream.tags?.[0] && (
                    <div className="text-[10px] bg-white/10 px-2 py-0.5 rounded inline-block text-white/60 mt-1 uppercase">
                      {stream.tags[0]}
                    </div>
                  )}
                </div>
              </div>
            ))}
          </div>
        </>
      )}
    </div>
  );
}
