// Import necessary components and types from React and other files
import { createContext, useReducer, Dispatch, ReactNode } from "react";
import { Health, ModeType } from "../constants/style";

// Define the shape of the state
interface State {
  health: Health;
  mode: ModeType;
}

// Set initial state values
const initialState: State = {
  health: Health.NONE,
  mode: ModeType.Light,
};

// Define possible actions that can modify the state
type Action = {
  type: "SET_STATE";
  state: Partial<State>;
};

// Create an interface that extends the State interface and adds the themeDispatch function
interface Context extends State {
  themeDispatch: Dispatch<Action>;
}

// Create a context with the initial state
const ThemeContext = createContext<Context>(initialState as Context);

// Destructure the Provider from ThemeContext
const { Provider } = ThemeContext;

// Define the provider component
export const ThemeProvider: React.FC<{ children: ReactNode }> = (
  props
) => {
  // Define a reducer function to handle state updates
  const reducer = (state: State, action: Action): State => {
    switch (action.type) {
      case "SET_STATE":
        // Merge the existing state with the new state from the action
        return { ...state, ...action.state };
      default:
        // Return the current state for unknown actions
        return { ...state };
    }
  };

  // Use the useReducer hook to manage state using the reducer function and initial state
  const [state, themeDispatch] = useReducer(reducer, initialState);

  // Provide the state and themeDispatch function to child components
  return <Provider value={{ ...state, themeDispatch }}>{props.children}</Provider>;
};

// Export the ThemeContext
export default ThemeContext;
