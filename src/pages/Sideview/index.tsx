import { useContext } from 'react';
import AppContext from '../../context/AppContext';
import SideviewComponent from './SideviewComponent';


const Sideview = () => {
    // Accessing lastTransactionsUpdate from AppContext and isLogin from LoginContext
    const { lastTransactionsUpdate } = useContext(AppContext);

    return (
        <>
            <SideviewComponent
                lastUpdate={
                    lastTransactionsUpdate.toLocaleDateString() + ' ' + lastTransactionsUpdate.toLocaleTimeString()
                }
            />
        </>
    );
};

export default Sideview;
