import { useState, useEffect } from 'react';

interface GoogleUser {
  email: string;
  name: string;
  picture: string;
}

export const useGoogleAuth = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState<GoogleUser | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Check for existing authentication on mount
    const token = localStorage.getItem('google_token');
    if (token) {
      try {
        // Decode JWT token to get user info
        const payload = JSON.parse(atob(token.split('.')[1]));
        setUser({
          email: payload.email,
          name: payload.name,
          picture: payload.picture,
        });
        setIsAuthenticated(true);
      } catch (error) {
        console.error('Error parsing token:', error);
        localStorage.removeItem('google_token');
      }
    }
    setIsLoading(false);

    // Listen for authentication changes
    const handleAuthChange = (event: any) => {
      if (event.detail.authenticated) {
        const token = event.detail.token;
        try {
          const payload = JSON.parse(atob(token.split('.')[1]));
          setUser({
            email: payload.email,
            name: payload.name,
            picture: payload.picture,
          });
          setIsAuthenticated(true);
        } catch (error) {
          console.error('Error parsing token:', error);
        }
      } else {
        setUser(null);
        setIsAuthenticated(false);
      }
    };

    window.addEventListener('auth-changed', handleAuthChange);

    return () => {
      window.removeEventListener('auth-changed', handleAuthChange);
    };
  }, []);

  const signOut = () => {
    localStorage.removeItem('google_token');
    setUser(null);
    setIsAuthenticated(false);
    
    // Trigger auth change event
    window.dispatchEvent(new CustomEvent('auth-changed', { 
      detail: { authenticated: false } 
    }));
  };

  return {
    isAuthenticated,
    user,
    isLoading,
    signOut,
  };
}; 