'use client';

import { AppBar, Toolbar, Typography, Box, Tabs, Tab, Avatar, Card, CardContent, Button, IconButton, useTheme, CssBaseline, Grid, Tooltip } from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import ChatIcon from '@mui/icons-material/Chat';
import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Header from '@/components/Header';
import MonthlyContributionCalendar from '@/components/ContributionCalendar';

function DashboardCard({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <Card sx={{ bgcolor: 'background.paper', borderRadius: 3, boxShadow: 3, mb: 3 }}>
      <CardContent>
        <Typography variant="h6" color="text.primary" gutterBottom>
          {title}
        </Typography>
        {children}
      </CardContent>
    </Card>
  );
}

function ActivityCalendar() {
  return (
    <Box display="flex" flexDirection="row" gap={2}>
      <MonthlyContributionCalendar year={2025} month={0} />
      <MonthlyContributionCalendar year={2025} month={1} />
      <MonthlyContributionCalendar year={2025} month={2} />
      <MonthlyContributionCalendar year={2025} month={3} />
      <MonthlyContributionCalendar year={2025} month={4} />
      <MonthlyContributionCalendar year={2025} month={5} />
      <MonthlyContributionCalendar year={2025} month={6} />
      <MonthlyContributionCalendar year={2025} month={7} />
      <MonthlyContributionCalendar year={2025} month={8} />
      <MonthlyContributionCalendar year={2025} month={9} />
      <MonthlyContributionCalendar year={2025} month={10} />
      <MonthlyContributionCalendar year={2025} month={11} />
    </Box>
  );
}

function AIAssistantCard() {
  const router = useRouter();
  
  return (
    <Card sx={{ 
      bgcolor: 'background.paper', 
      borderRadius: 3, 
      boxShadow: 3, 
      mb: 3,
      background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
      color: 'white'
    }}>
      <CardContent>
        <Box display="flex" alignItems="center" gap={2} mb={2}>
          <SmartToyIcon sx={{ fontSize: 32 }} />
          <Box>
            <Typography variant="h6" gutterBottom sx={{ mb: 0 }}>
              AI Fitness Assistant
            </Typography>
            <Typography variant="body2" sx={{ opacity: 0.9 }}>
              Get personalized workout advice and nutrition tips
            </Typography>
          </Box>
        </Box>
        
        <Typography variant="body2" sx={{ mb: 3, opacity: 0.9 }}>
          Ask questions about workouts, nutrition, and fitness goals. Get instant, personalized advice from our AI assistant.
        </Typography>
        
        <Box display="flex" gap={2}>
          <Button
            variant="contained"
            startIcon={<ChatIcon />}
            onClick={() => router.push('/chat')}
            sx={{
              bgcolor: 'rgba(255, 255, 255, 0.2)',
              color: 'white',
              '&:hover': {
                bgcolor: 'rgba(255, 255, 255, 0.3)',
              },
              backdropFilter: 'blur(10px)',
            }}
          >
            Start Chat
          </Button>
          <Button
            variant="outlined"
            sx={{
              borderColor: 'rgba(255, 255, 255, 0.5)',
              color: 'white',
              '&:hover': {
                borderColor: 'white',
                bgcolor: 'rgba(255, 255, 255, 0.1)',
              },
            }}
          >
            Learn More
          </Button>
        </Box>
      </CardContent>
    </Card>
  );
}

export default function DashboardPage() {
  const [tab, setTab] = useState(0);
  const theme = useTheme();

  return (
    <Box sx={{ 
      display: 'flex', 
      flexDirection: 'column', 
      minHeight: '100vh',
      position: 'relative'
    }}>
      <Box sx={{ 
        flex: 1,
        mt: '64px', // Height of the header
        mb: '64px', // Height of the footer
        overflow: 'auto',
        height: 'calc(100vh - 128px)', // Viewport height minus header and footer
        padding: 3
      }}>
        <AIAssistantCard />
        
        <DashboardCard title="Activity Calendar">
          <ActivityCalendar />
        </DashboardCard>
      </Box>

      <Box sx={{ 
        position: 'fixed',
        bottom: 0,
        left: 0,
        right: 0,
        zIndex: theme.zIndex.appBar
      }}>

      </Box>
    </Box>
  );
} 