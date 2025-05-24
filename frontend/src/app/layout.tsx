'use client';

//!This is the global layout for the app. It is used to provide the theme and the providers to the app.


import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { lightTheme } from '../theme/theme';
import { Providers } from '../components/providers';
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

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <Providers>          
          <ThemeProvider theme={lightTheme}>
            <CssBaseline />
            <ThemedContainer>
              {children}
            </ThemedContainer>
          </ThemeProvider>
        </Providers>
      </body>
    </html>
  );
}
