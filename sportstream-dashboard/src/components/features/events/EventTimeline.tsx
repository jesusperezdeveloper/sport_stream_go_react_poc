import { useState } from 'react';
import { ChevronDown } from 'lucide-react';
import { EventCard } from './EventCard';
import { useEvents } from '../../../hooks/useEvents';
import type { EventStatus } from '../../../types/event';

const statuses: EventStatus[] = ['upcoming', 'live', 'completed', 'cancelled'];

export function EventTimeline() {
  const [statusFilter, setStatusFilter] = useState<EventStatus | ''>('');
  const [sportFilter, setSportFilter] = useState('');

  const filters = {
    ...(statusFilter ? { status: statusFilter } : {}),
    ...(sportFilter ? { sport: sportFilter } : {}),
  };

  const { data: events, isLoading, error } = useEvents(
    Object.keys(filters).length > 0 ? filters : undefined,
  );

  if (error) {
    return (
      <div className="rounded-2xl border border-error/30 bg-error-container p-6 text-center text-sm text-on-error-container">
        Failed to load events: {error.message}
      </div>
    );
  }

  return (
    <div className="space-y-4 md:space-y-6">
      {/* Filters */}
      <div className="flex flex-wrap items-center gap-2 md:gap-3">
        <div className="relative">
          <select
            value={statusFilter}
            onChange={(e) => setStatusFilter(e.target.value as EventStatus | '')}
            className="appearance-none bg-surface-container-lowest border border-outline-variant/30 rounded-full px-4 py-2 pr-8 text-sm font-medium hover:bg-surface-container-low transition-all cursor-pointer focus:outline-none focus:ring-2 focus:ring-primary/20 min-h-[44px]"
          >
            <option value="">All Statuses</option>
            {statuses.map((s) => (
              <option key={s} value={s}>
                {s.charAt(0).toUpperCase() + s.slice(1)}
              </option>
            ))}
          </select>
          <ChevronDown size={16} className="absolute right-3 top-1/2 -translate-y-1/2 text-on-surface-variant pointer-events-none" />
        </div>

        <input
          type="text"
          placeholder="Filter by sport..."
          value={sportFilter}
          onChange={(e) => setSportFilter(e.target.value)}
          className="bg-surface-container-lowest border border-outline-variant/30 rounded-full px-4 py-2 text-sm font-medium placeholder-outline focus:outline-none focus:ring-2 focus:ring-primary/20 min-h-[44px]"
        />
      </div>

      {/* Events list */}
      {isLoading ? (
        <div className="space-y-4">
          {Array.from({ length: 4 }).map((_, i) => (
            <div key={i} className="h-20 animate-pulse rounded-xl bg-surface-container-lowest" />
          ))}
        </div>
      ) : !events || events.length === 0 ? (
        <div className="rounded-2xl bg-surface-container-lowest p-12 text-center text-sm text-on-surface-variant shadow-sm">
          No events found
        </div>
      ) : (
        <div className="bg-surface-container-low rounded-2xl p-3 md:p-6 space-y-2 md:space-y-3">
          {events.map((event) => (
            <EventCard key={event.id} event={event} />
          ))}
        </div>
      )}
    </div>
  );
}
