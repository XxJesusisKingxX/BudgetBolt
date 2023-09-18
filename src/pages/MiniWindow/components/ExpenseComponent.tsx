import { FC } from 'react';

interface Props {
    name: string;    // The name for the expense
    limit: string;  // Amount of the expense
    spent: string;   // The amount spent
}

const ExpenseComponent: FC<Props> = ({ name, limit, spent }) => {

    return (
        <div className='miniwindow__budget__view__item'>
            <span>{name}</span>---------------------<span>${limit}</span>--------------------------<span>${spent}</span>
        </div>
    );
}
export default ExpenseComponent