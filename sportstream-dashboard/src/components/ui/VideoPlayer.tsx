import { useRef, useEffect, useState } from 'react';
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

  useEffect(() => {
    const video = videoRef.current;
    if (!video || !src) return;

    if (src.endsWith('.m3u8') || src.includes('.m3u8')) {
      if (Hls.isSupported()) {
        const hls = new Hls({
          enableWorker: true,
          lowLatencyMode: true,
        });
        hlsRef.current = hls;
        hls.loadSource(src);
        hls.attachMedia(video);
        hls.on(Hls.Events.MANIFEST_PARSED, () => {
          if (autoPlay) {
            video.play().catch(() => {});
          }
        });
        hls.on(Hls.Events.ERROR, (_, data) => {
          if (data.fatal) {
            setError(true);
          }
        });
      } else if (video.canPlayType('application/vnd.apple.mpegurl')) {
        // Safari native HLS
        video.src = src;
        if (autoPlay) video.play().catch(() => {});
      }
    } else {
      // Regular video
      video.src = src;
      if (autoPlay) video.play().catch(() => {});
    }

    return () => {
      if (hlsRef.current) {
        hlsRef.current.destroy();
        hlsRef.current = null;
      }
    };
  }, [src, autoPlay]);

  if (error) {
    return (
      <div className={`flex items-center justify-center bg-slate-900 ${className}`}>
        <div className="text-center">
          <div className="text-white/40 text-sm">Stream unavailable</div>
        </div>
      </div>
    );
  }

  return (
    <div className={`relative group ${className}`}>
      <video
        ref={videoRef}
        poster={poster}
        muted={muted}
        playsInline
        controls
        className="w-full h-full object-cover rounded-xl"
        onPlay={() => setIsPlaying(true)}
        onPause={() => setIsPlaying(false)}
      />
      {!isPlaying && !autoPlay && (
        <div className="absolute inset-0 flex items-center justify-center bg-black/30 rounded-xl cursor-pointer"
          onClick={() => videoRef.current?.play()}>
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
