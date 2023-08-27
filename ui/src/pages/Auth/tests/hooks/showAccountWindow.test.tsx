import '@testing-library/jest-dom'
import { initState, mockDispatch, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import AuthContainer from '../..';
import { fireEvent, cleanup } from '@testing-library/react';

// Global States Intialization
const showAccountWindow = initState.showAccountWindow;

afterEach(() => {
    cleanup();
    initState.showAccountWindow = showAccountWindow;
})

describe("Show LoginWindow:", () => {
    test("click login button", async () => {
        initState.showAccountWindow = true;
        const { getByTestId } = renderWithLoginContext(<AuthContainer/>);
        // Plaid button exists
        expect(getByTestId("plaid-button")).toBeTruthy();
        // Click plaid button
        fireEvent.click(getByTestId("plaid-button"));
        // No need to test mock trigger because a succesful sign up will test it
    });
});
