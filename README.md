# About project
This project is for my own knowledgement about use cases of golang, and nothing more useful that implementing ecommerce for apply all my experience in golang.


# Basic usage of migrations
This application contains into path: `/migrations/migrate.go` the functions for create migrations using the package [golang-migrate](https://github.com/golang-migrate/migrate/releases).

## Create migration Up and Down files into `db/migrations` folder:
```bash
$ go run main.go migrate create create_name_of_table
```

## Up specific migration:
```bash
$ go run main.go migrate up 1
```

## Up all new migrations:
```bash
$ go run main.go migrate up
```

## Down specific migration:
```bash
$ go run main.go migrate down 1
```

## Down all new migrations:
```bash
$ go run main.go migrate down
```
