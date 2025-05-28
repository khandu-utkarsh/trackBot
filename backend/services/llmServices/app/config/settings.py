from pydantic_settings import BaseSettings
from functools import lru_cache
import os

class Settings(BaseSettings):
    # API Settings
    API_V1_STR: str = "/api/v1"
    PROJECT_NAME: str = "LLM Agents Service Workout App"
    
    # Model Settings
    MODEL_NAME: str = os.getenv("MODEL_NAME")
    OPENAI_API_KEY: str = os.getenv("OPENAI_API_KEY")
    
    # Workout App Service Settings
    WORKOUT_APP_URL: str = "http://workout-app:8080"

    class Config:
        case_sensitive = False

@lru_cache()
def get_settings() -> Settings:
    return Settings() 