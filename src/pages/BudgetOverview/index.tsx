import Sideview from '../Sideview';
import Dashboard from '../Dashboard';
import '../../assets/BudgetOverview.css'

// DashboardContainer component
const BudgetOverview = () => {
    return (
        <div className='budget-overview'>
            <Dashboard/>
            <Sideview/>
        </div>
    );
};

export default BudgetOverview;
