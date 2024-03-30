from django.db import models

#TODO: Provide the verbosity information to all the models here

#Following the convention, not capitalizing the first letter in the verbose name


# Create your models here.

# Contains user properties
class Users(models.Model):
    name = models.CharField(max_length=100, verbose_name = "name of the user")
    age = models.IntegerField(blank=True, null=True, verbose_name = "age of the user")  # Allowing age to be optional
    height = models.IntegerField(blank=True, null=True, verbose_name = "height of the user")  # Allowing height to be optional
    email = models.EmailField(max_length=254, verbose_name = "email of the user")
    username = models.CharField(max_length=20, verbose_name = "username of the user")  # Assuming username max length is 20 characters

#Can't be none
class Weight(models.Model):
    time_stamp = models.DateTimeField(auto_now=False, auto_now_add=False, verbose_name = "timestamp")  # Auto now is responsible for auto-saving the time object is created.
    weight = models.IntegerField(verbose_name = "weight")
    user = models.ForeignKey(Users, on_delete=models.CASCADE)  # On deleting the user, all weight data should also be deleted.

#Can't be none
class BodyFatPercentage(models.Model):
    time_stamp = models.DateTimeField(auto_now=True, auto_now_add=False, verbose_name = "timestamp")
    body_fat_percentage = models.DecimalField(max_digits=4, decimal_places=2, verbose_name = "body fat percentage")  # Corrected field name
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
        verbose_name='exercise Type'
    )


class Exercise(models.Model):
    name = models.CharField(max_length = 100, verbose_name = "name of the exercise")
    type = models.ForeignKey(ExerciseTypes, on_delete=models.CASCADE)

class ExerciseStrengthLog(models.Model):
    #Reference to exercise
    reps = models.IntegerField(verbose_name = "reps of the exercise")
    weight = models.IntegerField(verbose_name = "weight of the exercise")
    is_body_weight = models.BooleanField(default=False, verbose_name = "marker for converting it to a body weight exercises") #Set the default value to be false. #If true, means only body weight exercise was performed like planks, pull ups etc.
    time_stamp = models.DateTimeField(auto_now=False, auto_now_add=False, verbose_name = "timestamp")
    exercise = models.ForeignKey(Exercise, on_delete=models.CASCADE)

class ExerciseCardioLog(models.Model):
    exercise = models.ForeignKey(Exercise, on_delete=models.CASCADE)
    distance = models.IntegerField(verbose_name = "distance")
    time = models.DurationField(verbose_name = "duration")
    time_stamp = models.DateTimeField(auto_now=False, auto_now_add=False, verbose_name = "timestamp")
