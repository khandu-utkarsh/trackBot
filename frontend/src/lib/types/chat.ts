

//!Types for the chat API.

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
    updated_at: string;
    last_message?: string;
}
  