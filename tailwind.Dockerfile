# tailwind.Dockerfile (Based on the Proven Rails Example)

# Start from a stable, minimal Debian image.
FROM debian:bookworm-slim

# Install all necessary dependencies in a single, efficient command.
# This includes 'curl' for downloading and 'watchman' for file watching.
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    ca-certificates \
    watchman \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory.
WORKDIR /app

# Download and install the Tailwind CSS standalone CLI, and then verify it.
# This entire block will fail the build if the download or tool is broken.
RUN curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o /usr/local/bin/tailwindcss && \
    chmod +x /usr/local/bin/tailwindcss && \
    /usr/local/bin/tailwindcss --help > /dev/null

# NOTE: There is NO CMD here. We will provide it from docker-compose.yml.
