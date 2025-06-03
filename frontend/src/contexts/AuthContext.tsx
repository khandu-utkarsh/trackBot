'use client';

import React, { createContext, useContext, ReactNode, useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { Conversation } from '@/lib/types/chat';
import {User} from '@/lib/types/users';

interface AuthContextType {
  isAuthenticated: boolean;
  user: User | null;
  isLoading: boolean;
  signOut: () => void;
  conversations: Map<number, Conversation>;
  setConversations: (conversations: Map<number, Conversation>) => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);


export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [conversations, setConversations] = useState<Map<number, Conversation>>(new Map());

  // Check auth status on mount by calling backend
  useEffect(() => {
    const checkAuth = async () => {
      try {
        const response = await fetch('/api/auth/me', {
          credentials: 'include' // Include cookies
        });
        
        if (response.ok) {
          const userData = await response.json();
          setUser({
            id: userData.user_id,
            email: userData.email,
            name: userData.name,
            picture: userData.picture
          });
          setIsAuthenticated(true);
        }
      } catch (error) {
        console.error('Auth check failed:', error);
      } finally {
        setIsLoading(false);
      }
    };

    checkAuth();

    // Listen for auth changes
    const handleAuthChange = (event: Event) => {
      const customEvent = event as CustomEvent<{ authenticated: boolean; user: User }>;
      if (customEvent.detail.authenticated) {
        setUser(customEvent.detail.user);
        setIsAuthenticated(true);
      } else {
        setUser(null);
        setIsAuthenticated(false);
      }
    };

    window.addEventListener('auth-changed', handleAuthChange as EventListener);
    return () => window.removeEventListener('auth-changed', handleAuthChange as EventListener);
  }, []);

  const signOut = async () => {
    try {
      await fetch('/api/auth/logout', {
        method: 'POST',
        credentials: 'include'
      });
    } catch (error) {
      console.error('Logout error:', error);
    }
    
    setUser(null);
    setIsAuthenticated(false);
    window.dispatchEvent(new CustomEvent('auth-changed', { 
      detail: { authenticated: false, user: null } 
    }));
  };

  const value: AuthContextType = {
    isAuthenticated,
    user,
    isLoading,
    signOut,
    conversations,
    setConversations
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

// Hook to check if user is authenticated (useful for conditional rendering)
export const useRequireAuth = (): AuthContextType => {
  const auth = useAuth();
  const router = useRouter();  
  React.useEffect(() => {
    if (!auth.isLoading && !auth.isAuthenticated) {        
      //!Better to redirect to the home page. And it automatically takes care of the authentication to conditionally render the page.
      router.push('/');
      console.warn('User not authenticated');
    }
  }, [auth.isAuthenticated, auth.isLoading, router]); //!I need this since it is used by the hook. Some clouse thing with the react.

  return auth;
}; 

export const useConversations = (): [Map<number, Conversation>, (conversations: Map<number, Conversation>) => void] => {
  const auth = useAuth();
  return [auth.conversations, auth.setConversations];
};