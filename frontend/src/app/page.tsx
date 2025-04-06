'use client';

import { useState, useEffect } from 'react';
import { 
  Box, 
  Button, 
  Card, 
  CardContent, 
  Typography, 
  Grid,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  IconButton
} from '@mui/material';
import { Add as AddIcon, Edit as EditIcon, Delete as DeleteIcon } from '@mui/icons-material';
import Layout from '@/components/Layout';

interface Workout {
  id: number;
  name: string;
  description: string;
  date: string;
  duration: number;
}

export default function Home() {
  const [workouts, setWorkouts] = useState<Workout[]>([]);
  const [open, setOpen] = useState(false);
  const [newWorkout, setNewWorkout] = useState({
    name: '',
    description: '',
    date: new Date().toISOString().split('T')[0],
    duration: 30
  });

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleAddWorkout = async () => {
    // TODO: Implement API call to add workout
    setWorkouts([...workouts, { ...newWorkout, id: workouts.length + 1 }]);
    handleClose();
  };

  return (
    <Layout>
      <Box sx={{ mb: 4 }}>
        <Button
          variant="contained"
          color="primary"
          startIcon={<AddIcon />}
          onClick={handleClickOpen}
        >
          Add Workout
        </Button>
      </Box>

      <Grid container spacing={3}>
        {workouts.map((workout) => (
          <Grid item xs={12} sm={6} md={4} key={workout.id}>
            <Card>
              <CardContent>
                <Typography variant="h5" component="div">
                  {workout.name}
                </Typography>
                <Typography color="text.secondary" sx={{ mb: 1.5 }}>
                  {workout.date}
                </Typography>
                <Typography variant="body2">
                  {workout.description}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Duration: {workout.duration} minutes
                </Typography>
                <Box sx={{ mt: 2, display: 'flex', justifyContent: 'flex-end' }}>
                  <IconButton size="small">
                    <EditIcon />
                  </IconButton>
                  <IconButton size="small">
                    <DeleteIcon />
                  </IconButton>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>

      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Add New Workout</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            margin="dense"
            label="Workout Name"
            fullWidth
            variant="outlined"
            value={newWorkout.name}
            onChange={(e) => setNewWorkout({ ...newWorkout, name: e.target.value })}
          />
          <TextField
            margin="dense"
            label="Description"
            fullWidth
            variant="outlined"
            multiline
            rows={4}
            value={newWorkout.description}
            onChange={(e) => setNewWorkout({ ...newWorkout, description: e.target.value })}
          />
          <TextField
            margin="dense"
            label="Date"
            type="date"
            fullWidth
            variant="outlined"
            value={newWorkout.date}
            onChange={(e) => setNewWorkout({ ...newWorkout, date: e.target.value })}
          />
          <TextField
            margin="dense"
            label="Duration (minutes)"
            type="number"
            fullWidth
            variant="outlined"
            value={newWorkout.duration}
            onChange={(e) => setNewWorkout({ ...newWorkout, duration: parseInt(e.target.value) })}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleAddWorkout} variant="contained">
            Add
          </Button>
        </DialogActions>
      </Dialog>
    </Layout>
  );
}
