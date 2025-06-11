from pydantic_settings import BaseSettings
from functools import lru_cache
import os

class Settings(BaseSettings):
    # API Settings
    API_STR: str = "/api"
    PROJECT_NAME: str = "TrackBot_Agent_Backend"
    COOKIE_NAME: str = "trackbot_auth_token"
    HOST: str = "0.0.0.0"
    PORT: str = 8080
    # Model Settings
    MODEL_NAME: str = os.getenv("MODEL_NAME")
    OPENAI_API_KEY: str = os.getenv("OPENAI_API_KEY")
    DATABASE_URL: str = "postgresql://postgres:postgres@localhost:5432/workout_app"
    GOOGLE_CLIENT_ID: str = os.getenv("GOOGLE_CLIENT_ID")
    TRACKBOT_JWT_SECRET_KEY: str = os.getenv("TRACKBOT_JWT_SECRET_KEY")
    ACCESS_TOKEN_EXPIRE_MINUTES: int = 300
    class Config:
        case_sensitive = False

@lru_cache()
def get_settings() -> Settings:
    return Settings() 