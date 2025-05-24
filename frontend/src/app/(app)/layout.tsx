'use client';

import Header from "@/components/Header";
//!This is the global layout for the app. It is used to provide the theme and the providers to the app.

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <Header/>
      {children}
    </>
  );
}
