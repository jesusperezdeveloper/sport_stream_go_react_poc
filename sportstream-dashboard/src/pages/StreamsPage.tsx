import { Plus, ChevronRight } from 'lucide-react';
import { Header } from '../components/layout/Header';
import { StreamTable } from '../components/features/streams/StreamTable';

export function StreamsPage() {
  return (
    <div>
      <Header title="Streams" />
      <div className="p-8 space-y-8 bg-background min-h-screen">
        {/* Header Section */}
        <div className="flex flex-col md:flex-row md:items-end justify-between gap-4">
          <div>
            <nav className="flex items-center gap-2 text-xs font-bold text-outline uppercase tracking-widest mb-1">
              <span>Management</span>
              <ChevronRight size={10} />
              <span className="text-primary font-black">Streams</span>
            </nav>
            <h2 className="font-black text-4xl tracking-tight text-on-surface">Streams</h2>
          </div>
          <button className="flex items-center gap-2 px-4 py-2.5 bg-gradient-to-br from-primary to-primary-container text-white font-bold rounded-xl shadow-lg shadow-primary/20 hover:scale-[1.02] active:scale-95 transition-all">
            <Plus size={20} />
            <span>New Stream</span>
          </button>
        </div>

        <StreamTable />
      </div>
    </div>
  );
}
