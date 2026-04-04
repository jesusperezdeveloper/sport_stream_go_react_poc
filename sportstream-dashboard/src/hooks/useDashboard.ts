import { useQuery } from '@tanstack/react-query';
import { fetchAPI } from '../api/client';
import type { DashboardSummary } from '../types/event';

export function useDashboard() {
  return useQuery({
    queryKey: ['dashboard'],
    queryFn: () => fetchAPI<DashboardSummary>('/api/v1/dashboard/summary'),
    refetchInterval: 30_000,
  });
}
