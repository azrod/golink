#!/bin/bash

# This script is used to generate sample data for the application.

# Set variable CMD
CMD="go run cmd/main.go"

$CMD add ns demo
$CMD add ns tool

$CMD -n demo add li google /google https://www.google.com
$CMD -n tool add li golink /golink https://github.com/azrod/golink