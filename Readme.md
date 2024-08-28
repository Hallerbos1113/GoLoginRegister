## Description

To get started with Gin, you'll typically create a new Go project, import the Gin package, define your routes, middleware, and handlers, and then run the application.

I develope the user registration for api for login and update the user details with using database Postgres.

## GO Installation

[GO Install] -- Follow this link to install golang in machine.

## Run

Run our go API:

```sh
go run main.go || go run .
```

** To build and run: **

Build the script:

```sh
    go build -o build/bin
```

Run the build file:

```sh
    ./build/bin
```

## APIs

| Usage    | API                                    | Type              |
| -------- | -------------------------------------- | ----------------- |
| register | https://localhost:8080/api/v1/register | post              |
| Login    | https://localhost:8080/api/v1/login    | post              |
| users    | https://localhost:8080/api/v1/users    | GET, DELETE, POST |

HOW TO PLAY THIS TEST?

Please run go server as code

> go run .

And open "ZZZ" folder and double click "html.html"
Please follow your requirement step by step.
Thank you.

This backend project needs that you have set up Postgres on your machine.
Postgres has to contain the "testdb" database including "tb_test" table.
You may create database and table manually.

    host			= "localhost"
    port			= 5432
    user			= "postgres"
    password	=	"1234"
    dbname		= "testdb"

My skill is not good so that it couldn't give you a pleasure, I'm afraid.

Give an advicement, thank you.

If you need me to modify code, contact me and update my code in Github.
