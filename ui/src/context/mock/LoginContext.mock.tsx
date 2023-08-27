import LoginContext from '../LoginContext'
import { render } from '@testing-library/react'

// Mock dispatch function for testing
export const mockDispatch = jest.fn();

// Initial state for the LoginContext mock
export const initState = {
    showAccountWindow: false,
    showSignUpWindow: false,
    showLoginWindow: false,
    isLogin: false,
    loginDispatch: mockDispatch
};

/**
 * Renders a component with a mocked LoginContext for testing.
 *
 * @param {JSX.Element} component - The component to be wrapped and tested
 * @returns {RenderResult} The result of rendering the component
 */
export const renderWithLoginContext = (component: JSX.Element) => {
  const element = render(
    // Provide the mock initial state to the LoginContext.Provider
    <LoginContext.Provider value={initState}>
        {component}
    </LoginContext.Provider>
  );
  return element;
}
