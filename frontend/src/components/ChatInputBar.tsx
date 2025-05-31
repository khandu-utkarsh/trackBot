import { Box, TextField, IconButton } from '@mui/material';
import SendIcon from '@mui/icons-material/Send';
import { useRef, useState, useLayoutEffect } from 'react';

//!This component is only repsonible for sending the message to the backend.
export default function ChatInputBar({inputMessage, setInputMessage, handleSendMessage}: {inputMessage: string, setInputMessage: (inputMessage: string) => void, handleSendMessage: (inputMessage: string) => void}) {

    const inputRef = useRef<HTMLInputElement>(null);

    const sendMessageAndRefocus = () => {
        handleSendMessage(inputMessage);
        setInputMessage(''); // Clear the input after sending
        inputRef.current?.focus();
    }

    const handleKeyPress = (event: React.KeyboardEvent) => {
        if (event.key === 'Enter' && !event.shiftKey) {
          event.preventDefault();
          sendMessageAndRefocus();
        }
    };

    return (
        <Box className="chat-input-container" 
            sx={{
                maxWidth: '1200px',
                mx: 'auto',
                p: 3,
            }}
        >
            <Box display="flex" gap={2} alignItems="flex-end" mb={2}>
                <TextField
                    ref={inputRef}
                    fullWidth
                    multiline
                    maxRows={4}
                    value={inputMessage}
                    onChange={(e) => setInputMessage(e.target.value)}
                    onKeyDown={handleKeyPress}
                    placeholder="Type your fitness question here..."
                    variant="outlined"
                    sx={{
                        '& .MuiOutlinedInput-root': {
                            borderRadius: 3,
                            bgcolor: 'background.default',
                            fontSize: '1rem',
                            '& fieldset': {
                            borderColor: 'divider',
                            },
                            '&:hover fieldset': {
                            borderColor: 'primary.main',
                            },
                        },
                    }}
                />
                <IconButton
                    onClick={sendMessageAndRefocus}
                    disabled={!inputMessage.trim()}
                    color="primary"
                    sx={{
                        bgcolor: 'primary.main',
                        color: 'white',
                        '&:hover': {
                        bgcolor: 'primary.dark',
                        },
                        '&:disabled': {
                        bgcolor: 'action.disabled',
                        },
                        width: 56,
                        height: 56,
                    }}
                >
                    <SendIcon />
                </IconButton>
            </Box>
        </Box>
    );
}   