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
}

export const userAPI = new UserAPI(); 