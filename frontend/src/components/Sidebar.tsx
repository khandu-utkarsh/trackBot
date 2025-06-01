'use client';

import { useState, useEffect } from 'react';
import {
  Box,
  Drawer,
  Divider,
  useTheme,
  Typography,
  Button,
  CircularProgress,
} from '@mui/material';
import {
  Add as AddIcon,
} from '@mui/icons-material';
import { useRouter, useParams} from 'next/navigation';
import { chatAPI } from '@/lib/api/chat';
import { Conversation } from '@/lib/types/chat';
import { useRequireAuth, useConversations } from '@/contexts/AuthContext';
import { GoogleUser } from '@/hooks/useGoogleAuth';
import SidebarList from './SidebarList';

const DRAWER_WIDTH = 240;

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

export default function Sidebar() {
  const theme = useTheme();
  const { user } = useRequireAuth();
  const router = useRouter();
  const params = useParams();

  let currentConversationId: number | null = null;
  if (params.conversationId) {
    currentConversationId = parseInt(params.conversationId as string);
  }
  console.log("currentConversationId: ", currentConversationId);

  const [conversations, setConversations] = useConversations();
  const [isLoading, setIsLoading] = useState(true);

  // Load conversations on component mount and when the user changes
  useEffect(() => {
    setIsLoading(true);
    if (user) {
      loadConversations(user);
      setIsLoading(false);
    } 
  }, [user]);

  //!For loading the conversations from the backend
  const loadConversations = async (user: GoogleUser) => {
    let data : Conversation[] = [];
    try {
      const userId = 2;
      data = await chatAPI.getConversations(userId);
    } catch (err) {
      console.error('Unable to fetch the conversations for the  user:', user.name, " error: ", err);
    }

    const newMap = new Map(conversations);
    data.forEach(conversation => {
      newMap.set(conversation.id, conversation);
    });
    setConversations(newMap);
  };

  //!When user clicks on the new chat button
  const handleNewChat = async () => {
    router.push('/');
  };

  //!When user deletes a conversation
  const handleDeleteConversation = async (user: GoogleUser | null, conversationId: number, event: React.MouseEvent, currentConversationId: number | null) => {
    event.stopPropagation(); // Prevent chat selection
    
    const userId = 2;
    try {
      await chatAPI.deleteConversation(userId, conversationId);
      //!Doesn't make sense to reload everything, simply delete from the map using the id.
      if (user) {
        const newMap = new Map(conversations);
        newMap.delete(conversationId);
        setConversations(newMap);
      }
      
      // If we deleted the current conversation, navigate to new chat
      if (currentConversationId && currentConversationId === conversationId) {
        router.push('/');
      }
    } catch (err) {
      console.error('Error deleting conversation with the id: ', conversationId, ' error: ', err);
    }
  };

  //!When user selects a conversation, routing to the correct chat page
  const handleChatSelect = (conversationId: number) => {
    router.push(`/chat/${conversationId}`);
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
            onClick={() => handleNewChat()}
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
          ) : conversations.size === 0 ? (
            <Box sx={{ p: 2, textAlign: 'center' }}>
              <Typography variant="body2" color="text.secondary">
                No conversations yet. Start a new chat!
              </Typography>
            </Box>
          ) : (
            <SidebarList
              conversations={conversations}
              currentConversationId={currentConversationId}
              onChatSelect={handleChatSelect}
              onDeleteConversation={handleDeleteConversation}
              user={user}
            />
          )}
        </Box>
      </Box>
    </Drawer>
  );
}



