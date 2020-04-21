#!/bin/bash
set -e
go run cmd/generate/main.go \
    --queries  "./internal/integration/**/*.graphql" \
    --schema internal/integration/graph/schema.graphqls \
    --destination "./internal/integration/graphqlclient.go" \
    --packageName "integration"