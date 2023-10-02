import './Menu.css';
import { useContext, FC } from 'react';
import ThemeContext from '../../context/ThemeContext';


// Props interface for Menu component
interface Props {
    showDropdown: boolean // Whether to display the dropdown menu
    onMouseOut: Function  // Event handler for mouse out event
    onMouseOver: Function // Event handler for mouse over event
}

// Menu component definition
const Menu: FC<Props> = ({ showDropdown, onMouseOut, onMouseOver }) => {
    // Accessing theme mode from ThemeContext
    const { mode } = useContext(ThemeContext);

    // Dynamic icons based on theme mode
    const settingsIcon = `/images/${mode}/menu/settings.png`;
    const helpIcon = `/images/${mode}/menu/help.png`;
    const reminderIcon = `/images/${mode}/menu/reminder.png`;
    const goalIcon = `/images/${mode}/menu/goal.png`;
    const incomeIcon = `/images/${mode}/menu/income.png`;
    const expenseIcon = `/images/${mode}/menu/expense.png`;
    const transactionIcon = `/images/${mode}/menu/transaction.png`;
    const budgetIcon = `/images/${mode}/menu/budget.png`;

    return (
        <>
            {/* Display the menu if the user is logged in */}
            <div id='menu' className='menu'>
                {/* Menu lines for decoration */}
                <div className='menu__lines'>
                    <div className='menu__lines__line'/>
                    <div className='menu__lines__line'/>
                    <div className='menu__lines__line'/>
                </div>
                {/* Dropdown menu with conditional classes */}
                <div
                    onMouseOut={() => onMouseOut()}
                    onMouseOver={() => onMouseOver()}
                    className={showDropdown ? 'menu__list menu__list--default menu__list--show' : 'menu__list menu__list--hide'}
                >
                    {/* Menu items with icons and links */}
                    <ul className='menu__list__item'><img alt='budget icon' className='menu__list__item__icon' src={budgetIcon}/><a href='/'>Budget Overview</a></ul>
                    <ul className='menu__list__item'><img alt='transactions icon' className='menu__list__item__icon' src={transactionIcon}/><a href='/'>Transactions</a></ul>
                    <ul className='menu__list__item'><img alt='expenses icon' className='menu__list__item__icon' src={expenseIcon}/><a href='/'>Expenses</a></ul>
                    <ul className='menu__list__item'><img alt='income icon' className='menu__list__item__icon' src={incomeIcon}/><a href='/'>Income</a></ul>
                    <ul className='menu__list__item'><img alt='goals and savings icon' className='menu__list__item__icon' src={goalIcon}/><a href='/'>Goals and Savings</a></ul>
                    <ul className='menu__list__item'><img alt='remninders and alerts icon' className='menu__list__item__icon' src={reminderIcon}/><a href='/'>Reminders and Alerts</a></ul>
                    <ul className='menu__list__item'><img alt='settings icon' className='menu__list__item__icon' src={settingsIcon}/><a href='/'>Settings</a></ul>
                    <ul className='menu__list__item menu__list__item--last'><img alt='help icon' className='menu__list__item__icon' src={helpIcon}/><a href='/'>Help and Support</a></ul>
                </div>
            </div>
        </>
    );
};

export default Menu;