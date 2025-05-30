# Authentication Context Provider

This guide explains how to use the authentication context provider that centralizes session management throughout your TrackBot application.

## Overview

The authentication context provides a centralized way to manage user authentication state, making it easy to:
- Access user information anywhere in your app
- Check authentication status
- Make authenticated API calls
- Handle sign-out functionality
- Protect routes based on authentication

## Components Created

### 1. `AuthContext` (`src/contexts/AuthContext.tsx`)
The main context provider that wraps your existing `useGoogleAuth` hook and provides authentication state globally.

### 2. `Header` (`src/components/Header.tsx`)
A sample header component that shows user info and provides sign-out functionality.

### 3. `Dashboard` (`src/components/Dashboard.tsx`)
A sample dashboard that demonstrates accessing user data from the context.

### 4. `ProtectedRoute` (`src/components/ProtectedRoute.tsx`)
A wrapper component for protecting routes that require authentication.

### 5. `API Client` (`src/lib/api.ts`)
Utility functions for making authenticated API calls using the token from context.

## Setup

The authentication provider is already set up in your `layout.tsx`:

```tsx
<ThemeProvider theme={lightTheme}>
  <CssBaseline />
  <AuthProvider>
    <ThemedContainer>
      {children}
    </ThemedContainer>
  </AuthProvider>
</ThemeProvider>
```

## Usage Examples

### Basic Usage - Accessing Authentication State

```tsx
import { useAuth } from '../contexts/AuthContext';

function MyComponent() {
  const { isAuthenticated, user, isLoading, signOut, token } = useAuth();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (!isAuthenticated) {
    return <div>Please sign in</div>;
  }

  return (
    <div>
      <h1>Welcome, {user?.name}!</h1>
      <p>Email: {user?.email}</p>
      <button onClick={signOut}>Sign Out</button>
    </div>
  );
}
```

### Protected Routes

```tsx
import { ProtectedRoute } from '../components/ProtectedRoute';

function App() {
  return (
    <ProtectedRoute>
      <Dashboard />
    </ProtectedRoute>
  );
}

// Or with a custom fallback
<ProtectedRoute fallback={<CustomLoginPage />}>
  <PrivateContent />
</ProtectedRoute>
```

### Making Authenticated API Calls

```tsx
import { useAuth } from '../contexts/AuthContext';
import { apiClient } from '../lib/api';

function WorkoutComponent() {
  const { token } = useAuth();

  const fetchWorkouts = async () => {
    if (token) {
      const response = await apiClient.authenticatedGet('/api/workouts', token);
      if (response.data) {
        console.log('Workouts:', response.data);
      } else {
        console.error('Error:', response.error);
      }
    }
  };

  const createWorkout = async (workoutData: any) => {
    if (token) {
      const response = await apiClient.authenticatedPost('/api/workouts', workoutData, token);
      if (response.data) {
        console.log('Workout created:', response.data);
      }
    }
  };

  return (
    <div>
      <button onClick={fetchWorkouts}>Load Workouts</button>
      <button onClick={() => createWorkout({ name: 'Morning Run' })}>
        Create Workout
      </button>
    </div>
  );
}
```

### Conditional Rendering Based on Auth Status

```tsx
import { useAuth } from '../contexts/AuthContext';

function Navigation() {
  const { isAuthenticated, user } = useAuth();

  return (
    <nav>
      {isAuthenticated ? (
        <>
          <span>Hello, {user?.name}</span>
          <Link href="/dashboard">Dashboard</Link>
          <Link href="/workouts">Workouts</Link>
        </>
      ) : (
        <Link href="/login">Sign In</Link>
      )}
    </nav>
  );
}
```

### Force Authentication (Redirect if Not Authenticated)

```tsx
import { useRequireAuth } from '../contexts/AuthContext';

function PrivateComponent() {
  // This hook will log a warning if user is not authenticated
  // You can extend it to redirect to login page
  const { user } = useRequireAuth();

  return (
    <div>
      <h1>Private content for {user?.name}</h1>
    </div>
  );
}
```

## Available Context Properties

- `isAuthenticated: boolean` - Whether the user is authenticated
- `user: GoogleUser | null` - User information (name, email, picture)
- `isLoading: boolean` - Whether authentication check is in progress
- `signOut: () => void` - Function to sign out the user
- `token: string | null` - The authentication token for API calls

## API Client Usage

The API client automatically handles authorization headers when you pass a token:

```tsx
// Public endpoints (no authentication)
const publicData = await apiClient.get('/api/public/data');

// Authenticated endpoints
const { token } = useAuth();
if (token) {
  const userData = await apiClient.authenticatedGet('/api/user/profile', token);
  const newWorkout = await apiClient.authenticatedPost('/api/workouts', workoutData, token);
  const updatedWorkout = await apiClient.authenticatedPut('/api/workouts/123', updates, token);
  const deleted = await apiClient.authenticatedDelete('/api/workouts/123', token);
}
```

## Benefits

1. **Centralized State**: Authentication state is managed in one place
2. **Easy Access**: Use `useAuth()` hook anywhere in your app
3. **Type Safety**: Full TypeScript support
4. **Automatic Token Management**: Token is automatically included in API calls
5. **Loading States**: Built-in loading state management
6. **Route Protection**: Easy to protect routes with `ProtectedRoute` component
7. **Consistent UI**: User info and sign-out functionality can be consistently displayed

## Next Steps

You can extend this context to include:
- User preferences and settings
- Role-based access control
- Session persistence
- Token refresh logic
- Multiple authentication providers
- User profile management

The context provider pattern makes it easy to add these features while keeping them accessible throughout your application. 