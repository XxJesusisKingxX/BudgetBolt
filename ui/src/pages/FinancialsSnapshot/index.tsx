import FinancialsSnapshotComponent from './FinancialsSnapshotComponent';

const FinancialsSnapshot = () => {
    // Positive sign (for trend) declaration
    const pos = '+';

    // TODO: Add logic to determine financials and the trend direction (positive/negative)
    // For now, we're using placeholders in the provided example.

    return (
        // Render the FinancialsPeek component with data
        <FinancialsSnapshotComponent
            
            income={1200.00.toFixed(2)}      // Total income amount formatted to 2 decimal places
            expenses={5200.00.toFixed(2)}    // Total expenses amount formatted to 2 decimal places
            savings={300.00.toFixed(2)}      // Total savings amount formatted to 2 decimal places
            trend={pos + 112.00.toFixed(2)}  // Trend indicator with positive sign and formatted to 2 decimal places
        />
    );
}

export default FinancialsSnapshot;
