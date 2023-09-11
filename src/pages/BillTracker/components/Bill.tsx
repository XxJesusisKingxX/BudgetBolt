import { useContext, useState } from 'react';
import ThemeContext from '../../../context/ThemeContext';
import { useCreate, Bills } from '../useCreate';

const Bill = () => {
    const { mode } = useContext(ThemeContext);
    const { createBill } = useCreate();
    // eslint-disable-next-line
    const [loading, setLoading] = useState(false)
    // *TODO* actually fetch real data and test
    // eslint-disable-next-line
    const [bills, setBills] = useState<Bills[]>([]);
    
    const loadingIcon = `/images/${mode}/loading.png`;
    return (
        <>
            {!loading ? createBill(bills) : <img className='loading loading--bills' src={loadingIcon} alt='Loading'/>
            }
        </>
    );
};

export default Bill