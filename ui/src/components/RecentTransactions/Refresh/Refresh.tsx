import './Refresh.css';
import refresh from './icons/refresh.png'
import { useContext } from 'react';
import Context from '../../../context/Context'

const Refresh = () => {
    const {isTransactionsRefresh, dispatch} = useContext(Context);
    const handleRefreshClick = () => {
        dispatch({ type: "SET_STATE", state: { lastTransactionsUpdate: new Date(), isTransactionsRefresh: !isTransactionsRefresh }});
    };
    return (
        <>
            <img className={`sideview_refresh${isTransactionsRefresh ? "_load" : ""}`} onClick={handleRefreshClick} src={refresh}/>
        </>
    );
};

export default Refresh;