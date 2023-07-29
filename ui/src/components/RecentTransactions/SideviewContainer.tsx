import Context from '../../context/Context';
import Sideview from './Sideview';
import { useContext } from 'react';

const SideviewContainer = () => {
    const {lastTransactionsUpdate} = useContext(Context)
    return (
        <>
            <Sideview lastUpdate={lastTransactionsUpdate.toLocaleDateString() + " " + lastTransactionsUpdate.toLocaleTimeString()} />
        </>
    );
};

export default SideviewContainer;