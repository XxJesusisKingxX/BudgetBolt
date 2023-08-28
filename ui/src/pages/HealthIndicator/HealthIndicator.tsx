import React from 'react';
import './HealthIndicator.css'; // Importing the CSS file for styling

// Props interface for the HealthIndicator component
interface Props {
    healthClassName: string; // CSS class name to define the health indicator's appearance
}

// HealthIndicator component definition
const HealthIndicator: React.FC<Props> = ({ healthClassName }) => {
    return (
        <div className='healthind'> {/* Container div for the health indicator */}
            <span className='healthind__txt'>Health Indicator</span> {/* Text for the health indicator */}
            <div className={healthClassName}></div> {/* The actual health indicator element */}
        </div>
    );
};

export default HealthIndicator;
