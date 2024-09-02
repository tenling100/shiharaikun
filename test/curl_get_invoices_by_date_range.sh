#!/bin/bash

curl -X POST http://localhost:8080/invoices/date-range \
-H "Content-Type: application/json" \
-d '{
        "start_date": "2024-08-01T00:00:00Z",
        "end_date": "2024-08-30T00:00:00Z"
}'
