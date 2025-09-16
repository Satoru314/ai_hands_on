#!/bin/bash

set -e

echo "Setting up development environment..."

# Check if .env.example exists
if [ ! -f "../.env.example" ]; then
    echo "Error: .env.example not found in parent directory"
    exit 1
fi

# Copy .env.example to .env if .env doesn't exist
if [ ! -f "../.env" ]; then
    echo "Copying .env.example to .env..."
    cp ../.env.example ../.env
    echo ".env file created"
else
    echo ".env file already exists, skipping copy"
fi

# Initialize database
make run-dev
echo "Initializing database..."
make init-db

echo "Setup completed successfully!"

# Start development server
echo "Starting development server..."