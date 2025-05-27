import { NextRequest, NextResponse } from 'next/server';

export async function POST(request: NextRequest) {
  try {
    // This route is deprecated - the frontend now communicates directly with the backend
    // This is kept for backward compatibility or as a proxy if needed
    
    return NextResponse.json({
      error: 'This API route is deprecated. Please use the backend API directly.',
      message: 'The chat functionality has been moved to the backend service. Please update your client to use the new API endpoints.',
    }, { status: 410 }); // 410 Gone - indicates the resource is no longer available

  } catch (error) {
    console.error('Chat API error:', error);
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 }
    );
  }
} 