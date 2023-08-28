import { useState } from "react";
import Bill from "../pages/Bill/Bill";
import { ModeType } from "../constants/style";

// Interface for the shape of a bill
interface Bill {
    ID: number;       // Unique identifier for the bill
    From: string;     // Source of the bill
    Amount: number;   // Amount of the bill
    Vendor: string;   // Vendor associated with the bill
    Category: string; // Category of the bill
    Date: string;     // Date of the bill
}

export const useCreate = () => {
    const bills = [
        {
           ID: 1,
           From: "Discover",
           Amount: 111.11,
           Vendor: "Discover",
           Category: "CreditCard",
           Date: "Jan 12, 2023"
        },
        {
           ID: 1,
           From: "Discover",
           Amount: 111.11,
           Vendor: "Discover",
           Category: "CreditCard",
           Date: "Jan 12, 2023"
        }
    ]
    // const [bills, setBills] = useState<Bill[]>([]);
    const maxPeek = 6;
    const spacing = 2;

    const createBill = (mode: ModeType) => {
        bills.slice(0, maxPeek).map((bill, idx) => (
            <Bill
                key={bill.ID}
                name={bill.Vendor}                         // Vendor associated with the bill
                price={bill.Amount.toFixed(2)}             // Amount of the bill
                daysLeft={0}                               // Number of days left (TODO: Update)
                dueDate={bill.Date}                        // Date of the bill
                category={bill.Category}                   // Category of the bill
                spacing={idx * spacing}                    // Spacing based on index
                mode={mode}
                // icon={handleIcon(bill.Category)}        // TODO: Determine correct icon
            />
        ))
    };

    return {
        createBill
    };
};