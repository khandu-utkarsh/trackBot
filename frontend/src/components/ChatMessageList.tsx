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
        {messages
        .filter((message) => message.message_type !== MessageType.Other)
        .map((message) => (
            <ChatMessage key={message.id} message={message} />
        ))}
        {isLoading && ( <AIThinkingMessage />)}
        <div ref={messagesEndRef} />
        </Box>
    );
}