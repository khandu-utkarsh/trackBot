'use client';

import { useState } from 'react';
import {
  Box,
  Button,
  Typography,
  Paper,
  Alert,
  CircularProgress,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
} from '@mui/material';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import ErrorIcon from '@mui/icons-material/Error';
import PlayArrowIcon from '@mui/icons-material/PlayArrow';
import { testClient } from '@/lib/api/test-client';

interface TestResult {
  health: boolean;
  noAuth: boolean;
  protected: boolean;
  message: boolean;
}

export default function TestAuthPage() {
  const [isRunning, setIsRunning] = useState(false);
  const [results, setResults] = useState<TestResult | null>(null);
  const [logs, setLogs] = useState<string[]>([]);

  const runTests = async () => {
    setIsRunning(true);
    setResults(null);
    setLogs([]);

    // Capture console logs
    const originalLog = console.log;
    const originalError = console.error;
    const capturedLogs: string[] = [];

    console.log = (...args) => {
      const message = args.join(' ');
      capturedLogs.push(message);
      setLogs([...capturedLogs]);
      originalLog(...args);
    };

    console.error = (...args) => {
      const message = `ERROR: ${args.join(' ')}`;
      capturedLogs.push(message);
      setLogs([...capturedLogs]);
      originalError(...args);
    };

    try {
      const testResults = await testClient.runAllTests();
      setResults(testResults);
    } catch (error) {
      console.error('Test execution failed:', error);
    } finally {
      // Restore original console methods
      console.log = originalLog;
      console.error = originalError;
      setIsRunning(false);
    }
  };

  const getTestDescription = (testName: string) => {
    const descriptions = {
      health: 'Health endpoint (no auth required)',
      noAuth: 'Protected endpoint without auth (should fail)',
      protected: 'Protected endpoint with dummy token',
      message: 'Message creation with dummy token',
    };
    return descriptions[testName as keyof typeof descriptions] || testName;
  };

  return (
    <Box sx={{ maxWidth: 800, mx: 'auto', p: 3 }}>
      <Typography variant="h4" gutterBottom>
        🧪 Authentication Flow Test
      </Typography>
      
      <Alert severity="info" sx={{ mb: 3 }}>
        This page tests the authentication flow using dummy tokens. 
        Make sure your Go backend is running on localhost:8080.
      </Alert>

      <Paper sx={{ p: 3, mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 2, mb: 2 }}>
          <Button
            variant="contained"
            startIcon={isRunning ? <CircularProgress size={20} /> : <PlayArrowIcon />}
            onClick={runTests}
            disabled={isRunning}
            size="large"
          >
            {isRunning ? 'Running Tests...' : 'Run Authentication Tests'}
          </Button>
        </Box>

        {results && (
          <Box sx={{ mt: 3 }}>
            <Typography variant="h6" gutterBottom>
              Test Results:
            </Typography>
            <List>
              {Object.entries(results).map(([testName, passed]) => (
                <ListItem key={testName}>
                  <ListItemIcon>
                    {passed ? (
                      <CheckCircleIcon color="success" />
                    ) : (
                      <ErrorIcon color="error" />
                    )}
                  </ListItemIcon>
                  <ListItemText
                    primary={getTestDescription(testName)}
                    secondary={passed ? 'PASSED' : 'FAILED'}
                  />
                </ListItem>
              ))}
            </List>
            
            <Alert 
              severity={Object.values(results).every(Boolean) ? 'success' : 'warning'}
              sx={{ mt: 2 }}
            >
              {Object.values(results).every(Boolean) 
                ? '🎉 All tests passed! Authentication flow is working correctly.'
                : '⚠️ Some tests failed. Check the logs below for details.'
              }
            </Alert>
          </Box>
        )}
      </Paper>

      {logs.length > 0 && (
        <Paper sx={{ p: 3 }}>
          <Typography variant="h6" gutterBottom>
            Test Logs:
          </Typography>
          <Box
            sx={{
              bgcolor: 'grey.100',
              p: 2,
              borderRadius: 1,
              fontFamily: 'monospace',
              fontSize: '0.875rem',
              maxHeight: 400,
              overflow: 'auto',
            }}
          >
            {logs.map((log, index) => (
              <Box key={index} sx={{ mb: 0.5 }}>
                {log}
              </Box>
            ))}
          </Box>
        </Paper>
      )}

      <Paper sx={{ p: 3, mt: 3 }}>
        <Typography variant="h6" gutterBottom>
          What This Tests:
        </Typography>
        <List dense>
          <ListItem>
            <ListItemText primary="✅ Health endpoint works without authentication" />
          </ListItem>
          <ListItem>
            <ListItemText primary="✅ Protected endpoints reject requests without tokens" />
          </ListItem>
          <ListItem>
            <ListItemText primary="✅ Protected endpoints accept requests with dummy tokens" />
          </ListItem>
          <ListItem>
            <ListItemText primary="✅ Message creation triggers LLM service integration" />
          </ListItem>
        </List>
      </Paper>
    </Box>
  );
} 