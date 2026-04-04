import { MapPin, Clock, Radio } from 'lucide-react';
import type { SportEvent } from '../../../types/event';

interface EventCardProps {
  event: SportEvent;
}

const SPORT_BADGE: Record<string, string> = {
  football: 'bg-primary-container/30 text-on-primary-container',
  tennis: 'bg-primary-container/30 text-on-primary-container',
  basketball: 'bg-secondary-container text-on-secondary-container',
  esports: 'bg-secondary-container text-on-secondary-container',
  volleyball: 'bg-surface-variant text-on-surface-variant',
  swimming: 'bg-surface-variant text-on-surface-variant',
};

function getMonthDay(dateStr: string): { month: string; day: string } {
  const d = new Date(dateStr);
  return {
    month: d.toLocaleDateString('en-US', { month: 'short' }).toUpperCase(),
    day: d.getDate().toString(),
  };
}

function formatTime(dateStr: string): string {
  const d = new Date(dateStr);
  return d.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit', hour12: false });
}

export function EventCard({ event }: EventCardProps) {
  const { month, day } = getMonthDay(event.startTime);
  const sportKey = event.sport.toLowerCase();
  const badgeStyle = SPORT_BADGE[sportKey] ?? 'bg-surface-variant text-on-surface-variant';

  return (
    <div className="flex gap-3 md:gap-6 items-start p-3 md:p-4 bg-surface-container-lowest rounded-xl hover:shadow-sm transition-shadow">
      <div className="shrink-0 text-center">
        <div className="text-[10px] md:text-xs font-bold text-on-surface-variant">{month}</div>
        <div className="text-xl md:text-2xl font-black text-on-surface">{day}</div>
      </div>
      <div className="flex-1 min-w-0">
        <div className="flex justify-between mb-1">
          <h4 className="font-bold text-on-surface truncate text-sm md:text-base">{event.title}</h4>
          <span className={`text-[10px] px-2 py-0.5 ${badgeStyle} rounded font-black uppercase shrink-0 ml-2`}>
            {event.sport}
          </span>
        </div>
        <div className="flex items-center gap-2 md:gap-4 text-[11px] md:text-xs text-on-surface-variant font-medium flex-wrap">
          <span className="flex items-center gap-1">
            <Clock size={12} className="shrink-0" />
            {formatTime(event.startTime)}
          </span>
          <span className="flex items-center gap-1">
            <MapPin size={12} className="shrink-0" />
            <span className="truncate">{event.venue}</span>
          </span>
          {event.streamId && (
            <span className="flex items-center gap-1 text-primary">
              <Radio size={12} />
              Linked
            </span>
          )}
        </div>
      </div>
    </div>
  );
}
