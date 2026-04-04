export type EventStatus = 'upcoming' | 'live' | 'completed' | 'cancelled';

export interface SportEvent {
  id: string;
  clubId: string;
  title: string;
  description: string;
  venue: string;
  sport: string;
  startTime: string;
  endTime: string;
  status: EventStatus;
  streamId: string;
  createdAt: string;
  updatedAt: string;
}

export interface DashboardSummary {
  totalClubs: number;
  totalStreams: number;
  liveStreams: number;
  upcomingEvents: number;
  totalViews: number;
  streamsByType: Record<string, number>;
  streamsByStatus: Record<string, number>;
  topClubsByViews: Array<{ clubId: string; clubName: string; totalViews: number }>;
}
