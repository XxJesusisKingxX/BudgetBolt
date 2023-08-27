import React from 'react';
import './Sideview.css';
import Transaction from '../Transaction/TransactionContainer';
import Refresh from '../Refresh/RefreshContainer';

// Props interface for the Sideview component
interface Props {
    lastUpdate: string; // Last update time for the sidebar
}

// Sideview component definition
const Sideview: React.FC<Props> = ({ lastUpdate }) => {
    return (
        <div id='sidebar' className='sidebar'> {/* Container div for the sidebar */}
            <div className='sidebar__border'> {/* Top border section */}
                <span className='sidebar__border__title'>Recent Transaction</span> {/* Title for the recent transaction section */}
            </div>
            <span className='sidebar__footer'>Last updated: {lastUpdate}</span> {/* Last update information */}
            <Refresh /> {/* Refresh button */}
            <Transaction /> {/* Recent transaction content */}
        </div>
    );
};

export default Sideview;
