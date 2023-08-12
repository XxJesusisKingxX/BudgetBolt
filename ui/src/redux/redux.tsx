import { useContext } from "react";
import { ModeType } from "../enums/style";
import Context from "../context/Context";

export const useAppStateActions = () => {
    const { dispatch } = useContext(Context);

    const stateActions = {
        setLoadingState: (state: boolean) => {
            if (state) {
                dispatch({ type: "SET_STATE", state: { isLoading: true } });
            } else {
                dispatch({ type: "SET_STATE", state: { isLoading: false } });
            }
        },
        setLoginState: (state: boolean) => {
            if (state) {
                dispatch({ type: "SET_STATE", state: { isLogin: true } });
            } else {
                dispatch({ type: "SET_STATE", state: { isLogin: false } });
            }
        },
        setTransactionsUpdatedState: (state: boolean) => {
            if (state) {
                dispatch({ type: "SET_STATE", state: { isTransactionsUpdated: true } });
            } else {
                dispatch({ type: "SET_STATE", state: { isTransactionsUpdated: false } });
            }
        },
        setTransactionsRefreshState: (state: boolean) => {
            if (state) {
                dispatch({ type: "SET_STATE", state: { isTransactionsRefresh: true } });
            } else {
                dispatch({ type: "SET_STATE", state: { isTransactionsRefresh: false } });
            }
        },
        setProfileState: (state: string) => {
            dispatch({ type: "SET_STATE", state: { profile: state } });
        },
        setLinkTokenState: (state: string | null) => {
            dispatch({ type: "SET_STATE", state: { linkToken: state } });
        },
        setModeState: (state: ModeType) => {
            switch (state) {
                case 1:
                    dispatch({ type: "SET_STATE", state: { profile: "light" } });
                    break;
                case 2:
                    dispatch({ type: "SET_STATE", state: { profile: "dark" } });
                    break;
                default:
                    break;
            }
        },
        setLastTransactionsUpdateState: (state: Date) => {
            dispatch({ type: "SET_STATE", state: { lastTransactionsUpdate: state } });
        }
    };

    return stateActions;
};
