export interface Club {
  id: string;
  name: string;
  slug: string;
  country: string;
  league: string;
  logoUrl: string;
  sport: string;
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
}

export type CreateClubPayload = Omit<Club, 'id' | 'createdAt' | 'updatedAt'>;
export type UpdateClubPayload = Partial<CreateClubPayload>;
