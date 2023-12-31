import { screen, waitFor } from '@testing-library/react';
import { initLoginState, renderAllContext } from '../../../context/mock/Context.mock';
import Transaction from '../index';
import { mockLocalStorage } from '../../../utils/test';
import { deleteCookie } from '../../../utils/cookie';

afterEach(() => {
    window.localStorage.clear();
    deleteCookie("UID");
    jest.clearAllMocks();
});

describe("Transactions", () => {
    mockLocalStorage();
    test('renders all transactions', async () => {
        // Add fake transactions as needed
        const transactionsData = [
            { transaction_id: 1, from_account: 'Account1', net_amount: 100, vendor: 'Vendor1' },
            { transaction_id: 2, from_account: 'Account2', net_amount: 200, vendor: 'Vendor2' },

        ];
        // Add mocks
        const mockElement = jest.spyOn(document, 'getElementById');
        const mockSetInterval = jest.spyOn(global,'setInterval')
        jest.spyOn(global, 'fetch').mockResolvedValue(
            new Response (
                JSON.stringify({
                    transactions: transactionsData
                }),
                {
                    status: 200,
                }
        ));
        document.cookie = "UID=123";
        await waitFor(() => {
            renderAllContext(<Transaction />);
        });
        // Assertions
        // Wait for transaction to load
        expect(screen.queryByRole('img', { name: "Loading"})).toBeFalsy();
        expect(screen.getAllByRole('img', { name: "Transaction Icon" })).toBeTruthy();
        expect(screen.getByText("Account1")).toBeTruthy();
        expect(screen.getByText("Vendor1...$100")).toBeTruthy();
        expect(screen.getByText("Account2")).toBeTruthy();
        expect(screen.getByText("Vendor2...$200")).toBeTruthy();
        expect(mockSetInterval).toBeCalledWith(expect.any(Function), 3600000);
    });
    test('no transactions', async () => {
        initLoginState.isLogin = true
        await waitFor(() => {
            renderAllContext(<Transaction />);
        });
        expect(screen.getByRole('img', { name: "Loading"})).toBeTruthy();
    });
    // checkhourly isnt tested
})
