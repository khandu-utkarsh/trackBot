'use client';

//!This is the global layout for the app. It is used to provide the theme and the providers to the app.

import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { lightTheme } from '../theme/theme';
import { Box, useTheme } from '@mui/material';
import { AuthProvider } from '../contexts/AuthContext';
import Header  from "@/components/Header";
import Sidebar from "@/components/Sidebar";
import Footer from "@/components/Footer";
import { useRequireAuth } from '../contexts/AuthContext';



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

function ChatAppLayout({children}: Readonly<{children: React.ReactNode;}>) {
  const theme = useTheme();

  return (
    <Box sx={{ display: 'flex', minHeight: '100vh', flexDirection: 'row', overflow: 'hidden' }}>
      {/* Sidebar */}
      <Sidebar />
      {/* Main Area */}
      <Box sx={{ flex: 1, display: 'flex', flexDirection: 'column', overflow: 'hidden' }}>
        <Header />
        <Box className="app-layout-main-container" component="main" 
             sx={{ flex: 1, 
             overflow: 'auto', 
             display: 'flex', 
             flexDirection: 'column', 
             height: '100vh', 
             width: '100%',
             justifyContent: 'center'}}>
          {children}
        </Box>
        <Footer />
      </Box>
    </Box>
  );
}

const AuthenticatedLayout = ({ children }: { children: React.ReactNode }) => {
  const { isAuthenticated } = useRequireAuth();

  if (!isAuthenticated) {
    return (
      <>
      {children}
      </>
    );
  }
  return (
    <ChatAppLayout>
      {children}
    </ChatAppLayout>
  );
};


//!This is the main layout for entire website. All other layouts will be nested inside this one.
export default function RootLayout({  children,}: Readonly<{  children: React.ReactNode;}>) {
  return (
    <html lang="en">
      <body>
        <ThemeProvider theme={lightTheme}>
          <CssBaseline />
          <AuthProvider>
            <ThemedContainer>
              <AuthenticatedLayout>
                {children}
              </AuthenticatedLayout>
            </ThemedContainer>
          </AuthProvider>
        </ThemeProvider>
      </body>
    </html>
  );
}
