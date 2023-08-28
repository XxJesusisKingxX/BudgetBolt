import '@testing-library/jest-dom'
import { fireEvent, cleanup, render, getByTestId, queryByTestId, getByText, within, queryByText } from '@testing-library/react';
import { useCreate } from '../../useCreate';
import Bill from '../../components/Bill';

afterEach(() => {
    cleanup();
})

describe("Render Bill", () => {
    test("show bill", () => {
        // Create bill array of Bills type
        const billList = [
            {
               ID: 1,
               Amount: 111.11,
               Vendor: "Discover",
               Category: "CreditCard",
               DueDate: "Jan 12, 2023"
            }
        ]
        // Render bills with custom hook
        const { createBill } = useCreate();
        const bills = createBill(billList)
        // Get element
        const element = render(bills[0]).container;
        // Bill renders once
        expect(getByTestId(element, 'bill')).toBeTruthy();
    });
    test("show no bill", () => {
        const element = render(<Bill/>).container;
        // Bill does not render
        expect(queryByTestId(element, 'bill')).toBeFalsy();
    });
    test("show all bill props", () => {
        // Create bill array of Bills type
        const billList = [
            {
               ID: 1,
               Amount: 120,
               Vendor: "Walmart",
               Category: "Credit Card",
               DueDate: "Aug 26, 2023"
            }
        ]
        // Render bills with custom hook
        const { createBill } = useCreate();
        const bills = createBill(billList)
        // Get element
        const { getByText, queryByText, queryByTestId, getByTestId } = render(bills[0]);
        // Bill renders all props available
        expect(getByText("Walmart:")).toBeTruthy();
        expect(getByText("$120.00")).toBeTruthy();
        expect(getByTestId('bill-daysleft')).toBeTruthy();
        expect(getByText("Walmart")).toBeTruthy();
        expect(getByText("Category:Credit Card")).toBeTruthy();
        expect(getByText("Due Date:Aug 26, 2023")).toBeTruthy();
        expect(queryByTestId("bill-icon")).toBeTruthy();
        expect(queryByText("NaN")).toBeFalsy(); // Make sure no field is empty
    });
    // no need to test useCreate hook since already tested here
});
