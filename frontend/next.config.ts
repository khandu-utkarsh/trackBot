import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: 'standalone',
  // Enable environment variables to be available at build time
  env: {
    REACT_APP_API_URL: process.env.REACT_APP_API_URL,
    REACT_APP_WS_URL: process.env.REACT_APP_WS_URL,
  },
};

export default nextConfig;
