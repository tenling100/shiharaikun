
# About the API Coding Test

## Introduction

To all candidates,

Thank you for proceeding with the selection process for UPSIDER.

The purpose of this test is to design and implement a REST API to launch a fictitious web service for corporations called "Super Payment-kun.com". This web service allows users to register invoice data for future payment dates, and automatically make bank transfers even if there is no balance on the due date. This enables users to delay cash outflows by up to a month, making it a convenient web service.

In this coding test, you are required to design and implement a REST API that meets the following business requirements. Specifically, you will develop REST API endpoints based on the following user stories:

* "As a user, I want to create new invoice data so that I can ensure payment processing on the due date."
* "As a user, I want to retrieve a list of invoice data that will incur payments within a specified period so that I can check what payments have been registered."

Please use Golang as the programming language.
There are no specific framework requirements, and you are free to use other libraries as well.

Also, for this coding test, there is no specified database, but it is assumed that an RDBMS will be used instead of NoSQL.

The test is expected to take about 2-3 hours.

You will be asked to manage the source code using GitHub and submit the URL by pushing it as a Public Repository.

Please showcase your full capabilities and develop excellent software!

## Task/Conditions

### Design of REST API Endpoints

* POST /api/invoices
  * Create new invoice data
    * The `invoice amount` should be calculated automatically
      * The `payment amount` plus a 4% fee and the tax on the fee should be considered the invoice amount
      * Example: For a payment amount of `10,000`, the invoice amount would be `10,000 + (10,000 * 0.04 * 1.10) = 10,440`
* GET /api/invoices
  * Retrieve a list of invoice data that will incur payments within a specified period

### Data Model

* Company
  * Company Name
  * Representative Name
  * Phone Number
  * Postal Code
  * Address
* User (linked to a company)
  * Name
  * Email Address
  * Password
* Client (linked to a company)
  * Company Name
  * Representative Name
  * Phone Number
  * Postal Code
  * Address
* Client Bank Account (linked to a client)
  * Client ID
  * Bank Name
  * Branch Name
  * Account Number
  * Account Name
* Invoice Data (linked to a company and client)
  * Issue Date
  * Payment Amount
  * Fee
  * Fee Rate
  * Tax
  * Tax Rate
  * Invoice Amount
  * Payment Due Date
  * Status (Unprocessed, Processing, Paid, Error)

### Test Cases

* User can create new invoice data
  * Request succeeds and HTTP status code 200 is returned
  * The response contains the created invoice data
* User can retrieve a list of invoice data that will incur payments within a specified period
  * Request succeeds and HTTP status code 200 is returned
  * The response contains the invoice data within the specified period

## Coding Test Evaluation Criteria

* Consideration of the responsibilities and roles of classes, methods, structs, functions, etc.
* Understanding of SOLID principles and best practices in architecture
* Consideration of coding style and readability
* Whether the code is managed in appropriate commit units and is in a state that makes code review easy
* Consideration of authentication and authorization
* Consideration of handling sensitive information
* Whether appropriate test code is implemented
* Whether it can deliver sufficient performance when there is a large amount of test data or high concurrent access
* Whether appropriate error handling is performed when errors or exceptions occur
* Whether knowledge necessary for other developers to code is documented as needed
