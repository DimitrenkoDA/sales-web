## Core API

## Installation

- Install Go 1.13: https://golang.org/dl


## Development

First of all you should create the database:

```
psql -d postgres -c 'CREATE DATABASE dimon_sales_web_core_api_development'
psql -d dimon_sales_web_core_api_development -f db/schema.sql
```

Now we can build and start the server. This project requires Go 1.13. In order to build and run this project execute following command:


```
go run cmd/main.go
```

Check using:

```
curl -v -XGET http://localhost:8080/dealers | json_pp
```