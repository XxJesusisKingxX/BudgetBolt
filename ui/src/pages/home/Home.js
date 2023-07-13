import './Home.css'
import { getUser } from './profile/Profile';
import { formatOverviewDate } from './utils/FormatDate';

function Home() {
    const date = formatOverviewDate(new Date())
    const user = getUser()
    return (
        <>
            <div className="recent_transactions_container">
                <span className="recent_transaction_title">Recent Transactions</span>
                <div className="recent_transaction_border"></div>
            </div>
            <div className="budget_overview_container">
                <span className="budget_overview_user">Welcome, {user}</span>
                <span className="budget_overview_date">~ Today is {date} ~</span>
                <span className="budget_overview_title">Budget Overview</span>
                <div className="budget_overview_border"></div>
                <div className="budget_overview_divider"></div>
            </div>
        </>
    );
}
  
export default Home;