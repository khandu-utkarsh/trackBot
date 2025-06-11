import { Box, Paper, Typography } from '@mui/material';
import { Message } from '@/lib/api';

/*
export interface Message {
    id: string;
    content: string;
    role: 'user' | 'assistant';
    timestamp: Date;
}
*/
  
const formatTime = (date: Date) => {
    return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
};


export default function ChatMessage({ message }: { message: Message }) {


    const messageContent = JSON.parse(message.langchain_message);
    return (
    <Box
            key={message.id}
            sx={{
                    display: 'flex',
                    flexDirection: messageContent.type === 'human' ? 'row-reverse' : 'row',
                    gap: 2,
                    alignItems: 'flex-start',
                }}
    >
        <Paper
                elevation={2}
                sx={{
                        p: 2.5,
                        maxWidth: '70%',
                        bgcolor: messageContent.type === 'human' ? 'primary.main' : 'background.paper',
                        color: messageContent.type === 'human' ? 'primary.contrastText' : 'text.primary',
                        borderRadius: 3,
                        position: 'relative',
                        ...(messageContent.type === 'human' ? {borderTopRightRadius: 8,} : {borderTopLeftRadius: 8,}),
                }}
        >
            <Typography 
                variant="body1" 
                sx={{ 
                    whiteSpace: 'pre-wrap', 
                    lineHeight: 1.6,
                    fontSize: '1rem',
                    textAlign: messageContent.type === 'human' ? 'right' : 'left',
                }}
            >
                {messageContent.content}
            </Typography>
            <Typography 
                variant="caption" 
                sx={{ 
                    display: 'block', 
                    mt: 1.5, 
                    opacity: 0.7,
                    textAlign: messageContent.type === 'human' ? 'right' : 'left',
                    fontSize: '0.75rem'
                }}
            >
                {message.created_at ? formatTime(new Date(message.created_at)) : ''}
            </Typography>
        </Paper>
    </Box>
    );
}