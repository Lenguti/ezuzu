# EZUZU

### Requirements:
```
You will need docker installed on your machine to run the application locally.
https://docs.docker.com/desktop/install/mac-install/ (mac link)

I've utilzied goose for handling db migrations.
https://github.com/pressly/goose
brew isntall goose.

Copy the .env.sample file to .env and replace the values with your db configuration.
Copy over the .env.service.sample file to:
  .env.payments
  .env.property
  .env.tennant_portal
and replace the values with your db configuration.

Once the you have docker running you can spin up the web server and db by running:
`make run`

Once the services are running you can validate each servics status by curl'ing:
PAYMENTS:       `curl -vvv -X GET 'http://localhost:8000/v1/status'`
PROPERTY:       `curl -vvv -X GET 'http://localhost:8001/v1/status'`
TENNANT_PORTAL: `curl -vvv -X GET 'http://localhost:8002/v1/status'`

Once the webservers have been validated you can run migrations by running:
```

### Routes

### MODELS
```
API Error
{
  "error": {
    "code": "string ENUM", (BAD_REQUEST, INTERNAL_SERVER_ERROR, NOT_FOUND)
    "message": "string",
    "status_code": int,
    "details": {
      "string": "string"
    }
  }
}
```
