import { FC } from 'react';
import '../../../assets/FinancialsSnapshot.css';
import FinancialsSnapshotChart from './FinancialsSnapshotChart';

// Props interface describing the expected props for the FinancialsPeek component
interface Props {
    income: string;    // Total income amount as a string
    expenses: string;  // Total expenses amount as a string
    savings: string;   // Total savings amount as a string
    level: string;     // The current level expenses are at compared to savings
    percentage: number // The current level percentage
}

const FinancialsSnapshotComponent: FC<Props> = ({ income, expenses, savings, level, percentage}) => {
    return (
        <div className='financials-snapshot'>
            <div className='financials-snapshot__titles'>
                {/* Total Income */}
                <span className='financials-snapshot__titles__header'>Total Income:<span className='financials-snapshot__titles__income'>${income}</span></span>
                {/* Total Expenses */}
                <span className='financials-snapshot__titles__header'>Total Expenses:<span className='financials-snapshot__titles__expenses'>${expenses}</span></span>
                {/* Total Savings */}
                <span className='financials-snapshot__titles__header'>Total Savings:<span className={`financials-snapshot__titles__savings ${expenses > income ? 'financials-snapshot__titles__savings--overbought' : ''}`}>${savings}</span></span>
            </div>
            <FinancialsSnapshotChart level={level} percentage={percentage} />
        </div>
    );
}

export default FinancialsSnapshotComponent;
