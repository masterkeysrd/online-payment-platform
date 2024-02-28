# Online Payment Platform: Guide

This guide is intended to provide a quick how to use the Online Payment Platform.

## Merchant Registration

To use the Online Payment Platform, you need to register as a merchant. To do this, you need to provide the following information:

- Your business name
- Your business address
- Your business phone number
- Your business email address
- Your business country
- Your business currency (only USD is supported at the moment)
- Your business website
- Your application webhook (to receive events)

The endpoint to register a merchant is `POST /api/v1/merchants`. The request body should be a JSON object with the following fields:

```json
{
  "name": "Your Business Name",
  "address": "Your Business Address",
  "phone": "Your Business Phone Number",
  "email": "Your Business Email Address",
  "country": "Your Business Country",
  "currency": "Your Business Currency",
  "website": "Your Business Website",
  "webhook_url": "Your Application Webhook"
}
```

Curl Example:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "name": "Your Business Name",
  "address": "Your Business Address",
  "phone": "Your Business Phone Number",
  "email": "Your Business Email Address",
  "country": "Your Business Country",
  "currency": "Your Business Currency",
  "website": "Your Business Website",
  "webhook_url": "Your Application Webhook"
}' "<API_URL>/api/v1/merchants"
```

You will receive a response with the merchant ID and a secret key. You will need to use this secret key to authenticate your requests.

The secret key should be passed in the `Authorization` header as a Bearer token for subsequent requests:

```text
Authorization: Bearer <secret_key>
```

> Note: The secret key is only shown once. Make sure to store it securely.

## Customer Registration

To receive payments, you need to register your customers. To do this, you need to provide the following information:

- Your security key
- Your customer's name
- Your customer's email address
- Your customer's phone number
- Your customer's country
- Your customer's currency (only USD is supported at the moment)

The endpoint to register a customer is `POST /api/v1/customers`. The request body should be a JSON object with the following fields:

```json
{
  "name": "Your Customer's Name",
  "email": "Your Customer's Email Address",
  "phone": "Your Customer's Phone Number",
  "country": "Your Customer's Country",
  "currency": "Your Customer's Currency"
}
```

> Note: Pass the security key in the `Authorization` header as a Bearer token.

Curl Example:

```bash
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer <secret_key>" -d '{
  "name": "Your Customer's Name",
  "email": "Your Customer's Email Address",
  "phone": "Your Customer's Phone Number",
  "country": "Your Customer's Country",
  "currency": "Your Customer's Currency"
}' "<API_URL>/api/v1/customers"
```

You will receive a response with the customer ID.

## Payment Methods

To receive payments, you need to add payment methods to your customers. The Online Payment Platform supports the following payment methods:

- Credit Card (only this method is supported at the moment)

The endpoint to add a payment method is `POST /api/v1/payment_methods`. The request body should be a JSON object with the following fields:

```json
{
    "customer": "Customer ID",
    "type": "card",
    "card": {
        "number": "Card Number",
        "exp_month": "Expiry Month",
        "exp_year": "Expiry Year",
        "cvc": "CVC"
    }
}
```

> Note: Pass the security key in the `Authorization` header as a Bearer token.

Curl Example:

```bash
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer <secret_key>" -d '{
    "customer": "Customer ID",
    "type": "card",
    "card": {
        "number": "Card Number",
        "exp_month": "Expiry Month",
        "exp_year": "Expiry Year",
        "cvc": "CVC"
    }
}' "<API_URL>/api/v1/payment_methods"
```

You will receive a response with the payment method ID, this ID will be used to make payments.

## Payments

To make a payment, you need to provide the following information:

- Security key
- Customer ID
- Payment method ID
- Description (optional)
- Amount
- Currency (only USD is supported at the moment)
- Metadata (optional, not supported at the moment)

The endpoint to make a payment is `POST /api/v1/payments`. The request body should be a JSON object with the following fields:

```json
{
    "customer": "Customer ID",
    "payment_method": "Payment Method ID",
    "description": "Payment Description",
    "amount": "Payment Amount",
    "currency": "Payment Currency",
    "metadata": {
        "key": "value"
    }
}
```

> Note: Pass the security key in the `Authorization` header as a Bearer token.

Curl Example:

```bash
curl -X POST -H "Content-Type: application:json" -H "Authorization: Bearer <secret_key>" -d '{
    "customer": "Customer ID",
    "payment_method": "Payment Method ID",
    "description": "Payment Description",
    "amount": "Payment Amount",
    "currency": "Payment Currency",
    "metadata": {
        "key": "value"
    }
}' "<API_URL>/api/v1/payments"
```

You will receive a response with the payment ID and the payment status. The status will be updated when the payment is processed by the inquiring bank.

To check the status of a payment, you can use the following endpoint `GET /api/v1/payments/{payment_id}`.

> Note: Pass the security key in the `Authorization` header as a Bearer token.

Curl Example:

```bash
curl -X GET -H "Authorization: Bearer <secret_key>" "<API_URL>/api/v1/payments/{payment_id}"
```

The response will contain the payment status and the payment details.

## Webhooks

The Online Payment Platform sends events to your application using webhooks. You need to provide a webhook URL when you register as a merchant. The platform will send the following events:

- Payment Created
- Payment Processed
- Payment Failed
- Payment Refunded

> Note: This feature is not implemented at the moment and some events might change.
