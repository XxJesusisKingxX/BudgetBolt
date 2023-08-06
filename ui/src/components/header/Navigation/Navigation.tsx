import './Navigation.css';
import settingsIcon from './icons/settings.png';
import helpIcon from './icons/help.png';
import reminderIcon from './icons/reminder.png';
import goalIcon from './icons/goal.png';
import incomeIcon from './icons/income.png';
import expenseIcon from './icons/expense.png';
import transactionIcon from './icons/transaction.png';
import budgetIcon from './icons/budget.png';
import { useContext } from 'react';
import Context from '../../../context/Context';

function Navigation() {
    const { isLogin } = useContext(Context)
    return (
        <>
            {isLogin ? (
            <div className="navigation">
                <div className="top_line"></div>
                <div className="middle_line"></div>
                <div className="bottom_line"></div>
                <div className="hidden_text">Menu</div>
                <div className="hidden_menu">
                    {/* TODO: Add links for items */}
                    <ul><img src={budgetIcon}></img><a href="">Budget Overview</a></ul>
                    <ul><img src={transactionIcon}></img><a href="">Transactions</a></ul>
                    <ul><img src={expenseIcon}></img><a href="">Expenses</a></ul>
                    <ul><img src={incomeIcon}></img><a href="">Income</a></ul>
                    <ul><img src={goalIcon}></img><a href="">Goals and Savings</a></ul>
                    <ul><img src={reminderIcon}></img><a href="">Reminders and Alerts</a></ul>
                    <ul><img src={settingsIcon}></img><a href="">Settings</a></ul>
                    <ul><img src={helpIcon}></img><a href="">Help and Support</a></ul>
                </div>
            </div>
            ) : (
                null
            )}
        </>
    );
};
  
export default Navigation;
