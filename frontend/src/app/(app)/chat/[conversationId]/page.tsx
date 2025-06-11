'use client';

import { useState, useEffect } from 'react';
import {
  Box,
  Paper,
  useTheme,
  Alert,
} from '@mui/material';
import { useParams} from 'next/navigation';
import { chatAPI, Message, User } from '@/lib/api';
import { useRequireAuth} from '@/contexts/AuthContext';
import ChatInputBar from '@/components/ChatInputBar';
import Chatbox from '@/components/Chatbox';
import { CreateMessageRequest, ListMessagesResponse, MessageType } from '@/lib/types/generated';

export default function ChatPageContent() {

  console.log("ChatPageContent from page rendered.");

  const theme = useTheme();  
  const params = useParams();
  const conversationId: number = parseInt(params.conversationId as string);
  const { user, isAuthenticated } = useRequireAuth();
  
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputMessage, setInputMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [apiError, setApiError] = useState<string | null>(null);

  // Load conversation and messages when conversationId changes
  useEffect(() => {
    if (conversationId && user?.id) {
      loadConversationData(conversationId);
    }
  }, [conversationId, user]);

  const loadConversationData = async (convId: number) => {
    if (!user?.id) {
      setApiError('Authentication required');
      return;
    }

    try {
      setApiError(null);
      setIsLoading(true);
      
      // Load conversation details
     
      // Load messages
      const apiMessages: ListMessagesResponse = await chatAPI.getMessages(user.id, convId, 100, 0);      
      setMessages(apiMessages.messages);
    } catch (err) {
      setApiError('Failed to load conversation');
      console.error('Error loading conversation:', err);
    } finally {
      setIsLoading(false);
    }
  };

  // Start polling after sending message
  const pollForAIResponse = async (user: User) => {
    let retries = 0;
    const maxRetries = 10;
    const interval = 1000;

    while (retries < maxRetries) {
      const updatedMessages: ListMessagesResponse = await chatAPI.getMessages(user.id, conversationId, 100, 0);

      //!I want to append the new messages to the existing messages.
      setMessages(updatedMessages.messages);
      return;
      

      retries++;
      await new Promise(resolve => setTimeout(resolve, interval));
    }

    console.warn('AI response polling timed out');
  };



  const handleSendMessage = async () => {
    if (!inputMessage.trim() || isLoading || !user?.id) return;

    const userMessage: CreateMessageRequest = {
      langchain_message: inputMessage.trim(),
      message_type: 'user',
    };

    setInputMessage('');
    setIsLoading(true);
    setApiError(null);

    try {
      // Send message to backend
      const outputMessage: ListMessagesResponse = await chatAPI.createMessage(user.id, conversationId, userMessage);
      setMessages(outputMessage.messages);


      // The backend will automatically generate an AI response
      // We need to poll for new messages or implement WebSocket
      // For now, let's poll after a short delay
      //await pollForAIResponse(user);
      setIsLoading(false);

    } catch (error) {
      console.error('Error sending message:', error);
      setApiError('Failed to send message. Please try again.');
      const errorMessage: Message = {
        id: 0,
        conversation_id: conversationId,
        user_id: user.id,
        langchain_message: "I'm sorry, I'm having trouble responding right now. Please try again in a moment.",
        message_type: MessageType.Assistant,
      };
      setMessages(prev => [...prev, errorMessage]);
      setIsLoading(false);
    }
  };





  
  if (!isAuthenticated || !user) {
    return (
      <Box sx={{ 
        height: 'calc(100vh - 128px)',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
      }}>
        <Alert severity="warning">
          Please sign in to access the chat feature.
        </Alert>
      </Box>
    );
  }

  return (
    <Box sx={{ 
      height: 'calc(100vh - 128px)', // Account for header + footer
      display: 'flex', 
      flexDirection: 'column',
      width: '100%',
      bgcolor: 'background.default',
      overflow: 'hidden'
    }}>
      {/* Error Alert */}
      {apiError && (
        <Alert severity="error" sx={{ m: 2 }} onClose={() => setApiError(null)}>
          {apiError}
        </Alert>
      )}

      {/* Messages Container */}
      <Chatbox messages={messages} isLoading={isLoading} />
      <Paper 
        elevation={0} 
        sx={{ 
          borderRadius: 0,
          bgcolor: 'background.paper',
          borderTop: `1px solid ${theme.palette.divider}`,
          flexShrink: 0
        }}
      >
          <ChatInputBar 
            inputMessage={inputMessage} 
            setInputMessage={setInputMessage} 
            handleSendMessage={handleSendMessage} 
            isLoading={isLoading} />
      </Paper>
    </Box>
  );
}