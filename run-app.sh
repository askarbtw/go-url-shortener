#!/bin/bash

# Function to stop processes on Ctrl+C
cleanup() {
    echo "Stopping all processes..."
    kill $BACKEND_PID $FRONTEND_PID 2>/dev/null
    exit 0
}

# Trap Ctrl+C
trap cleanup SIGINT

# Start the Go backend
echo "Starting Go backend..."
cd "$(dirname "$0")"
go run main.go &
BACKEND_PID=$!

# Wait a moment for backend to start
sleep 2

# Start the React frontend
echo "Starting React frontend..."
cd frontend
npm run dev &
FRONTEND_PID=$!

# Wait for processes to finish
wait $BACKEND_PID $FRONTEND_PID 