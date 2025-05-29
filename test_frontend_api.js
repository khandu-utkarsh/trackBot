// Simple test script to verify API connectivity
const API_BASE_URL = 'http://localhost:8080';

async function testAPI() {
  console.log('🧪 Testing API connectivity from Node.js');
  console.log('===========================================');

  // Test 1: Health endpoint (no auth)
  try {
    console.log('\n1. Testing health endpoint...');
    const healthResponse = await fetch(`${API_BASE_URL}/health`);
    console.log(`✅ Health: ${healthResponse.status} - ${await healthResponse.text()}`);
  } catch (error) {
    console.error('❌ Health test failed:', error.message);
  }

  // Test 2: Protected endpoint without auth (should fail)
  try {
    console.log('\n2. Testing protected endpoint without auth...');
    const noAuthResponse = await fetch(`${API_BASE_URL}/api/users`);
    console.log(`✅ No auth: ${noAuthResponse.status} - ${await noAuthResponse.text()}`);
  } catch (error) {
    console.error('❌ No auth test failed:', error.message);
  }

  // Test 3: Protected endpoint with auth (should work)
  try {
    console.log('\n3. Testing protected endpoint with auth...');
    const authResponse = await fetch(`${API_BASE_URL}/api/users`, {
      headers: {
        'Authorization': 'Bearer dummy-token-123',
        'Content-Type': 'application/json',
      },
    });
    console.log(`✅ With auth: ${authResponse.status} - ${await authResponse.text()}`);
  } catch (error) {
    console.error('❌ Auth test failed:', error.message);
  }

  // Test 4: Message creation
  try {
    console.log('\n4. Testing message creation...');
    const messageResponse = await fetch(`${API_BASE_URL}/api/users/1/conversations/1/messages`, {
      method: 'POST',
      headers: {
        'Authorization': 'Bearer dummy-token-123',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        content: 'Hello from Node.js test!',
        message_type: 'user'
      }),
    });
    console.log(`✅ Message: ${messageResponse.status} - ${await messageResponse.text()}`);
  } catch (error) {
    console.error('❌ Message test failed:', error.message);
  }
}

testAPI().catch(console.error); 