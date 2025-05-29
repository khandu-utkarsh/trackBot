'use client';

import React, { useState, useRef, useEffect } from 'react';
import {
  Box,
  TextField,
  IconButton,
  Paper,
  Typography,
  Avatar,
  useTheme,
  Divider,
  Chip,
  CircularProgress,
  Alert,
} from '@mui/material';
import SendIcon from '@mui/icons-material/Send';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import PersonIcon from '@mui/icons-material/Person';
import { useGoogleAuth } from '@/hooks/useGoogleAuth';
import { useRouter, useSearchParams } from 'next/navigation';
import { chatAPI, Message as APIMessage, Conversation } from '@/lib/api/chat';

interface Message {
  id: string;
  content: string;
  role: 'user' | 'assistant';
  timestamp: Date;
}

export default function ChatPageComponent() {
    const { user } = useGoogleAuth();
    const theme = useTheme();
    const router = useRouter();
    const searchParams = useSearchParams();
    const conversationId = searchParams.get('conversationId');
    
    const [messages, setMessages] = useState<Message[]>([]);
    const [inputMessage, setInputMessage] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [currentConversation, setCurrentConversation] = useState<Conversation | null>(null);
    const messagesEndRef = useRef<HTMLDivElement>(null);
    const inputRef = useRef<HTMLInputElement>(null);
  
    // Mock user ID - in a real app, this would come from session
    const userId = 1;
  
    const scrollToBottom = () => {
      messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
    };
  
    useEffect(() => {
      scrollToBottom();
    }, [messages]);
  
    // Load conversation and messages when conversationId changes
    useEffect(() => {
      if (conversationId) {
        loadConversationData(parseInt(conversationId));
      } else {
        // Create a new conversation
        createNewConversation();
      }
    }, [conversationId]);
  
    const createNewConversation = async () => {
      try {
        setError(null);
        const conversation = await chatAPI.createConversation(userId, {
          title: 'New Chat',
          is_active: true,
        });
        setCurrentConversation(conversation);
        setMessages([
          {
            id: '1',
            content: "Hello! I'm your AI fitness assistant. I can help you with workout planning, nutrition advice, and fitness-related questions. How can I assist you today?",
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
      try {
        setError(null);
        setIsLoading(true);
        
        // Load conversation details
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
      if (!inputMessage.trim() || isLoading || !currentConversation) return;
  
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
        await chatAPI.createMessage(userId, currentConversation.id, {
          content: userMessage.content,
          message_type: 'user',
        });
  
        // The backend will automatically generate an AI response
        // We need to poll for new messages or implement WebSocket
        // For now, let's poll after a short delay
        setTimeout(async () => {
          try {
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
              <Box
                key={message.id}
                sx={{
                  display: 'flex',
                  flexDirection: message.role === 'user' ? 'row-reverse' : 'row',
                  gap: 2,
                  alignItems: 'flex-start',
                }}
              >
                <Avatar
                  sx={{
                    bgcolor: message.role === 'user' ? 'primary.main' : 'secondary.main',
                    width: 48,
                    height: 48,
                    flexShrink: 0,
                  }}
                >
                  {message.role === 'user' ? (
                    user?.picture ? (
                      <img 
                        src={user.picture} 
                        alt="User" 
                        style={{ width: '100%', height: '100%', borderRadius: '50%' }}
                      />
                    ) : (
                      <PersonIcon />
                    )
                  ) : (
                    <SmartToyIcon />
                  )}
                </Avatar>
                
                <Paper
                  elevation={2}
                  sx={{
                    p: 2.5,
                    maxWidth: '70%',
                    bgcolor: message.role === 'user' 
                      ? 'primary.main' 
                      : 'background.paper',
                    color: message.role === 'user' 
                      ? 'primary.contrastText' 
                      : 'text.primary',
                    borderRadius: 3,
                    position: 'relative',
                    ...(message.role === 'user' ? {
                      borderTopRightRadius: 8,
                    } : {
                      borderTopLeftRadius: 8,
                    }),
                  }}
                >
                  <Typography variant="body1" sx={{ 
                    whiteSpace: 'pre-wrap', 
                    lineHeight: 1.6,
                    fontSize: '1rem'
                  }}>
                    {message.content}
                  </Typography>
                  <Typography 
                    variant="caption" 
                    sx={{ 
                      display: 'block', 
                      mt: 1.5, 
                      opacity: 0.7,
                      textAlign: message.role === 'user' ? 'right' : 'left',
                      fontSize: '0.75rem'
                    }}
                  >
                    {formatTime(message.timestamp)}
                  </Typography>
                </Paper>
              </Box>
            ))}
            
            {isLoading && (
              <Box sx={{ display: 'flex', gap: 2, alignItems: 'flex-start' }}>
                <Avatar sx={{ bgcolor: 'secondary.main', width: 48, height: 48 }}>
                  <SmartToyIcon />
                </Avatar>
                <Paper
                  elevation={2}
                  sx={{
                    p: 2.5,
                    bgcolor: 'background.paper',
                    borderRadius: 3,
                    borderTopLeftRadius: 8,
                    display: 'flex',
                    alignItems: 'center',
                    gap: 2,
                  }}
                >
                  <CircularProgress size={20} />
                  <Typography variant="body1" color="text.secondary">
                    AI is thinking...
                  </Typography>
                </Paper>
              </Box>
            )}
            
            <div ref={messagesEndRef} />
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
          <Box 
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