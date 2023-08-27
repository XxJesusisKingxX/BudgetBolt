import LoginContext from '../LoginContext'
import { render } from '@testing-library/react'

export const mockDispatch = jest.fn();

export const initState = {
    showAccountWindow: false,
    showSignUpWindow: false,
    showLoginWindow: false,
    loginDispatch: mockDispatch
};

/**
 * Testing with LoginContext mock wrapper
 *
 * @param {JSX.Element} component - The component to isolate and wrap for testing
 * @returns {string} The element
 */
export const renderWithLoginContext = (component: JSX.Element) => {
  const element = render(
    <LoginContext.Provider value={initState}>
        {component}
    </LoginContext.Provider>
  );
  return element;
}