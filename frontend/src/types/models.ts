// Base interface for all entities with an ID
export interface BaseEntity {
  id: number;
}

// Workout related interfaces
export interface Workout extends BaseEntity {
  user_id: number;
  name: string;
  description: string;
  date: string;
  duration: number;
}

// Exercise related interfaces
export interface BaseExercise extends BaseEntity {
  workout_id: number;
  name: string;
}

export interface CardioExercise extends BaseExercise {
  duration: number;
  distance: number;
  calories: number;
}

export interface WeightExercise extends BaseExercise {
  sets: number;
  reps: number;
  weight: number;
}

// User related interfaces
export interface User extends BaseEntity {
  username: string;
  email: string;
  created_at: string;
}

// API Response interfaces
export interface ApiResponse<T> {
  data: T;
  message?: string;
  error?: string;
}

export interface PaginatedResponse<T> extends ApiResponse<T[]> {
  total: number;
  page: number;
  limit: number;
  total_pages: number;
}

// Request interfaces for creating/updating entities
export type CreateWorkoutRequest = Omit<Workout, 'id'>;
export type CreateCardioExerciseRequest = Omit<CardioExercise, 'id'>;
export type CreateWeightExerciseRequest = Omit<WeightExercise, 'id'>;
export type CreateUserRequest = Omit<User, 'id' | 'created_at'>;

// Update request interfaces
export type UpdateWorkoutRequest = Partial<Omit<Workout, 'id' | 'user_id'>>;
export type UpdateCardioExerciseRequest = Partial<Omit<CardioExercise, 'id' | 'workout_id'>>;
export type UpdateWeightExerciseRequest = Partial<Omit<WeightExercise, 'id' | 'workout_id'>>;
export type UpdateUserRequest = Partial<Omit<User, 'id' | 'created_at'>>; 