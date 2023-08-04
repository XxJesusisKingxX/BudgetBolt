import { useContext } from 'react';
import { formatOverviewDate } from '../../utils/FormatDate';
import Overview from './Overview';
import Context from '../../context/Context';

const OverviewContainer = () => {
    const { isLogin, profile, dispatch } = useContext(Context);
    return (
        <>
            {isLogin ? (
            <Overview
                user={profile}
                date={formatOverviewDate(new Date)}
            />
            ) : (
                null
            )}
        </>
    );
};

export default OverviewContainer;