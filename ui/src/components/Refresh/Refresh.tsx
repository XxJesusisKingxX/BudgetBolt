import { FC } from "react";

interface Props {
    mode: string
    isRefresh: boolean
    refresh: Function
};

const Refresh: FC<Props> = ({ mode, isRefresh, refresh }) => {
    const refreshIcon = `/images/${mode}/refresh.png`;
    return (
        <img className={isRefresh ? "sidebar__refresh sidebar__refresh--load" : "sidebar__refresh sidebar__refresh--loadalt"} onClick={() => refresh()} src={refreshIcon}/>
    );
};

export default Refresh;