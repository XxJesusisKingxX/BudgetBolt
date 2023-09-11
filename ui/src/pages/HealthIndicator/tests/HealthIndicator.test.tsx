import '@testing-library/jest-dom'
import { initThemeState, renderWithThemeContext } from '../../../context/mock/ThemeContext.mock';
import HealthIndicator from '..';

// Global States Intialization
const health = initThemeState.health;

afterEach(() => {
    initThemeState.health = health;
})

describe("Render HealthIndicator:", () => {
    test("show red", () => {
        initThemeState.health = 1
        const { container } = renderWithThemeContext(<HealthIndicator/>);
        const healthDot = container.querySelector('.healthind__dot--red');
        expect(healthDot).toBeInTheDocument();
    });
    test("show yellow", () => {
        initThemeState.health = 2
        const { container } = renderWithThemeContext(<HealthIndicator/>);
        const healthDot = container.querySelector('.healthind__dot--yellow');
        expect(healthDot).toBeInTheDocument();
    });
    test("show green", () => {
        initThemeState.health = 3
        const { container } = renderWithThemeContext(<HealthIndicator/>);
        const healthDot = container.querySelector('.healthind__dot--green');
        expect(healthDot).toBeInTheDocument();
    });
    test("show none", () => {
        initThemeState.health = 0
        const { container } = renderWithThemeContext(<HealthIndicator/>);
        const healthDot = container.querySelector('.healthind__dot--default');
        expect(healthDot).toBeInTheDocument();
    });
});