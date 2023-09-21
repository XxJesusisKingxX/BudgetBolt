import '@testing-library/jest-dom'
import { renderHook } from '@testing-library/react-hooks';
import { act, render, screen, waitFor } from '@testing-library/react';
import { mockingFetch } from '../../../utils/test';
import MiniWindowComponent from '../MiniWindowComponent';
import userEvent from '@testing-library/user-event';

describe("BudgetWindow", () => {
    test("empty expenses", async () => {
        mockingFetch(200, {})
        render(<MiniWindowComponent/>)
        await waitFor(() => {
            expect(screen.getByRole('img', { name: "Loading" })).toBeTruthy();
        })
    });
    test("add expenses", async () => {
        render(<MiniWindowComponent/>)
        userEvent.click(screen.getByRole('button', { name: "+ Create Expense" }))
        userEvent.type(screen.getByLabelText('expense-name'), "TestExpense")
        userEvent.type(screen.getByLabelText('expense-limit'), "123")
        userEvent.click(screen.getByRole('button', { name: "Save" }))
        await waitFor(() => {

            expect(screen.getByRole('button', { name: "Loading" })).toBeTruthy();
        })
    });
    test("update expenses", async () => {
        // TODO
    });
});
