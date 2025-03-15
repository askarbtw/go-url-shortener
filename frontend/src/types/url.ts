export interface URL {
  id: string;
  url: string;
  shortCode: string;
  createdAt: string;
  updatedAt: string;
}

export interface URLStats extends URL {
  accessCount: number;
}

export interface CreateURLRequest {
  url: string;
}

export interface UpdateURLRequest {
  url: string;
}

export interface APIResponse<T> {
  data: T;
  error?: string;
} 