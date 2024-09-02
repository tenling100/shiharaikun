#!/bin/bash

curl -X POST http://localhost:8080/api/invoice \
-H "Content-Type: application/json" \
-d '{
        "status": "UNPAID",
        "company_id": 1,
        "creator_id": 1,
        "client_id": 1,
        "issue_date": "2024-09-01T00:00:00Z",
        "invoice_amount": 10000,
        "payment_due_date": "2024-10-01T00:00:00Z"
    }'