'use client';

import { useState, useEffect } from 'react';
import {
  Box,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Divider,
  IconButton,
  useTheme,
  Avatar,
  Typography,
  Button,
  Chip,
  CircularProgress,
  Alert,
} from '@mui/material';
import {
  Add as AddIcon,
  Chat as ChatIcon,
  MoreVert as MoreVertIcon,
  Delete as DeleteIcon,
} from '@mui/icons-material';
import { useSession } from 'next-auth/react';
import { useRouter, usePathname, useSearchParams } from 'next/navigation';
import { chatAPI, Conversation } from '@/lib/api/chat';

const DRAWER_WIDTH = 240;

export default function Sidebar() {
  const theme = useTheme();
  const { data: session } = useSession();
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const currentConversationId = searchParams.get('conversationId');
  
  const [conversations, setConversations] = useState<Conversation[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Mock user ID - in a real app, this would come from session
  const userId = 1;

  // Load conversations on component mount
  useEffect(() => {
    loadConversations();
  }, []);

  const loadConversations = async () => {
    try {
      setError(null);
      setIsLoading(true);
      const data = await chatAPI.getConversations(userId);
      setConversations(data);
    } catch (err) {
      setError('Failed to load conversations');
      console.error('Error loading conversations:', err);
    } finally {
      setIsLoading(false);
    }
  };

  const handleNewChat = () => {
    // Navigate to new chat (without conversationId)
    router.push('/chat');
  };

  const handleChatSelect = (conversationId: number) => {
    router.push(`/chat?conversationId=${conversationId}`);
  };

  const handleDeleteConversation = async (conversationId: number, event: React.MouseEvent) => {
    event.stopPropagation(); // Prevent chat selection
    
    try {
      await chatAPI.deleteConversation(userId, conversationId);
      // Reload conversations
      await loadConversations();
      
      // If we deleted the current conversation, navigate to new chat
      if (currentConversationId === conversationId.toString()) {
        router.push('/chat');
      }
    } catch (err) {
      console.error('Error deleting conversation:', err);
      setError('Failed to delete conversation');
    }
  };

  const formatTimestamp = (dateString: string) => {
    const date = new Date(dateString);
    const now = new Date();
    const diffInHours = (now.getTime() - date.getTime()) / (1000 * 60 * 60);
    
    if (diffInHours < 1) {
      return 'Just now';
    } else if (diffInHours < 24) {
      return `${Math.floor(diffInHours)} hours ago`;
    } else if (diffInHours < 48) {
      return '1 day ago';
    } else if (diffInHours < 168) {
      return `${Math.floor(diffInHours / 24)} days ago`;
    } else {
      return `${Math.floor(diffInHours / 168)} weeks ago`;
    }
  };

  return (
    <Drawer
      variant="permanent"
      sx={{
        width: DRAWER_WIDTH,
        flexShrink: 0,
        '& .MuiDrawer-paper': {
          width: DRAWER_WIDTH,
          boxSizing: 'border-box',
          backgroundColor: theme.palette.background.paper,
          borderRight: `1px solid ${theme.palette.divider}`,
        },
      }}
      open
    >
      <DrawerContent />
    </Drawer>
  );

  function DrawerContent() {
    return (
      <Box sx={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
        {/* Header */}
        <Box sx={{ p: 2, borderBottom: `1px solid ${theme.palette.divider}` }}>
          <Typography variant="h6" color="text.primary" fontWeight={600}>
            Fitness Chat
          </Typography>
          <Typography variant="body2" color="text.secondary">
            AI-powered fitness coaching
          </Typography>
        </Box>

        {/* New Chat Button */}
        <Box sx={{ p: 2 }}>
          <Button
            fullWidth
            variant="contained"
            startIcon={<AddIcon />}
            onClick={handleNewChat}
            sx={{
              borderRadius: 2,
              textTransform: 'none',
              fontWeight: 500,
            }}
          >
            New Chat
          </Button>
        </Box>

        <Divider />

        {/* Error Alert */}
        {error && (
          <Alert 
            severity="error" 
            sx={{ m: 1, fontSize: '0.75rem' }} 
            onClose={() => setError(null)}
          >
            {error}
          </Alert>
        )}

        {/* Conversations List */}
        <Box sx={{ flex: 1, overflow: 'auto' }}>
          {isLoading ? (
            <Box sx={{ display: 'flex', justifyContent: 'center', p: 3 }}>
              <CircularProgress size={24} />
            </Box>
          ) : conversations.length === 0 ? (
            <Box sx={{ p: 2, textAlign: 'center' }}>
              <Typography variant="body2" color="text.secondary">
                No conversations yet. Start a new chat!
              </Typography>
            </Box>
          ) : (
            <List sx={{ px: 1 }}>
              {conversations.map((conversation) => (
                <ListItem key={conversation.id} disablePadding sx={{ mb: 0.5 }}>
                  <ListItemButton
                    selected={currentConversationId === conversation.id.toString()}
                    onClick={() => handleChatSelect(conversation.id)}
                    sx={{
                      borderRadius: 1,
                      '&.Mui-selected': {
                        backgroundColor: theme.palette.action.selected,
                      },
                      '&:hover': {
                        backgroundColor: theme.palette.action.hover,
                      },
                    }}
                  >
                    <ListItemIcon sx={{ minWidth: 36 }}>
                      <ChatIcon fontSize="small" />
                    </ListItemIcon>
                    <ListItemText
                      primary={
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                          <Typography variant="body2" noWrap sx={{ flex: 1 }}>
                            {conversation.title}
                          </Typography>
                          {conversation.is_active && (
                            <Chip
                              label="Active"
                              size="small"
                              color="primary"
                              sx={{ height: 16, fontSize: '0.6rem' }}
                            />
                          )}
                        </Box>
                      }
                      secondary={
                        <Box>
                          <Typography variant="caption" color="text.secondary" noWrap>
                            {conversation.last_message || 'No messages yet'}
                          </Typography>
                          <Typography variant="caption" color="text.secondary" display="block">
                            {formatTimestamp(conversation.updated_at)}
                          </Typography>
                        </Box>
                      }
                    />
                    <IconButton
                      size="small"
                      onClick={(e) => handleDeleteConversation(conversation.id, e)}
                      sx={{ 
                        opacity: 0.6,
                        '&:hover': { opacity: 1 },
                        ml: 1
                      }}
                    >
                      <DeleteIcon fontSize="small" />
                    </IconButton>
                  </ListItemButton>
                </ListItem>
              ))}
            </List>
          )}
        </Box>

        {/* User Profile Section at Bottom */}
        <Box sx={{ 
          p: 2, 
          borderTop: `1px solid ${theme.palette.divider}`,
          backgroundColor: theme.palette.background.default 
        }}>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
            <Avatar 
              sx={{ width: 32, height: 32 }}
              src={session?.user?.image || undefined}
            >
              {session?.user?.name?.[0] || 'U'}
            </Avatar>
            <Box sx={{ flex: 1, minWidth: 0 }}>
              <Typography variant="body2" fontWeight={500} noWrap>
                {session?.user?.name || 'User'}
              </Typography>
              <Typography variant="caption" color="text.secondary" noWrap>
                {session?.user?.email || 'user@example.com'}
              </Typography>
            </Box>
            <IconButton size="small">
              <MoreVertIcon fontSize="small" />
            </IconButton>
          </Box>
        </Box>
      </Box>
    );
  }
}



