from django.http import HttpResponse
from django.views import View


def index(request):
    return HttpResponse("Hello, world. You're at Workout Tracker Application.")

class HomePageView(View):
    def get(self, request):
        #Some logical thing here
        return HttpResponse("This is the home page")




# Create your views here.


#Views that I have to create i basically,

#CRUD
#Add exercise to the database, delete exercise from the database, list all the exercises in the database
#Update the exercise in the database, basically is the types of exercise in the databse.




#Create a view where we are having four buttons for a user to create CRUD


#That would be the homepage



#Let's create class based views:


##TODO: Right now:
# 
# Class based views
# Homepage- four button, where giving user the option, to CRUD
# 
# On C page, show the form where, user can enter the exercise name and type
# One R page, simply clicking the Read button, display all the exercises listed as a list
# U --> On U page, basically show a form where we can edit the type of existing exericse. 
# Right now, only allow the user to change the type of exercise
#Functionality to go back to home page



#First view should be of home page