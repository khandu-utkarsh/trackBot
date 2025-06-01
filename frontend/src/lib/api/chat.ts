import { Conversation, Message } from "@/lib/types/chat";
import BaseHTTPRequest from "./baseHTTPRequest";


export interface CreateConversationRequest {
  title: string;
}

export interface CreateMessageRequest {
  content: string;
  message_type: 'user' | 'assistant' | 'system';
}

class ChatAPI extends BaseHTTPRequest {
  //!Conversation methods
  async getConversations(userId: number): Promise<Conversation[]> {
    return this.request<Conversation[]>(`/users/${userId}/conversations`);
  }

  async createConversation(userId: number, data: CreateConversationRequest): Promise<Conversation> {
    return this.request<Conversation>(`/users/${userId}/conversations`, {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getConversation(userId: number, conversationId: number): Promise<Conversation> {
    return this.request<Conversation>(`/users/${userId}/conversations/${conversationId}`);
  }

  async updateConversation(userId: number, conversationId: number, data: Partial<CreateConversationRequest>): Promise<Conversation> {
    return this.request<Conversation>(`/users/${userId}/conversations/${conversationId}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteConversation(userId: number, conversationId: number): Promise<void> {
    return this.request<void>(`/users/${userId}/conversations/${conversationId}`, {
      method: 'DELETE',
    });
  }

  //!Message methods
  async getMessages(userId: number, conversationId: number): Promise<Message[]> {
    return this.request<Message[]>(`/users/${userId}/conversations/${conversationId}/messages`);
  }

  async createMessage(userId: number, conversationId: number, data: CreateMessageRequest): Promise<Message> {
    return this.request<Message>(`/users/${userId}/conversations/${conversationId}/messages`, {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getMessage(userId: number, conversationId: number, messageId: number): Promise<Message> {
    return this.request<Message>(`/users/${userId}/conversations/${conversationId}/messages/${messageId}`);
  }

  async updateMessage(userId: number, conversationId: number, messageId: number, data: Partial<CreateMessageRequest>): Promise<Message> {
    return this.request<Message>(`/users/${userId}/conversations/${conversationId}/messages/${messageId}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  async deleteMessage(userId: number, conversationId: number, messageId: number): Promise<void> {
    return this.request<void>(`/users/${userId}/conversations/${conversationId}/messages/${messageId}`, {
      method: 'DELETE',
    });
  }
}

export const chatAPI = new ChatAPI();