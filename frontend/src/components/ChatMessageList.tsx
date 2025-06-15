import ChatMessage, { Message } from "./ChatMessage";
import AIThinkingMessage from "./AIThinkingMessage";
import { Box } from "@mui/material";
import { useRef, useEffect } from "react";
import { MessageType } from "@/lib/types/generated/api";

export default function ChatMessageList({ messages, isLoading }: { messages: Message[], isLoading: boolean}) {
    const messagesEndRef = useRef<HTMLDivElement>(null);

    const scrollToBottom = () => {
        messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
    };

    useEffect(() => {
        scrollToBottom();
    }, [messages]);


    const filteredMessages = messages.filter(
        (message) => {
            if (message.message_type === MessageType.Other) {
                return false;
            }

            const messageContent = JSON.parse(message.langchain_message);
            let content = messageContent.content;
            if(content?.length > 0) {
                return true;
            }
            return false;    
        }
    );

    console.log(filteredMessages);
    return (
        <Box 
            className="messages-list-container"
            sx={{
                width: '100%',
                maxWidth: '1200px',
                p: 3,
                display: 'flex',
                flexDirection: 'column',
                gap: 3,
            }}
        >
        {filteredMessages.map((message) => ( <ChatMessage key={message.id} message={message} /> ))}
        {isLoading && ( <AIThinkingMessage />)}
        <div ref={messagesEndRef} />
        </Box>
    );
}