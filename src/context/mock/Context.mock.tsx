import { render } from '@testing-library/react'
import LoginContext from '../LoginContext';
import AppContext from '../AppContext';
import ThemeContext from '../ThemeContext';
import { Health, ModeType } from '../../constants/style';
import { BudgetView } from '../../constants/view';

// Create mock Login Context
export const mockLoginDispatch = jest.fn();
export const initLoginState = {
    isLogin: false,
    showAccountWindow: false,
    showSignUpWindow: false,
    showLoginWindow: false,
    loginDispatch: mockLoginDispatch
};

// Create mock Theme Context
export const mockThemeDispatch = jest.fn();
export const initThemeState = {
    health: Health.NONE,
    mode: ModeType.LIGHT,
    themeDispatch: mockThemeDispatch
};

// Create mock App Context
export const mockDispatch = jest.fn();
export const initAppState = {
    profile: "",
    totalIncome: 0.00,
    totalExpenses: 0.00,
    budgetView: BudgetView.MONTHLY,
    isTransactionsUpdated: false,
    lastTransactionsUpdate: new Date(),
    isTransactionsRefresh: false,
    linkToken: "",
    dispatch: mockDispatch
};

// Create mock App Context wrapped around Login Context
export const renderAllContext = (component: JSX.Element) => {
  const element = render(
    <LoginContext.Provider value={initLoginState}>
      <AppContext.Provider value={initAppState}>
        <ThemeContext.Provider value={initThemeState}>
          {component}
        </ThemeContext.Provider>
      </AppContext.Provider>
    </LoginContext.Provider>
  );
  return element;
}