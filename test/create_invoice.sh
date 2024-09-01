#!/bin/bash

grpcurl -plaintext -d '{
        "status": "UNPAID",
        "company_id": 1,
        "creator_id": 2,
        "client_id": 2,
        "issue_date": "2024-09-01T00:00:00Z",
        "invoice_amount": 10,
        "payment_due_date": "2024-10-01T00:00:00Z"
}' localhost:50051 invoice.InvoiceService.CreateInvoice
