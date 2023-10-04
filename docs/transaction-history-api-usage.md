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
- `recurring` (string): The view to find bills within transaction

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
    ],
    "recurring": {
        "AT&T": {
            "name": "AT&T",
            "total_amount": 1000.0,
            "max_amount": 1500.0,
            "average_amount": 125.0,
            "due_date": "2023-12-15",
            "earliest_date_cycle": "2023-11-01",
            "previous_date_cycle": "2023-11-30",
            "last_date_cycle": "2023-12-30",
            "frequency": 12,
            "status": "ACTIVE",
            "degraded": 0,
            "category": "Telecommunications"
        },
        "McDonald's": {
            "name": "McDonald's",
            "total_amount": 500.0,
            "max_amount": 800.0,
            "average_amount": 75.0,
            "due_date": "2023-12-20",
            "earliest_date_cycle": "2023-11-05",
            "previous_date_cycle": "2023-11-25",
            "last_date_cycle": "2023-12-25",
            "frequency": 12,
            "status": "ACTIVE",
            "degraded": 0,
            "category": "Fast Food"
        }
    }
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

## Delete Pending Transaction

Delete the user's pending transactions

- **Endpoint URL:** `/transactions/pending/remove`
- **HTTP Method:** DELETE

### Parameters

None
