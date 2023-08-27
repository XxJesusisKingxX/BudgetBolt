// Import necessary components from React
import { createContext, useReducer, Dispatch, ReactNode } from "react";

// Define the shape of the state
interface State {
  showAccountWindow: boolean;
  showSignUpWindow: boolean;
  showLoginWindow: boolean;
  isLogin: boolean;
}

// Set initial state values
const initialState: State = {
  showAccountWindow: false,
  showSignUpWindow: false,
  showLoginWindow: false,
  isLogin: false
};

// Define possible actions that can modify the state
type Action = {
  type: "SET_STATE";
  state: Partial<State>;
};

// Create an interface that extends the State interface and adds the loginDispatch function
interface Context extends State {
  loginDispatch: Dispatch<Action>;
}

// Create a context with the initial state
const LoginContext = createContext<Context>(initialState as Context);

// Destructure the Provider from LoginContext
const { Provider } = LoginContext;

// Define the provider component
export const LoginProvider: React.FC<{ children: ReactNode }> = (
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
  const [state, loginDispatch] = useReducer(reducer, initialState);

  // Provide the state and loginDispatch function to child components
  return <Provider value={{ ...state, loginDispatch }}>{props.children}</Provider>;
};

// Export the LoginContext
export default LoginContext;
