import {
  Workout,
  CardioExercise,
  WeightExercise,
  CreateWorkoutRequest,
  CreateCardioExerciseRequest,
  CreateWeightExerciseRequest,
  ApiResponse,
  PaginatedResponse
} from '../types/models';

const API_BASE_URL = 'http://localhost:8080/api';

// Workout API calls
/*
export const workoutApi = {
  create: async (workout: CreateWorkoutRequest): Promise<ApiResponse<Workout>> => {
    const response = await fetch(`${API_BASE_URL}/workouts`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(workout),
    });
    if (!response.ok) throw new Error('Failed to create workout');
    return response.json();
  },

  get: async (id: number): Promise<ApiResponse<Workout>> => {
    const response = await fetch(`${API_BASE_URL}/workouts/${id}`);
    if (!response.ok) throw new Error('Failed to fetch workout');
    return response.json();
  },

  list: async (userId: number, page = 1, limit = 10): Promise<PaginatedResponse<Workout>> => {
    const response = await fetch(
      `${API_BASE_URL}/workouts?user_id=${userId}&page=${page}&limit=${limit}`
    );
    if (!response.ok) throw new Error('Failed to fetch workouts');
    return response.json();
  },
};
*/

// Exercise API calls
export const exerciseApi = {
    createCardio: async (exercise: CreateCardioExerciseRequest): Promise<ApiResponse<CardioExercise>> => {
        const response = await fetch(`${API_BASE_URL}/exercises/cardio`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(exercise),
        });
        if (!response.ok) throw new Error('Failed to create cardio exercise');
        return response.json();
    },

    createWeight: async (exercise: CreateWeightExerciseRequest): Promise<ApiResponse<WeightExercise>> => {
        const response = await fetch(`${API_BASE_URL}/exercises/weights`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(exercise),
        });
        if (!response.ok) throw new Error('Failed to create weight exercise');
        return response.json();
    },

    get: async (id: number): Promise<ApiResponse<CardioExercise | WeightExercise>> => {
        const response = await fetch(`${API_BASE_URL}/exercises/${id}`);
        if (!response.ok) throw new Error('Failed to fetch exercise');
        return response.json();
    },

    list: async (workoutId: number, page = 1, limit = 10): Promise<PaginatedResponse<CardioExercise | WeightExercise>> => {
        const response = await fetch(
        `${API_BASE_URL}/exercises?workout_id=${workoutId}&page=${page}&limit=${limit}`
        );
        if (!response.ok) throw new Error('Failed to fetch exercises');
        return response.json();
    },
}; 