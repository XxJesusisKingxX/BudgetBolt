import { createContext, useReducer, Dispatch, ReactNode } from "react";

interface State {
  isTransactionsUpdated: boolean
  lastTransactionsUpdate: Date,
  isTransactionsRefresh: boolean
  linkToken: string | null;
}

const initialState: State = {
  isTransactionsUpdated: false,
  lastTransactionsUpdate: new Date(),
  isTransactionsRefresh: false,
  linkToken: "", // Don't set to null or error message will show up briefly when site loads
};

type Action = {
  type: "SET_STATE";
  state: Partial<State>;
};

interface Context extends State {
  dispatch: Dispatch<Action>;
}

const AppContext = createContext<Context>(
  initialState as Context
);

const { Provider } = AppContext;
export const AppProvider: React.FC<{ children: ReactNode }> = (
  props
) => {
  const reducer = (
    state: State,
    action: Action
  ): State => {
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

export default AppContext;
