import { User } from "@/lib/types/users";
import BaseHTTPRequest from "./baseHTTPRequest";

class UserAPI extends BaseHTTPRequest {
    //!User methods
    async createUser(user: User): Promise<User> {
        return this.request<User>(`/users`, {
            method: 'POST',
            body: JSON.stringify(user),
        });
    }   
}

export const userAPI = new UserAPI(); 