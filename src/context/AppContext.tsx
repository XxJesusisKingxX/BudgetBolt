// Import necessary components from React
import { createContext, useReducer, Dispatch, ReactNode } from "react";
import { BudgetView } from "../constants/view";

// Define the shape of the state
interface State {
  profile: string
  totalIncome: number
  totalExpenses: number
  budgetView: BudgetView
  isTransactionsUpdated: boolean;
  lastTransactionsUpdate: Date;
  isTransactionsRefresh: boolean;
  linkToken: string | null;
}

// Set initial state values
const initialState: State = {
  profile: '',
  totalIncome: 0.00,
  totalExpenses: 0.00,
  budgetView: BudgetView.MONTHLY,
  isTransactionsUpdated: false,
  lastTransactionsUpdate: new Date(),
  isTransactionsRefresh: false,
  linkToken: "", // Don't set to null or error message will show up briefly when site loads
};

// Define possible actions that can modify the state
type Action = {
  type: "SET_STATE";
  state: Partial<State>;
};

// Create an interface that extends the State interface and adds the dispatch function
interface Context extends State {
  dispatch: Dispatch<Action>;
}

// Create a context with the initial state
const AppContext = createContext<Context>(initialState as Context);

// Destructure the Provider from AppContext
const { Provider } = AppContext;

// Define the provider component
export const AppProvider: React.FC<{ children: ReactNode }> = (props) => {
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
  const [state, dispatch] = useReducer(reducer, initialState);

  // Provide the state and dispatch function to child components
  return <Provider value={{ ...state, dispatch }}>{props.children}</Provider>;
};

// Export the AppContext
export default AppContext;
