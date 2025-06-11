# Stage 1: Builder
FROM python:3.9 as builder

WORKDIR /app

# Create and activate virtual environment
RUN python3 -m venv /app/venv
ENV PATH="/app/venv/bin:$PATH"

# Copy requirements file
COPY requirements.txt .

# Install dependencies in virtual environment
RUN pip install --no-cache-dir -r requirements.txt

# Stage 2: Slim Runtime
FROM python:3.9-slim

WORKDIR /app

COPY --from=builder /app/venv /app/venv
ENV PATH="/app/venv/bin:$PATH"

COPY . .

# Expose port 5000
EXPOSE 5000

# Command to run the application
CMD ["python3", "app/main.py"] 



