import './Menu.css';
import { useContext, FC }from 'react';
import Context from '../../context/Context';

interface Props {
    showDropdown: boolean
    onMouseOut: Function
    onMouseOver: Function
}

const Menu: FC<Props> = ({ showDropdown, onMouseOut, onMouseOver }) => {
    const { mode, isLogin } = useContext(Context)

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
            {!isLogin ?
                <div id="menu" className="menu">
                    <div className="menu__lines">
                        <div className="menu__lines__line"/>
                        <div className="menu__lines__line"/>
                        <div className="menu__lines__line"/>
                    </div>
                    <div onMouseOut={() => onMouseOut()} onMouseOver={() => onMouseOver()} className={showDropdown ? "menu__list menu__list--default menu__list--show" : "menu__list menu__list--hide"}>
                        {/* TODO: Add links for items */}
                        <ul className="menu__list__item"><img className="menu__list__item__icon"src={budgetIcon}/><a href="">Budget Overview</a></ul>
                        <ul className="menu__list__item"><img className="menu__list__item__icon"src={transactionIcon}/><a href="">Transactions</a></ul>
                        <ul className="menu__list__item"><img className="menu__list__item__icon"src={expenseIcon}/><a href="">Expenses</a></ul>
                        <ul className="menu__list__item"><img className="menu__list__item__icon"src={incomeIcon}/><a href="">Income</a></ul>
                        <ul className="menu__list__item"><img className="menu__list__item__icon"src={goalIcon}/><a href="">Goals and Savings</a></ul>
                        <ul className="menu__list__item"><img className="menu__list__item__icon"src={reminderIcon}/><a href="">Reminders and Alerts</a></ul>
                        <ul className="menu__list__item"><img className="menu__list__item__icon"src={settingsIcon}/><a href="">Settings</a></ul>
                        <ul className="menu__list__item"><img className="menu__list__item__icon"src={helpIcon}/><a href="">Help and Support</a></ul>
                    </div>
                </div>
                :
                null
            }
        </>
    );
};

export default Menu;
