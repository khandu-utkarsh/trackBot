import { NextRequest, NextResponse } from 'next/server';

export async function POST(request: NextRequest) {
  try {
    const body = await request.json();
    const { message, conversation } = body;

    if (!message || typeof message !== 'string') {
      return NextResponse.json({ error: 'Message is required' }, { status: 400 });
    }

    // Simulate AI response with fitness-focused replies
    const fitnessResponses = [
      {
        pattern: /workout|exercise|training|gym/i,
        responses: [
          "Great question about workouts! Here's what I recommend: Start with compound movements like squats, deadlifts, and push-ups. These exercises work multiple muscle groups and are very effective for building strength and muscle.",
          "For an effective workout routine, aim for 3-4 sessions per week. Include both strength training and cardio. Remember to progress gradually and listen to your body!",
          "Consistency is key in any workout routine. I'd suggest starting with full-body workouts 3 times per week, allowing rest days between sessions for recovery."
        ]
      },
      {
        pattern: /nutrition|diet|eat|food|protein|calories/i,
        responses: [
          "Nutrition is crucial for fitness success! Focus on whole foods: lean proteins (chicken, fish, legumes), complex carbohydrates (oats, quinoa, sweet potatoes), and healthy fats (avocado, nuts, olive oil).",
          "For muscle building, aim for 0.8-1g of protein per pound of body weight. For weight loss, maintain a moderate caloric deficit while keeping protein high to preserve muscle mass.",
          "Meal timing can be important. Try to eat a balanced meal with protein and carbs within 2 hours after your workout to optimize recovery."
        ]
      },
      {
        pattern: /cardio|running|cycling|swimming/i,
        responses: [
          "Cardio is excellent for heart health and fat loss! Mix steady-state cardio (like jogging) with high-intensity intervals (HIIT) for best results. Start with 20-30 minutes, 3 times per week.",
          "For beginners, I recommend starting with low-impact cardio like walking, swimming, or cycling. Gradually increase intensity and duration as your fitness improves.",
          "HIIT workouts are very effective! Try 30 seconds of high intensity followed by 90 seconds of rest, repeated for 15-20 minutes."
        ]
      },
      {
        pattern: /weight loss|lose weight|fat loss/i,
        responses: [
          "Weight loss requires a caloric deficit - burning more calories than you consume. Combine strength training, cardio, and a balanced diet. Aim for 1-2 pounds of loss per week for sustainable results.",
          "Focus on building muscle while losing fat. Strength training is crucial because muscle tissue burns more calories at rest. Don't just focus on the scale - body composition matters more!",
          "Create a moderate caloric deficit (300-500 calories below maintenance), prioritize protein, and include both cardio and strength training in your routine."
        ]
      },
      {
        pattern: /muscle|strength|building|gain/i,
        responses: [
          "For muscle building, focus on progressive overload - gradually increasing weight, reps, or sets over time. Compound exercises like squats, deadlifts, and bench press are most effective.",
          "Muscle building requires adequate protein (0.8-1g per lb bodyweight), progressive resistance training, and proper rest. Don't neglect sleep - that's when muscle growth happens!",
          "Aim for 6-12 reps per set for muscle hypertrophy. Train each muscle group 2-3 times per week with adequate rest between sessions."
        ]
      }
    ];

    // Find matching response based on message content
    let response = "I'm here to help with your fitness journey! Feel free to ask me about workouts, nutrition, weight loss, muscle building, or any other fitness-related topics. I can provide personalized advice based on your goals.";
    
    for (const category of fitnessResponses) {
      if (category.pattern.test(message)) {
        const randomResponse = category.responses[Math.floor(Math.random() * category.responses.length)];
        response = randomResponse;
        break;
      }
    }

    return NextResponse.json({
      response,
      timestamp: new Date().toISOString(),
    });

  } catch (error) {
    console.error('Chat API error:', error);
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 }
    );
  }
} 