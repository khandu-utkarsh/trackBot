// Test API client with dummy authentication
const API_BASE_URL = 'http://localhost:8080';

export class TestApiClient {
  private dummyToken = 'dummy-token-123';

  private getHeaders() {
    return {
      'Authorization': `Bearer ${this.dummyToken}`,
      'Content-Type': 'application/json',
    };
  }

  async testHealthEndpoint() {
    console.log('🧪 Testing health endpoint (no auth required)...');
    try {
      const response = await fetch(`${API_BASE_URL}/health`);
      console.log(`✅ Health check: ${response.status} - ${await response.text()}`);
      return response.ok;
    } catch (error) {
      console.error('❌ Health check failed:', error);
      return false;
    }
  }

  async testProtectedEndpoint() {
    console.log('🧪 Testing protected endpoint with dummy token...');
    try {
      const response = await fetch(`${API_BASE_URL}/api/users`, {
        headers: this.getHeaders(),
      });
      console.log(`✅ Protected endpoint: ${response.status}`);
      if (response.ok) {
        const data = await response.json();
        console.log('📊 Response data:', data);
      } else {
        const error = await response.text();
        console.log('❌ Error response:', error);
      }
      return response.ok;
    } catch (error) {
      console.error('❌ Protected endpoint failed:', error);
      return false;
    }
  }

  async testMessageCreation() {
    console.log('🧪 Testing message creation...');
    try {
      const response = await fetch(`${API_BASE_URL}/api/users/1/conversations/1/messages`, {
        method: 'POST',
        headers: this.getHeaders(),
        body: JSON.stringify({
          content: 'Hello from frontend test!',
          message_type: 'user'
        }),
      });
      
      console.log(`✅ Message creation: ${response.status}`);
      if (response.ok) {
        const data = await response.json();
        console.log('📊 Created message:', data);
      } else {
        const error = await response.text();
        console.log('❌ Error response:', error);
      }
      return response.ok;
    } catch (error) {
      console.error('❌ Message creation failed:', error);
      return false;
    }
  }

  async testWithoutAuth() {
    console.log('🧪 Testing protected endpoint without auth (should fail)...');
    try {
      const response = await fetch(`${API_BASE_URL}/api/users`);
      console.log(`✅ No auth test: ${response.status} (should be 401)`);
      const error = await response.text();
      console.log('📊 Error message:', error);
      return response.status === 401;
    } catch (error) {
      console.error('❌ No auth test failed:', error);
      return false;
    }
  }

  async runAllTests() {
    console.log('🚀 Starting API Authentication Tests');
    console.log('=====================================');
    
    const results = {
      health: await this.testHealthEndpoint(),
      noAuth: await this.testWithoutAuth(),
      protected: await this.testProtectedEndpoint(),
      message: await this.testMessageCreation(),
    };

    console.log('\n📊 Test Results:');
    console.log('================');
    Object.entries(results).forEach(([test, passed]) => {
      console.log(`${passed ? '✅' : '❌'} ${test}: ${passed ? 'PASSED' : 'FAILED'}`);
    });

    const allPassed = Object.values(results).every(Boolean);
    console.log(`\n🎯 Overall: ${allPassed ? 'ALL TESTS PASSED' : 'SOME TESTS FAILED'}`);
    
    return results;
  }
}

// Export for use in browser console or components
export const testClient = new TestApiClient(); 