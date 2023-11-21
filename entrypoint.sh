#!/bin/bash

# Function to start the application
start_app() {
    echo "Starting the application..."
    ./main
}

# Retry loop
max_retries=10
retry_interval=10

for ((i=1; i<=$max_retries; i++)); do
    start_app
    exit_code=$?

    if [ $exit_code -eq 0 ]; then
        echo "Application started successfully."
        exit 0
    else
        echo "Error starting application (attempt $i/$max_retries). Retrying in $retry_interval seconds..."
        sleep $retry_interval
    fi
done

echo "Failed to start the application after $max_retries attempts. Exiting."
exit 1
