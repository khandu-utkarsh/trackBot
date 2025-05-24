# Workout Web App Frontend

A modern Next.js application for tracking workouts and getting AI-powered fitness advice. Built with Next.js 13+ App Router, Material-UI, and NextAuth.js.

## Features

- 🔐 **Authentication**: Secure Google OAuth integration
- 💬 **AI Chat**: Interactive chat with AI fitness assistant
- 📊 **Dashboard**: User activity tracking and workout management
- 🎨 **Modern UI**: Clean, responsive design with Material-UI
- 🔒 **Type Safety**: Full TypeScript support
- 🚀 **Performance**: Optimized with Next.js App Router

## Tech Stack

- **Framework**: Next.js 13+ (App Router)
- **UI Library**: Material-UI (MUI)
- **Authentication**: NextAuth.js v5
- **Language**: TypeScript
- **Styling**: Emotion (via MUI)
- **State Management**: React Hooks
- **API**: Next.js API Routes

## Project Structure

```
frontend/
├── src/
│   ├── app/                # Next.js App Router
│   │   ├── (app)/         # Protected routes
│   │   │   ├── chat/      # AI chat interface
│   │   │   └── layout.tsx # Protected layout
│   │   ├── (auth)/        # Auth routes
│   │   │   ├── signin/    # Sign-in page
│   │   │   └── error/     # Auth error page
│   │   ├── (public)/      # Public routes
│   │   │   └── landing/   # Landing page
│   │   └── api/           # API routes
│   ├── components/        # React components
│   │   ├── Header.tsx     # Navigation header
│   │   └── Layout.tsx     # Layout component
│   ├── theme/            # MUI theme config
│   └── auth.ts           # NextAuth config
├── public/              # Static assets
└── [config files]       # Various config files
```

## Getting Started

1. **Clone the repository**
   ```bash
   git clone [repository-url]
   cd frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Set up environment variables**
   Create a `.env.local` file:
   ```env
   # Google OAuth
   GOOGLE_CLIENT_ID=your_google_client_id
   GOOGLE_CLIENT_SECRET=your_google_client_secret

   # NextAuth
   NEXTAUTH_URL=http://localhost:3000
   NEXTAUTH_SECRET=your_random_secret_key
   ```

4. **Start development server**
   ```bash
   npm run dev
   ```

5. **Open** [http://localhost:3000](http://localhost:3000)

## Development

### Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run start` - Start production server
- `npm run lint` - Run ESLint

### Code Style

- TypeScript for type safety
- ESLint for code linting
- Prettier for code formatting

## Authentication Flow

1. User visits landing page
2. Clicks "Get Started" to go to sign-in page
3. Authenticates with Google
4. Redirected to chat interface
5. Session persists across page reloads

## API Integration

The frontend communicates with the backend through:
- Next.js API Routes for server-side operations
- Direct API calls for client-side operations

## Deployment

### Vercel (Recommended)
1. Push to GitHub
2. Import project in Vercel
3. Set environment variables
4. Deploy

### Other Platforms
- Netlify
- AWS Amplify
- Custom server

## Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
