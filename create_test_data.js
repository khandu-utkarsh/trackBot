// Script to create test data for API testing
const API_BASE_URL = 'http://localhost:8080';

async function createTestData() {
  console.log('🔧 Creating test data...');
  console.log('========================');

  const headers = {
    'Authorization': 'Bearer dummy-token-123',
    'Content-Type': 'application/json',
  };

  try {
    // Step 1: Create a test user
    console.log('\n1. Creating test user...');
    const userResponse = await fetch(`${API_BASE_URL}/api/users`, {
      method: 'POST',
      headers,
      body: JSON.stringify({
        email: 'testuser@example.com'
      }),
    });
    
    if (userResponse.ok) {
      const user = await userResponse.json();
      console.log(`✅ User created: ID ${user.id}, Email: ${user.email}`);
      
      // Step 2: Create a test conversation for this user
      console.log('\n2. Creating test conversation...');
      const conversationResponse = await fetch(`${API_BASE_URL}/api/users/${user.id}/conversations`, {
        method: 'POST',
        headers,
        body: JSON.stringify({
          title: 'Test Conversation',
          is_active: true
        }),
      });
      
      if (conversationResponse.ok) {
        const conversation = await conversationResponse.json();
        console.log(`✅ Conversation created: ID ${conversation.id}, Title: ${conversation.title}`);
        
        // Step 3: Create a test message
        console.log('\n3. Creating test message...');
        const messageResponse = await fetch(`${API_BASE_URL}/api/users/${user.id}/conversations/${conversation.id}/messages`, {
          method: 'POST',
          headers,
          body: JSON.stringify({
            content: 'Hello! This is a test message.',
            message_type: 'user'
          }),
        });
        
        if (messageResponse.ok) {
          const message = await messageResponse.json();
          console.log(`✅ Message created: ID ${message.id}, Content: "${message.content}"`);
          
          console.log('\n🎉 Test data created successfully!');
          console.log(`📊 Use these IDs for testing:`);
          console.log(`   User ID: ${user.id}`);
          console.log(`   Conversation ID: ${conversation.id}`);
          console.log(`   Message ID: ${message.id}`);
          
          return { user, conversation, message };
        } else {
          const error = await messageResponse.text();
          console.error(`❌ Message creation failed: ${messageResponse.status} - ${error}`);
        }
      } else {
        const error = await conversationResponse.text();
        console.error(`❌ Conversation creation failed: ${conversationResponse.status} - ${error}`);
      }
    } else {
      const error = await userResponse.text();
      console.error(`❌ User creation failed: ${userResponse.status} - ${error}`);
    }
  } catch (error) {
    console.error('❌ Error creating test data:', error.message);
  }
}

createTestData().catch(console.error); 