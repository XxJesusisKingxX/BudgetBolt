export enum EndPoint {
    GET_TRANSACTIONS = "api/transactions/get",
    CREATE_TRANSACTIONS = "api/transactions/create",
    REMOVE_PENDING = "api/transactions/pending/remove",
    GET_EXPENSES = "api/expenses/get",
    CREATE_EXPENSES = "api/expenses/create",
    UPDATE_EXPENSES = "api/expenses/update",
    UPDATE_ALL_EXPENSES = "api/expenses/update/all",
    GET_PROFILE = "api/profile/get",
    CREATE_PROFILE = "api/profile/create",
    GET_INCOMES = "api/incomes/get",
    UPSERT_INCOMES = "api/incomes/upsert",
    CREATE_ACCOUNTS = "api/accounts/create",
    CREATE_LINK_TOKEN = "api/link_token/create",
    CREATE_ACCESS_TOKEN = "api/access_token/create"
}