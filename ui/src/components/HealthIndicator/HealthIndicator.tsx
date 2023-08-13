import { Health } from '../../enums/style';
import './HealthIndicator.css';

interface Props {
    healthClassName: string
};

const HealthIndicator: React.FC<Props> = ({ healthClassName }) => {
    return (
        <div className="healthind">
            <span className="healthind__txt">Health Indicator</span>
            <div className={healthClassName}></div>
        </div>
    );
};

export default HealthIndicator;