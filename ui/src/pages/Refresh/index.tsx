import { useContext } from 'react';
import AppContext from '../../context/AppContext';
import ThemeContext from '../../context/ThemeContext';
import RefreshComponent from './RefreshComponent';

// Refresh component
const Refresh = () => {
    // Accessing mode from ThemeContext and isTransactionsRefresh, dispatch from AppContext
    const { mode } = useContext(ThemeContext);
    const { isTransactionsRefresh, dispatch } = useContext(AppContext);

    // Click event handler for the refresh button
    const handleRefreshClick = () => {
        // Toggling the isTransactionsRefresh state and updating the lastTransactionsUpdate time
        dispatch({
            type: "SET_STATE",
            state: {
                lastTransactionsUpdate: new Date(),
                isTransactionsRefresh: !isTransactionsRefresh,
            },
        });
    };

    return (
        // Render the Refresh component with necessary props
        <RefreshComponent
            isRefresh={isTransactionsRefresh} // Current refreshing state
            refresh={() => {
                handleRefreshClick();
            }} // Function to handle refresh button click
            mode={mode} // Theme mode for image path
        />
    );
};

export default Refresh;
