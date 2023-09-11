import { ModeType } from '../../constants/style';
import BillComponent from './components/BillComponent';
import entertainment from '../../assets/view/images/billtracker/categories/entertainment.svg'

// Interface for the shape of a bill
export interface Bills {
    ID: number;       // Unique identifier for the bill
    Amount: number;   // Amount of the bill
    Vendor: string;   // Vendor associated with the bill
    Category: string; // Category of the bill
    DueDate: string;     // Date of the bill
}

export const useCreate = () => {
    
    const createBill = (bills: Array<Bills>, mode: ModeType = ModeType.LIGHT) => {
        const maxPeek = 6; // max amount of bills to show
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

        return bills.slice(0, maxPeek).map((bill, idx) => (
            <BillComponent
                key={bill.ID}
                name={bill.Vendor}                         // Vendor associated with the bill
                price={bill.Amount.toFixed(2)}             // Amount of the bill
                daysLeft={handleDaysLeft(bill.DueDate)}    // Number of days left (TODO: Update)
                dueDate={bill.DueDate}                     // Date of the bill
                category={bill.Category}                   // Category of the bill
                spacing={idx * spacing}                    // Spacing based on index
                icon={getCategoryIcon(bill.Category.toLowerCase())}  // Show icon based on category
            />
        ))
    };

    return {
        createBill
    };
};