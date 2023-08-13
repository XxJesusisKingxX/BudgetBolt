import { useContext } from "react";
import Context from "../../context/Context"
import Refresh from "./Refresh";
import { useAppStateActions } from "../../redux/redux";

const RefreshContainer = () => {
    const { mode, isTransactionsRefresh } = useContext(Context);
    const { setLastTransactionsUpdateState, setTransactionsRefreshState } = useAppStateActions();
    const handleRefreshClick = () => {
        setLastTransactionsUpdateState(new Date())
        setTransactionsRefreshState(!isTransactionsRefresh)
    };
    return (
        <Refresh
            isRefresh={isTransactionsRefresh}
            refresh={() => {
                handleRefreshClick();
            }}
            mode={mode}
        />
    );
};

export default RefreshContainer;