export type StreamType = 'live' | 'vod' | 'highlight' | 'behind_the_scenes';
export type StreamStatus = 'scheduled' | 'live' | 'ended' | 'archived';

export interface Stream {
  id: string;
  clubId: string;
  title: string;
  description: string;
  type: StreamType;
  status: StreamStatus;
  streamUrl: string;
  thumbnailUrl: string;
  viewCount: number;
  duration: number;
  scheduledAt: string;
  startedAt: string;
  endedAt: string;
  tags: string[];
  createdAt: string;
  updatedAt: string;
}

export type CreateStreamPayload = Omit<Stream, 'id' | 'createdAt' | 'updatedAt' | 'viewCount'>;
export type UpdateStreamStatusPayload = { status: StreamStatus };
