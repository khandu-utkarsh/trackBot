//!This is the main layout for entire website. All other layouts will be nested inside this one.
'use client';

import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { lightTheme } from '../theme/theme';
import { Box, useTheme } from '@mui/material';
import { AuthProvider } from '../contexts/AuthContext';



function ThemedContainer({ children }: { children: React.ReactNode }) {
    const theme = useTheme();
    return (
      <Box
        sx={{
          bgcolor: theme.palette.background.default,
          minHeight: '100vh',
          maxWidth: theme.layout.maxWidth,
          mx: 'auto',
        }}
      >
        {children}
      </Box>
    );
  }

export default function RootLayout({  children,}: Readonly<{  children: React.ReactNode;}>) {
    return (
      <html lang="en">
        <body>
          <ThemeProvider theme={lightTheme}>
            <CssBaseline />
            <AuthProvider>
              <ThemedContainer>
                  {children}
              </ThemedContainer>
            </AuthProvider>
          </ThemeProvider>
        </body>
      </html>
    );
  }
  