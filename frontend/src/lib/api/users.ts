import { User } from "@/lib/types/users";
import BaseHTTPRequest from "./baseHTTPRequest";

class UserAPI extends BaseHTTPRequest {
    //!User methods
    async createUser(googleToken: string): Promise<User> {
        return this.request<User>(`/auth/google`, {
            method: 'POST',
            body: JSON.stringify({ googleToken }),
        });
    }


    async getUser(): Promise<User> {

        console.log("getUser");
        const userOut = await this.request<User>(`/auth/me`, {
            method: 'GET',
        });
        console.log("userOut", userOut);
        return userOut;        
    }


    async logout(): Promise<void> {
        return this.request<void>(`/auth/logout`, {
            method: 'POST',
        });
    }
}

export const userAPI = new UserAPI(); 