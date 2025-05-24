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
  const theme = useTheme();
  const { data: session } = useSession();
  const router = useRouter();
  const pathname = usePathname();

  const handleNavigation = (path: string) => {
    router.push(path);
  };

  return (
    <Drawer
      variant="permanent"
      sx={{
        width: DRAWER_WIDTH,
        flexShrink: 0,
        '& .MuiDrawer-paper': {
          width: DRAWER_WIDTH,
          boxSizing: 'border-box',
          backgroundColor: theme.palette.background.paper,
          borderRight: `1px solid ${theme.palette.divider}`,
        },
      }}
      open
    >
      <DrawerContent />
    </Drawer>
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



