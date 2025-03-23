#!/bin/bash

# Database migration script

# Set environment variables
export DATABASE_URL="your_database_url_here"
export MIGRATION_DIR="./migrations"

# Run migrations
echo "Running database migrations..."
migrate -path $MIGRATION_DIR -database $DATABASE_URL up

echo "Migrations completed."