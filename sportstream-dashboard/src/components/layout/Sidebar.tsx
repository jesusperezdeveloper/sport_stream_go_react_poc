import { NavLink } from 'react-router-dom';
import { LayoutDashboard, Tv, Building2, Calendar, Zap, Settings, HelpCircle, PlusCircle } from 'lucide-react';

const navItems = [
  { to: '/', icon: LayoutDashboard, label: 'Dashboard' },
  { to: '/streams', icon: Tv, label: 'Streams' },
  { to: '/clubs', icon: Building2, label: 'Clubs' },
  { to: '/events', icon: Calendar, label: 'Events' },
] as const;

export function Sidebar() {
  return (
    <aside className="fixed left-0 top-0 z-50 flex h-screen w-64 flex-col bg-sidebar-bg">
      {/* Logo */}
      <div className="px-6 py-8">
        <div className="flex items-center gap-2 mb-10">
          <Zap className="text-sidebar-accent" size={28} fill="currentColor" />
          <div>
            <h1 className="text-sidebar-accent font-black text-xl tracking-tighter italic">
              SportStream
            </h1>
            <p className="text-white/40 text-[10px] uppercase tracking-widest font-bold">
              Management Console
            </p>
          </div>
        </div>

        {/* Navigation */}
        <nav className="space-y-1">
          {navItems.map(({ to, icon: Icon, label }) => (
            <NavLink
              key={to}
              to={to}
              end={to === '/'}
              className={({ isActive }) =>
                `flex items-center gap-3 py-3 px-6 transition-colors duration-200 ${
                  isActive
                    ? 'bg-sidebar-accent/10 text-sidebar-accent border-l-2 border-sidebar-accent'
                    : 'text-white/70 hover:text-white hover:bg-white/5 rounded-lg'
                }`
              }
            >
              <Icon size={20} />
              <span className="font-semibold uppercase tracking-widest text-[10px]">
                {label}
              </span>
            </NavLink>
          ))}
        </nav>
      </div>

      {/* Bottom section */}
      <div className="mt-auto p-6 space-y-4">
        <button className="w-full bg-gradient-to-br from-primary to-primary-container text-white py-3 rounded-full font-bold flex items-center justify-center gap-2 active:scale-95 transition-transform shadow-lg shadow-primary/20">
          <PlusCircle size={20} />
          Go Live
        </button>

        <div className="pt-6 border-t border-white/5">
          <a href="#" className="flex items-center gap-3 text-white/70 hover:text-white px-6 py-3 transition-colors duration-200">
            <Settings size={18} />
            <span className="font-semibold uppercase tracking-widest text-[10px]">Settings</span>
          </a>
          <a href="#" className="flex items-center gap-3 text-white/70 hover:text-white px-6 py-3 transition-colors duration-200">
            <HelpCircle size={18} />
            <span className="font-semibold uppercase tracking-widest text-[10px]">Support</span>
          </a>
          <div className="px-6 py-4 text-white/30 text-[10px] font-medium tracking-tighter">
            v0.1.0 ENGINE
          </div>
        </div>
      </div>
    </aside>
  );
}
