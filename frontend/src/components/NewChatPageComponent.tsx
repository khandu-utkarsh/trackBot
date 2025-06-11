'use client';

import React, { useState, useRef, useEffect } from 'react';
import {
  Box,
  useTheme,
} from '@mui/material';
import { useRouter } from 'next/navigation';
import { chatAPI, Conversation } from '@/lib/api';
import ChatInputBar from './ChatInputBar';
import { useRequireAuth, useConversations } from '@/contexts/AuthContext';
import { CreateConversationResponse, ListMessagesResponse, MessageType } from '@/lib/types/generated';

let chatIdTemp : number = 0;
function ChatPageComponent() {
  const [inputMessage, setTextMessage] = useState('');
  const [conversations, setConversations] = useConversations();

  console.log("ChatPageComponent rendered.");
  const { user } = useRequireAuth();
  if(!user || !user.id) {
    console.log("user", user);    
    return <div>User is not authenticated</div>;
  }
  else{
    console.log("user is null");
  }
  const router = useRouter();
  
  const createNewConversation = async (inputMessage: string) => {
    let conversation : Conversation | null = null;
    try {
          let conversationResponse : CreateConversationResponse = await chatAPI.createConversation(user.id, {
            title: 'New Chat ' + chatIdTemp.toString(),
          });

          if(conversationResponse.id){
            chatIdTemp++;
            conversation = {
              id: conversationResponse.id,
              title: conversationResponse.title,
              user_id: conversationResponse.user_id,
              updated_at: conversationResponse.updated_at,
            };
          }
          else{
            console.error('Error creating conversation: Conversation is null');
          }
        }
        catch(err){
          console.error('Error creating conversation:', err);
        }
        return conversation;
      };  

      const handleSendMessage = async (inputMessage: string) => {
        const conversationCreated = await createNewConversation(inputMessage);
        if(conversationCreated){
          const newMap = new Map(conversations);
          newMap.set(conversationCreated.id, conversationCreated);
          setConversations(newMap);

          //!Send message to the conversation.
          const message : ListMessagesResponse = await chatAPI.createMessage(user.id, conversationCreated.id, {
            langchain_message: inputMessage,
            message_type: MessageType.User,                
          });
          if(message.messages.length > 0){
            console.log("Message sent to the conversation.");
          }
          else{
            console.error('Error sending message: Message is null');
          }

          //!Message sent, route to the conversation page.
          router.replace(`/chat/${conversationCreated.id}`);
        }
        else{
          console.error('Error creating conversation: Conversation is null');
        }
      }

  return (
      <Box className="chat-page-container" sx={{ width: '100%' }}>
        <ChatInputBar inputMessage={inputMessage} setInputMessage={setTextMessage} handleSendMessage={handleSendMessage} />
      </Box>
  );
}

export default function ChatApp() {
  return (
      <ChatPageComponent />
  );
}





 