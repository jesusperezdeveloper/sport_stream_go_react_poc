import { ChevronRight } from 'lucide-react';
import { Header } from '../components/layout/Header';
import { ClubGrid } from '../components/features/clubs/ClubGrid';

export function ClubsPage() {
  return (
    <div>
      <Header title="Clubs" />
      <div className="p-8 space-y-8 bg-background min-h-screen">
        <div>
          <nav className="flex items-center gap-2 text-xs font-bold text-outline uppercase tracking-widest mb-1">
            <span>Management</span>
            <ChevronRight size={10} />
            <span className="text-primary font-black">Clubs</span>
          </nav>
          <h2 className="font-black text-4xl tracking-tight text-on-surface">Clubs</h2>
        </div>
        <ClubGrid />
      </div>
    </div>
  );
}
