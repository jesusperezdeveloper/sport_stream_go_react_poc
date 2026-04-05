import { useRef, useEffect, useState, useCallback } from 'react';
import Hls from 'hls.js';

interface VideoPlayerProps {
  src: string;
  poster?: string;
  autoPlay?: boolean;
  muted?: boolean;
  className?: string;
}

export function VideoPlayer({ src, poster, autoPlay = false, muted = true, className = '' }: VideoPlayerProps) {
  const videoRef = useRef<HTMLVideoElement>(null);
  const hlsRef = useRef<Hls | null>(null);
  const [isPlaying, setIsPlaying] = useState(false);
  const [error, setError] = useState(false);
  const [loading, setLoading] = useState(true);

  const destroyHls = useCallback(() => {
    if (hlsRef.current) {
      hlsRef.current.destroy();
      hlsRef.current = null;
    }
  }, []);

  useEffect(() => {
    const video = videoRef.current;
    if (!video || !src) return;

    setError(false);
    setLoading(true);

    const isHls = src.includes('.m3u8');

    if (isHls) {
      // Safari: prefer native HLS (more reliable)
      if (video.canPlayType('application/vnd.apple.mpegurl')) {
        video.src = src;
        video.addEventListener('loadedmetadata', () => {
          setLoading(false);
          if (autoPlay) video.play().catch(() => {});
        }, { once: true });
        video.addEventListener('error', () => setError(true), { once: true });
      } else if (Hls.isSupported()) {
        const hls = new Hls({
          enableWorker: true,
          lowLatencyMode: true,
          fragLoadingMaxRetry: 3,
          manifestLoadingMaxRetry: 3,
          levelLoadingMaxRetry: 3,
        });
        hlsRef.current = hls;
        hls.loadSource(src);
        hls.attachMedia(video);

        hls.on(Hls.Events.MANIFEST_PARSED, () => {
          setLoading(false);
          if (autoPlay) video.play().catch(() => {});
        });

        hls.on(Hls.Events.ERROR, (_, data) => {
          if (data.fatal) {
            switch (data.type) {
              case Hls.ErrorTypes.NETWORK_ERROR:
                hls.startLoad();
                break;
              case Hls.ErrorTypes.MEDIA_ERROR:
                hls.recoverMediaError();
                break;
              default:
                setError(true);
                break;
            }
          }
        });
      } else {
        setError(true);
        setLoading(false);
      }
    } else {
      video.src = src;
      video.addEventListener('loadedmetadata', () => {
        setLoading(false);
        if (autoPlay) video.play().catch(() => {});
      }, { once: true });
      video.addEventListener('error', () => setError(true), { once: true });
    }

    return destroyHls;
  }, [src, autoPlay, destroyHls]);

  if (error) {
    return (
      <div className={`flex items-center justify-center bg-slate-900 rounded-xl ${className}`}>
        <div className="text-center p-8">
          <div className="text-white/60 text-sm font-medium mb-2">Stream unavailable</div>
          <button
            onClick={() => { setError(false); setLoading(true); }}
            className="text-xs text-primary hover:underline"
          >
            Retry
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className={`relative group ${className}`}>
      {loading && (
        <div className="absolute inset-0 flex items-center justify-center bg-slate-900 rounded-xl z-10">
          <div className="flex flex-col items-center gap-2">
            <div className="w-8 h-8 border-2 border-primary/30 border-t-primary rounded-full animate-spin" />
            <span className="text-white/40 text-xs">Loading stream...</span>
          </div>
        </div>
      )}
      <video
        ref={videoRef}
        poster={poster}
        muted={muted}
        playsInline
        controls
        className="w-full h-full object-cover rounded-xl"
        onPlay={() => setIsPlaying(true)}
        onPause={() => setIsPlaying(false)}
        onCanPlay={() => setLoading(false)}
      />
      {!isPlaying && !autoPlay && !loading && (
        <div
          className="absolute inset-0 flex items-center justify-center bg-black/30 rounded-xl cursor-pointer"
          onClick={() => videoRef.current?.play()}
        >
          <div className="w-16 h-16 flex items-center justify-center rounded-full bg-white/20 backdrop-blur-sm">
            <svg className="w-8 h-8 text-white ml-1" fill="currentColor" viewBox="0 0 24 24">
              <path d="M8 5v14l11-7z" />
            </svg>
          </div>
        </div>
      )}
    </div>
  );
}
