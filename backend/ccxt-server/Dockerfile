FROM python:3.12-slim

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Install uv package manager
RUN pip install uv

# Copy dependency files
COPY pyproject.toml uv.lock* ./

# Install Python dependencies
RUN uv sync --frozen || uv sync

# Copy source code
COPY . .

# Expose gRPC port
EXPOSE 50052

# Run the gRPC server
CMD ["uv", "run", "python", "-m", "src.server.grpc_server"]
