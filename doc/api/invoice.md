# Invoice API Spec

## Create Invoice

Endpoint : POST /api/v1/invoice

Request Body :

```json
{
  "issue_date": "2022-01-02 ",
  "due_date": "2022-01-07",
  "subject": "Spring Marketing Campaign",
  "customer_id": 2,
  "items": [
    {
      "id": 1,
      "quantity": "41.00",
      "unit_price": "230.00"
    },
    {
      "id": 1,
      "quantity": "41.00",
      "unit_price": "230.00"
    }
  ]
}
```

Response Body (Success) :

```json
{
  "data": {
    "invoice_id": 001,
    "issue_date": "2022-01-02 ",
    "due_date": "2022-01-07",
    "subject": "Spring Marketing Campaign",
    "customer": {
      "id": 1,
      "customer_name": "fandi nur",
      "customer_address": "ST.Teluk Gong,North Jakarta Indonesia"
    },
    "items": [
      {
        "id": 1,
        "items_name": "Design",
        "quantity": "41.00",
        "unit_price": "230.00"
      },
      {
        "id": 2,
        "items_name": "Design",
        "quantity": "41.00",
        "unit_price": "230.00"
      }
    ]
  }
}
```

Response Body (Failed) :

```json
{
  "errors": "Items is required"
}
```

```json
{
  "errors": "Customer is required"
}
```

## Get invoice

Endpoint : GET /api/v1/invoice/:id

Response Body (Success) :

```json
{
  "data": {
    "invoice_id": 001,
    "issue_date": "2022-01-02 ",
    "due_date": "2022-01-07",
    "subject": "Spring Marketing Campaign",
    "customer": {
      "id": 1,
      "customer_name": "fandi nur",
      "customer_address": "ST.Teluk Gong,North Jakarta Indonesia"
    },
    "items": [
      {
        "id": 1,
        "items_name": "Design",
        "quantity": "41.00",
        "unit_price": "230.00"
      },
      {
        "id": 1,
        "items_name": "Design",
        "quantity": "41.00",
        "unit_price": "230.00"
      }
    ]
  }
}
```

Response Body (Failed) :

```json
{
  "errors": "Not Found"
}
```

## Update invoice

Endpoint : PATCH /api/v1/invoice/:id

Response Body (Success) :

```json
{
  "issue_date": "2022-01-02 ",
  "due_date": "2022-01-07",
  "subject": "Spring Marketing Campaign",
  "customer_id": 2,
  "items": [
    {
      "id": 1,
      "quantity": "41.00",
      "unit_price": "230.00"
    },
    {
      "id": 3,
      "quantity": "41.00",
      "unit_price": "230.00"
    }
  ]
}
```

Response Body (Failed) :

```json
{
  "errors": "Items is required"
}
```

```json
{
  "errors": "Customer is required"
}
```

## Search Invoice

Endpoint : GET /api/contacts

Query Parameter :

- invoice_id : string, invoice id, optional
- issue_date : string, invoice issue date, optional
- due_date : string, invoice due date, optional
- subject : string, invoice subject, optional
- status : string, invoice status, optional
- customer_name : string, customer name, optional
- total_items : string, items total, optional
- page : number, default 1
- size : number, default 10

Request Header :

- X-API-TOKEN : token

Response Body (Success) :

```json
{
  "data": [
    {
      "invoice_id": 001,
      "issue_date": "2022-01-02 ",
      "subject": "Spring Marketing Campaign",
      "total_items": 3,
      "customer_name": "fandi nur",
      "due_date": "2022-01-07",
      "status": "paid"
    },
    {
      "invoice_id": 002,
      "issue_date": "2022-01-02 ",
      "subject": "Spring Marketing Campaign Sudirman",
      "total_items": 2,
      "customer_name": "fandi nur",
      "due_date": "2022-01-07",
      "status": "unpaid"
    }
  ],
  "paging": {
    "current_page": 1,
    "size": 10,
    // "total_page": 10, not yet
  }
}
```
