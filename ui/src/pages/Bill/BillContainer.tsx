import { useContext, useState } from "react";
import Context from "../../context/UserContext";
import Bill from "./Bill"

interface Bill {
    ID: number
    From: string
    Amount: number
    Vendor: string
    Category: string
    Date: string
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
    const { mode } = useContext(Context);
    const [isLoading, setIsLoading] = useState(true);
    const maxPeek = 6;
    const spacing = 2;

    const handleIcon = (category: string) => {
        //TODO figure logic to handle icon
    }

    const loading = `/images/${mode}/loading.png`;
    return (
        <>
            {isLoading ? bills.slice(0, maxPeek).map((bill,idx) => (
                <Bill
                    key={bill.ID}
                    name={bill.Vendor}
                    price={bill.Amount.toFixed(2)}
                    daysLeft={0}
                    dueDate={bill.Date}
                    category={bill.Category}
                    spacing={idx*spacing}
                    mode={mode}
                    // icon={handleIcon(bill.Category)} // TODO determine correct icon
                />
            )) : (
                <img className="loading loading--bills" src={loading}/>
            )}
        </>
    );
}

export default BillContainer