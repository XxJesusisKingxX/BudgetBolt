import { formatOverviewDate } from '../../utils/formatDate';
import DashboardComponent from './DashboardComponent';

// DashboardContainer component
const Dashboard = () => {
    return (
        <>
            <DashboardComponent
                user={localStorage.getItem('profile')?.toLocaleUpperCase() || ''}          // Displaying the user's name in uppercase
                date={formatOverviewDate(new Date())} // Formatting the current date using the utility function
            />
        </>
    );
};

export default Dashboard;
