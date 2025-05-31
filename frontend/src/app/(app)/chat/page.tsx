'use client';

import { useState, useRef, useEffect } from 'react';
import {
  Box,
  TextField,
  IconButton,
  Paper,
  Typography,
  useTheme,
  Alert,
} from '@mui/material';
import SendIcon from '@mui/icons-material/Send';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import { useRouter, useSearchParams } from 'next/navigation';
import { chatAPI, Message as APIMessage, Conversation } from '@/lib/api/chat';
import { useAuth } from '@/contexts/AuthContext';
import ChatMessage from '@/components/ChatMessage';
import AIThinkingMessage from '@/components/AIThinkingMessage';


export default function ChatPageContent() {

  console.log("ChatPageContent from page rendered.");

  const theme = useTheme();  
  const router = useRouter();
  const searchParams = useSearchParams();
  const conversationId = searchParams.get('conversationId');
  const { user, token, isAuthenticated } = useAuth();
  
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputMessage, setInputMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
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
      loadConversationData(parseInt(conversationId));
    } else if (userId) {
      // Create a new conversation
      createNewConversation();
    }
  }, [conversationId, userId]);

  const createNewConversation = async () => {
    if (!userId || !token) {
      setError('Authentication required');
      return;
    }

    try {
      setError(null);
      const userId = 2;
      const conversation = await chatAPI.createConversation(userId, {
        title: 'New Chat',
        is_active: true,
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
      setError('Failed to create new conversation');
      console.error('Error creating conversation:', err);
    }
  };

  const loadConversationData = async (convId: number) => {
    if (!userId || !token) {
      setError('Authentication required');
      return;
    }

    try {
      setError(null);
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
      setError('Failed to load conversation');
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
    setError(null);

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
      setError('Failed to send message. Please try again.');
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

  const handleKeyPress = (event: React.KeyboardEvent) => {
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault();
      handleSendMessage();
    }
  };

  const formatTime = (date: Date) => {
    return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
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
      {error && (
        <Alert severity="error" sx={{ m: 2 }} onClose={() => setError(null)}>
          {error}
        </Alert>
      )}

      {/* Messages Container */}
      <Box 
        className="prior-messages-container-class-for-debugging"
        sx={{ 
          flex: 1,
          overflow: 'auto',
          display: 'flex',
          justifyContent: 'center',
          '&::-webkit-scrollbar': {
            width: '8px',
          },
          '&::-webkit-scrollbar-track': {
            background: 'rgba(0,0,0,0.1)',
            borderRadius: '4px',
          },
          '&::-webkit-scrollbar-thumb': {
            background: 'rgba(0,0,0,0.3)',
            borderRadius: '4px',
            '&:hover': {
              background: 'rgba(0,0,0,0.5)',
            },
          },
        }}
      >
        <Box 
          className="messages-container-class-for-debugging"
          sx={{
            width: '100%',
            maxWidth: '1200px',
            p: 3,
            display: 'flex',
            flexDirection: 'column',
            gap: 3,
          }}
        >
          {messages.map((message) => (
            <ChatMessage key={message.id} message={message} />
          ))}
          
          {isLoading && ( <AIThinkingMessage />)}
          
        </Box>
      </Box>

      {/* Input Area */}
      <Paper 
        elevation={4} 
        sx={{ 
          borderRadius: 0,
          bgcolor: 'background.paper',
          borderTop: `1px solid ${theme.palette.divider}`,
          flexShrink: 0
        }}
      >
        <Box className="chat-input-container" 
          sx={{
            maxWidth: '1200px',
            mx: 'auto',
            p: 3,
          }}
        >
          <Box display="flex" gap={2} alignItems="flex-end" mb={2}>
            <TextField
              ref={inputRef}
              fullWidth
              multiline
              maxRows={4}
              value={inputMessage}
              onChange={(e) => setInputMessage(e.target.value)}
              onKeyPress={handleKeyPress}
              placeholder="Type your fitness question here..."
              variant="outlined"
              disabled={isLoading || !currentConversation}
              sx={{
                '& .MuiOutlinedInput-root': {
                  borderRadius: 3,
                  bgcolor: 'background.default',
                  fontSize: '1rem',
                  '& fieldset': {
                    borderColor: 'divider',
                  },
                  '&:hover fieldset': {
                    borderColor: 'primary.main',
                  },
                },
              }}
            />
            <IconButton
              onClick={handleSendMessage}
              disabled={!inputMessage.trim() || isLoading || !currentConversation}
              color="primary"
              sx={{
                bgcolor: 'primary.main',
                color: 'white',
                '&:hover': {
                  bgcolor: 'primary.dark',
                },
                '&:disabled': {
                  bgcolor: 'action.disabled',
                },
                width: 56,
                height: 56,
              }}
            >
              <SendIcon />
            </IconButton>
          </Box>
        </Box>
      </Paper>
    </Box>
  );
}