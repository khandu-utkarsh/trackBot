from langgraph.checkpoint.postgres import PostgresSaver
from config.settings import get_settings
import logging


logger = logging.getLogger(__name__)
settings = get_settings()

async def init_langgraph_checkpointing():
    """
    Initialize the LangGraph PostgreSQL checkpointing tables.
    This should be run once when setting up the application.
    """
    try:
        logger.info("Initializing LangGraph PostgreSQL checkpointing...")
        
        # Use the sync database URL for initial setup
        DB_URI = settings.DATABASE_URL
        
        # Initialize the checkpointer synchronously for setup
        with PostgresSaver.from_conn_string(DB_URI) as checkpointer:
            logger.info(f"Checkpointer type: {type(checkpointer)}")  # <- Add this line
            checkpointer.setup()
        
        logger.info("LangGraph PostgreSQL checkpointing initialized successfully")
            
    except Exception as e:
        logger.error(f"Error initializing LangGraph PostgreSQL checkpointing: {e}")
        raise 