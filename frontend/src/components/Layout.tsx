import { AppBar, Toolbar, Typography, Button, Container } from '@mui/material';
import { FitnessCenter, Timeline, Person } from '@mui/icons-material';
import Link from 'next/link';

interface LayoutProps {
  children: React.ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  return (
    <>
      <AppBar position="static">
        <Toolbar>
          <FitnessCenter sx={{ mr: 2 }} />
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Workout Tracker
          </Typography>
          <Button color="inherit" component={Link} href="/">
            <Timeline sx={{ mr: 1 }} />
            Workouts
          </Button>
          <Button color="inherit" component={Link} href="/profile">
            <Person sx={{ mr: 1 }} />
            Profile
          </Button>
        </Toolbar>
      </AppBar>
      <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
        {children}
      </Container>
    </>
  );
} 