import { createContext, useReducer, Dispatch, ReactNode } from "react";

interface QuickstartState {
  mode: string,
  profile: string,
  isLoading: boolean
  isLogin: boolean
  showCreateWindow: boolean,
  showLoginWindow: boolean,
  isTransactionsUpdated: boolean
  lastTransactionsUpdate: Date,
  isTransactionsRefresh: boolean
  linkToken: string | null;
}

const initialState: QuickstartState = {
  mode: "light",
  profile: "",
  isLoading: false,
  isLogin: false,
  showCreateWindow: false,
  showLoginWindow: false,
  isTransactionsUpdated: false,
  lastTransactionsUpdate: new Date(),
  isTransactionsRefresh: false,
  linkToken: "", // Don't set to null or error message will show up briefly when site loads
};

type QuickstartAction = {
  type: "SET_STATE";
  state: Partial<QuickstartState>;
};

interface QuickstartContext extends QuickstartState {
  dispatch: Dispatch<QuickstartAction>;
}

const Context = createContext<QuickstartContext>(
  initialState as QuickstartContext
);

const { Provider } = Context;
export const QuickstartProvider: React.FC<{ children: ReactNode }> = (
  props
) => {
  const reducer = (
    state: QuickstartState,
    action: QuickstartAction
  ): QuickstartState => {
    switch (action.type) {
      case "SET_STATE":
        return { ...state, ...action.state };
      default:
        return { ...state };
    }
  };
  const [state, dispatch] = useReducer(reducer, initialState);
  return <Provider value={{ ...state, dispatch }}>{props.children}</Provider>;
};

export default Context;
