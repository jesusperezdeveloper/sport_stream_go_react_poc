import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { fetchAPI, mutateAPI } from '../api/client';
import type { Club, CreateClubPayload, UpdateClubPayload } from '../types/club';

export function useClubs() {
  return useQuery({
    queryKey: ['clubs'],
    queryFn: () => fetchAPI<Club[]>('/api/v1/clubs'),
  });
}

export function useClub(id: string) {
  return useQuery({
    queryKey: ['clubs', id],
    queryFn: () => fetchAPI<Club>(`/api/v1/clubs/${id}`),
    enabled: !!id,
  });
}

export function useCreateClub() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateClubPayload) =>
      mutateAPI<Club>('/api/v1/clubs', 'POST', payload),
    onSuccess: () => {
      void qc.invalidateQueries({ queryKey: ['clubs'] });
    },
  });
}

export function useUpdateClub(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: UpdateClubPayload) =>
      mutateAPI<Club>(`/api/v1/clubs/${id}`, 'PUT', payload),
    onSuccess: () => {
      void qc.invalidateQueries({ queryKey: ['clubs'] });
    },
  });
}

export function useDeleteClub() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      mutateAPI<void>(`/api/v1/clubs/${id}`, 'DELETE'),
    onSuccess: () => {
      void qc.invalidateQueries({ queryKey: ['clubs'] });
    },
  });
}
