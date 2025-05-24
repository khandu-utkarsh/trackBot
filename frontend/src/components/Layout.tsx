import { Box } from '@mui/material';
import Header from './Header';

interface LayoutProps {
  children: React.ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  return (
    <Box sx={{ 
      display: 'flex', 
      flexDirection: 'column', 
      minHeight: '100vh',
      position: 'relative'
    }}>
      <Header />
      <Box sx={{ 
        flex: 1,
        mt: '64px', // Height of the header
        overflow: 'auto',
        height: 'calc(100vh - 64px)', // Viewport height minus header
        padding: 3
      }}>
        {children}
      </Box>
    </Box>
  );
} 