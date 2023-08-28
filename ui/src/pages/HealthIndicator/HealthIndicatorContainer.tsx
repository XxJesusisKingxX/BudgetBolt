import React, { useContext } from 'react';
import HealthIndicator from './HealthIndicator'; // Importing the HealthIndicator component
import ThemeContext from '../../context/ThemeContext'; // Importing the ThemeContext

// HealthIndicatorContainer component
const HealthIndicatorContainer = () => {
    // Accessing the health value from ThemeContext
    const { health } = useContext(ThemeContext);

    let healthClassName = "";

    // Calculate the appropriate CSS class based on the health value
    if (health > 0) {
        if (health === 1) {
            healthClassName = "healthind__dot healthind__dot--red";
        } else if (health === 2) {
            healthClassName = "healthind__dot healthind__dot--yellow";
        } else if (health === 3) {
            healthClassName = "healthind__dot healthind__dot--green";
        }
    } else {
        healthClassName = "healthind__dot healthind__dot--default";
    }

    return (
        // Render the HealthIndicator component with the calculated class
        <HealthIndicator
            healthClassName={healthClassName}
        />
    );
};

export default HealthIndicatorContainer;
