#!/bin/bash

curl -X GET "http://localhost:8080/api/invoices?start_date=2024-09-01T00:00:00Z&end_date=2024-09-30T00:00:00Z" \
-H "Content-Type: application/json"
