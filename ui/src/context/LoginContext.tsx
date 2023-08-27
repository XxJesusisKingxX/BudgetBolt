import { createContext, useReducer, Dispatch, ReactNode } from "react";

interface State {
  showAccountWindow: boolean
  showSignUpWindow: boolean,
  showLoginWindow: boolean,
}

const initialState: State = {
    showAccountWindow: false,
    showSignUpWindow: false,
    showLoginWindow: false,
};

type Action = {
  type: "SET_STATE";
  state: Partial<State>;
};

interface Context extends State {
  loginDispatch: Dispatch<Action>;
}

const LoginContext = createContext<Context>(
  initialState as Context
);

const { Provider } = LoginContext;
export const LoginProvider: React.FC<{ children: ReactNode }> = (
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
  const [state, loginDispatch] = useReducer(reducer, initialState);
  return <Provider value={{ ...state, loginDispatch }}>{props.children}</Provider>;
};

export default LoginContext;
