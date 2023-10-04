import { FC } from 'react';
import '../../../assets/view/styles/billtracker/Bill.css';

interface Props {
    name: string;                 // The name associated with the bill.
    price: string;                // The price or amount of the bill.
    daysLeft: string;             // The number of days left until the bill's due date.
    dueDate: string;              // The due date of the bill.
    category: string;             // The category or type of the bill (e.g., utility, rent, subscription).
    spacing: number;              // The spacing value for positioning the component.
    icon?: string                  // An icon for the bill category.
}

const BillComponent: FC<Props> = ({ icon, name, price, daysLeft, dueDate, category, spacing }) => {
    return (
        <div data-testid='bill' className='bill'>
            {/* Display the icon if provided */}
            <img data-testid='bill-icon' className='bill__icon' src={icon} alt='ico'/>
            {/* Display the bill name, price, and days left */}
            <span className='bill__name'>{name}:</span>
            <span className='bill__amount'>${price}</span>
            <span data-testid='bill-daysleft' className="bill__daysleft">{daysLeft}</span>days
            {/* Expanded view */}
            <div style={{ top: `${7 + spacing}rem` }} className='bill__fullview'>
                {/* Placeholder for expanded view content */}
                <div className='bill__fullview__cont'/>
                <div className='bill__fullview__txt'>
                    <div className='bill__fullview__txt__title'>
                        {name}
                    </div>
                    <div className='bill__fullview__txt__cols'>
                        <div className='bill__fullview__txt__cols__lcol'>
                            Due Date:
                            <br/>
                            {dueDate}
                        </div>
                        <div className='bill__fullview__txt__cols__rcol'>
                            Category:
                            <br/>
                            {category}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default BillComponent;
