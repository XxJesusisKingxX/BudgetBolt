import { FC, useContext } from 'react';
import '../../../assets/Bill.css';
import ThemeContext from '../../../context/ThemeContext';

interface Props {
    name: string;                 // The name associated with the bill.
    shortName: string;                 // The shorten name associated with the bill.
    price: string;                // The price or amount of the bill..
    dueDate: string;              // The due date of the bill.
    categoryName: string;             // The category or type of the bill (e.g., utility, rent, subscription).
    category: string;             // The category to determine icon to display.
}

const BillComponent: FC<Props> = ({ shortName, name, price, dueDate, category, categoryName }) => {
    const { mode } = useContext(ThemeContext);
    return (
        <div data-testid='bill' className='bill'>
            <div className='bill__title'>
                <img data-testid='bill-icon' className='bill__title__icon' src={`/images/${mode}/transactions/${category}.png`} alt='ico'/>
                <span className='bill__title__name'>{shortName}:</span>
            </div>
            <div className='bill__amount'>
                <div className='bill__amount__text'>${price}</div>
            </div>
            {/* Expanded view */}
            <div className='bill__fullview'>
                {/* Placeholder for expanded view content */}
                <div className='bill__fullview__cont'>
                    <div className='bill__fullview__cont__txt'>
                        <div className='bill__fullview__cont__txt__title'>
                            {name}
                        </div>
                        <div className='bill__fullview__cont__txt__cols'>
                            <div className='bill__fullview__cont__txt__cols__lcol'>
                                Due Date:
                                <br/>
                                {dueDate}
                            </div>
                            <div className='bill__fullview__cont__txt__cols__rcol'>
                                Category:
                                <br/>
                                {categoryName}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default BillComponent;
