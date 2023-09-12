import { screen, waitFor } from '@testing-library/react';
import { initLoginState, renderAllContext } from '../../../context/mock/Context.mock';
import Transaction from '../index';

beforeEach(() => {
    jest.clearAllMocks();
});

describe("Transactions", () => {
    test('renders all transactions', async () => {
        // Add fake transactions as needed
        const transactionsData = [
            { ID: 1, From: 'Account1', Amount: 100, Vendor: 'Vendor1' },
            { ID: 2, From: 'Account2', Amount: 200, Vendor: 'Vendor2' },

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
        initLoginState.isLogin = true
        await waitFor(() => {
            renderAllContext(<Transaction />);
        });
        // Assertions
        // Wait for transaction to load
        expect(mockElement).toBeCalledWith('sidebar');
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
    // sidebar animation isnt tested or checkhourly
})
