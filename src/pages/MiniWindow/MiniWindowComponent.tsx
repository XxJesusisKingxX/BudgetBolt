import { ChangeEvent, FC, useEffect, useState } from 'react';
import './MiniWindow.css'
import { Expense, useCreate } from './useCreate';

interface Props {
    // TODO add props needed
}

const MiniWindowComponent: FC<Props> = () => {
    const { getExpenses, addExpenses, showExpenses } = useCreate();

    const [loading, setLoading] = useState(false);

    const [addExpense, setAddExpense] = useState(false);
    const [expense, setExpense] = useState<Expense[]>([]);

    const [name, setName] = useState('');
    const [amount, setAmount] = useState('0.00');
    const newExpense: Expense = {
        Name: name,
        Limit: amount,
        Spent: "0.00"
    }

    const handleNameOnChange = (event: ChangeEvent<HTMLInputElement>) => {
        setName(event.target.value);
    }
    const handleAmountOnChange = (event: any) => {
        setAmount(event.target.value);
    }

    useEffect(() => {
        getExpenses();
    },[])

    return (
        <div className='miniwindow'>
            <div className='miniwindow__view'>
                <h4 className='miniwindow__view__header'>Current View</h4>
                <select className='miniwindow__view__list'>
                    <option value='weekly'>Weekly</option>
                    <option value='monthly'>Monthly</option>
                    <option value='yearly'>Yearly</option>
                </select>
            </div>
            <div className='miniwindow__budget'>
                <div className='miniwindow__budget__header'>
                    <span className='miniwindow__budget__header__item'>Name</span>
                    <span className='miniwindow__budget__header__item'>Budgeted Amount</span>
                    <span className='miniwindow__budget__header__item'>Actual Spent</span>
                </div>
                <div className='miniwindow__budget__view'>
                    {showExpenses()}
                    <button className='miniwindow__budget__view__item miniwindow__budget__view__item__button' onClick={() => setAddExpense(!addExpense)}>+ Create Expense</button>
                    {addExpense ?
                    <div className='miniwindow__budget__view__item'>
                        <span><input className='miniwindow__budget__view__item__input' value={name} onChange={(event) => handleNameOnChange(event)}/></span>---------------------<span><input className='miniwindow__budget__view__item__input' value={amount} onChange={(event) => handleAmountOnChange(event)}/></span>--------------------------<span className='miniwindow__budget__view__item__input miniwindow__budget__view__item__input--filled'>$0.00</span>
                        <button className='miniwindow__budget__view__item__button miniwindow__budget__view__item__button--save' onClick={() => addExpenses(newExpense)}>Save</button>
                    </div>
                    :
                    null
                    }
                </div>
            </div>
        </div>
    );
}
export default MiniWindowComponent