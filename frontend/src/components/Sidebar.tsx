'use client';

import { useState } from 'react';
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
} from '@mui/material';
import {
  Add as AddIcon,
  Chat as ChatIcon,
  MoreVert as MoreVertIcon,
  Delete as DeleteIcon,
} from '@mui/icons-material';
import { useSession } from 'next-auth/react';
import { useRouter, usePathname } from 'next/navigation';

const DRAWER_WIDTH = 240;

// Mock chat data - replace with real data from your backend
const mockChats = [
  { id: '1', title: 'Workout Plan Discussion', lastMessage: 'Can you create a push-pull routine?', timestamp: '2 hours ago', isActive: false },
  { id: '2', title: 'Nutrition Advice', lastMessage: 'What should I eat post-workout?', timestamp: '1 day ago', isActive: false },
  { id: '3', title: 'Form Check', lastMessage: 'Is my deadlift form correct?', timestamp: '3 days ago', isActive: true },
  { id: '4', title: 'Recovery Tips', lastMessage: 'How long should I rest between sets?', timestamp: '1 week ago', isActive: false },
];

export default function Sidebar() {
  const theme = useTheme();
  const { data: session } = useSession();
  const router = useRouter();
  const pathname = usePathname();

  const handleNewChat = () => {
    // Navigate to new chat or create new chat
    router.push('/chat');
  };

  const handleChatSelect = (chatId: string) => {
    router.push(`/chat/${chatId}`);
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
      <>
        {/* Drawer Header */}
        <Box
          sx={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between',
            p: 2,
          }}
        >
          <Typography variant="h6" noWrap component="div">
            Chats
          </Typography>
        </Box>

        <Divider />

        {/* New Chat Button */}
        <Box sx={{ p: 2 }}>
          <Button
            variant="contained"
            fullWidth
            startIcon={<AddIcon />}
            onClick={handleNewChat}
            sx={{
              borderRadius: 2,
              textTransform: 'none',
            }}
          >
            New Chat
          </Button>
        </Box>

        <Divider />

        {/* Chat History */}
        <Box sx={{ px: 2, py: 1 }}>
          <Typography variant="subtitle2" color="text.secondary">
            Recent Chats
          </Typography>
        </Box>

        <List sx={{ px: 1 }}>
          {mockChats.map((chat) => (
            <ListItem key={chat.id} disablePadding sx={{ mb: 0.5 }}>
              <ListItemButton
                selected={pathname === `/chat/${chat.id}`}
                onClick={() => handleChatSelect(chat.id)}
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
                        {chat.title}
                      </Typography>
                      {chat.isActive && (
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
                        {chat.lastMessage}
                      </Typography>
                      <Typography variant="caption" color="text.secondary" display="block">
                        {chat.timestamp}
                      </Typography>
                    </Box>
                  }
                />
              </ListItemButton>
            </ListItem>
          ))}
        </List>

        {/* User Profile Section at Bottom */}
        <Box sx={{ mt: 'auto' }}>
          <Divider />
          {session?.user && (
            <Box
              sx={{
                p: 2,
                display: 'flex',
                alignItems: 'center',
                gap: 2,
              }}
            >
              <Avatar
                src={session.user.image ?? undefined}
                alt={session.user.name ?? 'User'}
                sx={{ width: 32, height: 32 }}
              />
              <Box sx={{ flex: 1, minWidth: 0 }}>
                <Typography variant="subtitle2" noWrap>
                  {session.user.name}
                </Typography>
                <Typography variant="caption" color="text.secondary" noWrap>
                  {session.user.email}
                </Typography>
              </Box>
            </Box>
          )}
        </Box>
      </>
    );
  }
}



