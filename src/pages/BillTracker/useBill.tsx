import BillComponent from './components/BillComponent';
import entertainment from '../../assets/view/images/billtracker/categories/entertainment.svg'
import { useContext, useState } from 'react';
import { EndPoint } from '../../constants/endpoints';
import { getDateView } from '../../utils/formatDate';
import AppContext from '../../context/AppContext';
import ThemeContext from '../../context/ThemeContext';

// Interface for the shape of a bill
export interface Bills {
    ID: number;       // Unique identifier for the bill
    Amount: number;   // Amount of the bill
    Vendor: string;   // Vendor associated with the bill
    Category: string; // Category of the bill
    DueDate: string;     // Date of the bill
}

export const useBill = () => {
    const [bills, setBills] = useState<Bills[]>([]);
    const [isLoading, setLoading] = useState(false);

    const { budgetView } = useContext(AppContext)
    const { mode } = useContext(ThemeContext);

    const getBills = async () => {
        const date = getDateView(new Date(), budgetView)
        const query = `?date=${date}&recurring=${budgetView}`
        setLoading(true)
        try {
            const response = await fetch(EndPoint.GET_TRANSACTIONS + query, {
                method: "GET"
            })
            if (response.ok) {
                const data = await response.json();
                if (data) {
                    setBills(data["recurring"]);
                    setLoading(false)
                }
                setLoading(false);
            }
        } catch(error) {
            console.log(error);
        }

    };

    const showBills = (loading = isLoading, billList = bills) => {
        console.log(billList)
        const loadingIcon = `/images/${mode}/loading.png`;
        const maxPeek = 6; // max amount of bills to show
        const spacing = 2; // spacing between bills every render

        const handleDaysLeft= (date: string) => {
            const currentDate = new Date();
            const targetDate = new Date(date);

            // Calculate the time difference in milliseconds
            const timeDiff = targetDate.getTime() - currentDate.getTime();

            // Convert milliseconds to days
            const daysLeft = Math.ceil(timeDiff / (1000 * 60 * 60 * 24));

            return String(daysLeft);
        };

        const getCategoryIcon = (category: string) => {
            // Dynamically select the image based on the category
            switch (category) {
                case 'entertainment':
                    return entertainment;
                default:
                    return ''; // Default icon or handle missing icons
            }
        };
        
        // update to get the correct due date
        // update to inlcude category
        billList = Array.from(billList)
        const rows = billList.slice(0, maxPeek).map((bill: any, idx) => (
            <BillComponent
                key={idx}
                name={bill.name}                         // Vendor associated with the bill
                price={bill.avgAmt.toFixed(2)}             // Amount of the bill
                daysLeft={handleDaysLeft(bill.lastDateCycle)}    // Number of days left (TODO: Update)
                dueDate={bill.lastDateCycle}                     // Date of the bill
                category={bill.Category}                   // Category of the bill
                spacing={idx * spacing}                    // Spacing based on index
                icon={getCategoryIcon(bill.Category.toLowerCase())}  // Show icon based on category
            />
        ));
        return (
        !loading && billList ?
        rows
        :
        <img className='loading loading--bills' src={loadingIcon} alt='Loading'/>
        );

    }

    return {
        showBills,
        getBills
    };
};