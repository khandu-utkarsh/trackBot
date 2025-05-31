'use client';

import { useState, useRef, useEffect } from 'react';
import {
  Box,
  Paper,
  Typography,
  useTheme,
  Alert,
} from '@mui/material';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import { useParams, useRouter, useSearchParams } from 'next/navigation';
import { chatAPI, Message as APIMessage, Conversation } from '@/lib/api/chat';
import { useRequireAuth} from '@/contexts/AuthContext';
import ChatInputBar from '@/components/ChatInputBar';
import Chatbox from '@/components/Chatbox';
import { Message } from '@/components/ChatMessage';


export default function ChatPageContent() {

  console.log("ChatPageContent from page rendered.");

  const theme = useTheme();  
  const router = useRouter();
  const params = useParams();
  const conversationId: number = parseInt(params.conversationId as string);
  const { user, token, isAuthenticated } = useRequireAuth();
  
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputMessage, setInputMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [apiError, setApiError] = useState<string | null>(null);
  const [currentConversation, setCurrentConversation] = useState<Conversation | null>(null);

  const inputRef = useRef<HTMLInputElement>(null);

  // Get user ID from authenticated user
  const getUserId = (email: string): number => {
    // Simple hash function to convert email to a consistent numeric ID
    let hash = 0;
    for (let i = 0; i < email.length; i++) {
      const char = email.charCodeAt(i);
      hash = ((hash << 5) - hash) + char;
      hash = hash & hash; // Convert to 32-bit integer
    }
    return Math.abs(hash);
  };

  const userId = user?.email ? getUserId(user.email) : null;


  // Load conversation and messages when conversationId changes
  useEffect(() => {
    if (conversationId && userId) {
      loadConversationData(conversationId);
    } else if (userId) {
      // Create a new conversation
      createNewConversation();
    }
  }, [conversationId, userId]);

  const createNewConversation = async () => {
    if (!userId || !token) {
      setApiError('Authentication required');
      return;
    }

    try {
      setApiError(null);
      const userId = 2;
      const conversation = await chatAPI.createConversation(userId, {
        title: 'New Chat',
      });
      setCurrentConversation(conversation);
      setMessages([
        {
          id: '1',
          content: `Hello ${user?.name}! I'm your AI fitness assistant. I can help you with workout planning, nutrition advice, and fitness-related questions. How can I assist you today?`,
          role: 'assistant',
          timestamp: new Date(),
        },
      ]);
      // Update URL with new conversation ID
      router.replace(`/chat?conversationId=${conversation.id}`);
    } catch (err) {
      setApiError('Failed to create new conversation');
      console.error('Error creating conversation:', err);
    }
  };

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
      const conversation = await chatAPI.getConversation(userId, convId);
      setCurrentConversation(conversation);
      
      // Load messages
      const apiMessages = await chatAPI.getMessages(userId, convId);
      const formattedMessages: Message[] = apiMessages.map((msg: APIMessage) => ({
        id: msg.id.toString(),
        content: msg.content,
        role: msg.message_type === 'user' ? 'user' : 'assistant',
        timestamp: new Date(msg.created_at),
      }));
      
      setMessages(formattedMessages);
    } catch (err) {
      setApiError('Failed to load conversation');
      console.error('Error loading conversation:', err);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSendMessage = async () => {
    if (!inputMessage.trim() || isLoading || !currentConversation || !userId || !token) return;

    const userMessage: Message = {
      id: Date.now().toString(),
      content: inputMessage.trim(),
      role: 'user',
      timestamp: new Date(),
    };

    setMessages(prev => [...prev, userMessage]);
    setInputMessage('');
    setIsLoading(true);
    setApiError(null);

    try {
      // Send message to backend
      const userId = 2;
      await chatAPI.createMessage(userId, currentConversation.id, {
        content: userMessage.content,
        message_type: 'user',
      });

      // The backend will automatically generate an AI response
      // We need to poll for new messages or implement WebSocket
      // For now, let's poll after a short delay
      setTimeout(async () => {
        try {
          const userId = 2;
          const updatedMessages = await chatAPI.getMessages(userId, currentConversation.id);
          const formattedMessages: Message[] = updatedMessages.map((msg: APIMessage) => ({
            id: msg.id.toString(),
            content: msg.content,
            role: msg.message_type === 'user' ? 'user' : 'assistant',
            timestamp: new Date(msg.created_at),
          }));
          setMessages(formattedMessages);
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
        id: (Date.now() + 1).toString(),
        content: "I'm sorry, I'm having trouble responding right now. Please try again in a moment.",
        role: 'assistant',
        timestamp: new Date(),
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
      {/* Header */}
      <Paper 
        elevation={1} 
        sx={{ 
          borderRadius: 0,
          bgcolor: 'background.paper',
          borderBottom: `1px solid ${theme.palette.divider}`,
          flexShrink: 0
        }}
      >
        <Box 
          sx={{ 
            maxWidth: '1200px', 
            mx: 'auto', 
            p: 3,
            display: 'flex',
            alignItems: 'center',
            gap: 2
          }}
        >
          <SmartToyIcon color="primary" sx={{ fontSize: 32 }} />
          <Box>
            <Typography variant="h5" color="text.primary" fontWeight={600}>
              {currentConversation?.title || 'AI Fitness Assistant'}
            </Typography>
            <Typography variant="body1" color="text.secondary">
              Your personal workout and nutrition advisor
            </Typography>
          </Box>
        </Box>
      </Paper>

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