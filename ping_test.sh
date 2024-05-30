#!/bin/bash

# Loop to run the curl command 10 times to test the counter is increasing
for i in {1..100}
do
  # Print the command to the terminal
  echo "$i. Executing command: curl localhost:8081/ping"

  # Execute the curl command
  curl muktoacademy.com
done
