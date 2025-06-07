'use client';

import { useState, useEffect } from 'react';
import {
  Box,
  Paper,
  useTheme,
  Alert,
} from '@mui/material';
import { useParams} from 'next/navigation';
import { chatAPI, Message } from '@/lib/api';
import { useRequireAuth} from '@/contexts/AuthContext';
import ChatInputBar from '@/components/ChatInputBar';
import Chatbox from '@/components/Chatbox';

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
      const apiMessages = await chatAPI.getMessages(user.id, convId, 100, 0);      
      setMessages(apiMessages);
    } catch (err) {
      setApiError('Failed to load conversation');
      console.error('Error loading conversation:', err);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSendMessage = async () => {
    if (!inputMessage.trim() || isLoading || !user?.id) return;

    const userMessage: Message = {
      id: 0,
      content: inputMessage.trim(),
      message_type: 'user',
      created_at: new Date().toISOString(),
      conversation_id: conversationId,
      user_id: user.id,
    };

    setMessages(prev => [...prev, userMessage]);
    setInputMessage('');
    setIsLoading(true);
    setApiError(null);

    try {
      // Send message to backend
      await chatAPI.createMessage(user.id, conversationId, {
        content: userMessage.content,
        message_type: 'user',
      });

      // The backend will automatically generate an AI response
      // We need to poll for new messages or implement WebSocket
      // For now, let's poll after a short delay
      setTimeout(async () => {
        try {
          const updatedMessages = await chatAPI.getMessages(user.id, conversationId, 100, 0);
          setMessages(updatedMessages);
        } catch (err) {
          console.error('Error fetching updated messages:', err);
        } finally {
          setIsLoading(false);
        }
      }, 2000); // Wait 2 seconds for AI response

    } catch (error) {
      console.error('Error sending message:', error);
      setApiError('Failed to send message. Please try again.');
      const errorMessage: Message = {
        id: 0,
        content: "I'm sorry, I'm having trouble responding right now. Please try again in a moment.",
        message_type: 'assistant',
        created_at: new Date().toISOString(),
        conversation_id: conversationId,
        user_id: user.id,
      };
      setMessages(prev => [...prev, errorMessage]);
      setIsLoading(false);
    }
  };





  // Show message if user is not authenticated
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