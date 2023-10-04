import '@testing-library/jest-dom'
import { render, screen } from '@testing-library/react';
import { useBill } from '../../useBill';
import Bill from '../../components/Bill';

describe("Render Bill", () => {
    // need to test for loading screen after bill implementation
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
        const { showBills } = useBill();

        // Assertions
        expect(screen.getByTestId('bill')).toBeTruthy();
    });
    test("show no bill", () => {
        render(<Bill/>);
        expect(screen.queryByTestId('bill')).toBeFalsy();
    });
    test("show all bill props", async () => {
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
        const { showBills } = useBill();
        
        // Assertions
        expect(screen.getByText("$120.00")).toBeTruthy();
        expect(screen.queryByTestId('bill-daysleft')).toBeTruthy();
        expect(screen.getByText("Walmart")).toBeTruthy();
        expect(screen.getByText("Category:Credit Card")).toBeTruthy();
        expect(screen.getByText("Due Date:Aug 26, 2023")).toBeTruthy();
        expect(screen.queryByTestId("bill-icon")).toBeTruthy();
        expect(screen.queryByText("NaN")).toBeFalsy(); // Make sure no field is empty
    });
    // no need to test useCreate hook since already tested here
});
