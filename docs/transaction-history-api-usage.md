# Budgeting API Documentation

This document provides detailed information about the Budgeting API endpoints.

## Base URL

All API endpoints are relative to the base URL: http://localhost:8000/api

## Retrieve Transactions

Store the user's plaid transactions

- **Endpoint URL:** `/transactions/get`
- **HTTP Method:** GET

### Parameters

- `date` (string): The date range to get transactions from
- `uid` (float, required): The user's id
- `category` (string): The primary category

**Response Example:**

```json
{
    "transactions": [
        {
            "transaction_id": "12345abcde",
            "transaction_date": "20230929",
            "net_amount": 100.50,
            "payment_method": "Credit Card",
            "vendor": "Example Vendor",
            "is_recurring": true,
            "description": "Purchase of goods",
            "primary_category": "Shopping",
            "secondary_category": "Clothing",
            "profile_id": 9876,
            "from_account": "Savings Account"
        },
        {
            "transaction_id": "67890fghij",
            "transaction_date": "20230930",
            "net_amount": 75.25,
            "payment_method": "Debit Card",
            "vendor": "Another Vendor",
            "is_recurring": false,
            "description": "Dining out with friends",
            "primary_category": "Dining",
            "secondary_category": "Restaurant",
            "profile_id": 5432,
            "from_account": "Checking Account"
        }
    ]
}
```

## Store Transactions

Store the user's plaid transactions

- **Endpoint URL:** `/transactions/store`
- **HTTP Method:** POST

### Parameters

- `id` (int, required): the user's id.

**Request Example:**

```json
{
    "id": 1,
}
```

## Store Accounts

Store the user's plaid accounts

- **Endpoint URL:** `/accounts/store`
- **HTTP Method:** POST

### Parameters

- `id` (int, required): the user's id.

**Request Example:**

```json
{
    "id": 1,
}
```