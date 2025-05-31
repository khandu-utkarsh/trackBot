// API service for chat functionality
const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

export interface Message {
  id: number;
  conversation_id: number;
  user_id: number;
  content: string;
  message_type: 'user' | 'assistant' | 'system';
  created_at: string;
  updated_at: string;
}

export interface Conversation {
  id: number;
  user_id: number;
  title: string;
  is_active: boolean;
  updated_at: string;
  last_message?: string;
}

export interface CreateConversationRequest {
  title: string;
  is_active?: boolean;
}

export interface CreateMessageRequest {
  content: string;
  message_type: 'user' | 'assistant' | 'system';
}

class ChatAPI {
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

  private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    const url = `${API_BASE_URL}/api${endpoint}`;
    
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

  // Conversation methods
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

  // Message methods
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