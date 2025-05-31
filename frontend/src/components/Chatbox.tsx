import { Box } from '@mui/material';
import ChatMessageList from './ChatMessageList';

export default function Chatbox() {
    return (
        <Box
            className="chatbox-container"
            sx={{
                flex: 1,
                display: 'flex',
                justifyContent: 'center',
                overflow: 'auto',
                '&::-webkit-scrollbar': {
                    width: '8px',
                },
                '&::-webkit-scrollbar-track': {
                    background: 'rgba(0,0,0,0.1)',
                    borderRadius: '4px',
                },
                '&::-webkit-scrollbar-thumb': {
                    background: 'rgba(0,0,0,0.3)',
                    borderRadius: '4px',
                    '&:hover': {
                        background: 'rgba(0,0,0,0.5)',
                    },
                },
            }}
        >
            <ChatMessageList messages={messages} isLoading={isLoading}/>
     </Box>
 
    );
}