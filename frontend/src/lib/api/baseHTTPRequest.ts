class BaseHTTPRequest {
    private API_BASE_URL: string;

    constructor() {
        //!Hardcoding it here as well so to find the base url.
        this.API_BASE_URL = process.env.NEXT_PUBLIC_GO_BACKEND_BASE_API_URL || 'http://localhost:8080';
    }

    private getAuthHeaders(): Record<string, string> {
      const token = localStorage.getItem('google_token');
      const headers: Record<string, string> = {
        'Content-Type': 'application/json',
      };
      
      if (token) {
        headers['Authorization'] = `Bearer ${token}`;
      }
      
      return headers;
    }
  
    public async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
      const url = `${this.API_BASE_URL}/api${endpoint}`;
      
      const response = await fetch(url, {
        headers: {
          ...this.getAuthHeaders(),
          ...options.headers,
        },
        ...options,
      });
  
      if (!response.ok) {
        const errorData = await response.json().catch(() => ({ error: 'Unknown error' }));
        throw new Error(errorData.error || `HTTP ${response.status}`);
      }  
      return response.json();
    }
}

export default BaseHTTPRequest;