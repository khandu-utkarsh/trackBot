'use client';

import { useEffect, useRef } from 'react';
import {
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  IconButton,
  useTheme,
  Typography,
  Chip,
} from '@mui/material';
import {
  Chat as ChatIcon,
  Delete as DeleteIcon,
} from '@mui/icons-material';
import { Conversation, User } from '@/lib/api';

interface SidebarListItemProps {
  conversation: Conversation;
  currentConversationId: number | null;
  onChatSelect: (conversationId: number) => void;
  onDeleteConversation: (
    user: User | null, 
    conversationId: number, 
    event: React.MouseEvent, 
    currentConversationId: number | null
  ) => void;
  user: User | null;
}

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

export default function SidebarListItem({
  conversation,
  currentConversationId,
  onChatSelect,
  onDeleteConversation,
  user,
}: SidebarListItemProps) {
  const theme = useTheme();
  const listItemRef = useRef<HTMLLIElement>(null);
  const isActive = currentConversationId === conversation.id;

  // Scroll the active item into view when it becomes active
  useEffect(() => {
    if (isActive && listItemRef.current) {
      listItemRef.current.scrollIntoView({
        behavior: 'smooth',
        block: 'nearest',
        inline: 'nearest'
      });
    }
  }, [isActive]);

  return (
    <ListItem 
      ref={listItemRef}
      key={conversation.id} 
      disablePadding 
      sx={{ mb: 0.5 }}
    >
      <ListItemButton
        selected={isActive}
        onClick={() => onChatSelect(conversation.id)}
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
              {isActive && (
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
                No messages yet
              </Typography>
              <Typography variant="caption" color="text.secondary" display="block">
                {formatTimestamp(conversation.created_at)}
              </Typography>
            </>
          }
        />
        <IconButton
          size="small"
          onClick={(e) => onDeleteConversation(user, conversation.id, e, currentConversationId)}
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
  );
} 