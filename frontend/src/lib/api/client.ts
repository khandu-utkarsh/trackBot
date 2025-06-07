import { Configuration, ConversationsApi, MessagesApi, UsersApi } from '@/lib/types/generated';

// Create the configuration for the API client
const createApiConfiguration = (): Configuration => {
  const baseURL = process.env.NEXT_PUBLIC_GO_BACKEND_BASE_API_URL || 'http://localhost:8080';
  
  return new Configuration({
    basePath: `${baseURL}/api`,
    baseOptions: {
      withCredentials: true,  // axios equivalent of credentials: 'include'
      headers: {
        'Content-Type': 'application/json',
      },
    },
  });
};

// Create API instances
const config = createApiConfiguration();

export const conversationsApi = new ConversationsApi(config);
export const messagesApi = new MessagesApi(config);
export const usersApi = new UsersApi(config);

// Export the configuration for custom usage
export { config as apiConfiguration }; 