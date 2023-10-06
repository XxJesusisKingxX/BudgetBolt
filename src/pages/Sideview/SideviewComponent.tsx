import React from 'react';
import '../../assets/Sideview.css';
import Transaction from '../Transaction';
import Refresh from '../Refresh';

// Props interface for the Sideview component
interface Props {
    lastUpdate: string; // Last update time for the sidebar
}

const SideviewComponent: React.FC<Props> = ({ lastUpdate }) => {
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

export default SideviewComponent;
