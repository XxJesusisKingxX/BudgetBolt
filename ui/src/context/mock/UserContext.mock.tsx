import { Health, ModeType } from '../../constants/style';
import UserContext from '../UserContext'
import { render } from '@testing-library/react'

export const mockDispatch = jest.fn();

export const initState = {
  health: Health.NONE,
  mode: ModeType.Light,
  profile: "",
  isLoading: false,
  isLogin: false,
  isTransactionsUpdated: false,
  lastTransactionsUpdate: new Date(),
  isTransactionsRefresh: false,
  linkToken: "", // Don't set to null or error message will show up briefly when site loads
  dispatch: mockDispatch
};

/**
 * Testing with UserContext mock wrapper
 *
 * @param {JSX.Element} component - The component to isolate and wrap for testing
 * @returns {string} The element
 */
export const renderWithLoginContext = (component: JSX.Element) => {
  const element = render(
    <UserContext.Provider value={initState}>
        {component}
    </UserContext.Provider>
  );
  return element;
}