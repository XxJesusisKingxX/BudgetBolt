import './Sideview.css';
import Transaction from './Transaction/TransactionContainer';
import Refresh from './Refresh/Refresh';
import Context from '../../context/Context';
import { useContext } from 'react';

interface Props {
    lastUpdate: string
}
const Sideview: React.FC<Props> = ({ lastUpdate }) => {
    return (
        <>
            <div className="sideview-container">
                <div className="sideview-title">Recent Transaction<Refresh/></div>
                <div className="sideview-border"></div>
                <div className="sideview-footer">Last updated: {lastUpdate}</div>
                <Transaction />
            </div>
        </>
    );
};

export default Sideview;