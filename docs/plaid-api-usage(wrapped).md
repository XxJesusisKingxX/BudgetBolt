# Budgeting API Documentation

This document provides detailed information about the Budgeting API endpoints.

## Base URL

All API endpoints are relative to the base URL: http://localhost:8000/api

## Create Link Token

Create a Plaid Link token for connecting bank accounts.

- **Endpoint URL:** `/link_token/create`
- **HTTP Method:** POST

### Parameters

None

**Response Example:**

```json
{
    "link_token": "link-sandbox-12345678-90ab-cdef-ghij-klmnopqrstuv"
}

```

## Create Access Token

Create a Plaid access token from a public token.

- **Endpoint URL:** `/access_token/create`
- **HTTP Method:** POST

### Parameters

- `public_token` (string): plaid public token needed for workflow oauth.

**Request Example:**

```json
{
    "public_token": "public-sandbox-12345678-90ab-cdef-ghij-klmnopqrstuv"
}
```

## Create Accounts

Create and store user accounts from Plaid data.

- **Endpoint URL:** `/accounts/create`
- **HTTP Method:** POST

### Parameters

None

## Create Transactions

Create and store user transactions from Plaid data.

- **Endpoint URL:** `/transactions/create`
- **HTTP Method:** POST

### Parameters

None

