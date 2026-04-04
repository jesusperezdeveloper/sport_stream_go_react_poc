import { Outlet } from 'react-router-dom';
import { Sidebar } from './Sidebar';

export function Layout() {
  return (
    <div className="flex min-h-screen bg-background">
      <Sidebar />
      <main className="flex-1 min-h-screen lg:ml-64 pb-20 lg:pb-0">
        <Outlet />
      </main>
    </div>
  );
}
