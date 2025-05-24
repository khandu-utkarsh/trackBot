'use client';

import { AppBar, Toolbar, Typography, Tabs, Tab, Box, Avatar, Menu, MenuItem, IconButton, useTheme } from '@mui/material';
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import { useSession, signOut } from 'next-auth/react';
import { useState } from 'react';
import Link from 'next/link';

const navLinks = [
  { label: 'Dashboard', href: '/dashboard' },
  { label: 'AI Chat', href: '/chat' },
//  { label: 'Workouts', href: '/workouts' },
 // { label: 'Nutrition', href: '/nutrition' },
  //{ label: 'Settings', href: '/settings' },
];

//!Component 1 -- Name and Logo
function Logo() {
  return (
    <Typography variant="h5" sx={{ flex: 1, minWidth: 0, fontWeight: 600 }}>
      Tracker
    </Typography>
  );
}

//!Component 2 -- Navigation Tabs
function NavigationTabs() {
  return (
    <Box sx={{ display: 'flex', alignItems: 'center', flex: 2, justifyContent: 'center' }}>
        <Tabs value={0}>
          {navLinks.map((link) => (
            <Tab
              key={link.label}
              label={link.label}
              component={Link}
              href={link.href}
              disableRipple
            />
          ))}
        </Tabs>
    </Box>
  );
}

//!Component 3 -- User Info Dashboard
function UserInfoDashboard() {
  const { data: session } = useSession();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

  const handleMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  // Since this component is only rendered in authenticated routes,
  // we can safely assert that session exists
  const user = session?.user;

  return (
    <Box sx={{ display: 'flex', alignItems: 'center', gap: 2, flex: 1, justifyContent: 'flex-end', minWidth: 0 }}>
      <Box sx={{ 
        display: 'flex',
        alignItems: 'center', 
        bgcolor: 'rgba(255, 255, 255, 0.1)', 
        borderRadius: '999px', 
        px: 1.5, 
        py: 0.5 
      }}>
        <Avatar src={user?.image ?? undefined} sx={{ width: 36, height: 36, mr: 1 }} />
        <Box sx={{ textAlign: 'left', mr: 1 }}>
          <Typography sx={{ color: 'white', fontWeight: 600, fontSize: 16 }}>
            {user?.name}
          </Typography>
          <Typography sx={{ color: 'rgba(255, 255, 255, 0.7)', fontSize: 13 }}>
            {user?.email}
          </Typography>
        </Box>
        <IconButton 
          size="small" 
          onClick={handleMenuOpen} 
          sx={{ 
            color: 'white',
            '&:hover': { bgcolor: 'rgba(255, 255, 255, 0.2)' }
          }}
        >
          <ArrowDropDownIcon />
        </IconButton>
        <Menu
          anchorEl={anchorEl}
          open={Boolean(anchorEl)}
          onClose={handleMenuClose}
          anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
          transformOrigin={{ vertical: 'top', horizontal: 'right' }}
          PaperProps={{
            sx: {
              bgcolor: '#18181b',
              color: 'white',
              '& .MuiMenuItem-root': {
                '&:hover': { bgcolor: 'rgba(255, 255, 255, 0.1)' }
              }
            }
          }}
        >
          <MenuItem onClick={() => { signOut(); handleMenuClose(); }}>Sign out</MenuItem>
        </Menu>
      </Box>
    </Box>
  );
}

export default function Header() {
  return (
    <>
      <AppBar position="static" color="transparent" sx={{ boxShadow: 'none', maxWidth: 'md'}}>
        <Toolbar sx={{ display: 'flex', alignItems: 'center', px: 2}}>
          <Logo />
          <NavigationTabs />
          <UserInfoDashboard />
        </Toolbar>
      </AppBar>
    </>
  );
}
