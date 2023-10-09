import { ChangeEvent, FC, useEffect, useRef, useState } from 'react';
import '../../../assets/Button.css';

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
    const inputRef = useRef<HTMLInputElement>(null);
  
    // Add a click outside listener when 'edit' is true
    useEffect(() => {
      const handleOutsideClick = (event: MouseEvent) => {
        if (inputRef.current && !inputRef.current.contains(event.target as Node)) {
          setEdit(false);
          setLimit("");
        }
      };
  
      if (edit) {
        document.addEventListener('click', handleOutsideClick);
      }
  
      return () => {
        document.removeEventListener('click', handleOutsideClick);
      };
    }, [edit]);
  
    const handleLimitOnChange = (event: ChangeEvent<HTMLInputElement>) => {
      setLimit(event.target.value);
    };
  
    return (
      <div className='expense-grid'>
        <span className={`expense-grid__item expense-grid__item--name ${spent > limit ? 'expense-grid__item--overbudget' : ''}`}>{name}</span>
        <div className='expense-grid__item'>
          {/* Hide edit input box if not editing...must click 'Done' */}
          {edit ? (
            <input
              ref={inputRef}
              aria-label='expense-edit-limit'
              className='expense-grid__item expense-grid__item--input'
              placeholder={limit}
              value={editedLimit}
              onChange={handleLimitOnChange}
            />
          ) : (
            <span>${limit}</span>
          )}
        </div>
        <span className={'expense-grid__item '}>${spent}</span>
        {/* Edit/Done button */}
        {!edit ? (
          <button
            className=' btn btn--edit'
            onClick={() => setEdit(true)}
          >
            Edit
          </button>
        ) : (
          <button
            className=' btn btn--done'
            onClick={() => {
              update(id, editedLimit);
            }}
          >
            Done
          </button>
        )}
      </div>
    );
};

export default ExpenseComponent