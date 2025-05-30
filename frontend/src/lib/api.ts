// API utility functions that use authentication context

interface ApiResponse<T = any> {
  data?: T;
  error?: string;
  status: number;
}

class ApiClient {
  private baseURL: string;

  constructor(baseURL: string = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001') {
    this.baseURL = baseURL;
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {},
    token?: string | null
  ): Promise<ApiResponse<T>> {
    const url = `${this.baseURL}${endpoint}`;
    
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
      ...(options.headers as Record<string, string>),
    };

    // Add authorization header if token is provided
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }

    try {
      const response = await fetch(url, {
        ...options,
        headers,
      });

      const data = await response.json();

      return {
        data: response.ok ? data : undefined,
        error: response.ok ? undefined : data.message || 'An error occurred',
        status: response.status,
      };
    } catch (error) {
      return {
        error: error instanceof Error ? error.message : 'Network error',
        status: 0,
      };
    }
  }

  // Public endpoints (no auth required)
  async get<T>(endpoint: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'GET' });
  }

  async post<T>(endpoint: string, data: any): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  // Authenticated endpoints
  async authenticatedGet<T>(endpoint: string, token: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'GET' }, token);
  }

  async authenticatedPost<T>(endpoint: string, data: any, token: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, {
      method: 'POST',
      body: JSON.stringify(data),
    }, token);
  }

  async authenticatedPut<T>(endpoint: string, data: any, token: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, {
      method: 'PUT',
      body: JSON.stringify(data),
    }, token);
  }

  async authenticatedDelete<T>(endpoint: string, token: string): Promise<ApiResponse<T>> {
    return this.request<T>(endpoint, { method: 'DELETE' }, token);
  }
}

// Create a singleton instance
export const apiClient = new ApiClient();

// Example usage with authentication context:
/*
import { useAuth } from '../contexts/AuthContext';

function SomeComponent() {
  const { token } = useAuth();
  
  const fetchUserData = async () => {
    if (token) {
      const response = await apiClient.authenticatedGet('/api/user/profile', token);
      if (response.data) {
        console.log('User data:', response.data);
      } else {
        console.error('Error:', response.error);
      }
    }
  };

  const createWorkout = async (workoutData: any) => {
    if (token) {
      const response = await apiClient.authenticatedPost('/api/workouts', workoutData, token);
      if (response.data) {
        console.log('Workout created:', response.data);
      } else {
        console.error('Error:', response.error);
      }
    }
  };

  return (
    // Your component JSX
  );
}
*/ 