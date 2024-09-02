# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [invoice.proto](#invoice-proto)
    - [DateRangeRequest](#invoice-DateRangeRequest)
    - [Invoice](#invoice-Invoice)
    - [InvoiceRequest](#invoice-InvoiceRequest)
    - [InvoiceResponse](#invoice-InvoiceResponse)
    - [InvoicesResponse](#invoice-InvoicesResponse)
    - [company](#invoice-company)
    - [user](#invoice-user)
  
    - [InvoiceStatus](#invoice-InvoiceStatus)
  
    - [InvoiceService](#invoice-InvoiceService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="invoice-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## invoice.proto



<a name="invoice-DateRangeRequest"></a>

### DateRangeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| end_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="invoice-Invoice"></a>

### Invoice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [google.protobuf.UInt64Value](#google-protobuf-UInt64Value) |  |  |
| company | [company](#invoice-company) |  |  |
| creator | [user](#invoice-user) |  |  |
| client | [company](#invoice-company) |  |  |
| issue_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| payment_amount | [float](#float) |  |  |
| fee | [float](#float) |  |  |
| fee_rate | [float](#float) |  |  |
| tax | [float](#float) |  |  |
| tax_rate | [float](#float) |  |  |
| invoice_amount | [float](#float) |  |  |
| payment_due_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| status | [InvoiceStatus](#invoice-InvoiceStatus) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="invoice-InvoiceRequest"></a>

### InvoiceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| company_id | [uint64](#uint64) |  |  |
| creator_id | [uint64](#uint64) |  |  |
| client_id | [uint64](#uint64) |  |  |
| issue_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| invoice_amount | [float](#float) |  |  |
| payment_due_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| status | [InvoiceStatus](#invoice-InvoiceStatus) |  |  |






<a name="invoice-InvoiceResponse"></a>

### InvoiceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| invoice | [Invoice](#invoice-Invoice) |  |  |






<a name="invoice-InvoicesResponse"></a>

### InvoicesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| invoices | [Invoice](#invoice-Invoice) | repeated |  |






<a name="invoice-company"></a>

### company



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [google.protobuf.UInt64Value](#google-protobuf-UInt64Value) |  |  |
| name | [string](#string) |  |  |
| representative | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| postal_code | [string](#string) |  |  |
| address | [string](#string) |  |  |






<a name="invoice-user"></a>

### user



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [google.protobuf.UInt64Value](#google-protobuf-UInt64Value) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| company_id | [uint64](#uint64) |  |  |





 


<a name="invoice-InvoiceStatus"></a>

### InvoiceStatus
invoice status enum

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| UNPAID | 1 |  |
| PAID | 2 |  |
| REJECTED | 3 |  |
| PENDING | 4 |  |
| CANCELLED | 5 |  |
| DELETED | 6 |  |


 

 


<a name="invoice-InvoiceService"></a>

### InvoiceService
Invoice service definitions

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateInvoice | [InvoiceRequest](#invoice-InvoiceRequest) | [InvoiceResponse](#invoice-InvoiceResponse) |  |
| GetInvoicesByDateRange | [DateRangeRequest](#invoice-DateRangeRequest) | [InvoicesResponse](#invoice-InvoicesResponse) |  |
| CreateCompany | [company](#invoice-company) | [company](#invoice-company) |  |
| CreateUser | [user](#invoice-user) | [user](#invoice-user) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

