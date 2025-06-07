#!/bin/bash

# Black Box Challenge - Go Implementation
# This script takes three parameters and outputs the reimbursement amount
# Usage: ./run.sh <trip_duration_days> <miles_traveled> <total_receipts_amount>

# Check if correct number of arguments provided
if [ $# -ne 3 ]; then
    echo "Usage: $0 <trip_duration_days> <miles_traveled> <total_receipts_amount>" >&2
    exit 1
fi

# Run the Go CLI app with the provided arguments
go run main.go "$1" "$2" "$3"
