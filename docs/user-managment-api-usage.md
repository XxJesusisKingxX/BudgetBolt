# User Profile and Token Management API Documentation

This document provides detailed information about the User Profile and Token Management API endpoints.

## Base URL

All API endpoints are relative to the base URL: http://localhost:8000/api

## Create Profile

Create a user profile.

- **Endpoint URL:** `/profile/create`
- **HTTP Method:** POST

**Parameters:**

- `username` (string, required): The username of the user.
- `password` (string, required): The user's password.

**Request Example:**

```json
{
    "username": "john_doe",
    "password": "secretpassword"
}
```

## Retrieve Profile

Retrieve a user's profile and authenticate with the provided password or UID.

- **Endpoint URL:** `/profile/get`
- **HTTP Method:** POST

### Parameters

You can use one of the following sets of parameters:

1. To retrieve a profile using a username and password:

   - `username` (string, required): The username of the user.
   - `password` (string, required): The user's password.

   **Request Example (Username and Password):**

   ```json
   {
       "username": "john_doe",
       "password": "secretpassword"
   }
   ```

2. To retrieve a profile using a username and password:

   - `username` (string, required): The username of the user.
   - `password` (string, required): The user's password.

   **Request Example (UID):**

   ```json
   {
       "uid": "fgtg435dviukgfdert645hgr5hgfghetyryhh566hyrt4htrhhb54",
   }

## Create Token

Create an access token to use plaid products.

- **Endpoint URL:** `/token/create`
- **HTTP Method:** POST

**Parameters:**

- `id` (int, required): The ID of the user profile.
- `itemId` (string, required): The ID of the item for which the plaid token is created.
- `token` (string, required): The plaid access token.

**Request Example:**

```json
{
    "id": 123,
    "itemId": "wvpEqaR98pfNJbAraz6JhoP7Peb6QyTrLz7pE",
    "token": "access-production-e193aaa7-b4ca-4130-933b-ecbd59e0e0ab"
}
```

## Retrieve Token

Retrieve a user's plaid token.

- **Endpoint URL:** `/token/get`
- **HTTP Method:** GET

**Parameters:**

- `uid` (string, required): The ID of the user profile.

**Response Example:**

```json
{
    "access_token":"access-production-e193aaa7-b4ca-4130-933b-ecbd59e0e0ab"
}
```