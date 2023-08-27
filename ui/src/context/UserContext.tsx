// Import necessary components from React
import { createContext, useReducer, Dispatch, ReactNode } from "react";

// Define the shape of the state
interface State {
  profile: string;
}

// Set initial state values
const initialState: State = {
  profile: "",
};

// Define possible actions that can modify the state
type Action = {
  type: "SET_STATE";
  state: Partial<State>;
};

// Create an interface that extends the State interface and adds the userDispatch function
interface Context extends State {
  userDispatch: Dispatch<Action>;
}

// Create a context with the initial state
const UserContext = createContext<Context>(initialState as Context);

// Destructure the Provider from UserContext
const { Provider } = UserContext;

// Define the provider component
export const UserProvider: React.FC<{ children: ReactNode }> = (
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
  const [state, userDispatch] = useReducer(reducer, initialState);

  // Provide the state and userDispatch function to child components
  return <Provider value={{ ...state, userDispatch }}>{props.children}</Provider>;
};

// Export the UserContext
export default UserContext;
