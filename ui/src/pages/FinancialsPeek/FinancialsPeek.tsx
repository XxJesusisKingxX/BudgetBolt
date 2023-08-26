import { FC } from "react"
import "./FinancialsPeek.css"

interface Props {
    income: string
    expenses: string
    savings: string
    trend: string
}

const FinancialsPeek: FC<Props> = ({ income, expenses, savings, trend }) => {
    return (
        <div className="financialspeek">
            <div className="financialspeek__titles">
                <h3 className="financialspeek__titles__header">Total Income:<span className="financialspeek__titles__income">${income}</span></h3>
                <h3 className="financialspeek__titles__header">Total Expenses:<span className="financialspeek__titles__expenses">${expenses}</span></h3>
                <h3 className="financialspeek__titles__header">Total Savings:<span className="financialspeek__titles__savings">${savings}</span><span className="financialspeek__titles__trend">{trend}</span></h3>
            </div>
        </div>
    );
}

export default FinancialsPeek;