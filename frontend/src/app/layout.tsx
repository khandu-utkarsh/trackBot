'use client';

//!This is the global layout for the app. It is used to provide the theme and the providers to the app.

import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { lightTheme } from '../theme/theme';
import { Box, useTheme } from '@mui/material';

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

//!This is the main layout for entire website. All other layouts will be nested inside this one.
export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <ThemeProvider theme={lightTheme}>
          <CssBaseline />
          <ThemedContainer>
            {children}
          </ThemedContainer>
        </ThemeProvider>
      </body>
    </html>
  );
}
