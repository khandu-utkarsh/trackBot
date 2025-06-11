import { conversationsApi, messagesApi } from './client';
import { 
  Conversation, 
  Message, 
  CreateConversationRequest, 
  CreateConversationResponse,
  CreateMessageRequest,
  ListMessagesResponse,
  UpdateConversationRequest,
  DeleteConversationRequest,
  DeleteConversationResponse,
  ListConversationsResponse,
} from '@/lib/types/generated';

class ChatAPI {
  // Conversation methods
  async getConversations(userId: number): Promise<ListConversationsResponse> {
    const response = await conversationsApi.listConversations(userId);
    return response.data;
  }

  async createConversation(userId: number, data: CreateConversationRequest): Promise<CreateConversationResponse> {
    const response = await conversationsApi.createConversation(userId, data);
    return response.data;
  }

  async getConversation(userId: number, conversationId: number): Promise<Conversation> {
    const response = await conversationsApi.getConversationById(userId, conversationId);
    return response.data;
  }

  async updateConversation(userId: number, conversationId: number, data: UpdateConversationRequest): Promise<Conversation> {
    const response = await conversationsApi.updateConversation(userId, conversationId, data);
    return response.data;
  }

  async deleteConversation(userId: number, conversationId: number, confirm: boolean = true): Promise<DeleteConversationResponse> {
    const deleteRequest: DeleteConversationRequest = { confirm };
    const response = await conversationsApi.deleteConversation(userId, conversationId, deleteRequest);
    return response.data;
  }

  // Message methods
  async getMessages(userId: number, conversationId: number, limit?: number, offset?: number): Promise<ListMessagesResponse> {
    const response = await messagesApi.listMessages(userId, conversationId, limit, offset);
    return response.data;
  }

  async createMessage(userId: number, conversationId: number, data: CreateMessageRequest): Promise<ListMessagesResponse> {
    const response = await messagesApi.createMessage(userId, conversationId, data);
    return response.data;
  }

  async getMessage(userId: number, conversationId: number, messageId: number): Promise<Message> {
    const response = await messagesApi.getMessageById(userId, conversationId, messageId);
    return response.data;
  }

  async deleteMessage(userId: number, conversationId: number, messageId: number, confirm: boolean = true): Promise<void> {
    const options = { params: { confirm } };
    await messagesApi.deleteMessage(userId, conversationId, messageId, options);
  }
}

export const chatAPI = new ChatAPI();