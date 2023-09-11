import { useContext } from 'react';
import HealthIndicatorComponent from './HealthIndicatorComponent'; // Importing the HealthIndicator component
import ThemeContext from '../../context/ThemeContext'; // Importing the ThemeContext

const HealthIndicator = () => {
    // Accessing the health value from ThemeContext
    const { health } = useContext(ThemeContext);

    let healthClassName = "";

    // Calculate the appropriate CSS class based on the health value
    if (health === 1) {
        healthClassName = "healthind__dot healthind__dot--red";
    } else if (health === 2) {
        healthClassName = "healthind__dot healthind__dot--yellow";
    } else if (health === 3) {
        healthClassName = "healthind__dot healthind__dot--green";
    } else if (health === 0) {
        healthClassName = "healthind__dot healthind__dot--default";
    }

    return (
        // Render the HealthIndicator component with the calculated class
        <HealthIndicatorComponent
            healthClassName={healthClassName}
        />
    );
};

export default HealthIndicator;
