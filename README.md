# Graphql Client

Another graphql written in go

## Why

Coming from using graphql in node and python I found writing integration tests always very tedious and error prone.

This library is inspired by relay and other go code generation libraries like ent and gqlgen

Given a graphql schema and graphql queries this project will generate a type safe client

See this [test](https://github.com/stackworx-go/graphql/blob/master/internal/integration/client_test.go) for example use

## Example

Given the Query

```
query TodosQuery {
  todos {
    id
    text
    done
    user {
      id
      name
    }
  }
}
```

Client Usage

```
client := Client{
	Url: "https://someserver.com",
}

data, err := client.TodosQuery()

... Consume the data
```

## Quickstart

```
go get github.com/stackworx-go/graphql
```

```
go run github.com/stackworx-go/graphql/generate \
    --queries  "./**/*.graphql" \
    --schema schema.graphqls \
    --destination "./client.go" \
    --packageName "pkgname"
```

## Known Issues and Limitations

- One Query per file
- Query names must be unique
- Mutation Queries must end in Mutation, Queries must end in Mutation
- Non inline fragments are not supported
