import './Sideview.css';
import Transaction from './Transaction/TransactionContainer';

const Sideview = () => {
    return (
        <>
            <div className="sideview_container">
                <span className="sideview_title">Recent Transactions</span>
                <div className="sideview_border"></div>
                <Transaction />
            </div>
        </>
    );
};

export default Sideview;