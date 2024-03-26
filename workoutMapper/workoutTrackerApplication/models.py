from django.db import models

#TODO: Provide the verbosity information to all the models here

# Create your models here.

# Contains user properties
class Users(models.Model):
    name = models.CharField(max_length=100, verbose_name = "Name of the user")
    age = models.IntegerField(blank=True, null=True, verbose_name = "Age of the user")  # Allowing age to be optional
    height = models.IntegerField(blank=True, null=True, verbose_name = "Height of the user")  # Allowing height to be optional
    email = models.EmailField(max_length=254, verbose_name = "Email of the user")
    username = models.CharField(max_length=20, verbose_name = "Username of the user")  # Assuming username max length is 20 characters

# Find a way to update the date time from the manual data provided by the user

#Can't be none
class Weight(models.Model):
    time_stamp = models.DateTimeField(auto_now=False, auto_now_add=False, verbose_name = "Timestamp")  # Auto now is responsible for auto-saving the time object is created.
    weight = models.IntegerField(verbose_name = "Weight")
    user = models.ForeignKey(Users, on_delete=models.CASCADE)  # On deleting the user, all weight data should also be deleted.

#Can't be none
class BodyFatPercentage(models.Model):
    time_stamp = models.DateTimeField(auto_now=True, auto_now_add=False, verbose_name = "Timestamp")
    body_fat_percentage = models.DecimalField(max_digits=4, decimal_places=2, verbose_name = "Body fat percentage")  # Corrected field name
    user = models.ForeignKey(Users, on_delete=models.CASCADE)

class ExerciseTypes(models.Model):
    CARDIO = "CD"
    STRENGTH = "ST"
    #Right now, only supporting two types:
    EXERCISE_CHOICES = {
        CARDIO: 'Cardio',
        STRENGTH: 'Strength'
    }
    
    exercise_type = models.CharField(
        max_length=2,
        choices=EXERCISE_CHOICES,
        default=STRENGTH,  # Assuming default is strength
        verbose_name='Exercise Type'
    )


class Exercise(models.Model):
    name = models.CharField(max_length = 100, verbose_name = "Name of the exercise")
    type = models.ForeignKey(ExerciseTypes, on_delete=models.CASCADE)

class ExerciseStrengthLog(models.Model):
    #Reference to exercise
    reps = models.IntegerField(verbose_name = "Reps of the exercise")
    weight = models.IntegerField(verbose_name = "Weight of the exercise")
    is_body_weight = models.BooleanField(default=False, verbose_name = "Marker for converting it to a body weight exercises") #Set the default value to be false. #If true, means only body weight exercise was performed like planks, pull ups etc.
    time_stamp = models.DateTimeField(auto_now=False, auto_now_add=False, verbose_name = "Timestamp")
    exercise = models.ForeignKey(Exercise, on_delete=models.CASCADE)

class ExerciseCardioLog(models.Model):
    exercise = models.ForeignKey(Exercise, on_delete=models.CASCADE)
    distance = models.IntegerField(verbose_name = "Distance")
    time = models.DurationField(verbose_name = "Duration")
    time_stamp = models.DateTimeField(auto_now=False, auto_now_add=False, verbose_name = "Timestamp")
