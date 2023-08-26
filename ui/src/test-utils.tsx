import UserContext from '../context/UserContext';
import { render } from '@testing-library/react'

export const mockDispatch = jest.fn();
export const initState = {
  isLogin: false,
  profile: "",
  dispatch: mockDispatch
};

/**
 * Testing with context mock wrapper
 *
 * @param {JSX.Element} component - The component to isolate and wrap for testing
 * @returns {string} The element
 */
export const renderWithContext = (component: JSX.Element) => {
  const element = render(
    <UserContext.Provider value={initState}>
        {component}
    </UserContext.Provider>
  );
  return element;
}