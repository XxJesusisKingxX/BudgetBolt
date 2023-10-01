import { FC } from 'react';
import './FinancialsSnapshotChart.css';

// Props interface describing the expected props for the FinancialsPeek component
interface Props {
    level: string;
    per: string; 
}

const FinancialsSnapshotChart: FC<Props> = ({ level, per }) => {
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
                    <div className='financials-snapshot-chart__inner-circle--toptext'>{100-Number(per)}%</div>
                    <div className='financials-snapshot-chart__inner-circle--btmtext'>{per}%</div>
                </div>
            </div>
        </div>
    );
}

export default FinancialsSnapshotChart;
