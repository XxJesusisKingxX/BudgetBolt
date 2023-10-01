import '@testing-library/jest-dom'
import fetchMock  from 'jest-fetch-mock';
import { act, render, screen, waitFor } from '@testing-library/react';
import MiniWindowComponent from '../MiniWindowComponent';
import userEvent from '@testing-library/user-event';
import * as Create from '../useExpense';
import { EndPoint } from '../../../constants/endpoints';
import { BudgetView } from '../../../constants/view';
import { mockDispatch, renderWithAppContext } from '../../../context/mock/AppContext.mock';

beforeEach(() => {
    fetchMock.resetMocks();
    jest.restoreAllMocks();
});

describe("Expenses", () => {

    test('user should be able to add expenses and get expenses on success', async () => {
        // Create mocks
        const mockGetExpenses = jest.fn();
        const mockAddExpenses = jest.fn();
        const mockShowExpenses = jest.fn();
        const mockUpdateExpenses = jest.fn();
        const mockUpdateAllExpenses = jest.fn();
        jest.spyOn(Create,'useCreate').mockReturnValue({
            getExpenses: mockGetExpenses,
            addExpenses: mockAddExpenses,
            showExpenses: mockShowExpenses,
            updateExpense: mockUpdateExpenses,
            updateAllExpenses: mockUpdateAllExpenses,
            expTotal: 0,
            isLoading: true,
        });

        // Render
        render(<MiniWindowComponent/>)

        // Reflect user actions
        userEvent.click(screen.getByRole('button', { name: "+ Create Expense"}))
        userEvent.type(screen.getByLabelText('expense-name'), "TestExpense")
        userEvent.type(screen.getByLabelText('expense-limit'), "100.00")
        userEvent.click(screen.getByRole('button', { name: "Save"}))

        // Assert
        expect(mockGetExpenses).toHaveBeenCalledTimes(1);
        await waitFor( async ()=> {
            expect(mockAddExpenses).toHaveBeenCalledTimes(1);
        })
        expect(mockAddExpenses).toHaveBeenCalledWith({"ID": "0", "Limit": "100.00", "Name": "TestExpense", "Spent": "0.00"});
    });

    test("user should see expenses at initial render after loading spinner", async () => {
        // Mock
        fetchMock.enableMocks();
        fetchMock.mockResponseOnce(JSON.stringify({"expenses":[{"expense_id":"1","expense_name":"Test","expense_limit":"100.00","expense_spent":"150.00"}]}), {status: 200})

        // Render
        render(<MiniWindowComponent/>)

        // Assert for loading spinner
        expect(screen.getByRole('img', { name: "Loading" })).toBeTruthy();
        // Assert for all expenses displayed
        await waitFor(() => {
            expect(screen.getByText("Test")).toBeTruthy();
        })
        expect(screen.getByText("100.00")).toBeTruthy();
        expect(screen.getByText("$150.00")).toBeTruthy();
    });

    test("user should be able to edit expenses and save changes", async () => {
        // Mock
        fetchMock.enableMocks();
        fetchMock.mockResponseOnce(JSON.stringify({"expenses":[{"expense_id":"1","expense_name":"Test","expense_limit":"100.00","expense_spent":"150.00"}]}), {status: 200})
        // Render
        render(<MiniWindowComponent/>)
        // Recreate user actions
        await waitFor(() => {
            userEvent.click(screen.getByRole('button', { name: "Edit" }));
        })
        userEvent.type(screen.getByRole('textbox', { name: "expense-edit-limit" }),"120.00");
        act(() => {
            userEvent.click(screen.getByRole('button', { name: "Done" }));
        })
        // Assert
        await waitFor(() => {
            expect(fetchMock).toBeCalledWith(
                EndPoint.GET_EXPENSES,
                {
                    method: 'GET',
                }
            );
            expect(fetchMock).toBeCalledWith(
                EndPoint.UPDATE_EXPENSES,
                {
                    method: 'POST',
                    body: new URLSearchParams({
                        id: "1",
                        limit: "120.00",
                    }),
                }
            );
        })
    });

    test("user should be able to change view and update locally the view", async () => {
        // Create mocks
        const mockGetExpenses = jest.fn();
        const mockAddExpenses = jest.fn();
        const mockShowExpenses = jest.fn();
        const mockUpdateExpenses = jest.fn();
        const mockUpdateAllExpenses = jest.fn();
        jest.spyOn(Create,'useCreate').mockReturnValue({
            getExpenses: mockGetExpenses,
            addExpenses: mockAddExpenses,
            showExpenses: mockShowExpenses,
            updateExpense: mockUpdateExpenses,
            updateAllExpenses: mockUpdateAllExpenses,
            expTotal: 0,
            isLoading: true,
        });
        fetchMock.enableMocks();
        fetchMock.mockResponseOnce(JSON.stringify({"expenses":[{"ID":"1","Name":"Test","Limit":"100.00","Spent":"150.00"}]}), {status: 200})
        // Render
        renderWithAppContext(<MiniWindowComponent/>)
        // Recreate user actions
        act(() => {
            userEvent.selectOptions(screen.getByRole('combobox', { name: "" }), BudgetView.WEEKLY);
        })
        // Assert
        await waitFor(() => {
            expect(mockDispatch).toBeCalledTimes(1)
        })
        expect(mockDispatch).toBeCalledWith({ type:"SET_STATE", state: { budgetView: BudgetView.WEEKLY }})
        expect(mockUpdateAllExpenses).toBeCalledWith(BudgetView.WEEKLY)
    });
});
