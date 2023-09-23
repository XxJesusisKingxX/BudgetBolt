import { ChangeEvent, FC, useState } from 'react';

interface Props {
    update: Function // Update expense limit amount
    name: string;    // The name for the expense
    limit: string;   // Amount of the expense
    spent: string;   // The amount spent
    id: string;      // The unique id for the expense
}

const ExpenseComponent: FC<Props> = ({ update, id, name, limit, spent }) => {
    const [edit, setEdit] = useState(false);
    const [editedLimit, setLimit] = useState("");
    const handleLimitOnChange = (event: ChangeEvent<HTMLInputElement>) => {
        setLimit(event.target.value);
    }

    return (
        <div className='miniwindow__budget__view__item'>
            <span>{name}</span>------------${edit ? <input aria-label='expense-edit-limit' className='miniwindow__budget__view__item__input' value={editedLimit} onChange={handleLimitOnChange}/> : <span>{limit}</span>}------------<span className='miniwindow__budget__view__item__text'>${spent}</span>
            {!edit ? <button onClick={() => setEdit(true)}>Edit</button> : <button onClick={() => {setEdit(false); update(id, editedLimit)}}>Done</button>}
        </div>
    );
}
export default ExpenseComponent