import { FC } from 'react';
import './FinancialsSnapshot.css';

// Props interface describing the expected props for the FinancialsPeek component
interface Props {
    income: string;   // Total income amount as a string
    expenses: string; // Total expenses amount as a string
    savings: string;  // Total savings amount as a string
    trend: string;    // Trend indicator as a string
}

const FinancialsSnapshotComponent: FC<Props> = ({ income, expenses, savings, trend }) => {
    return (
        <div className='financialspeek'>
            <div className='financialspeek__titles'>
                {/* Total Income */}
                <h3 className='financialspeek__titles__header'>Total Income:<span className='financialspeek__titles__income'>${income}</span></h3>
                {/* Total Expenses */}
                <h3 className='financialspeek__titles__header'>Total Expenses:<span className='financialspeek__titles__expenses'>${expenses}</span></h3>
                {/* Total Savings */}
                <h3 className='financialspeek__titles__header'>Total Savings:<span className='financialspeek__titles__savings'>${savings}</span><span className='financialspeek__titles__trend'>{trend}</span></h3>
            </div>
        </div>
    );
}

export default FinancialsSnapshotComponent;
