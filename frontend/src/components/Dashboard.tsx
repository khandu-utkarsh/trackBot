'use client';

import React from 'react';
import { 
  Box, 
  Typography, 
  Card, 
  CardContent, 
  Avatar, 
  Chip, 
  Button,
  Paper
} from '@mui/material';
import { useAuth } from '../contexts/AuthContext';

export const Dashboard: React.FC = () => {
  const { user, isAuthenticated, signOut, token } = useAuth();

  if (!isAuthenticated || !user) {
    return (
      <Box sx={{ p: 4 }}>
        <Typography variant="h4">
          Please sign in to access the dashboard
        </Typography>
      </Box>
    );
  }

  return (
    <Box sx={{ p: 4 }}>
      <Typography variant="h3" gutterBottom>
        Dashboard
      </Typography>
      
      <Box sx={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
        {/* User Info and Session Cards Row */}
        <Box sx={{ 
          display: 'flex', 
          gap: 3, 
          flexDirection: { xs: 'column', md: 'row' } 
        }}>
          {/* User Info Card */}
          <Box sx={{ flex: 1 }}>
            <Card>
              <CardContent>
                <Box sx={{ display: 'flex', alignItems: 'center', gap: 2, mb: 2 }}>
                  <Avatar 
                    src={user.picture} 
                    alt={user.name}
                    sx={{ width: 64, height: 64 }}
                  />
                  <Box>
                    <Typography variant="h5">
                      {user.name}
                    </Typography>
                    <Typography variant="body2" color="textSecondary">
                      {user.email}
                    </Typography>
                    <Chip 
                      label="Authenticated" 
                      color="success" 
                      size="small" 
                      sx={{ mt: 1 }}
                    />
                  </Box>
                </Box>
                <Button 
                  variant="outlined" 
                  color="error" 
                  onClick={signOut}
                  fullWidth
                >
                  Sign Out
                </Button>
              </CardContent>
            </Card>
          </Box>

          {/* Session Info Card */}
          <Box sx={{ flex: 1 }}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  Session Information
                </Typography>
                <Paper sx={{ p: 2, bgcolor: 'grey.50' }}>
                  <Typography variant="body2" component="div">
                    <strong>Token Available:</strong> {token ? 'Yes' : 'No'}
                  </Typography>
                  <Typography variant="body2" component="div" sx={{ mt: 1 }}>
                    <strong>Authentication Status:</strong> {isAuthenticated ? 'Authenticated' : 'Not Authenticated'}
                  </Typography>
                  {token && (
                    <Typography variant="body2" component="div" sx={{ mt: 1 }}>
                      <strong>Token (first 20 chars):</strong> {token.substring(0, 20)}...
                    </Typography>
                  )}
                </Paper>
              </CardContent>
            </Card>
          </Box>
        </Box>

        {/* Workout Stats Placeholder */}
        <Box>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Your Workout Stats
              </Typography>
              <Typography variant="body2" color="textSecondary">
                This is where your workout tracking data would go. 
                The authentication context makes it easy to:
              </Typography>
              <Box component="ul" sx={{ mt: 2 }}>
                <li>Access user information across all components</li>
                <li>Make authenticated API calls using the token</li>
                <li>Implement role-based access control</li>
                <li>Track user sessions and preferences</li>
              </Box>
            </CardContent>
          </Card>
        </Box>
      </Box>
    </Box>
  );
}; 