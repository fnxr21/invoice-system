# Items API Spec

## Create Item

Endpoint : POST /api/v1/item

Request Body :

```json
{
  "item_name": "design",
  "item_type": "service",
}
```

Response Body (Success) :

```json
{
  "data": {
    "item_name": "design",
    "item_type": "service",
  }
}
```

Response Body (Failed) :

```json
{
  "errors": "customer name is required"
}
```

## List Item

Endpoint : GET /api/v1/items

Response Body (Success) :

```json
{
  "data": [
    {
      "id": "1",
      "item_name": "design",
      "item_type": "service",
    },
    {
      "id": "2",
      "item_name": "developlement",
      "item_type": "service",
    }
  ]
}
```

Response Body (Failed) :

```json
{
  "errors": "customer name is required"
}
```
