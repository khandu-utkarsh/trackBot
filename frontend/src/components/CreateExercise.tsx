'use client';

import { useState, useEffect } from 'react';
import { exerciseApi } from '../services/api';
import { CreateCardioExerciseRequest, CreateWeightExerciseRequest } from '../types/models';
import { 
  Box, 
  Button, 
  TextField, 
  Select, 
  MenuItem, 
  FormControl, 
  InputLabel, 
  Typography, 
  CircularProgress, 
  Alert 
} from '@mui/material';

type ExerciseType = 'cardio' | 'weight';

// Assume a workout ID is available (e.g., passed as prop or from context)
// Replace 1 with the actual workout ID
const DEFAULT_WORKOUT_ID = 1;

export default function CreateExercise() {
  const [exerciseType, setExerciseType] = useState<ExerciseType>('cardio');
  const [formData, setFormData] = useState<Partial<CreateCardioExerciseRequest | CreateWeightExerciseRequest>>({
    workout_id: DEFAULT_WORKOUT_ID,
    name: '',
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  // Reset form when exercise type changes
  useEffect(() => {
    setFormData({
      workout_id: DEFAULT_WORKOUT_ID,
      name: '',
      ...(exerciseType === 'cardio' ? { duration: 30, distance: 5, calories: 300 } : { sets: 3, reps: 10, weight: 50 })
    });
    setError(null);
    setSuccess(null);
  }, [exerciseType]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: name === 'name' ? value : parseFloat(value) || 0 // Ensure numeric fields are numbers
    }));
  };

  const handleTypeChange = (event: any) => {
    setExerciseType(event.target.value as ExerciseType);
  };

  const handleSubmit = async (e: React.FormEvent) => {

    console.log("Handle submit called.");

    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      let response;
      if (exerciseType === 'cardio') {
        response = await exerciseApi.createCardio(formData as CreateCardioExerciseRequest);
      } else {
        response = await exerciseApi.createWeight(formData as CreateWeightExerciseRequest);
      }
      
      if (response.error) {
        throw new Error(response.error);
      }

      setSuccess(`Successfully created ${exerciseType} exercise: ${formData.name}`);
      // Reset form partially (keep workout_id and type)
      setFormData(prev => ({
        workout_id: DEFAULT_WORKOUT_ID,
        name: '',
        ...(exerciseType === 'cardio' ? { duration: 30, distance: 5, calories: 300 } : { sets: 3, reps: 10, weight: 50 })
      }));

    } catch (err) {
      setError(err instanceof Error ? err.message : `Failed to create ${exerciseType} exercise`);
    } finally {
      setLoading(false);
    }
  };
  
  // Auto-clear success/error messages
  useEffect(() => {
      if (success || error) {
          const timer = setTimeout(() => {
              setSuccess(null);
              setError(null);
          }, 5000); // Clear after 5 seconds
          return () => clearTimeout(timer);
      }
  }, [success, error]);

  return (
    <Box component="form" onSubmit={handleSubmit} sx={{ maxWidth: 500, mx: 'auto', p: 3, border: '1px solid grey', borderRadius: 2 }}>
      <Typography variant="h5" gutterBottom>
        Add New Exercise
      </Typography>

      {error && <Alert severity="error" sx={{ mb: 2 }}>{error}</Alert>}
      {success && <Alert severity="success" sx={{ mb: 2 }}>{success}</Alert>}

      <FormControl fullWidth margin="normal">
        <InputLabel id="exercise-type-label">Exercise Type</InputLabel>
        <Select
          labelId="exercise-type-label"
          value={exerciseType}
          label="Exercise Type"
          onChange={handleTypeChange}
        >
          <MenuItem value="cardio">Cardio</MenuItem>
          <MenuItem value="weight">Weight Lifting</MenuItem>
        </Select>
      </FormControl>

      <TextField
        label="Exercise Name"
        name="name"
        value={formData.name}
        onChange={handleChange}
        fullWidth
        required
        margin="normal"
      />

      {exerciseType === 'cardio' && (
        <>
          <TextField
            label="Duration (minutes)"
            name="duration"
            type="number"
            value={(formData as CreateCardioExerciseRequest).duration ?? 30}
            onChange={handleChange}
            fullWidth
            required
            margin="normal"
            inputProps={{ min: 1 }}
          />
          <TextField
            label="Distance (km)"
            name="distance"
            type="number"
            value={(formData as CreateCardioExerciseRequest).distance ?? 5}
            onChange={handleChange}
            fullWidth
            required
            margin="normal"
            inputProps={{ min: 0, step: 0.1 }}
          />
          <TextField
            label="Calories Burned"
            name="calories"
            type="number"
            value={(formData as CreateCardioExerciseRequest).calories ?? 300}
            onChange={handleChange}
            fullWidth
            required
            margin="normal"
            inputProps={{ min: 0 }}
          />
        </>
      )}

      {exerciseType === 'weight' && (
        <>
          <TextField
            label="Sets"
            name="sets"
            type="number"
            value={(formData as CreateWeightExerciseRequest).sets ?? 3}
            onChange={handleChange}
            fullWidth
            required
            margin="normal"
            inputProps={{ min: 1 }}
          />
          <TextField
            label="Reps"
            name="reps"
            type="number"
            value={(formData as CreateWeightExerciseRequest).reps ?? 10}
            onChange={handleChange}
            fullWidth
            required
            margin="normal"
            inputProps={{ min: 1 }}
          />
          <TextField
            label="Weight (kg)"
            name="weight"
            type="number"
            value={(formData as CreateWeightExerciseRequest).weight ?? 50}
            onChange={handleChange}
            fullWidth
            required
            margin="normal"
            inputProps={{ min: 0, step: 0.5 }}
          />
        </>
      )}

      <Button 
        type="submit" 
        variant="contained" 
        color="primary" 
        disabled={loading} 
        fullWidth
        sx={{ mt: 2 }}
      >
        {loading ? <CircularProgress size={24} /> : 'Add Exercise'}
      </Button>
    </Box>
  );
} 