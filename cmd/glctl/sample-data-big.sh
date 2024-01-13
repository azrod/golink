#!/bin/bash

# This script is used to generate sample data for the application.

# Set variable CMD
CMD="./glctl --host http://localhost:8081"

# $CMD add ns default
# $CMD add ns myproject

# $CMD add li grafana /grafana https://grafana.example.com

# Generate 20 random namespaces
# for i in {1..20}
# do
#   $CMD add ns ns$i
# done

# Generate 100 random links by randomly selecting a namespace
for i in {1..10000}
do
  # Generate a random number between 1 and 20
  n=$(( ( RANDOM % 20 )  + 1 ))
  # Generate a random link name
  linkName=$(openssl rand -hex 10)
  # Generate a random link URL
  linkURL=$(openssl rand -hex 10)
  # Generate a random link description
  linkDescription=$(openssl rand -hex 15)
  # Generate a random namespace
  namespace="ns$n"
  # Add the link
  $CMD -n $namespace add li $linkName /$linkURL https://$linkDescription.example.com
done
