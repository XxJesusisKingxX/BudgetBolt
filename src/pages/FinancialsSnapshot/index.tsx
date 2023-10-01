import { useContext, useEffect } from 'react';
import FinancialsSnapshotComponent from './FinancialsSnapshotComponent';
import AppContext from '../../context/AppContext';

const FinancialsSnapshot = () => {
    // TODO: Add logic to determine income
    const { totalExpenses, totalIncome} = useContext(AppContext)

    // Calculate income and expenses metrics
    const maxLevel = 4;
    const percentage = (totalExpenses / 5000)
    const percentageStr = ((totalExpenses / 3000) * 100).toFixed();
    const currentLevel = (percentage * maxLevel).toString()

    useEffect(() => {
    },[totalExpenses, totalIncome])

    return (
        // Render the FinancialsPeek component with data
        <div>
            <FinancialsSnapshotComponent
                income={totalIncome.toFixed(2)}                     // Total income amount formatted to 2 decimal places
                expenses={totalExpenses.toFixed(2)}                 // Total expenses amount formatted to 2 decimal places
                savings={(totalIncome - totalExpenses).toFixed(2)}  // Total savings amount formatted to 2 decimal places
                level={currentLevel}
                per={percentageStr}
            />
        </div>
    );
}

export default FinancialsSnapshot;
