import { createContext, useReducer, Dispatch, ReactNode } from "react";

interface State {
  profile: string,
  isLogin: boolean
}

const initialState: State = {
  profile: "",
  isLogin: false,
};

type Action = {
  type: "SET_STATE";
  state: Partial<State>;
};

interface Context extends State {
  userDispatch: Dispatch<Action>;
}

const UserContext = createContext<Context>(
  initialState as Context
);

const { Provider } = UserContext;
export const UserProvider: React.FC<{ children: ReactNode }> = (
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
  const [state, userDispatch] = useReducer(reducer, initialState);
  return <Provider value={{ ...state, userDispatch }}>{props.children}</Provider>;
};

export default UserContext;
