#!/bin/bash

curl -X POST http://localhost:8080/company \
-H "Content-Type: application/json" \
-d '{
        "name": "Example Company",
        "representative": "John Doe",
        "phone": "123-456-7890",
        "postal_code": "12345",
        "address": "123 Main St, Anytown, USA"
    }'
