import LoggedInLayout from '@/components/LoggedInLayout';

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return <LoggedInLayout>{children}</LoggedInLayout>;
}