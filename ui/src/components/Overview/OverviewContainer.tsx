import { formatOverviewDate } from '../../utils/FormatDate';
import { getUser } from '../../utils/Profile';
import Overview from './Overview';


const OverviewContainer = () => {

    return (
        <Overview user={getUser()} date={formatOverviewDate(new Date)}/>
    );
};

export default OverviewContainer;