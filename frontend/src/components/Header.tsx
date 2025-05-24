'use client';

import { AppBar, Toolbar, Typography, Tabs, Tab, Box, Avatar, Menu, MenuItem, IconButton, useTheme } from '@mui/material';
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import { useSession, signOut } from 'next-auth/react';
import { useState } from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';

const navLinks = [
  { label: 'Dashboard', href: '/dashboard' },
  { label: 'AI Chat', href: '/chat' }
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
  const pathname = usePathname();
  
  // Find current tab index
  const currentTabIndex = navLinks.findIndex(link => pathname === link.href);
  const tabValue = currentTabIndex !== -1 ? currentTabIndex : false;

  return (
    <Box sx={{ display: 'flex', alignItems: 'center', flex: 2, justifyContent: 'center' }}>
      <Tabs 
        value={tabValue}
        sx={{
          '& .MuiTabs-indicator': {
            backgroundColor: 'white',
          },
          '& .MuiTab-root': {
            color: 'rgba(255, 255, 255, 0.8)',
            fontWeight: 500,
            '&.Mui-selected': {
              color: 'white',
            },
          },
        }}
      >
        {navLinks.map((link, index) => (
          <Tab
            key={link.label}
            label={link.label}
            component={Link}
            href={link.href}
            value={index}
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
  const theme = useTheme();

  return (
    <AppBar 
      position="fixed" 
      sx={{ 
        zIndex: theme.zIndex.drawer + 1,
        width: 'calc(100% - 240px)',
        ml: '240px',
        background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
      }}
    >
      <Toolbar sx={{ justifyContent: 'space-between', alignItems: 'center', gap: 2 }}>
        <Logo />
        <NavigationTabs />
        <UserInfoDashboard />
      </Toolbar>
    </AppBar>
  );
}
