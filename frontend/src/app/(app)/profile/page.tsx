'use client';

import { useState } from 'react';
import {
  Box,
  Card,
  CardContent,
  Typography,
  Grid,
  Avatar,
  TextField,
  Button,
  List,
  ListItem,
  ListItemText,
  Divider
} from '@mui/material';
import { Person as PersonIcon, FitnessCenter as FitnessIcon } from '@mui/icons-material';
import Layout from '@/app/(auth)/layout';

export default function Profile() {
  const [user, setUser] = useState({
    name: 'John Doe',
    email: 'john@example.com',
    joinDate: '2024-01-01',
    totalWorkouts: 0,
    totalDuration: 0,
    favoriteExercise: 'Bench Press'
  });

  return (
      <Grid container spacing={3}>
        <Grid item xs={12} md={4}>
          <Card>
            <CardContent sx={{ textAlign: 'center' }}>
              <Avatar
                sx={{
                  width: 100,
                  height: 100,
                  margin: '0 auto 20px',
                  bgcolor: 'primary.main'
                }}
              >
                <PersonIcon sx={{ fontSize: 50 }} />
              </Avatar>
              <Typography variant="h5" gutterBottom>
                {user.name}
              </Typography>
              <Typography color="text.secondary" gutterBottom>
                {user.email}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Member since {user.joinDate}
              </Typography>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={8}>
          <Card>
            <CardContent>
              <Typography variant="h6" gutterBottom>
                Statistics
              </Typography>
              <List>
                <ListItem>
                  <FitnessIcon sx={{ mr: 2 }} />
                  <ListItemText
                    primary="Total Workouts"
                    secondary={user.totalWorkouts}
                  />
                </ListItem>
                <Divider />
                <ListItem>
                  <FitnessIcon sx={{ mr: 2 }} />
                  <ListItemText
                    primary="Total Duration"
                    secondary={`${user.totalDuration} minutes`}
                  />
                </ListItem>
                <Divider />
                <ListItem>
                  <FitnessIcon sx={{ mr: 2 }} />
                  <ListItemText
                    primary="Favorite Exercise"
                    secondary={user.favoriteExercise}
                  />
                </ListItem>
              </List>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
  );
} 