import type { Club } from '../../../types/club';

interface ClubCardProps {
  club: Club;
}

const SPORT_COLORS: Record<string, string> = {
  football: 'bg-primary/10 text-primary',
  basketball: 'bg-secondary-container text-on-secondary-container',
  tennis: 'bg-primary-container/30 text-on-primary-container',
  baseball: 'bg-tertiary-container/30 text-on-tertiary-container',
  hockey: 'bg-surface-variant text-on-surface-variant',
  rugby: 'bg-error-container/30 text-on-error-container',
  cricket: 'bg-surface-container-high text-on-surface-variant',
  volleyball: 'bg-secondary-container/30 text-on-secondary-container',
};

const INITIALS_COLORS: Record<string, string> = {
  football: 'bg-blue-100 text-blue-700',
  basketball: 'bg-amber-100 text-amber-700',
  tennis: 'bg-green-100 text-green-700',
  volleyball: 'bg-yellow-100 text-yellow-700',
  default: 'bg-surface-variant text-on-surface-variant',
};

function getInitials(name: string): string {
  return name
    .split(' ')
    .map((w) => w[0])
    .filter(Boolean)
    .slice(0, 2)
    .join('')
    .toUpperCase();
}

export function ClubCard({ club }: ClubCardProps) {
  const initialsColor = INITIALS_COLORS[club.sport.toLowerCase()] ?? INITIALS_COLORS.default;
  const sportStyle = SPORT_COLORS[club.sport.toLowerCase()] ?? 'bg-surface-container text-on-surface-variant';

  return (
    <div className="rounded-2xl bg-surface-container-lowest p-6 shadow-sm hover:shadow-md transition-shadow border border-outline-variant/10">
      <div className="mb-4 flex items-center gap-4">
        {club.logoUrl ? (
          <img src={club.logoUrl} alt={club.name} className="w-12 h-12 rounded-xl object-cover" />
        ) : (
          <div className={`w-12 h-12 rounded-xl flex items-center justify-center font-black text-lg ${initialsColor}`}>
            {getInitials(club.name)}
          </div>
        )}
        <div className="min-w-0 flex-1">
          <h3 className="truncate text-base font-bold text-on-surface">
            {club.name}
          </h3>
          <p className="truncate text-xs text-on-surface-variant font-medium">{club.league}</p>
        </div>
      </div>

      <div className="flex items-center gap-2 text-xs">
        <span className="text-on-surface-variant font-medium">{club.country}</span>
        <span className="text-outline-variant">|</span>
        <span className={`rounded-full px-2.5 py-0.5 text-[10px] font-black tracking-widest uppercase ${sportStyle}`}>
          {club.sport}
        </span>
      </div>

      <div className="mt-4">
        {club.isActive ? (
          <span className="inline-flex items-center rounded-full bg-primary/10 px-2.5 py-0.5 text-[10px] font-bold uppercase tracking-widest text-primary">
            Active
          </span>
        ) : (
          <span className="inline-flex items-center rounded-full bg-surface-container-high px-2.5 py-0.5 text-[10px] font-bold uppercase tracking-widest text-outline">
            Inactive
          </span>
        )}
      </div>
    </div>
  );
}
