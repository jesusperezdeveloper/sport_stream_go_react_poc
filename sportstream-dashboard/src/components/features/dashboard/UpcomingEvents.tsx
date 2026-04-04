import { Clock, MapPin } from 'lucide-react';
import { useUpcomingEvents } from '../../../hooks/useEvents';

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

const SPORT_COLORS: Record<string, { bg: string; text: string }> = {
  football: { bg: 'bg-primary-container/30', text: 'text-on-primary-container' },
  tennis: { bg: 'bg-primary-container/30', text: 'text-on-primary-container' },
  basketball: { bg: 'bg-secondary-container/30', text: 'text-on-secondary-container' },
  esports: { bg: 'bg-secondary-container', text: 'text-on-secondary-container' },
  volleyball: { bg: 'bg-surface-variant', text: 'text-on-surface-variant' },
  swimming: { bg: 'bg-surface-variant', text: 'text-on-surface-variant' },
};

export function UpcomingEventsPanel() {
  const { data: events, isLoading } = useUpcomingEvents(5);

  return (
    <div className="bg-surface-container-low rounded-2xl p-8">
      <h2 className="text-xl font-bold tracking-tight mb-8 text-on-surface">Upcoming Events</h2>

      {isLoading ? (
        <div className="space-y-4">
          {Array.from({ length: 3 }).map((_, i) => (
            <div key={i} className="h-20 animate-pulse rounded-xl bg-surface-container" />
          ))}
        </div>
      ) : !events || events.length === 0 ? (
        <p className="py-8 text-center text-sm text-on-surface-variant">
          No upcoming events
        </p>
      ) : (
        <div className="space-y-4">
          {events.map((event, index) => {
            const { month, day } = getMonthDay(event.startTime);
            const sportKey = event.sport.toLowerCase();
            const sportColor = SPORT_COLORS[sportKey] ?? { bg: 'bg-surface-variant', text: 'text-on-surface-variant' };

            return (
              <div
                key={event.id}
                className={`flex gap-6 items-start p-4 bg-surface-container-lowest rounded-xl ${
                  index % 2 !== 0 ? 'opacity-80' : ''
                }`}
              >
                <div className={`shrink-0 text-center ${index % 2 !== 0 ? 'opacity-60' : ''}`}>
                  <div className="text-xs font-bold text-on-surface-variant">{month}</div>
                  <div className="text-2xl font-black text-on-surface">{day}</div>
                </div>
                <div className="flex-1 min-w-0">
                  <div className="flex justify-between mb-1">
                    <h4 className="font-bold text-on-surface truncate">{event.title}</h4>
                    <span className={`text-[10px] px-2 py-0.5 ${sportColor.bg} ${sportColor.text} rounded font-black uppercase shrink-0 ml-2`}>
                      {event.sport}
                    </span>
                  </div>
                  <div className="flex items-center gap-4 text-xs text-on-surface-variant font-medium">
                    <span className="flex items-center gap-1">
                      <Clock size={14} />
                      {formatTime(event.startTime)}
                    </span>
                    <span className="flex items-center gap-1">
                      <MapPin size={14} />
                      {event.venue}
                    </span>
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
