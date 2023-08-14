import FinancialsPeek from "./FinancialsPeek"

const FinancialsPeekContainer = () => {
    const pos = "+"
    // TODO add logic to determine finanacials and determine if trend is pos and if so show pos sign
    return (
        <FinancialsPeek
            income={1200.00.toFixed(2)}
            expenses={5200.00.toFixed(2)}
            savings={300.00.toFixed(2)}
            trend={pos + 112.00.toFixed(2)}
        />
    );
}

export default FinancialsPeekContainer;