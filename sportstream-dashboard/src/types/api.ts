export interface ApiResponse<T> {
  data: T;
  meta: {
    total: number;
    timestamp: string;
  };
}

export interface ApiError {
  error: {
    code: string;
    message: string;
  };
}
