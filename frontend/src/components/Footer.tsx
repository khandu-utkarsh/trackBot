'use client';

import { Box, Container, Typography, Link, IconButton, useTheme } from '@mui/material';
import GitHubIcon from '@mui/icons-material/GitHub';
import LinkedInIcon from '@mui/icons-material/LinkedIn';
import TwitterIcon from '@mui/icons-material/Twitter';

export default function Footer() {
  const theme = useTheme();
  const currentYear = new Date().getFullYear();

  return (
    <Box
      component="footer"
      sx={{
        py: 2,
        px: 2,
        backgroundColor: theme.palette.background.paper,
        borderTop: `1px solid ${theme.palette.divider}`,

        flexShrink: 0
      }}
    >
      <Container maxWidth="lg">
        <Box
          sx={{
            display: 'flex',
            flexDirection: { xs: 'column', sm: 'row' },
            justifyContent: 'space-between',
            alignItems: 'center',
            gap: 2,
          }}
        >
          {/* Copyright */}
          <Typography variant="body2" color="text.secondary">
            © {currentYear} Workout Tracker. All rights reserved.
          </Typography>

          {/* Links */}
          <Box
            sx={{
              display: 'flex',
              gap: 3,
              flexWrap: 'wrap',
              justifyContent: 'center',
            }}
          >
            <Link href="/about" color="text.secondary" underline="hover">
              About
            </Link>
            <Link href="/privacy" color="text.secondary" underline="hover">
              Privacy
            </Link>
            <Link href="/terms" color="text.secondary" underline="hover">
              Terms
            </Link>
          </Box>

          {/* Social Media Icons */}
          <Box sx={{ display: 'flex', gap: 1 }}>
            <IconButton
              href="https://github.com"
              target="_blank"
              rel="noopener noreferrer"
              color="inherit"
              size="small"
            >
              <GitHubIcon />
            </IconButton>
            <IconButton
              href="https://linkedin.com"
              target="_blank"
              rel="noopener noreferrer"
              color="inherit"
              size="small"
            >
              <LinkedInIcon />
            </IconButton>
            <IconButton
              href="https://twitter.com"
              target="_blank"
              rel="noopener noreferrer"
              color="inherit"
              size="small"
            >
              <TwitterIcon />
            </IconButton>
          </Box>
        </Box>
      </Container>
    </Box>
  );
}
