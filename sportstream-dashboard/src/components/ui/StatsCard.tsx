import type { LucideIcon } from 'lucide-react';

interface StatsCardProps {
  icon: LucideIcon;
  iconColor: string;
  borderColor: string;
  label: string;
  value: number | string;
  extraLabel?: string;
  isLive?: boolean;
}

export function StatsCard({
  icon: Icon,
  iconColor,
  borderColor,
  label,
  value,
  extraLabel,
  isLive,
}: StatsCardProps) {
  return (
    <div
      className="bg-surface-container-lowest rounded-2xl p-4 md:p-6 shadow-sm hover:shadow-md transition-shadow"
      style={{ borderLeft: `4px solid ${borderColor}` }}
    >
      <div className="flex justify-between items-start mb-2 md:mb-4">
        {isLive ? (
          <div className="flex items-center gap-2">
            <div className="w-3 h-3 rounded-full kinetic-pulse" style={{ backgroundColor: borderColor }} />
            <span
              className="px-2 py-0.5 text-white text-[10px] font-black rounded-sm"
              style={{ backgroundColor: borderColor }}
            >
              LIVE
            </span>
          </div>
        ) : (
          <div
            className="p-2 md:p-3 rounded-xl"
            style={{ backgroundColor: `${iconColor}18` }}
          >
            <Icon size={18} className="md:w-[22px] md:h-[22px]" style={{ color: iconColor }} />
          </div>
        )}
        {extraLabel && (
          <span className="text-[10px] font-bold uppercase tracking-widest text-on-surface-variant">
            {extraLabel}
          </span>
        )}
      </div>

      <div className="text-2xl md:text-4xl font-black tracking-tight text-on-surface mb-1">
        {value}
      </div>
      <div className="text-[10px] md:text-xs font-semibold uppercase tracking-widest text-on-surface-variant">
        {label}
      </div>
    </div>
  );
}
