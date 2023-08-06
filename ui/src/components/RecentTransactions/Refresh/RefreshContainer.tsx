import "./Refresh.css";
import { FC, useContext } from "react";
import Context from "../../../context/Context"
import Refresh from "./Refresh";


const RefreshContainer = () => {
    const {mode, isTransactionsRefresh, dispatch} = useContext(Context);
    const handleRefreshClick = () => {
        dispatch({ type: "SET_STATE", state: { lastTransactionsUpdate: new Date(), isTransactionsRefresh: !isTransactionsRefresh }});
    };
    return (
        <>
            <Refresh
                isRefresh={isTransactionsRefresh}
                refresh={() => {
                    handleRefreshClick();
                }}
                mode={mode}
            />
        </>
    );
};

export default RefreshContainer;