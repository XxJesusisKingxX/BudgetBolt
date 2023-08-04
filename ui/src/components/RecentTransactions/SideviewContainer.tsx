import Context from '../../context/Context';
import Sideview from './Sideview';
import { useContext } from 'react';

const SideviewContainer = () => {
    const {lastTransactionsUpdate, isLogin} = useContext(Context)
    return (
        <>
            {isLogin ? (
            <Sideview
                lastUpdate={lastTransactionsUpdate.toLocaleDateString() + " " + lastTransactionsUpdate.toLocaleTimeString()} 
            />
            ) : (
                null
            )}
        </>
    );
};

export default SideviewContainer;