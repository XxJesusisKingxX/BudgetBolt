import { useContext, useEffect } from 'react';
import FinancialsSnapshotComponent from './component/FinancialsSnapshotComponent';
import AppContext from '../../context/AppContext';
import { useIncome } from './useIncome';

const FinancialsSnapshot = () => {
    const { totalExpenses, totalIncome } = useContext(AppContext)

    const { upsertIncome } = useIncome();

    // Calculate income and expenses metrics
    const maxLevel = 4;
    let percentage: number = totalExpenses == 0 ? 0 : 100
    if (totalExpenses > 0 && totalIncome > 0) {
        percentage = Number(((totalExpenses / totalIncome) * 100).toFixed())
    }
    // if totalExpenses > 0 
    const level = (totalExpenses / totalIncome) * maxLevel // level current
    const currentLevel = level > maxLevel ? maxLevel : level // determine if level has surpassed max

    useEffect(() => {
        upsertIncome();
        // eslint-disable-next-line
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
