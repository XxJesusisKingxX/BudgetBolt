import '@testing-library/jest-dom'
import HealthIndicator from '..';
import { initAppState, mockDispatch, renderWithAppContext } from '../../../context/mock/AppContext.mock';



// Global States Intialization
const health = initAppState.health;


afterEach(() => {
    initAppState.health = health;
})

describe("Render HealthIndicator:", () => {
    test("show red", () => {
        initAppState.health = 1
        const { container } = renderWithAppContext(<HealthIndicator/>);
        const healthDot = container.querySelector('.healthind__dot--red');
        expect(healthDot).toBeInTheDocument();
    });
    test("show yellow", () => {
        initAppState.health = 2
        const { container } = renderWithAppContext(<HealthIndicator/>);
        const healthDot = container.querySelector('.healthind__dot--yellow');
        expect(healthDot).toBeInTheDocument();
    });
    test("show green", () => {
        initAppState.health = 3
        const { container } = renderWithAppContext(<HealthIndicator/>);
        const healthDot = container.querySelector('.healthind__dot--green');
        expect(healthDot).toBeInTheDocument();
    });
    test("show none", async () => {
        initAppState.health = 0
        const { container } = renderWithAppContext(<HealthIndicator/>);
        const healthDot = container.querySelector('.healthind__dot--default');
        expect(healthDot).toBeInTheDocument();
    });
});