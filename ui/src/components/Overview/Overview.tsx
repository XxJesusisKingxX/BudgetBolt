import BudgetTrender from '../BudgetTrender/BudgetTrender';
import FinancialsPeek from '../FinancialsPeek/FinancialsPeekContainer';
import HealthIndicator from '../HealthIndicator/HealthIndicatorContainer';
import MiniWindow from '../MiniWindow/MiniWindowContainer';
import Upcoming from '../Upcoming/Upcoming';
import './Overview.css';

interface Props {
    user: string
    date: string
};

const Overview: React.FC<Props> = ({ user, date }) => {
    return (
        <div className="overview">
            <div className="overview__header">
                <span className="overview__header__user">Welcome, {user}</span>
                <span className="overview__header__date">~ Today is {date} ~</span>
                <span className="overview__header__title">Budget Overview</span>
            </div>
            <div className="overview__widgets">
                <div className="overview__widgets__top">
                    <div className="overview__widgets__top__left">
                        <HealthIndicator/>
                        <Upcoming/>
                    </div>
                    <div className="overview__widgets__top__right">
                        <FinancialsPeek/>
                        <MiniWindow/>
                    </div>
                </div>
                <div className="overview__widgets__bottom">
                    <BudgetTrender/>
                </div>
            </div>
        </div>
    );
};

export default Overview;