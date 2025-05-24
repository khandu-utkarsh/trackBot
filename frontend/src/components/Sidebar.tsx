'use client';

import { useState } from 'react';
import {
  Box,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Divider,
  IconButton,
  useTheme,
  Avatar,
  Typography,
} from '@mui/material';
import {
  Dashboard as DashboardIcon,
  Chat as ChatIcon,
  FitnessCenter as WorkoutIcon,
  Restaurant as NutritionIcon,
  Settings as SettingsIcon,
  Menu as MenuIcon,
  ChevronLeft as ChevronLeftIcon,
} from '@mui/icons-material';
import { useSession } from 'next-auth/react';
import { useRouter, usePathname } from 'next/navigation';

const DRAWER_WIDTH = 240;

const menuItems = [
  { text: 'Dashboard', icon: <DashboardIcon />, path: '/dashboard' },
  { text: 'AI Chat', icon: <ChatIcon />, path: '/chat' },
  { text: 'Workouts', icon: <WorkoutIcon />, path: '/workouts' },
  { text: 'Nutrition', icon: <NutritionIcon />, path: '/nutrition' },
  { text: 'Settings', icon: <SettingsIcon />, path: '/settings' },
];

export default function Sidebar() {
  const [open, setOpen] = useState(false);
  const theme = useTheme();
  const { data: session } = useSession();
  const router = useRouter();
  const pathname = usePathname();

  const handleDrawerToggle = () => {
    setOpen(!open);
  };

  const handleNavigation = (path: string) => {
    router.push(path);
    // Close sidebar after navigation on mobile
    if (typeof window !== 'undefined' && window.innerWidth < 600) {
      setOpen(false);
    }
  };

  return (
    <>
      {/* Mobile and Desktop Drawer - Now using temporary for both */}
      <Drawer
        variant="temporary"
        open={open}
        onClose={handleDrawerToggle}
        ModalProps={{
          keepMounted: true, // Better open performance on mobile
        }}
        sx={{
          '& .MuiDrawer-paper': {
            boxSizing: 'border-box',
            width: DRAWER_WIDTH,
            backgroundColor: theme.palette.background.paper,
            borderRight: `1px solid ${theme.palette.divider}`,
          },
        }}
      >
        <DrawerContent />
      </Drawer>
      
      {/* Floating Menu Button */}
      <IconButton
        onClick={handleDrawerToggle}
        sx={{
          position: 'fixed',
          top: 80, // Below header
          left: 16,
          zIndex: theme.zIndex.speedDial,
          bgcolor: 'primary.main',
          color: 'white',
          '&:hover': {
            bgcolor: 'primary.dark',
          },
          boxShadow: 3,
        }}
      >
        <MenuIcon />
      </IconButton>
    </>
  );

  function DrawerContent() {
    return (
      <>
        {/* Drawer Header */}
        <Box
          sx={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between',
            p: 2,
          }}
        >
          <Typography variant="h6" noWrap component="div">
            Menu
          </Typography>
          <IconButton onClick={handleDrawerToggle}>
            <ChevronLeftIcon />
          </IconButton>
        </Box>

        <Divider />

        {/* User Profile Section */}
        {session?.user && (
          <Box
            sx={{
              p: 2,
              display: 'flex',
              alignItems: 'center',
              gap: 2,
            }}
          >
            <Avatar
              src={session.user.image ?? undefined}
              alt={session.user.name ?? 'User'}
            />
            <Box>
              <Typography variant="subtitle1" noWrap>
                {session.user.name}
              </Typography>
              <Typography variant="body2" color="text.secondary" noWrap>
                {session.user.email}
              </Typography>
            </Box>
          </Box>
        )}

        <Divider />

        {/* Navigation Items */}
        <List>
          {menuItems.map((item) => (
            <ListItem key={item.text} disablePadding>
              <ListItemButton
                selected={pathname === item.path}
                onClick={() => handleNavigation(item.path)}
                sx={{
                  '&.Mui-selected': {
                    backgroundColor: theme.palette.action.selected,
                  },
                }}
              >
                <ListItemIcon
                  sx={{
                    color: pathname === item.path ? 'primary.main' : 'inherit',
                  }}
                >
                  {item.icon}
                </ListItemIcon>
                <ListItemText primary={item.text} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>
      </>
    );
  }
}



