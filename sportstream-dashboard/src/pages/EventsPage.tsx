import { ChevronRight } from 'lucide-react';
import { Header } from '../components/layout/Header';
import { EventTimeline } from '../components/features/events/EventTimeline';

export function EventsPage() {
  return (
    <div>
      <Header title="Events" />
      <div className="p-4 md:p-8 space-y-4 md:space-y-8 bg-background min-h-screen">
        <div>
          <nav className="flex items-center gap-2 text-xs font-bold text-outline uppercase tracking-widest mb-1">
            <span>Management</span>
            <ChevronRight size={10} />
            <span className="text-primary font-black">Events</span>
          </nav>
          <h2 className="font-black text-2xl md:text-4xl tracking-tight text-on-surface">Events</h2>
        </div>
        <EventTimeline />
      </div>
    </div>
  );
}
