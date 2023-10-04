import '@testing-library/jest-dom'
import { act, render, screen, waitFor } from '@testing-library/react';
import { Bills } from '../../useBill';
import * as Create from '../../useBill';
import Bill from '../../components/Bill';
import { renderHook } from '@testing-library/react-hooks';

describe("Render Bill", () => {
    // need to test for loading screen after bill implementation
    test("show bill", async () => {
        // Create bill array of Bills type
        const billList: Bills = {
            "AMF": {
                average_amount: 24.8,
                category: "ENTERTAINMENT",
                degraded: 0,
                due_date: "2023-10-18",
                earliest_date_cycle: "2023-09-18",
                frequency: 1,
                last_date_cycle: "2023-09-18",
                max_amount: 24.8,
                name: "AMF",
                previous_date_cycle: "2023-09-18",
                status: "UNKNOWN",
                total_amount: 24.8
            }
        }
        // Render bills with custom hook
        let endResult;
        const { result } = renderHook(() => Create.useBill());
        act(() => {
            endResult = result.current.showBills(false, billList);
        });
        render(<div>{endResult}</div>);
        // Assertions
        await waitFor(() => {
            expect(screen.getByText("AMF")).toBeTruthy();
        })
        expect(screen.getByText("$24.80")).toBeTruthy();
        expect(screen.queryByTestId('bill-daysleft')).toBeTruthy();
        expect(screen.getByText("Category:ENTERTAINMENT")).toBeTruthy();
        expect(screen.getByText("Due Date:2023-10-18")).toBeTruthy();
        expect(screen.queryByTestId("bill-icon")).toBeTruthy();
        expect(screen.queryByText("NaN")).toBeFalsy(); // Make sure no field is empty
    });
    test("show no bill", () => {
        render(<Bill/>);
        expect(screen.queryByTestId('bill')).toBeFalsy();
    });
});
