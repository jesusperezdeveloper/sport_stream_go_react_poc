import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { fetchAPI, mutateAPI } from '../api/client';
import type { Stream, StreamStatus, StreamType, CreateStreamPayload, UpdateStreamStatusPayload } from '../types/stream';

interface StreamFilters {
  status?: StreamStatus;
  type?: StreamType;
}

function buildStreamQuery(filters?: StreamFilters): string {
  const params = new URLSearchParams();
  if (filters?.status) params.set('status', filters.status);
  if (filters?.type) params.set('type', filters.type);
  const qs = params.toString();
  return qs ? `?${qs}` : '';
}

export function useStreams(filters?: StreamFilters) {
  return useQuery({
    queryKey: ['streams', filters],
    queryFn: () => fetchAPI<Stream[]>(`/api/v1/streams${buildStreamQuery(filters)}`),
  });
}

export function useStream(id: string) {
  return useQuery({
    queryKey: ['streams', id],
    queryFn: () => fetchAPI<Stream>(`/api/v1/streams/${id}`),
    enabled: !!id,
  });
}

export function useCreateStream() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateStreamPayload) =>
      mutateAPI<Stream>('/api/v1/streams', 'POST', payload),
    onSuccess: () => {
      void qc.invalidateQueries({ queryKey: ['streams'] });
    },
  });
}

export function useUpdateStreamStatus(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: UpdateStreamStatusPayload) =>
      mutateAPI<Stream>(`/api/v1/streams/${id}/status`, 'PATCH', payload),
    onSuccess: () => {
      void qc.invalidateQueries({ queryKey: ['streams'] });
      void qc.invalidateQueries({ queryKey: ['dashboard'] });
    },
  });
}
