import { Clock, CheckCircle, Archive, Video } from 'lucide-react';

interface StatusBadgeProps {
  status: string;
}

const statusConfig: Record<string, { bg: string; text: string; icon?: React.ReactNode; dot?: boolean }> = {
  live: {
    bg: 'bg-error-container/20',
    text: 'text-error',
    dot: true,
  },
  scheduled: {
    bg: 'bg-tertiary-container/10',
    text: 'text-tertiary',
    icon: <Clock size={12} className="md:w-3.5 md:h-3.5" />,
  },
  ended: {
    bg: 'bg-surface-container-high',
    text: 'text-outline',
    icon: <CheckCircle size={12} className="md:w-3.5 md:h-3.5" />,
  },
  completed: {
    bg: 'bg-surface-container-high',
    text: 'text-outline',
    icon: <CheckCircle size={12} className="md:w-3.5 md:h-3.5" />,
  },
  archived: {
    bg: 'bg-surface-container-highest',
    text: 'text-secondary',
    icon: <Archive size={12} className="md:w-3.5 md:h-3.5" />,
  },
  vod: {
    bg: 'bg-blue-50',
    text: 'text-blue-600',
    icon: <Video size={12} className="md:w-3.5 md:h-3.5" />,
  },
  upcoming: {
    bg: 'bg-tertiary-container/10',
    text: 'text-tertiary',
    icon: <Clock size={12} className="md:w-3.5 md:h-3.5" />,
  },
  cancelled: {
    bg: 'bg-error-container/20',
    text: 'text-error',
    icon: <CheckCircle size={12} className="md:w-3.5 md:h-3.5" />,
  },
};

export function StatusBadge({ status }: StatusBadgeProps) {
  const config = statusConfig[status] ?? { bg: 'bg-surface-container', text: 'text-on-surface-variant' };

  return (
    <div className={`flex items-center gap-1.5 md:gap-2 ${config.text} text-[10px] md:text-[11px] font-bold ${config.bg} px-2 py-1 rounded-full w-fit`}>
      {config.dot && (
        <span className="relative flex h-2 w-2">
          <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-error opacity-75" />
          <span className="relative inline-flex rounded-full h-2 w-2 bg-error" />
        </span>
      )}
      {config.icon}
      {status.toUpperCase()}
    </div>
  );
}
