import "./Refresh.css";
import { FC } from "react";

interface Props {
    mode: string
    isRefresh: boolean
    refresh: Function
};

const Refresh: FC<Props> = ({ mode, isRefresh, refresh}) => {
    const refreshIcon = `/images/${mode}/refresh.png`;
    return (
        <>
            <img className={`sideview_refresh${isRefresh ? "_load" : ""}`} onClick={() => refresh()} src={refreshIcon}/>
        </>
    );
};

export default Refresh;