'use client';

import { useRequireAuth } from '../contexts/AuthContext';
import LandingPageComponent from '../components/Landing';
import { Box, CircularProgress, Typography } from '@mui/material';
import ChatApp from '../components/ChatPage';
import LoggedInLayout from '@/components/LoggedInLayout';

export default function HomePage() {
  const { isAuthenticated, isLoading } = useRequireAuth();

  // Show loading state while checking authentication
  if (isLoading) {
    return (
      <Box 
        sx={{ 
          display: 'flex', 
          justifyContent: 'center', 
          alignItems: 'center', 
          minHeight: '100vh',
          flexDirection: 'column',
          gap: 2
        }}
      >
        <CircularProgress />
        <Typography variant="body2" color="textSecondary">
          Loading...
        </Typography>
      </Box>
    );
  }

  // Show different content based on authentication status
  return (
    <Box>
      {isAuthenticated ? (
        <LoggedInLayout>
          <ChatApp />
        </LoggedInLayout>
      ) : (
        <LandingPageComponent />
      )}
    </Box>
  );
}
