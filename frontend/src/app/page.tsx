'use client';

import Layout from '@/components/Layout';
import { Box, Typography, Button, Paper } from '@mui/material';
import { FitnessCenter, Timeline, TrendingUp } from '@mui/icons-material';
import Link from 'next/link';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

export default function Home() {
  const { data: session, status } = useSession();
  const router = useRouter();

  useEffect(() => {
    if (status === 'authenticated') {
      router.push('/dashboard');
    }
    
    //!Default to landing page
    
    else if (status === 'unauthenticated') {
      router.push('/landing');
    }
  }, [status, router]);

  // Show loading state while checking authentication
  if (status === 'loading') {
    return (
      <Layout>
        <Box sx={{ 
          minHeight: '80vh',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center'
        }}>
          <Typography>Loading...</Typography>
        </Box>
      </Layout>
    );
  }

  // The rest of the component will only render briefly before redirect
  return (
    <Layout>
      <Box sx={{ 
        minHeight: '80vh',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        textAlign: 'center',
        gap: 4
      }}>
        {/* Hero Section */}
        <Box sx={{ maxWidth: 800, px: 2 }}>
          <FitnessCenter sx={{ fontSize: 60, color: 'primary.main', mb: 2 }} />
          <Typography variant="h2" component="h1" gutterBottom sx={{ fontWeight: 'bold' }}>
            Track Your Fitness Journey
          </Typography>
          <Typography variant="h5" color="text.secondary" sx={{ mb: 4 }}>
            A simple and effective way to monitor your workouts and achieve your fitness goals
          </Typography>
          <Button
            component={Link}
            href="/workouts"
            variant="contained"
            size="large"
            sx={{ px: 4, py: 1.5 }}
          >
            Get Started
          </Button>
        </Box>

        {/* Features Section */}
        <Box sx={{ 
          maxWidth: 1200, 
          mt: 8, 
          px: 2,
          display: 'flex',
          flexDirection: { xs: 'column', md: 'row' },
          gap: 4
        }}>
          <Paper elevation={0} sx={{ p: 3, flex: 1, textAlign: 'center' }}>
            <Timeline sx={{ fontSize: 40, color: 'primary.main', mb: 2 }} />
            <Typography variant="h6" gutterBottom>
              Track Workouts
            </Typography>
            <Typography color="text.secondary">
              Log your exercises and monitor your progress over time
            </Typography>
          </Paper>
          <Paper elevation={0} sx={{ p: 3, flex: 1, textAlign: 'center' }}>
            <FitnessCenter sx={{ fontSize: 40, color: 'primary.main', mb: 2 }} />
            <Typography variant="h6" gutterBottom>
              Multiple Exercise Types
            </Typography>
            <Typography color="text.secondary">
              Support for both cardio and weight training exercises
            </Typography>
          </Paper>
          <Paper elevation={0} sx={{ p: 3, flex: 1, textAlign: 'center' }}>
            <TrendingUp sx={{ fontSize: 40, color: 'primary.main', mb: 2 }} />
            <Typography variant="h6" gutterBottom>
              Progress Tracking
            </Typography>
            <Typography color="text.secondary">
              Visualize your improvement and stay motivated
            </Typography>
          </Paper>
        </Box>
      </Box>
    </Layout>
  );
}
