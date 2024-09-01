#!/bin/bash

grpcurl -plaintext -d '{
        "name": "Example Company",
        "representative": "John Doe",
        "phone": "123-456-7890",
        "postal_code": "12345",
        "address": "123 Main St, Anytown, USA"
}' localhost:50051 invoice.InvoiceService.CreateCompany
