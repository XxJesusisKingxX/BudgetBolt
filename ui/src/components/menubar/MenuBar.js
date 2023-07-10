import './MenuBar.css'
import settingsIcon from '../menubar/icons/settings.png';
import helpIcon from '../menubar/icons/help.png';
import reminderIcon from '../menubar/icons/reminder.png';
import goalIcon from '../menubar/icons/goal.png';
import incomeIcon from '../menubar/icons/income.png';
import expenseIcon from '../menubar/icons/expense.png';
import transactionIcon from '../menubar/icons/transaction.png';
import budgetIcon from '../menubar/icons/budget.png';


function MenuBar() {
    return (
        <div className="menubar">
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
    );
}
  
export default MenuBar;
