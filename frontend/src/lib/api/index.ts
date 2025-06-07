// Export all API instances
export { chatAPI } from './chat';
export { userAPI } from './users';

// Export generated types directly (only conversation, message, and user related)
export type {
  User,
  Configuration,
  Conversation,
  Message,
  GoogleLoginRequest,
  CreateConversationRequest,
  CreateMessageRequest,
  UpdateConversationRequest,
  DeleteConversationRequest,
  MessageType,
  ListConversationsResponse,
  ListMessagesResponse,
  ModelError
} from '@/lib/types/generated';