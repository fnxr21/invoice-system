# Customer API Spec

## Create Customer

Endpoint : POST /api/v1/customer

Request Body :

```json
{
  "customer_name": "fandi nur",
  "customer_address": "ST.Teluk Gong,North Jakarta Indonesia"
}
```

Response Body (Success) :

```json
{
  "data": {
    "customer_name": "fandi nur",
    "customer_address": "ST.Teluk Gong,North Jakarta Indonesia"
  }
}
```

Response Body (Failed) :

```json
{
  "errors": "customer name is required"
}
```

```json
{
  "errors": "customer address is required"
}
```

## List Customer

Endpoint : GET /api/v1/customers

Response Body (Success) :

```json
{
  "data": {
    "customer_name": "fandi nur",
    "customer_address": "ST.Teluk Gong,North Jakarta Indonesia"
  },
  "data": {
    "customer_name": "fandi ",
    "customer_address": "ST.Sunter Agung,North Jakarta Indonesia"
  }
}
```
