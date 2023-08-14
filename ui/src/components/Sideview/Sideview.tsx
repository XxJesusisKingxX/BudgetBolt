import "./Sideview.css";
import Transaction from "../Transaction/TransactionContainer";
import Refresh from "../Refresh/RefreshContainer";

interface Props {
    lastUpdate: string
};
const Sideview: React.FC<Props> = ({ lastUpdate }) => {
    return (
        <div id="sidebar" className="sidebar">
            <div className="sidebar__border">
                <span className="sidebar__border__title">Recent Transaction</span>
            </div>
            <span className="sidebar__footer">Last updated: {lastUpdate}</span>
            <Refresh/>
            <Transaction />
        </div>
    );
};

export default Sideview;