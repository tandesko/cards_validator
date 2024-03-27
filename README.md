# Cards validator

## Description

Cards validator verifies credit cards by validating its number, month and year.
Number is validated with Luhn method and IIN(Issuer identification number).

### To run app in docker
```
docker-compose up
```

### Json example
```
{
    "data": {
        "attributes": {
            "cardnumber": "4098939612891",
            "month": "9",
            "year": "2024"
        }
    }
}
```

### Response example
```
{
    "data": {
        "id": "1",
        "type": "validate-card-response",
        "attributes": {
            "valid": true
        }
    },
    "included": []
}
```