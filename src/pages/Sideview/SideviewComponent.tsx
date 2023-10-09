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
        <div id='sidebar' className='sidebar'>
            <div className='sidebar__header'>
                <span className='sidebar__header__title'>Recent Transaction</span>
                <Refresh />
            </div>
            <Transaction />
            <span className='sidebar__footer'>Last updated: {lastUpdate}</span>
        </div>
    );
};

export default SideviewComponent;
