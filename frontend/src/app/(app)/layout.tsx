'use client';

import Header from "@/components/Header";
import Sidebar from "@/components/Sidebar";
import Footer from "@/components/Footer";
import { Box, CssBaseline, Toolbar } from '@mui/material';

//!This is the layout for the app. It will contain the header and the sidebar and the footer.

//!Header will be a simple navbar
//!Sidebar needs to be a collection of all the chats
//!Footer will be a simple line with some basic info


export default function LoggedInAppLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <Box sx={{ display: 'flex', minHeight: '100vh', flexDirection: 'column' }}>
      <CssBaseline />
      <Header />
      <Box sx={{ display: 'flex', flex: 1 }}>
        <Sidebar />
        <Box
          component="main"
          sx={{
            flexGrow: 1,
            width: '100%', // Full width since sidebar overlays
            mt: '64px', // Height of the header
            minHeight: 'calc(100vh - 64px)',
            display: 'flex',
            flexDirection: 'column'
          }}
        >
          {children}
        </Box>
      </Box>
      <Footer />
    </Box>
  );
}