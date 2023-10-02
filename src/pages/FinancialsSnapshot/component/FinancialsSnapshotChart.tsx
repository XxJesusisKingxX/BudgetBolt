import { FC, useContext } from 'react';
import './FinancialsSnapshotChart.css';
import AppContext from '../../../context/AppContext';

// Props interface describing the expected props for the FinancialsPeek component
interface Props {
    level: string;
    percentage: number;
}

const FinancialsSnapshotChart: FC<Props> = ({ level, percentage }) => {
    const maxPercentage = 100
    
    const { totalExpenses, totalIncome } = useContext(AppContext)

    return (
        <div className='financials-snapshot-chart'>
            <div className='financials-snapshot-chart__legend'>
                <div className='financials-snapshot-chart__legend__key'>
                    <span className='financials-snapshot-chart__legend__key__color financials-snapshot-chart__legend__key__color--white'/>Savings
                </div>
                <br/>
                <div className='financials-snapshot-chart__legend__key'>
                    <span className='financials-snapshot-chart__legend__key__color financials-snapshot-chart__legend__key__color--black'/>Expenses
                </div>
            </div>
            <div className='financials-snapshot-chart__outer-circle'>
                <div style={{ height:`${level}rem`}} className='financials-snapshot-chart__inner-circle'>
                    {totalExpenses === 0 && totalIncome === 0 ?
                    null
                    :
                    <>
                        <div data-testid='financials-snapshot-top-percent' className='financials-snapshot-chart__inner-circle--toptext'>{percentage > maxPercentage ? 0 : maxPercentage - percentage}%</div>
                        <div data-testid='financials-snapshot-btm-percent'  className='financials-snapshot-chart__inner-circle--btmtext'>{percentage}%</div>
                    </>
                    }
                </div>
            </div>
        </div>
    );
}

export default FinancialsSnapshotChart;
