import { useContext } from 'react';
import Refresh from './Refresh';
import AppContext from '../../context/AppContext';
import ThemeContext from '../../context/ThemeContext';

// RefreshContainer component
const RefreshContainer = () => {
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
        <Refresh
            isRefresh={isTransactionsRefresh} // Current refreshing state
            refresh={() => {
                handleRefreshClick();
            }} // Function to handle refresh button click
            mode={mode} // Theme mode for image path
        />
    );
};

export default RefreshContainer;
