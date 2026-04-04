import { ClubCard } from './ClubCard';
import { useClubs } from '../../../hooks/useClubs';

export function ClubGrid() {
  const { data: clubs, isLoading, error } = useClubs();

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
        {Array.from({ length: 6 }).map((_, i) => (
          <div key={i} className="h-44 animate-pulse rounded-2xl bg-surface-container-lowest shadow-sm" />
        ))}
      </div>
    );
  }

  if (error) {
    return (
      <div className="rounded-2xl border border-error/30 bg-error-container p-6 text-center text-sm text-on-error-container">
        Failed to load clubs: {error.message}
      </div>
    );
  }

  if (!clubs || clubs.length === 0) {
    return (
      <div className="rounded-2xl bg-surface-container-lowest p-12 text-center text-sm text-on-surface-variant shadow-sm">
        No clubs found
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
      {clubs.map((club) => (
        <ClubCard key={club.id} club={club} />
      ))}
    </div>
  );
}
