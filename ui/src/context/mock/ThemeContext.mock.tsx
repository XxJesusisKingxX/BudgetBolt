import { Health, ModeType } from '../../constants/style';
import ThemeContext from '../ThemeContext';
import { render, RenderResult } from '@testing-library/react';

export const mockThemeDispatch = jest.fn();

export const initThemeState = {
  health: Health.NONE,
  mode: ModeType.LIGHT,
  themeDispatch: mockThemeDispatch
};

/**
 * Testing with UserContext mock wrapper
 *
 * @param {JSX.Element} component - The component to isolate and wrap for testing
 * @returns {RenderResult} The result of rendering the component
 */
export const renderWithThemeContext = (component: JSX.Element): RenderResult => {
  const element = render(
    <ThemeContext.Provider value={initThemeState}>
        {component}
    </ThemeContext.Provider>
  );
  return element;
}
