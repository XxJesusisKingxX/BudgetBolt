import { useContext } from "react";
import { formatOverviewDate } from "../../utils/formatDate";
import Overview from "./Overview";
import Context from "../../context/Context";

const OverviewContainer = () => {
    const { isLogin, profile } = useContext(Context);
    return (
        <>
            {isLogin ? (
            <Overview
                user={profile.toLocaleUpperCase()}
                date={formatOverviewDate(new Date)}
            />
            ) : (
                null
            )}
        </>
    );
};

export default OverviewContainer;