from typing import Dict, Any
import httpx
from langchain_openai import ChatOpenAI
from langchain.prompts import ChatPromptTemplate
from langchain.output_parsers import PydanticOutputParser
from pydantic import BaseModel, Field
from config.settings import get_settings

settings = get_settings()

class WorkoutAction(BaseModel):
    action: str = Field(description="The action to perform (create_workout, get_history, update_workout, delete_workout, get_statistics)")
    parameters: Dict[str, Any] = Field(description="Parameters required for the action")

class LLMService:
    def __init__(self):
        self.llm = ChatOpenAI(
            model_name=settings.MODEL_NAME,
            temperature=settings.TEMPERATURE,
            max_tokens=settings.MAX_TOKENS,
            api_key=settings.OPENAI_API_KE
        )
        self.parser = PydanticOutputParser(pydantic_object=WorkoutAction)
        self.prompt = ChatPromptTemplate.from_messages([
            ("system", """You are a workout assistant that helps users interact with a workout application.
            Analyze the user's request and determine the appropriate action to take.
            
            Available actions:
            - create_workout: Create a new workout
            - get_history: Get workout history
            - update_workout: Update an existing workout
            - delete_workout: Delete a workout
            - get_statistics: Get workout statistics
            
            {format_instructions}
            """),
            ("user", "{input}")
        ])
        self.chain = self.prompt | self.llm | self.parser

    async def process_request(self, user_input: str) -> Dict[str, Any]:
        # Process the input through LangChain
        result = await self.chain.ainvoke({
            "input": user_input,
            "format_instructions": self.parser.get_format_instructions()
        })
        
        # Forward the result to the workout app service
        async with httpx.AsyncClient() as client:
            response = await client.post(
                f"{settings.WORKOUT_APP_URL}/api/v1/workouts/process",
                json=result.dict()
            )
            response.raise_for_status()
            return response.json() 