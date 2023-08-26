import { useContext } from 'react';
import HealthIndicator from './HealthIndicator';
import Context from '../../context/Context';

const HealthIndicatorContainer = () => {
    const { health } = useContext(Context)

    let healthClassName = "";
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
        <HealthIndicator
            healthClassName={healthClassName}
        />
    );
};

export default HealthIndicatorContainer;