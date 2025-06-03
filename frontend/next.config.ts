import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: 'standalone',
  // Enable environment variables to be available at build time
  env: {
    REACT_APP_API_URL: process.env.REACT_APP_API_URL,
    REACT_APP_WS_URL: process.env.REACT_APP_WS_URL,
  },
  // Proxy API requests to backend to avoid CORS issues
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://localhost:8080/api/:path*',
      },
      {
        source: '/health',
        destination: 'http://localhost:8080/health',
      },
    ];
  },
};

export default nextConfig;
