import NextAuth from "next-auth";
import Google from "next-auth/providers/google";
import { Session } from "next-auth";
import { JWT } from "next-auth/jwt";

/**
 * NextAuth configuration for the application
 * This file contains all authentication-related configuration including:
 * - Authentication providers (Google OAuth)
 * - Custom pages for authentication flows
 * - Session and JWT callbacks
 */
export const authConfig = {
  providers: [
    Google({
      clientId: process.env.GOOGLE_CLIENT_ID!,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET!,
    }),
  ],
  pages: {
    signIn: "/signin",
    error: "/error",
  },
  callbacks: {
    async session({ session, token }: { session: Session; token: JWT }) {
      return session;
    },
    async jwt({ token, user }: { token: JWT; user: any }) {
      return token;
    },
  },
};

// Export the configured NextAuth instance
export const { handlers: { GET, POST }, auth, signIn, signOut } = NextAuth(authConfig); 