import { useContext } from 'react';
import { formatOverviewDate } from '../../utils/formatDate';
import Overview from './Overview';
import UserContext from '../../context/UserContext';
import LoginContext from '../../context/LoginContext';

// OverviewContainer component
const OverviewContainer = () => {
    // Accessing the user's profile from UserContext
    const { profile } = useContext(UserContext);

    // Checking if the user is logged in from LoginContext
    const { isLogin } = useContext(LoginContext);

    return (
        <>
            {!isLogin ? (
                // Render the Overview component with user information if logged in
                <Overview
                    user={profile.toUpperCase()}          // Displaying the user's name in uppercase
                    date={formatOverviewDate(new Date())} // Formatting the current date using the utility function
                />
            ) : (
                null // If not logged in, display nothing
            )}
        </>
    );
};

export default OverviewContainer;
