import BillComponent from './components/BillComponent';
import entertainment from '../../assets/view/images/billtracker/categories/entertainment.svg'
import { useContext, useState } from 'react';
import { EndPoint } from '../../constants/endpoints';
import { getDateView } from '../../utils/formatDate';
import AppContext from '../../context/AppContext';
import ThemeContext from '../../context/ThemeContext';

// Interface for the shape of a bill
interface Bill {
    name: string
    total_amount: number
    max_amount: number
    average_amount: number
    due_date: string
    earliest_date_cycle: string
    previous_date_cycle: string
    last_date_cycle: string
    frequency: number
    status: string
    degraded: number
    category: string
}

export const useBill = () => {
    const [bills, setBills] = useState<Bill[]>([]);
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
                    setBills(data["bills"]);
                    console.log(data["bills"])
                }
                setLoading(false);
            }
        } catch(error) {
            console.log(error);
        }

    };

    const showBills = (loading = isLoading, billList = bills) => {
        const maxPeek = 6; // max amount of bills to show
        const maxChar = 4; // max amount characters to show for name
        const spacing = 2; // spacing between bills every render

        const handleDaysLeft= (date: string) => {
            const currentDate = new Date();
            const targetDate = new Date(date);

            // Calculate the time difference in milliseconds
            const timeDiff = targetDate.getTime() - currentDate.getTime();

            // Convert milliseconds to days
            const daysLeft = Math.ceil(timeDiff / (1000 * 60 * 60 * 24));

            return daysLeft;
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

        const rows = Object.keys(bills)
        .slice(0, maxPeek)
        .map((name: any, idx) => {
            const bill: Bill = bills[name];
            const billName = bill.name.length > maxChar ? bill.name.slice(0, 4) + "*" : bill.name;
            if (handleDaysLeft(bill.due_date) > 0) {
            return (
                <BillComponent
                    key={idx}
                    name={billName}
                    price={bill.average_amount.toFixed(2)}
                    daysLeft={handleDaysLeft(bill.due_date).toString()}
                    dueDate={bill.due_date}
                    category={bill.category}
                    spacing={idx * spacing}
                    icon={getCategoryIcon(bill.category.toLowerCase())}
                />
            );
            }
        });

        return rows;

    }

    return {
        showBills,
        getBills,
        bills
    };
};