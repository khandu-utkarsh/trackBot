'use client';

import { Box, Typography, Button, Container } from '@mui/material';
import Link from 'next/link';
import GoogleIcon from '@mui/icons-material/Google';
import { signIn } from "next-auth/react";

export default function LandingPage() {
  return (
    <Container maxWidth="md">
      <Box sx={{ 
        minHeight: '100vh',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        gap: 1,
        py: 8
      }}>
        <Typography variant="h1" component="h1">
          Tracker
        </Typography>

        <Typography variant="h5" color="text.secondary" >
          Track everything in natural language.
        </Typography>

        <Box sx={{ 
          display: 'flex', 
          gap: 2, 
          mt: 2,
          justifyContent: 'center',
          width: '100%'
        }}>
          <Button
            variant="outlined"
            size="large"
            startIcon={<GoogleIcon />}
            onClick={() => signIn("google", { callbackUrl: "/dashboard" })}
          >
            Login with Google
          </Button>
        </Box>
      </Box>
    </Container>
  );
}
