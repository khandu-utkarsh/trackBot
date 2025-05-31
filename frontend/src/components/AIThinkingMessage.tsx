import { Box, Avatar, Paper, CircularProgress, Typography } from '@mui/material';
import SmartToyIcon from '@mui/icons-material/SmartToy';

export default function AIThinkingPrompt() {
    return (
        <Box sx={{ display: 'flex', gap: 2, alignItems: 'flex-start' }}>
            <Avatar sx={{ bgcolor: 'secondary.main', width: 48, height: 48 }}>
                <SmartToyIcon />
            </Avatar>
            <Paper
                elevation={2}
                sx={{
                    p: 2.5,
                    bgcolor: 'background.paper',
                    borderRadius: 3,
                    borderTopLeftRadius: 8,
                    display: 'flex',
                    alignItems: 'center',
                    gap: 2,
                }}
            >
                <CircularProgress size={20} />
                <Typography variant="body1" color="text.secondary">
                    AI is thinking...
                </Typography>
            </Paper>
        </Box>
    );
}
