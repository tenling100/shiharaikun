#!/bin/bash

curl -X POST http://localhost:8080/user \
-H "Content-Type: application/json" \
-d '{
        "name": "tenling3",
        "email": "tenling3@gmail.com",
        "company_id": 1
}'
