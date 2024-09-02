#!/bin/sh

# Wait for MySQL to be ready
/wait-for-mysql.sh

# Run the migration script
./migrate

# Start the application
exec "$@"