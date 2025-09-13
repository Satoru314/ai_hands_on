#!/bin/bash

# Task Management API Test Script
# Usage: ./e2e/scripts/test_api.sh

BASE_URL="http://localhost:8080"

echo "🚀 Testing Task Management API"
echo "================================"

# Health Check
echo "1. Health Check"
curl -X GET "$BASE_URL/health" | jq .
echo -e "\n"

# Create Task
echo "2. Creating a task..."
TASK_ID=$(curl -s -X POST "$BASE_URL/api/v1/tasks" \
  -H "Content-Type: application/json" \
  -d '{"content": "Learn Go programming", "due_date": "2024-12-31T23:59:59Z"}' | jq -r '.id')
echo "Created task with ID: $TASK_ID"
echo -e "\n"

# Get All Tasks
echo "3. Getting all tasks..."
curl -s -X GET "$BASE_URL/api/v1/tasks" | jq .
echo -e "\n"

# Get Task by ID
echo "4. Getting task by ID: $TASK_ID"
curl -s -X GET "$BASE_URL/api/v1/tasks/$TASK_ID" | jq .
echo -e "\n"

# Update Task
echo "5. Updating task status to 'in_progress'..."
curl -s -X PUT "$BASE_URL/api/v1/tasks/$TASK_ID" \
  -H "Content-Type: application/json" \
  -d '{"content": "Learn Go programming - Updated", "status": "in_progress", "due_date": "2024-12-31T23:59:59Z"}' | jq .
echo -e "\n"

# Complete Task
echo "6. Completing the task..."
curl -s -X PUT "$BASE_URL/api/v1/tasks/$TASK_ID" \
  -H "Content-Type: application/json" \
  -d '{"content": "Learn Go programming - Completed!", "status": "completed", "due_date": "2024-12-31T23:59:59Z"}' | jq .
echo -e "\n"

# Delete Task
echo "7. Deleting task ID: $TASK_ID"
curl -s -X DELETE "$BASE_URL/api/v1/tasks/$TASK_ID"
echo "Task deleted"
echo -e "\n"

echo "✅ API testing completed!"