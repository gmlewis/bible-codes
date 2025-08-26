#!/bin/bash -ex

# # Check if exactly one argument is provided
# if [ $# -ne 1 ]; then
#     echo "Usage: $0 <N>"
#     echo "Where N is a positive integer"
#     exit 1
# fi

# Get the argument and validate it's a positive integer
N="999"

# Check if N is a valid positive integer
if ! [[ "$N" =~ ^[0-9]+$ ]] || [ "$N" -lt 2 ]; then
    echo "Error: Argument must be a positive integer >= 2"
    exit 1
fi

echo "Running commands for n=2 to n=$N"

# Loop from 2 to N
for ((n=2; n<=N; n++)); do
    echo "Running: go run ../cmd/gen-skip-data/main.go -skip $n"

    # Run the command and check if it fails
    if ! go run ../cmd/gen-skip-data/main.go -skip "$n"; then
        echo "Command failed with n=$n. Stopping execution."
        exit 1
    fi

    echo "Command completed successfully for n=$n"
done

echo "All commands completed successfully!"
