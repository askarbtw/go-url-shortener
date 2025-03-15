import axios from 'axios';
import { CreateURLRequest, UpdateURLRequest, URL, URLStats } from '../types/url';

// When using the proxy in development, we use relative URLs
// In production, the full URL should be used
const IS_DEV = true; // Force development mode for testing
const API_BASE_URL = IS_DEV ? '' : 'http://localhost:8080';

// Create axios instance
const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// URL APIs
export const urlAPI = {
  // Create a new short URL
  createURL: async (data: CreateURLRequest): Promise<URL> => {
    const response = await api.post<URL>('/shorten', data);
    return response.data;
  },

  // Get a URL by short code
  getURL: async (shortCode: string): Promise<URL> => {
    const response = await api.get<URL>(`/shorten/${shortCode}`);
    return response.data;
  },

  // Update an existing URL
  updateURL: async (shortCode: string, data: UpdateURLRequest): Promise<URL> => {
    const response = await api.put<URL>(`/shorten/${shortCode}`, data);
    return response.data;
  },

  // Delete a URL
  deleteURL: async (shortCode: string): Promise<void> => {
    await api.delete(`/shorten/${shortCode}`);
  },

  // Get URL statistics
  getURLStats: async (shortCode: string): Promise<URLStats> => {
    const response = await api.get<URLStats>(`/shorten/${shortCode}/stats`);
    return response.data;
  },
  
  // Get all URLs with stats
  getAllURLs: async (): Promise<URLStats[]> => {
    try {
      const response = await api.get<URLStats[]>('/shorten');
      return response.data;
    } catch (error) {
      console.error('Error fetching URLs:', error);
      return [];
    }
  },
};

// Generate the full redirect URL for a short code
export const getRedirectURL = (shortCode: string): string => {
  return `${API_BASE_URL}/r/${shortCode}`;
};

export default api; 