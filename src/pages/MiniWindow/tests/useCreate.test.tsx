import '@testing-library/jest-dom'
import { renderHook } from '@testing-library/react-hooks';
import { act, render, screen, waitFor } from '@testing-library/react';
import { mockingFetch } from '../../../utils/test';
import MiniWindowComponent from '../MiniWindowComponent';

describe("Expenses", () => {
    test("get expenses then show expenses", async () => {
        mockingFetch(200, {"expenses":[{"ID":"1","Name":"Test","Limit":"100.00","Spent":"150.00"}]})
        render(<MiniWindowComponent/>)
        // Assert for loading spinner
        expect(screen.getByRole('img', { name: "Loading" })).toBeTruthy();
        // Assert for all expenses text display
        await waitFor(() => {
            expect(screen.getByText("Test")).toBeTruthy();
        })
        expect(screen.getByText("100.00")).toBeTruthy();
        expect(screen.getByText("$150.00")).toBeTruthy();
    });
    // remainder functions are tested indireclty in budgetwindow
});
