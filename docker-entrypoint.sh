#!/bin/sh

# Load environment variables from .env file
if [ -f .env ]; then
  export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
fi
echo "env loaded"
# Start PostgreSQL
exec "$@"