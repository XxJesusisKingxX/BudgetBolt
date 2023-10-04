import { useContext, useEffect } from 'react';
import { useBill } from '../useBill';
import AppContext from '../../../context/AppContext';

const Bill = () => {
    const { getBills, showBills } = useBill();

    const { budgetView } = useContext(AppContext)
    
    useEffect(() => {
        getBills();
        // eslint-disable-next-line
    },[budgetView])

    return (
        <>
            {showBills()}
        </>
    );
};

export default Bill