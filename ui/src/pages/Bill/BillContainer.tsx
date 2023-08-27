import { useContext, useState } from 'react';
import Bill from './Bill';
import ThemeContext from '../../context/ThemeContext';

// Interface for the shape of a bill
interface Bill {
    ID: number;       // Unique identifier for the bill
    From: string;     // Source of the bill
    Amount: number;   // Amount of the bill
    Vendor: string;   // Vendor associated with the bill
    Category: string; // Category of the bill
    Date: string;     // Date of the bill
}

const BillContainer = () => {
    // *TODO* remove testcases and use live data
    const bills = [
        {
           ID: 1,
           From: "Discover",
           Amount: 111.11,
           Vendor: "Discover",
           Category: "CreditCard",
           Date: "Jan 12, 2023"
        },{
           ID: 2,
           From: "Plaid",
           Amount: 222.22,
           Vendor: "Plaid",
           Category: "Banking",
           Date: "Aug 1, 2023"
        },{
            ID: 3,
            From: "Region",
            Amount: 333.33,
            Vendor: "Region",
            Category: "Banking",
           Date: "Sep 13, 2023"
        },{
            ID: 4,
            From: "Rent",
            Amount: 444.44,
            Vendor: "Discover",
            Category: "Mortage",
           Date: "Dec 25, 2023"
         },{
            ID: 5,
            From: "At&t",
            Amount: 555.55,
            Vendor: "At&t",
            Category: "Phone",
            Date: "Jul 28, 2023"
         },{
             ID: 6,
             From: "Apple",
             Amount: 777.77,
             Vendor: "Apple",
             Category: "Entertainment",
             Date: "Feb 3, 2023"
         }
    ]
    // *TODO* add back when live data is ready
    // const [bills, setBills] = useState<Bill[]>([]);
    const { mode } = useContext(ThemeContext);
    const [isLoading, setIsLoading] = useState(true);
    const maxPeek = 6;
    const spacing = 2;

    // TODO: Implement logic to handle icons based on category
    const handleIcon = (category: string) => {
        // ...
    };

    const loading = `/images/${mode}/loading.png`;
    return (
        <>
            {isLoading ? bills.slice(0, maxPeek).map((bill, idx) => (
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
            )) : (
                <img className='loading loading--bills' src={loading} alt='Loading' />
            )}
        </>
    );
};

export default BillContainer