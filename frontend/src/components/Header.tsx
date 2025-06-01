'use client';

import React from 'react';
import { 
  AppBar, 
  Toolbar, 
  Typography, 
  Button, 
  Avatar, 
  Box, 
  Menu, 
  MenuItem,
  IconButton
} from '@mui/material';
import { useAuth } from '../contexts/AuthContext';
import { useRouter } from 'next/navigation';

export default function Header() {
  const { isAuthenticated, user, signOut, isLoading } = useAuth();
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

  const router = useRouter();
  const handleMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleSignOut = () => {
    signOut();
    handleMenuClose();
  };

  const handleDashboard = () => {
    router.push('/dashboard');
    handleMenuClose();
  };


  if (isLoading) {
    return (
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            TrackBot
          </Typography>
          <Typography>Loading...</Typography>
        </Toolbar>
      </AppBar>
    );
  }

  return (
    <AppBar position="static">
      <Toolbar>        
        {user ? (
          <Box sx={{ display: 'flex', flex: 1, alignItems: 'center', justifyContent: 'flex-end', gap: 2 }}>
            <Typography variant="body2">
              Welcome, {user.name}
            </Typography>
            <IconButton
              onClick={handleMenuOpen}
              sx={{ p: 0 }}
            >
              <Avatar 
                src={user.picture} 
                alt={user.name}
                sx={{ width: 32, height: 32 }}
              />
            </IconButton>
            <Menu
              anchorEl={anchorEl}
              open={Boolean(anchorEl)}
              onClose={handleMenuClose}
              transformOrigin={{ horizontal: 'right', vertical: 'top' }}
              anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
            >
              <MenuItem onClick={handleDashboard}>
                Dashboard
              </MenuItem>
              <MenuItem onClick={handleSignOut}>
                Sign Out
              </MenuItem>
            </Menu>
          </Box>
        ) : (
          <Button color="inherit">
            Sign In
          </Button>
        )}
      </Toolbar>
    </AppBar>
  );
};
