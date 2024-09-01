#!/bin/bash

grpcurl -plaintext -d '{
        "name": "tenling2",
        "email": "tenling2@gmail.com",
        "company_id": 1
}' localhost:50051 invoice.InvoiceService.CreateUser
