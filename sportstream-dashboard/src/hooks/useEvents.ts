import { useQuery } from '@tanstack/react-query';
import { fetchAPI } from '../api/client';
import type { SportEvent, EventStatus } from '../types/event';

interface EventFilters {
  sport?: string;
  status?: EventStatus;
}

function buildEventQuery(filters?: EventFilters): string {
  const params = new URLSearchParams();
  if (filters?.sport) params.set('sport', filters.sport);
  if (filters?.status) params.set('status', filters.status);
  const qs = params.toString();
  return qs ? `?${qs}` : '';
}

export function useEvents(filters?: EventFilters) {
  return useQuery({
    queryKey: ['events', filters],
    queryFn: () => fetchAPI<SportEvent[]>(`/api/v1/events${buildEventQuery(filters)}`),
  });
}

export function useEvent(id: string) {
  return useQuery({
    queryKey: ['events', id],
    queryFn: () => fetchAPI<SportEvent>(`/api/v1/events/${id}`),
    enabled: !!id,
  });
}

export function useUpcomingEvents(limit = 5) {
  return useQuery({
    queryKey: ['events', 'upcoming', limit],
    queryFn: () => fetchAPI<SportEvent[]>(`/api/v1/events?status=upcoming&limit=${limit}`),
  });
}
