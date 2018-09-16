# golang-todolist
A todolist app written in Golang for educational purposes.

I'm writing this application to educate myself with the Go programming language, and to develop an overall structure and style for developing future web applications in Go.

## Local Setup

Install all dependencies:

```go get ./...```

Add the following to a .env file in the main src directory of the project:

```
MODE=dev
PORT=8080
MYSQL_USERNAME=<username>
MYSQL_PASSWORD=<password>
MYSQL_DB=golang_todolist
AWS_S3_UPLOAD_BUCKET_NAME=
AWS_S3_REGION=
AWS_S3_ACCESS_KEY_ID=
AWS_S3_SECRET_ACCESS_KEY=
SESSION_MYSQL_USERNAME=
SESSION_MYSQL_PASSWORD=
SESSION_MYSQL_DB=
SESSION_MYSQL_HOST=
SESSION_MYSQL_TABLE=
SESSION_NAME=s
EMAIL_HOST=
EMAIL_FROM=
EMAIL_USERNAME=
EMAIL_PASSWORD=
```

Then load the database with test data:

```./reloaddb.sh```

And run it locally:

```./run.sh```

## Requirements

**Execution**
* The development environment will listen to a local port, and the production environment will run as an AWS Lambda function.
* Environment variables are read from a .env in the development environment, and entered directly onto the Lambda function in production.

**Functionality**
* CRUD functionality for todo items and todo lists.

## Dependencies

* Bootstrap 4.1 CSS Framework
* github.com/joho/godotenv
	To store dev environment variables in a .env file
* github.com/akrylysov/algnhsa
	To trigger the Lambda function from API Gateway while keeping w http.ResponseWriter, r *http.Request in handler functions.
* github.com/gorilla/mux
	Used for improved routing.
* upper.io/db.v2/mysql
	Used as data access layer.
* https://github.com/srinathgs/mysqlstore
	Gorilla's Session Store Implementation for MySQL
