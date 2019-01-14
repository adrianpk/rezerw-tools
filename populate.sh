#!/bin/bash
echo "Populating..."
AWS_REGION=eu-west-1 go run init-db.go
echo "DynamoDB database populated."
