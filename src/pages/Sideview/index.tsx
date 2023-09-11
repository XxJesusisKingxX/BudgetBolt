import { useContext } from 'react';
import AppContext from '../../context/AppContext';
import LoginContext from '../../context/LoginContext';
import SideviewComponent from './SideviewComponent';


const Sideview = () => {
    // Accessing lastTransactionsUpdate from AppContext and isLogin from LoginContext
    const { lastTransactionsUpdate } = useContext(AppContext);
    const { isLogin } = useContext(LoginContext);

    return (
        <>
            {isLogin ? (
                // Render the Sideview component with formatted last update time if logged in
                <SideviewComponent
                    lastUpdate={
                        lastTransactionsUpdate.toLocaleDateString() +
                        ' ' +
                        lastTransactionsUpdate.toLocaleTimeString()
                    }
                />
            ) : (
                null // If not logged in, display nothing
            )}
        </>
    );
};

export default Sideview;
