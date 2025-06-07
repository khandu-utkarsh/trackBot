'use client';

import { List } from '@mui/material';
import { Conversation, User } from '@/lib/api';
import SidebarListItem from './SidebarListItem';

interface SidebarListProps {
  conversations: Map<number, Conversation>;
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

export default function SidebarList({
  conversations,
  currentConversationId,
  onChatSelect,
  onDeleteConversation,
  user,
}: SidebarListProps) {
  const sortedConversations = Array.from(conversations.values()).sort(
    (a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
  );

  return (
    <List sx={{ px: 1 }}>
      {sortedConversations.map((conversation) => (
        <SidebarListItem
          key={conversation.id}
          conversation={conversation}
          currentConversationId={currentConversationId}
          onChatSelect={onChatSelect}
          onDeleteConversation={onDeleteConversation}
          user={user}
        />
      ))}
    </List>
  );
} 