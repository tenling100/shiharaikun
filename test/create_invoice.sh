#!/bin/bash

grpcurl -plaintext -d '{
        "status": "unpaid",
        "amount": 50000,
        "repayment_date": "2024-10-01T00:00:00Z"
}' localhost:50051 invoice.InvoiceService.CreateInvoice
