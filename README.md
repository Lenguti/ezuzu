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
and replace the values with your db configuration.

Once the you have docker running you can spin up the web server and db by running:
`make run`

Once the services are running you can validate each servics status by curl'ing:
PAYMENTS:       `curl -vvv -X GET 'http://localhost:8000/v1/status'`
PROPERTY:       `curl -vvv -X GET 'http://localhost:8001/v1/status'`

Once the webservers have been validated you can run migrations by running:
make migration SERVICE={{service}} DB_USER={{user}} DB_PASS={{pass}} DB_NAME={{name}} DB_PORT={{port}}
```

### NOTES
Documentation notes can be found here:
https://whimsical.com/H3CwGrFebiSWYdSRDK3YEr

To run tests please run `make test`

### Routes
PROPERTY SERVICE			
create manager	          POST	/v1/managers<br>
create property	          POST	/v1/managers/:id/properties<br>
update property	          PUT	/v1/managers/:id/properties/:id<br>
get property	            GET	/v1/managers/:id/properties/:id<br>
list properties	          GET	/v1/managers/:id/properties<br>
create tennant	          POST	/v1/managers/:id/properties/:id/tenants<br>
update tennant	          PUT	/v1/managers/:id/properties/:id/tenants/:id<br>
get tennant               GET	/v1/managers/:id/properties/:id/tenants/:id<br>
list tennants of property	GET	/v1/managers/:id/properties/:id/tenants<br>
			
PAYMENTS SERVICE			
create invoice	POST	/v1/managers/:id/properties/:id/invoices<br>
create payment	POST	/v1/tenants/:id/invoices/:id/payments<br>
payment history	GET	/v1/tenants/:id/payments/history<br>

Route Request / Responses can be viewed in ./docs/payments.http ./docs/property.http

### MODELS
```
PROPERTY MANAGER
{
  "manager": {
    "id": "uuid",
    "entity": "string",
    "createdAt": int,
    "updatedAt": int
  }
}

PROPERTY
{
  "property": {
    "id": "uuid",
    "address": "string",
    "name": "string",
    "rent": float,
    "type": "string enum", (APARTMENT, HOME)
    "unitNumber": int,
    "createdAt": int,
    "updatedAt": int
  }
}

TENANT
{
  "tenant": {
    "id": "uuid",
    "type": "string enum", (PRIMARY, SECONDARY)
    "firstName": "string",
    "lastName": "string",
    "dateOfBirth": "string", (format yyyy-mm-dd)
    "ssn": "string", (hashedout "####") 
    "createdAt": int,
    "updatedAt": int
  }
}

INVOICE
{
  "invoice": {
    "id": "uuid",
    "tenantId": "uuid",
    "amount": float,
    "dueDate": "string",
    "createdAt": int,
    "updatedAt": int
  }
}

PAYMENT
{
  "payment": {
    "id": "uuid",
    "invoiceId": "uuid",
    "amount": float,
    "createdAt": int,
    "updatedAt": int
  }
}

PAYMENT HISTORY
{
  "payment_history": [
    {
      "month": "string",
      "total": float,
      "timestamp": int
    }
  ]
}

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

In practice, there is another component of processing rent payments and managing accounts and leases. Let’s say that we want our system to be able to create a report showing how much revenue a given property generated within a historical period, considering payments made by all tenants. Describe how you would extend the service to store rent payments history and make queries that allow you to pull revenue history from a given property or apartment over a given time window. Note: Implementation of this part is not required, just a description and commentary in the README file.

```
I think I would go about this the same way as payment history. I would create another view table pulling from the properties, invoices, and payments to aggregate all the relevant data. The route would need to accept query params of a date range to be queried with and the
query would need to include a BETWEEN clause.

route GET /v1/properties/:id/revenue?start_date={{start}}&end_date={{end}}
```