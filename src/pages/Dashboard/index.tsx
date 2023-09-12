import { useContext } from 'react';
import { formatOverviewDate } from '../../utils/formatDate';
import AppContext from '../../context/AppContext';
import DashboardComponent from './DashboardComponent';

// DashboardContainer component
const Dashboard = () => {
    // Accessing the user's profile from AppContext
    const { profile } = useContext(AppContext);

    return (
        <>
            <DashboardComponent
                user={profile.toUpperCase()}          // Displaying the user's name in uppercase
                date={formatOverviewDate(new Date())} // Formatting the current date using the utility function
            />
        </>
    );
};

export default Dashboard;
