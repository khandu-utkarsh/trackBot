'use client';

import { SessionProvider } from "next-auth/react";

//!This is the provider for the app. It is used to provide the session to the app.

export function Providers({ children }: { children: React.ReactNode }) {
  return <SessionProvider>{children}</SessionProvider>;
} 