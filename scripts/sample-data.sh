#!/bin/bash

# This script is used to generate sample data for the application.

# Set variable CMD
CMD="go run cmd/glctl/main.go"

$CMD add ns default
$CMD add ns myproject

$CMD add li grafana /grafana https://grafana.example.com
$CMD -n myproject add li jira /jira https://jira.example.com
$CMD -n myproject add li git /git https://git.myproject.myorgnization.local
$CMD -n myproject add li doc /doc https://wiki.myproject.myorgnization.local/myproject
$CMD -n myproject add li jenkins /jenkins https://jenkins.myproject.myorgnization.local
