# Online Payment Platform: Requirements

## Introduction

The purpose of this document is to define the requirements for the Online Payment Platform. The Online Payment Platform is a service that processes payments on behalf of the merchant. It is responsible for validating the customer's card information, processing the payment, and sending the payment details to the acquiring bank.

### Requirements

#### REQ-001: Register Merchant

The system should allow the merchant to register with the Online Payment Platform. The merchant should provide the following information:

- Name: *required*
- Email: *required*
- Phone: *required*
- Address: *required*
- Country: *required*
- Currency: *required* (only USD are supported for now).
- Application URL: *required* (the URL of the merchant's e-commerce platform).
- Webhook URL: *required* (the URL where the Online Payment Platform will send notifications, just one URL is supported for now).

> The Application API Key and the Webhook API Key should be generated for the merchant and used to authenticate the requests.

#### REQ-002: Register Customer

The system should allow merchants to register customers. The customer should provide the following information:

- Name: *required*
- Email: *required*
- Phone: *required*
- Address: *required*
- Country: *required*

#### REQ-003: Manage Payment Methods - List Cards

The system should allow customers to list their payment methods. The system should return the following information:

- Cardholder Name
- Card Number (only the last 4 digits)
- Expiration Date
- Address
- Country
- Default (if the card is the default payment method).

This endpoint should be paginated and allow the customer to filter the cards by the default status.

#### REQ-004: Manage Payment Methods - Register Card

The system should allow customers to register payment methods. The customer should provide the following information:

- Cardholder Name: *required*
- Card Number: *required*
- Expiration Date: *required*
- Security Code: *required*
- Address: *required*
- Country: *required*
- Default: *optional* (if the customer wants to set this card as the default payment method).

In a real-world scenario, this should be managed by tokenization and encryption, but for the purpose of this project, we are not going to implement this feature. Implementing will make the system more secure and reliable but there is not enough time to implement it.

#### REQ-004: Update Payment Method

The system should allow customers to update their payment methods. The customer should provide the following information:

- Cardholder Name: *required*
- Card Number: *required*
- Expiration Date: *required*
- Security Code: *required*
- Address: *required*
- Country: *required*
- Default: *optional* (if the customer wants to set this card as the default payment method).

#### REQ-005: Delete Payment Method

The system should allow customers to delete their payment methods.

#### REQ-006: Process Payment - Request Payment

The system should allow the merchant to register a payment. The merchant should provide the following information:

- Customer ID: *required*
- Amount: *required*
- Currency: *required* (only USD are supported for now).
- Description: *required*
- Payment Method ID: *required*
- Context: *optional* (additional information about the payment).

Requested payments should created in a `PENDING` status.

#### REQ-007: Process Payment - Update Payment

When the acquiring bank processes the payment, the Online Payment Platform should update the payment status. The acquiring bank should provide the following information:

- Payment ID: *required*
- Status: *required* (one of: `APPROVED`, `DECLINED`, `ERROR`).
- Authorization Code: *optional* (the authorization code provided by the acquiring bank).
- Response Code: *optional* (the response code provided by the acquiring bank).
- Response Message: *optional* (the response message provided by the acquiring bank).
- Signature: *required* (a signature to validate the payment status).

> The Signature validation should be implemented for next versions.

#### REQ-008: Process Payment - Notify Merchant

When the payment status is updated, the Online Payment Platform should notify the merchant. The system should send a notification to the merchant's webhook URL with the following information:

- Event: *required* (one of: `PAYMENT_CREATED`, `PAYMENT_APPROVED`, `PAYMENT_DECLINED`, `PAYMENT_ERROR`).
- Data: *required* (the payment details).
- Timestamp: *required* (the timestamp when the notification was sent).
- Signature: *required* (a signature to validate the notification).
- Context: *optional* (additional information about the notification).

> The Signature validation should be implemented for next versions.

#### REQ-009: Payment History - List Payments

The system should allow the merchant to list the payments. The system should return the following information:

- Payment ID
- Customer ID
- Amount
- Currency
- Description
- Payment Method ID
- Status
- Authorization Code
- Status
- Timestamp

Filtering by Customer ID, Payment Method ID, Status, and Timestamp should be supported.

#### REQ-010: Payment History - Get Payment

The system should allow the merchant to get the payment details. The system should return the following information:

- Payment ID
- Customer ID
- Amount
- Currency
- Description
- Payment Method ID
- Status
- Authorization Code
- Status
- Timestamp
- Context
- Signature
- Response Code
- Response Message

#### REQ-011: Refunds - Request Refund

The system should allow to request a refund. The merchant should provide the following information:

- Payment ID: *required*
- Amount: *required*
- Description: *required*
- Context: *optional* (additional information about the refund).
- Payment Method ID: *optional* (if the refund should be processed to a different payment method).

The refund should be created in a `PENDING`, after the acquiring bank processes the refund, the Online Payment Platform should update the refund status.

#### REQ-012: Refunds - Update Refund

When the acquiring bank processes the refund, the Online Payment Platform should update the refund status. The acquiring bank should provide the following information:

- Refund ID: *required*
- Status: *required* (one of: `APPROVED`, `DECLINED`, `ERROR`).
- Authorization Code: *optional* (the authorization code provided by the acquiring bank).
- Response Code: *optional* (the response code provided by the acquiring bank).
- Response Message: *optional* (the response message provided by the acquiring bank).
- Signature: *required* (a signature to validate the refund status).

#### REQ-013: Refunds - Notify Merchant

When the refund status is updated, the Online Payment Platform should notify the merchant. The system should send a notification to the merchant's webhook URL with the following information:

- Event: *required* (one of: `REFUND_CREATED`, `REFUND_APPROVED`, `REFUND_DECLINED`, `REFUND_ERROR`).
- Data: *required* (the refund details).
- Timestamp: *required* (the timestamp when the notification was sent).
- Signature: *required* (a signature to validate the notification).
- Context: *optional* (additional information about the notification).

#### REQ-014: Refunds - List Refunds

- The system should allow the merchant to list the refunds. The system should return the following information:

- Refund ID
- Payment ID
- Amount
- Description
- Status
- Authorization Code
- Status
- Timestamp

Filtering by Payment ID, Status, and Timestamp should be supported.

#### REQ-015: Refunds - Get Refund

The system should allow the merchant to get the refund details. The system should return the following information:

- Refund ID
- Payment ID
- Amount
- Description
- Status
- Authorization Code
- Status
- Context
- Signature
- Response Code
- Response Message
- Timestamp

#### REQ-016: Security - Authentication

The system should use API keys to authenticate the requests. The API keys should be generated for the merchant and used to authenticate the requests.

#### REQ-017: Security - Signature Validation

The system should use signature validation to validate the payment status and the notifications.

#### REQ-018: Audit - Trail

The system should keep an audit trail of the requests and the responses.