import { useEffect } from 'react';
import { X, Users, Radio } from 'lucide-react';
import { VideoPlayer } from './VideoPlayer';
import { StatusBadge } from './StatusBadge';
import type { Stream } from '../../types/stream';

interface StreamModalProps {
  stream: Stream;
  clubName: string;
  onClose: () => void;
}

function formatViewers(n: number): string {
  if (n >= 1_000_000) return `${(n / 1_000_000).toFixed(1)}M`;
  if (n >= 1_000) return `${Math.round(n / 1_000)}K`;
  return n.toLocaleString();
}

export function StreamModal({ stream, clubName, onClose }: StreamModalProps) {
  // Close on Escape key
  useEffect(() => {
    function handleKeyDown(e: KeyboardEvent) {
      if (e.key === 'Escape') onClose();
    }
    document.addEventListener('keydown', handleKeyDown);
    // Prevent body scroll while modal is open
    document.body.style.overflow = 'hidden';
    return () => {
      document.removeEventListener('keydown', handleKeyDown);
      document.body.style.overflow = '';
    };
  }, [onClose]);

  return (
    <div
      className="fixed inset-0 z-50 flex items-center justify-center p-4 md:p-8"
      onClick={onClose}
    >
      {/* Backdrop */}
      <div className="absolute inset-0 bg-black/80 backdrop-blur-sm" />

      {/* Modal content */}
      <div
        className="relative w-full max-w-4xl bg-surface-container-lowest rounded-2xl shadow-2xl overflow-hidden border border-outline-variant/10"
        onClick={(e) => e.stopPropagation()}
      >
        {/* Close button */}
        <button
          onClick={onClose}
          className="absolute top-4 right-4 z-10 w-10 h-10 flex items-center justify-center rounded-full bg-black/50 text-white hover:bg-black/70 transition-colors"
        >
          <X size={20} />
        </button>

        {/* Video Player */}
        <VideoPlayer
          src={stream.streamUrl}
          poster={stream.thumbnailUrl}
          autoPlay={stream.status === 'live'}
          muted
          className="aspect-video w-full"
        />

        {/* Stream info */}
        <div className="p-4 md:p-6">
          <div className="flex items-start justify-between gap-4">
            <div className="flex-1 min-w-0">
              <h2 className="text-lg font-bold text-on-surface truncate">{stream.title}</h2>
              <p className="text-sm text-on-surface-variant font-medium mt-1">{clubName}</p>
              {stream.description && (
                <p className="text-sm text-on-surface-variant/70 mt-2">{stream.description}</p>
              )}
            </div>
            <div className="flex items-center gap-3 shrink-0">
              <StatusBadge status={stream.status} />
            </div>
          </div>

          <div className="flex items-center gap-4 mt-4 pt-4 border-t border-outline-variant/10">
            {stream.status === 'live' && (
              <div className="flex items-center gap-1.5 text-error">
                <Radio size={14} className="animate-pulse" />
                <span className="text-xs font-bold uppercase">Live</span>
              </div>
            )}
            {stream.viewCount > 0 && (
              <div className="flex items-center gap-1.5 text-on-surface-variant">
                <Users size={14} />
                <span className="text-xs font-semibold">{formatViewers(stream.viewCount)} viewers</span>
              </div>
            )}
            {stream.tags && stream.tags.length > 0 && (
              <div className="flex items-center gap-1.5 flex-wrap">
                {stream.tags.map((tag) => (
                  <span
                    key={tag}
                    className="text-[10px] bg-surface-container px-2 py-0.5 rounded-full text-on-surface-variant uppercase font-bold tracking-wider"
                  >
                    {tag}
                  </span>
                ))}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
