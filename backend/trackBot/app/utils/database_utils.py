import logging
from sqlalchemy import text
from sqlalchemy.exc import SQLAlchemyError

logger = logging.getLogger(__name__)

async def table_exists(db, table_name: str) -> bool:
    """
    Check if a table exists in the database
    """
    try:
        query = text("SELECT to_regclass(:table_name) IS NOT NULL")
        result = await db.execute(query, {"table_name": table_name})
        exists = result.scalar()
        logger.info(f"Table {table_name} exists: {exists}")
        return exists
    except SQLAlchemyError as e:
        logger.error(f"Error checking if table {table_name} exists: {str(e)}")
        raise

async def create_table(db, table_name: str, schema: str):
    """
    Create a table with the given schema if it doesn't exist
    """
    try:
        query = text(f"CREATE TABLE IF NOT EXISTS {table_name} ({schema})")
        await db.execute(query)
        logger.info(f"Table {table_name} created or already exists")
    except SQLAlchemyError as e:
        logger.error(f"Error creating table {table_name}: {str(e)}")
        raise 