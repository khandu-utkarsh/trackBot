'use client';

import { Box, Typography, Button } from '@mui/material';
import Script from 'next/script';

declare global {
  interface Window {
    google: any;
  }
}

export default function LandingPageComponent() {
  const handleCredentialResponse = (response: any) => {
    console.log('Google Auth Response:', response);
    localStorage.setItem('google_token', response.credential);
    window.dispatchEvent(new CustomEvent('auth-changed', {
      detail: { authenticated: true, token: response.credential }
    }));
  };

  const handleGoogleSignIn = () => {
    console.log('Google SDK available:', !!window.google);    
    if (window.google && window.google.accounts) {
      try {
        // Try the standard prompt first
        window.google.accounts.id.prompt((notification: any) => {
          console.log('Prompt notification:', notification);
          if (notification.isNotDisplayed() || notification.isSkippedMoment()) {
            console.log('Prompt not displayed, trying renderButton method');
            // Fallback: Create a temporary button element for Google to render
            const tempDiv = document.createElement('div');
            tempDiv.style.display = 'none';
            document.body.appendChild(tempDiv);
            
            window.google.accounts.id.renderButton(tempDiv, {
              type: 'standard',
              size: 'large',
              text: 'signin_with',
              theme: 'outline',
            });
            
            // Trigger the button click
            const googleButton = tempDiv.querySelector('div[role="button"]') as HTMLElement;
            if (googleButton) {
              googleButton.click();
            }
            
            // Clean up
            setTimeout(() => {
              document.body.removeChild(tempDiv);
            }, 1000);
          }
        });
      } catch (error) {
        console.error('Error with Google Sign-In:', error);
      }
    } else {
      console.error('Google SDK not loaded or accounts not available');
    }
  };

  return (
    <>
      <Script
        src="https://accounts.google.com/gsi/client"
        strategy="afterInteractive"
        onLoad={() => {
          console.log('Google Client ID:', process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID);
          const clientId = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID?.trim();
          
          if (!clientId || clientId === 'undefined') {
            console.error('Google Client ID is not set or invalid');
            return;
          }
          
          if (window.google && window.google.accounts) {
            try {
              // Only initialize, don't auto-prompt
              window.google.accounts.id.initialize({
                client_id: clientId,
                callback: handleCredentialResponse,
                auto_select: false,
                cancel_on_tap_outside: true,
                itp_support: true, // Intelligent Tracking Prevention support
              });
              
              // Disable automatic prompts
              window.google.accounts.id.disableAutoSelect();
              console.log('Google SDK initialized successfully');
            } catch (error) {
              console.error('Error initializing Google SDK:', error);
            }
          } else {
            console.error('Google SDK not available');
          }
        }}
        onError={(error) => {
          console.error('Failed to load Google SDK:', error);
        }}
      />
      <Box sx={{ 
        minHeight: '100vh',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        textAlign: 'center',
        py: 8
      }}>
        <Typography variant="h1" component="h1" gutterBottom sx={{ fontWeight: 'bold' }}>
          TrackBot
        </Typography>
        <Button
          variant="contained"
          size="large"
          onClick={handleGoogleSignIn}
          sx={{ px: 4, py: 1.5 }}
        >
          Sign in with Google
        </Button>
      </Box>
    </>
  );
}
