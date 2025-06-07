import axios from 'axios';
import { usersApi } from './client';
import { User, GoogleLoginRequest } from '@/lib/types/generated';

class UserAPI {
    // HTTP client method using axios for consistent error handling
    private async httpRequest<T>(url: string, method: 'GET' | 'POST' | 'PUT' | 'DELETE' = 'GET', data?: any): Promise<T> {
        try {
            const response = await axios({
                url,
                method,
                data,
                withCredentials: true,
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            return response.data;
        } catch (error: any) {
            const errorMessage = error.response?.data?.error || 
                               error.response?.data?.message || 
                               error.message || 
                               'An unknown error occurred';
            throw new Error(errorMessage);
        }
    }

    // User methods (using generated API clients)
    async createUser(input: GoogleLoginRequest): Promise<User> {
        const response = await usersApi.googleLogin(input);
        return response.data;
    }

    // Authentication methods (calling backend directly)
    async getUser(): Promise<User | null> {
        try {
            const response = await usersApi.getCurrentUser();
            return response.data;
        } catch (error) {
            console.error('Failed to get user:', error);
            return null;
        }
    }

    async logout(): Promise<void> {
        try {
            const backendUrl = process.env.NEXT_PUBLIC_GO_BACKEND_BASE_API_URL || 'http://localhost:8080';
            await this.httpRequest<void>(`${backendUrl}/api/auth/logout`, 'POST');
        } catch (error) {
            console.error('Logout error:', error);
            throw error;
        }
    }

    async login(redirectUrl?: string): Promise<void> {
        // Redirect to Google OAuth login
        const params = new URLSearchParams();
        if (redirectUrl) {
            params.append('redirect', redirectUrl);
        }
        
        window.location.href = `/api/auth/login?${params.toString()}`;
    }
}

export const userAPI = new UserAPI(); 