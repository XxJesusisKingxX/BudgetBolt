import { useContext } from 'react';
import AppContext from '../../context/AppContext';
import LoginContext from '../../context/LoginContext';
import Sideview from './Sideview';

// SideviewContainer component
const SideviewContainer = () => {
    // Accessing lastTransactionsUpdate from AppContext and isLogin from LoginContext
    const { lastTransactionsUpdate } = useContext(AppContext);
    const { isLogin } = useContext(LoginContext);

    return (
        <>
            {isLogin ? (
                // Render the Sideview component with formatted last update time if logged in
                <Sideview
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

export default SideviewContainer;
