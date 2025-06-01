'use client';

import React, { createContext, useContext, ReactNode, useState } from 'react';
import { GoogleUser, useGoogleAuth } from '../hooks/useGoogleAuth';
import { useRouter } from 'next/navigation';
import { Conversation } from '@/lib/types/chat';

interface AuthContextType {
  isAuthenticated: boolean;
  user: GoogleUser | null;
  isLoading: boolean;
  signOut: () => void;
  token: string | null;
  conversations: Map<number, Conversation>;
  setConversations: (conversations: Map<number, Conversation>) => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const { isAuthenticated, user, isLoading, signOut } = useGoogleAuth();
  const [conversations, setConversations] = useState<Map<number, Conversation>>(new Map());
  
  // Get token from localStorage
  const token = typeof window !== 'undefined' ? localStorage.getItem('google_token') : null;
  console.log("Token: ", token);

  const value: AuthContextType = {
    isAuthenticated,
    user,
    isLoading,
    signOut,
    token,
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
      // Could redirect to login page or show login modal
      console.warn('User not authenticated');
    }
  }, [auth.isAuthenticated, auth.isLoading, router]); //!I need this since it is used by the hook. Some clouse thing with the react.

  return auth;
}; 

export const useConversations = (): [Map<number, Conversation>, (conversations: Map<number, Conversation>) => void] => {
  const auth = useAuth();
  return [auth.conversations, auth.setConversations];
};