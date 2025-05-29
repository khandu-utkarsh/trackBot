'use client';

import LandingPageComponent from '@/components/Landing';
import ChatPageComponent from '@/components/ChatPage';
import { useGoogleAuth } from '@/hooks/useGoogleAuth';

export default function Home() {
  const { isAuthenticated, isLoading } = useGoogleAuth();
  console.log('Printing the state of the auth: ', 'isAuthenticated: ', isAuthenticated, 'isLoading: ', isLoading);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isAuthenticated) {
    return <ChatPageComponent />;
  } else {
    return <LandingPageComponent />;
  }
}
