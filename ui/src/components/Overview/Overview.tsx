import './Overview.css';

interface Props {
    user: string
    date: string
};

const Overview: React.FC<Props> = ({user, date}) => {
    return (
        <div className="overview-container">
            <span className="overview-user">Welcome, {user} </span>
            <span className="overview-date">~ Today is {date} ~</span>
            <span className="overview-title">Budget Overview</span>
            <div className="overview-border"></div>
            <div className="overview-divider"></div>
        </div>
    );
};

export default Overview;