import { useContext, useEffect } from 'react';
import FinancialsSnapshotComponent from './FinancialsSnapshotComponent';
import AppContext from '../../context/AppContext';
import { useIncome } from './useIncome';

const FinancialsSnapshot = () => {
    // TODO: Add logic to determine income
    const { totalExpenses, totalIncome } = useContext(AppContext)

    const { upsertIncome } = useIncome();

    // Calculate income and expenses metrics
    const maxLevel = 4;
    const percentage = Number(((totalExpenses / totalIncome) * 100).toFixed())
    const level = (totalExpenses / totalIncome) * maxLevel // level current
    const currentLevel = level > maxLevel ? maxLevel : level // determine if level has surpassed max

    useEffect(() => {
        upsertIncome();
    },[totalExpenses, totalIncome])

    return (
        // Render the FinancialsPeek component with data
        <div>
            <FinancialsSnapshotComponent
                income={totalIncome.toFixed(2)}                     // Total income amount formatted to 2 decimal places
                expenses={totalExpenses.toFixed(2)}                 // Total expenses amount formatted to 2 decimal places
                savings={(totalIncome - totalExpenses).toFixed(2)}  // Total savings amount formatted to 2 decimal places
                level={currentLevel.toString()}
                percentage={percentage}
            />
        </div>
    );
}

export default FinancialsSnapshot;
