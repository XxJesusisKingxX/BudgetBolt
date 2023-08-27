import { createContext, useReducer, Dispatch, ReactNode } from "react";
import { Health, ModeType } from "../constants/style";

interface State {
  health: Health
  mode: ModeType,
}

const initialState: State = {
  health: Health.NONE,
  mode: ModeType.Light,
};

type Action = {
  type: "SET_STATE";
  state: Partial<State>;
};

interface Context extends State {
    themeDispatch: Dispatch<Action>;
}

const ThemeContext = createContext<Context>(
  initialState as Context
);

const { Provider } = ThemeContext;
export const ThemeProvider: React.FC<{ children: ReactNode }> = (
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
  const [state, themeDispatch] = useReducer(reducer, initialState);
  return <Provider value={{ ...state, themeDispatch }}>{props.children}</Provider>;
};

export default ThemeContext;
