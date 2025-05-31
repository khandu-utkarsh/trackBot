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
  Delete as DeleteIcon,
} from '@mui/icons-material';
import { useRouter, usePathname, useSearchParams } from 'next/navigation';
import { chatAPI, Conversation } from '@/lib/api/chat';
import { useAuth } from '@/contexts/AuthContext';
import { GoogleUser } from '@/hooks/useGoogleAuth';


const DRAWER_WIDTH = 240;

const loadConversations = async (user: GoogleUser) => {
  let data : Conversation[] = [];
  try {
    const userId = 2;
    data = await chatAPI.getConversations(userId);
  } catch (err) {
    console.error('Unable to fetch the conversations for the  user:', user.name, " error: ", err);
  }
  return data;
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


const handleDeleteConversation = async (conversationId: number, event: React.MouseEvent) => {
  console.log("Yet to be implemented: Implementation is pending.");
  console.log("Yet to be implemented: Implementation is pending.");

  /*
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
*/
  };

export default function Sidebar() {
  const theme = useTheme();
  const { user } = useAuth();

  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const currentConversationId = searchParams.get('conversationId');
  
  const [conversations, setConversations] = useState<Conversation[]>([]);
  const [isLoading, setIsLoading] = useState(true);


  const handleNewChat = async () => {
    router.push('/');
  };

  // Load conversations on component mount and when the user changes
  useEffect(() => {
    setIsLoading(true);
    if (user) {
      loadConversations(user).then((data) => {
        setConversations(data);
        setIsLoading(false);
      });
    } 
  }, [user]);

  const handleChatSelect = (conversationId: number) => {
    console.log("Chat selected: ", conversationId);
    console.log("Yet to be implemented: Implementation is pending.");
  
    //!This should open up the chat box for the selected conversation
    router.push(`/chat?conversationId=${conversationId}`);
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
                        <span style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
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
                        </span>
                      }
                      secondary={
                        <>
                          <Typography variant="caption" color="text.secondary" noWrap>
                            {conversation.last_message || 'No messages yet'}
                          </Typography>
                          <Typography variant="caption" color="text.secondary" display="block">
                            {formatTimestamp(conversation.updated_at)}
                          </Typography>
                        </>
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
      </Box>
    );
  }
}



