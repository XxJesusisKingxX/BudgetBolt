import { useContext, useEffect } from 'react';
import '../../../assets/BudgetWindow.css'
import { useExpense } from '../useExpense';
import AppContext from '../../../context/AppContext';
import { BudgetView } from '../../../constants/view';

const BudgetWindowComponent = () => {
    const { showExpenses, updateAllExpenses, } = useExpense();
    const { dispatch } = useContext(AppContext);

    const handleViewOnChange = (event: any) => {
        updateAllExpenses(event.target.value)
        dispatch({ type:'SET_STATE', state:{ budgetView: event.target.value}})
    };

    useEffect(() => {
        updateAllExpenses();
        // eslint-disable-next-line
    },[])

    return (
        <div className='miniwindow'>
            <div className='miniwindow__view'>
                <span className='miniwindow__view__header'>Current View</span>
                <select onChange={(event) => handleViewOnChange(event)}className='miniwindow__view__list'>
                    <option value={BudgetView.MONTHLY}>Monthly</option>
                    <option value={BudgetView.WEEKLY}>Weekly</option>
                    <option value={BudgetView.YEARLY}>Yearly</option>
                </select>
            </div>
            <div className='miniwindow__budget'>
                <div className='miniwindow__budget__header'>
                    <span className='miniwindow__budget__header__item'>Name</span>
                    <span className='miniwindow__budget__header__item miniwindow__budget__header__item--limit'> Budget</span>
                    <span className='miniwindow__budget__header__item miniwindow__budget__header__item--spent'>Spent</span>
                </div>
                <div className='miniwindow__budget__view'>
                    {showExpenses()}
                </div>
            </div>
        </div>
    );
}
export default BudgetWindowComponent