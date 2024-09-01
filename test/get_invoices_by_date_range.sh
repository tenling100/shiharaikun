#!/bin/bash

grpcurl -plaintext -d '{
        "start_date": "2024-08-01T00:00:00Z",
        "end_date": "2024-08-30T00:00:00Z"
}' localhost:50051 invoice.InvoiceService.GetInvoicesByDateRange
