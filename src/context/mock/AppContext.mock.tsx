import AppContext from '../AppContext';
import { render, RenderResult } from '@testing-library/react';

export const mockDispatch = jest.fn();

export const initAppState = {
    profile: "",
    isTransactionsUpdated: false,
    lastTransactionsUpdate: new Date(),
    isTransactionsRefresh: false,
    linkToken: "",
    dispatch: mockDispatch
};

/**
 * Testing with AppContext mock wrapper
 *
 * @param {JSX.Element} component - The component to isolate and wrap for testing
 * @returns {RenderResult} The result of rendering the component
 */
export const renderWithAppContext = (component: JSX.Element): RenderResult => {
  const element = render(
    <AppContext.Provider value={initAppState}>
        {component}
    </AppContext.Provider>
  );
  return element;
}
