import { Search, Bell, User, Zap } from 'lucide-react';

interface HeaderProps {
  title: string;
}

export function Header({ title: _title }: HeaderProps) {
  return (
    <header className="sticky top-0 z-40 flex h-16 items-center justify-between bg-white/80 backdrop-blur-xl px-4 md:px-8">
      {/* Mobile: logo + title, Desktop: tabs */}
      <div className="flex items-center gap-4 md:gap-8">
        {/* Mobile logo - visible only on small screens */}
        <div className="flex items-center gap-2 md:hidden">
          <Zap className="text-primary" size={22} fill="currentColor" />
          <span className="font-black text-sm tracking-tighter italic text-primary">
            SportStream
          </span>
        </div>

        {/* Desktop tabs */}
        <div className="hidden md:flex gap-6">
          <a href="#" className="text-primary font-bold border-b-2 border-primary py-5 transition-all text-sm">
            Global Feed
          </a>
          <a href="#" className="text-on-surface-variant font-medium py-5 hover:text-primary transition-all text-sm">
            Network Status
          </a>
        </div>
      </div>

      <div className="flex items-center gap-2 md:gap-4">
        {/* Search bar - hidden on mobile */}
        <div className="relative hidden md:block">
          <Search size={20} className="absolute left-3 top-1/2 -translate-y-1/2 text-on-surface-variant" />
          <input
            className="bg-surface-container-lowest border border-outline-variant/30 rounded-xl pl-10 pr-4 py-2 w-64 text-sm focus:ring-2 focus:ring-primary/20 focus:outline-none"
            placeholder="Search streams, clubs..."
            type="text"
          />
        </div>

        <button className="p-2 hover:bg-surface-container-high rounded-full transition-all relative min-w-[44px] min-h-[44px] flex items-center justify-center">
          <Bell size={20} className="text-on-surface-variant" />
          <span className="absolute top-2 right-2 w-2 h-2 bg-error rounded-full border-2 border-white" />
        </button>

        <div className="hidden md:block h-8 w-px bg-outline-variant/30 mx-2" />

        <button className="hidden md:flex items-center gap-2 px-3 py-1.5 hover:bg-surface-container-high rounded-full transition-all">
          <User size={20} className="text-on-surface-variant" />
          <span className="text-sm font-semibold text-on-surface">Admin Portal</span>
        </button>
      </div>
    </header>
  );
}
