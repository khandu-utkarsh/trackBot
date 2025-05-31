'use client';

import React, { useState, useRef, useEffect } from 'react';
import {
  Box,
  useTheme,
} from '@mui/material';
import { useRouter } from 'next/navigation';
import { chatAPI } from '@/lib/api/chat';
import ChatInputBar from './ChatInputBar';
import { Conversation } from '@/lib/api/chat';
import { useRequireAuth, useConversations } from '@/contexts/AuthContext';


interface Message {
  id: string;
  content: string;
  role: 'user' | 'assistant';
  timestamp: Date;
}
let chatIdTemp : number = 0;
function ChatPageComponent() {
  const [inputMessage, setTextMessage] = useState('');
  const [conversations, setConversations] = useConversations();

  console.log("ChatPageComponent rendered.");
  const { user } = useRequireAuth();
  console.log("user", user);
  const router = useRouter();
  
  // Mock user ID - in a real app, this would come from session
  const userId = 2;

  const createNewConversation = async (inputMessage: string) => {
    let conversation : Conversation | null = null;
    try {
          const userId = 2;
          conversation = await chatAPI.createConversation(userId, {
            title: 'New Chat ' + chatIdTemp.toString(),
            is_active: true,
          });
          if(conversation){ 
            chatIdTemp++;

            //!Send message to the conversation.
            const message = await chatAPI.createMessage(userId, conversation.id, {
              content: inputMessage,
              message_type: 'user',                
            });
            if(message){
              console.log("Message sent to the conversation.");
            }
            else{
              console.error('Error sending message: Message is null');
            }
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
  const theme = useTheme();

  return (
      <ChatPageComponent />
  );
}





 