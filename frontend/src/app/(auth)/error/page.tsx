'use client';

import { Box, Container, Typography, Button } from '@mui/material';
import { useSearchParams } from 'next/navigation';
import Link from 'next/link';

export default function AuthError() {
  const searchParams = useSearchParams();
  const error = searchParams.get('error');

  return (
    <Container maxWidth="sm">
      <Box
        sx={{
          minHeight: '100vh',
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
          gap: 3,
          textAlign: 'center'
        }}
      >
        <Typography variant="h4" component="h1" gutterBottom>
          Authentication Error
        </Typography>
        
        <Typography color="error" gutterBottom>
          {error === 'AccessDenied' 
            ? 'You do not have permission to sign in.'
            : 'An error occurred during authentication.'}
        </Typography>

        <Button
          component={Link}
          href="/landing"
          variant="contained"
          size="large"
        >
          Return to Home
        </Button>
      </Box>
    </Container>
  );
} 