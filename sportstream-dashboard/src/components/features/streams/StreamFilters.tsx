import { ChevronDown } from 'lucide-react';
import type { StreamStatus, StreamType } from '../../../types/stream';

interface StreamFiltersProps {
  status: StreamStatus | '';
  type: StreamType | '';
  onStatusChange: (status: StreamStatus | '') => void;
  onTypeChange: (type: StreamType | '') => void;
}

const statuses: StreamStatus[] = ['live', 'scheduled', 'ended', 'archived'];
const types: StreamType[] = ['live', 'vod', 'highlight', 'behind_the_scenes'];

export function StreamFilters({
  status,
  type,
  onStatusChange,
  onTypeChange,
}: StreamFiltersProps) {
  return (
    <div className="flex items-center gap-2">
      <div className="relative">
        <select
          value={status}
          onChange={(e) => onStatusChange(e.target.value as StreamStatus | '')}
          className="appearance-none bg-surface-container-lowest border border-outline-variant/30 rounded-full px-4 py-2 pr-8 text-sm font-medium hover:bg-surface-container-low transition-all cursor-pointer focus:outline-none focus:ring-2 focus:ring-primary/20"
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

      <div className="relative">
        <select
          value={type}
          onChange={(e) => onTypeChange(e.target.value as StreamType | '')}
          className="appearance-none bg-surface-container-lowest border border-outline-variant/30 rounded-full px-4 py-2 pr-8 text-sm font-medium hover:bg-surface-container-low transition-all cursor-pointer focus:outline-none focus:ring-2 focus:ring-primary/20"
        >
          <option value="">All Types</option>
          {types.map((t) => (
            <option key={t} value={t}>
              {t.toUpperCase().replace('_', ' ')}
            </option>
          ))}
        </select>
        <ChevronDown size={16} className="absolute right-3 top-1/2 -translate-y-1/2 text-on-surface-variant pointer-events-none" />
      </div>
    </div>
  );
}
