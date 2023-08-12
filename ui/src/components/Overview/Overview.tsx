import HealthIndicator from '../HealthIndicator/HealthIndicatorContainer';
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
            <HealthIndicator/>
        </div>
    );
};

export default Overview;