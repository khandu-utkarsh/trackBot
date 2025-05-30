//!Try not to use any hardcoded values. 

// API service for chat functionality
const API_BASE_URL = process.env.NEXT_PUBLIC_GO_BACKEND_BASE_API_URL

export interface User {
	id?: number;
	email: string;  //!Right now, only dealing with the email address.
}


class UserAPI {

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
    

    async createUser(user: User): Promise<User> {
        const response = await fetch(`${API_BASE_URL}/api/users`, {
            method: 'POST',
            headers: this.getAuthHeaders(),
            body: JSON.stringify(user),
        });
        console.log("User created: ", response);
        return response.json();
    }   
}

export const userAPI = new UserAPI(); 