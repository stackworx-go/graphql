#!/bin/bash
set -e
go run cmd/main.go \
    --queries  "./internal/integration/**/*.graphql" \
    --schema internal/integration/graph/schema.graphqls \
    --destination "./internal/integration/client.go" \
    --packageName "integration"