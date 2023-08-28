import { Health, ModeType } from '../../constants/style';
import UserContext from '../UserContext';
import { render, RenderResult } from '@testing-library/react';

export const mockDispatch = jest.fn();

export const initState = {
  profile: "",
  userDispatch: mockDispatch
};

/**
 * Testing with UserContext mock wrapper
 *
 * @param {JSX.Element} component - The component to isolate and wrap for testing
 * @returns {RenderResult} The result of rendering the component
 */
export const renderWithUserContext = (component: JSX.Element): RenderResult => {
  const element = render(
    <UserContext.Provider value={initState}>
        {component}
    </UserContext.Provider>
  );
  return element;
}
