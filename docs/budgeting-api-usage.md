# Budgeting API Documentation

This document provides detailed information about the Budgeting API endpoints.

## Base URL

All API endpoints are relative to the base URL: http://localhost:8000/api

## Create Expense

Create a user's budgeted expenses.

- **Endpoint URL:** `/expenses/create`
- **HTTP Method:** POST

### Parameters

- `name` (string, required): The name of the expense.
- `limit` (float, required): The budgeted limit for the expense.
- `spent` (float, required): The amount already spent for the expense.

**Request Example:**

```json
{
    "name": "Rent",
    "limit": 1000.00,
    "spent": 750.00
}
```
**Response Example:**

```json
{
    // "message": "Expense created successfully."
}
```

## Update Expense

Update a user's budgeted expenses.

- **Endpoint URL:** `/expenses/update`
- **HTTP Method:** POST

### Parameters

- `uid` (string): User ID for whom to update expenses.
- `id` (int, required): The ID of the expense to update.
- `limit` (float, required): The updated budgeted limit for the expense.

**Request Example:**

```json
{
    "uid": "fgtg435dviukgfdert645hgr5hgfghetyryhh566hyrt4htrhhb54",
    "id": 1,
    "limit": 1100.00
}
```
**Response Example:**

```json
{
    // "message": "Expense updated successfully."
}
```

## Retrieve Expenses

Retrieve a user's budgeted expenses.

- **Endpoint URL:** `/expenses/get`
- **HTTP Method:** GET

### Parameters

- `uid` (string): User ID to retrieve expenses for.

**Response Example:**

```json
{
    "expenses": [
        {
            "id": 1,
            "name": "Rent",
            "limit": 1000.00,
            "spent": 750.00
        },
        {
            "id": 2,
            "name": "Groceries",
            "limit": 400.00,
            "spent": 300.00
        }
    ]
}
```

## Update All Expenses

Update all user's budgeted expenses based on transaction data.

- **Endpoint URL:** `/expenses/update/all`
- **HTTP Method:** POST

### Parameters

- `uid` (string): User ID for whom to update expenses.
- `date` (string): The date for which to update expenses based on transactions. Note: YYYY-MM-DD

## Request Example

```json
{
    "uid": "fgtg435dviukgfdert645hgr5hgfghetyryhh566hyrt4htrhhb54",
    "date": "2023-09-30"
}
```

**Response Example:**

```json
{
    // "message": "All expenses updated successfully."
}
```
