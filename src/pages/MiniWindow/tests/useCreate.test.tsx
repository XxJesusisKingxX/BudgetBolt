import '@testing-library/jest-dom'
import fetchMock  from 'jest-fetch-mock';
import { renderHook } from '@testing-library/react-hooks';
import { act, waitFor } from '@testing-library/react';
import * as Create from '../useCreate';
import { EndPoint } from '../../../constants/endpoints';


beforeEach(() => {
    fetchMock.resetMocks();
});

describe("useCreate", () => {

    test('should update an expense when valid id and limit are provided', async () => {
        // Assign
        const id = '1';
        const limit = '100.00';
        // Mock
        fetchMock.enableMocks();
        fetchMock.mockResponseOnce(JSON.stringify({}), { status: 200 });
        // Render
        const { result } = renderHook(Create.useCreate);
        act(() => {
            result.current.updateExpense(id,limit)
        })
        // Assert
        await waitFor(() => {
            expect(fetchMock).toHaveBeenCalledWith(EndPoint.UPDATE_EXPENSES, {
              method: 'POST',
              body: new URLSearchParams({
                id: id,
                limit: limit,
              }),
            });

            expect(fetchMock).toHaveBeenCalledWith(
                EndPoint.GET_EXPENSES,
                {
                    method: 'GET',
                }
            );
        });
    });

    test('should update an expense when valid id and limit are provided', async () => {
        // Assign
        const id = "";
        const limit = '100.00';
        // Mock
        fetchMock.enableMocks();
        fetchMock.mockResponseOnce(JSON.stringify({}), { status: 200 });
        // Render
        const { result } = renderHook(Create.useCreate);
        act(() => {
            result.current.updateExpense(id,limit)
        })
        // Assert
        await waitFor(() => {
            expect(fetchMock).not.toHaveBeenCalledWith(EndPoint.UPDATE_EXPENSES, {
              method: 'POST',
              body: new URLSearchParams({
                id: id,
                limit: limit,
              }),
            });

            expect(fetchMock).not.toHaveBeenCalledWith(
                EndPoint.GET_EXPENSES,
                {
                    method: 'GET',
                }
            );
        });
    });

    test('add expense should POST with correct URL Params and call get expenses', async () => {
        // Create mocks
        fetchMock.enableMocks();
        fetchMock.mockResponse(JSON.stringify({}), { status: 200 });

        // Render
        const expense: Create.Expense = {
            ID: "1",
            Name: "Test",
            Limit: "100.00",
            Spent: "0.00"
        }
        const { result } = renderHook(Create.useCreate);

        act(() => {
            result.current.addExpenses(expense)
        })

        // Assert
        await waitFor(() => {
            expect(fetchMock).toHaveBeenCalledWith(
                EndPoint.CREATE_EXPENSES,
                {
                    method: 'POST',
                    body: new URLSearchParams({
                        name: expense.Name,
                        limit: expense.Limit,
                        spent: expense.Spent,
                    }),
                }
            );
            expect(fetchMock).toHaveBeenCalledWith(
                EndPoint.GET_EXPENSES,
                {
                    method: 'GET',
                }
            );
        })
    });
});
