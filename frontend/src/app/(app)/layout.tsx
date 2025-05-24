'use client';

import Header from "@/components/Header";
import Sidebar from "@/components/Sidebar";
import Footer from "@/components/Footer";
import { Box, CssBaseline, Toolbar } from '@mui/material';

//!This is the layout for the app. It will contain the header and the sidebar and the footer.

//!Header will be a simple navbar
//!Sidebar needs to be a collection of all the chats
//!Footer will be a simple line with some basic info

const DRAWER_WIDTH = 240;

export default function LoggedInAppLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <Box sx={{ display: 'flex', minHeight: '100vh', flexDirection: 'row' }}>
      {/* Sidebar */}
      <Sidebar />
      {/* Main Area */}
      <Box sx={{ flex: 1, display: 'flex', flexDirection: 'column' }}>
        <Header />
        <Box component="main" sx={{ flex: 1, overflow: 'auto' }}>
          {children}
        </Box>
        <Footer />
      </Box>
    </Box>
  );
}