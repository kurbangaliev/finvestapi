#!/bin/bash
# Start the backend server
./backend-server &
# Start the frontend server
./frontend-server &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?