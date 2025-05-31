import { Box, Paper, Typography } from '@mui/material';

export interface Message {
    id: string;
    content: string;
    role: 'user' | 'assistant';
    timestamp: Date;
}
  
const formatTime = (date: Date) => {
    return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
};


export default function ChatMessage({ message }: { message: Message }) {
    return (
    <Box
            key={message.id}
            sx={{
                    display: 'flex',
                    flexDirection: message.role === 'user' ? 'row-reverse' : 'row',
                    gap: 2,
                    alignItems: 'flex-start',
                }}
    >
        <Paper
                elevation={2}
                sx={{
                        p: 2.5,
                        maxWidth: '70%',
                        bgcolor: message.role === 'user' ? 'primary.main' : 'background.paper',
                        color: message.role === 'user' ? 'primary.contrastText' : 'text.primary',
                        borderRadius: 3,
                        position: 'relative',
                        ...(message.role === 'user' ? {borderTopRightRadius: 8,} : {borderTopLeftRadius: 8,}),
                }}
        >
            <Typography 
                variant="body1" 
                sx={{ 
                    whiteSpace: 'pre-wrap', 
                    lineHeight: 1.6,
                    fontSize: '1rem',
                    textAlign: message.role === 'user' ? 'right' : 'left',
                }}
            >
                {message.content}
            </Typography>
            <Typography 
                variant="caption" 
                sx={{ 
                    display: 'block', 
                    mt: 1.5, 
                    opacity: 0.7,
                    textAlign: message.role === 'user' ? 'right' : 'left',
                    fontSize: '0.75rem'
                }}
            >
                {formatTime(message.timestamp)}
            </Typography>
        </Paper>
    </Box>
    );
}