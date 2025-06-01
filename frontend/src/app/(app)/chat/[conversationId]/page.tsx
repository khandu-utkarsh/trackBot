'use client';

import { useState, useEffect } from 'react';
import {
  Box,
  Paper,
  Typography,
  useTheme,
  Alert,
} from '@mui/material';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import { useParams} from 'next/navigation';
import { chatAPI } from '@/lib/api/chat';
import { Conversation, Message } from '@/lib/types/chat';
import { useRequireAuth} from '@/contexts/AuthContext';
import ChatInputBar from '@/components/ChatInputBar';
import Chatbox from '@/components/Chatbox';

export default function ChatPageContent() {

  console.log("ChatPageContent from page rendered.");

  const theme = useTheme();  
  const params = useParams();
  const conversationId: number = parseInt(params.conversationId as string);
  const { user, token, isAuthenticated } = useRequireAuth();
  
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputMessage, setInputMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [apiError, setApiError] = useState<string | null>(null);

  //!Hardcoding the user id for now.
  const userId = 2;

  // Load conversation and messages when conversationId changes
  useEffect(() => {
    if (conversationId && userId) {
      loadConversationData(conversationId);
    }
  }, [conversationId, userId]);

  const loadConversationData = async (convId: number) => {
    if (!userId || !token) {
      setApiError('Authentication required');
      return;
    }

    try {
      setApiError(null);
      setIsLoading(true);
      
      // Load conversation details
      const userId = 2;
     
      // Load messages
      const apiMessages = await chatAPI.getMessages(userId, convId);      
      setMessages(apiMessages);
    } catch (err) {
      setApiError('Failed to load conversation');
      console.error('Error loading conversation:', err);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSendMessage = async () => {
    if (!inputMessage.trim() || isLoading || !userId || !token) return;

    const userMessage: Message = {
      id: 0,
      content: inputMessage.trim(),
      message_type: 'user',
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
      conversation_id: conversationId,
      user_id: userId,
    };

    setMessages(prev => [...prev, userMessage]);
    setInputMessage('');
    setIsLoading(true);
    setApiError(null);

    try {
      // Send message to backend
      const userId = 2;
      await chatAPI.createMessage(userId, conversationId, {
        content: userMessage.content,
        message_type: 'user',
      });

      // The backend will automatically generate an AI response
      // We need to poll for new messages or implement WebSocket
      // For now, let's poll after a short delay
      setTimeout(async () => {
        try {
          const userId = 2;
          const updatedMessages = await chatAPI.getMessages(userId, conversationId);
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
        updated_at: new Date().toISOString(),
        conversation_id: conversationId,
        user_id: userId,
      };
      setMessages(prev => [...prev, errorMessage]);
      setIsLoading(false);
    }
  };





  // Show message if user is not authenticated
  if (!isAuthenticated || !user || !userId) {
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