import { FC } from 'react';
import './Bill.css';

interface Props {
    name: string;     // The name associated with the bill.
    price: string;    // The price or amount of the bill.
    daysLeft: number; // The number of days left until the bill's due date.
    dueDate: string;  // The due date of the bill.
    category: string; // The category or type of the bill (e.g., utility, rent, subscription).
    spacing: number;  // The spacing value for positioning the component.
    mode: string;     // The current mode of the component (e.g., light mode, dark mode).
    icon?: string;    // An optional icon for the bill.
}

const Bill: FC<Props> = ({ icon, mode, name, price, daysLeft, dueDate, category, spacing }) => {
    const pic = `/images/${mode}/bills/${icon}.png`;

    return (
        <div className='bill'>
            {/* Display the icon if provided */}
            {icon ? <img className='bill__icon' src={pic} alt={`${name} icon`} /> : ''}
            {/* Display the bill name, price, and days left */}
            {name}: ${price} ~ {daysLeft} days
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

export default Bill;
