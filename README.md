# Login api
A simple login api endpoint that follows MVC architecture. It only handles login and not signup.
Upon successful login the endpoint sends a jwt that is signed with a secret key that is known to other services that may need to verify it.

## Requirements
1. go `1.24.x`

## Setup
1. Install dependecnies with `go get .`
2. Start the postgres server with `docker-compose up`
3. build with `go build .`
4. run `./login-api`
5. or simply run `make start-process` instead
6. kill the go process with ctrl-c, to stop the container, run `make stop-process`

### Note
an environment file needs to be setup, check example `.env` file for context.


## Routes
```http
use for testing
Request
POST /login
Content-Type: application/json

{
  "id":"testid",
  "password":"test-pass"
}

Response
{
    "token": "JWT token with 24h expiry"
}
```


## Example env file
```bash
DSN="host=localhost user=appuser password=apppassword dbname=appdb port=5432 sslmode=disable"
JWT_SECRET="a very secret token"
```
