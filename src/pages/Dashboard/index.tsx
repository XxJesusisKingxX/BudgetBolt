import { useContext } from 'react';
import { formatOverviewDate } from '../../utils/formatDate';
import AppContext from '../../context/AppContext';
import LoginContext from '../../context/LoginContext';
import DashboardComponent from './DashboardComponent';

// DashboardContainer component
const Dashboard = () => {
    // Accessing the user's profile from AppContext
    const { profile } = useContext(AppContext);

    // Checking if the user is logged in from LoginContext
    const { isLogin } = useContext(LoginContext);

    return (
        <>
            {isLogin ? (
                // Render the Overview component with user information if logged in
                <DashboardComponent
                    user={profile.toUpperCase()}          // Displaying the user's name in uppercase
                    date={formatOverviewDate(new Date())} // Formatting the current date using the utility function
                />
            ) : (
                null // If not logged in, display nothing
            )}
        </>
    );
};

export default Dashboard;
