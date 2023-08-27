import '@testing-library/jest-dom'
import { initState, mockDispatch, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import AuthContainer from '../..';
import { fireEvent, cleanup } from '@testing-library/react';

// Global States Intialization
const showLoginWindow = initState.showLoginWindow;

afterEach(() => {
    cleanup();
    initState.showLoginWindow = showLoginWindow;
})

afterEach(() => {
    cleanup();
})

describe("Show SignupWindow:", () => {
    test("click login button", async () => {
        initState.showLoginWindow = true;
        const {getByTestId} = renderWithLoginContext(<AuthContainer/>);
        // Signup link button exists
        expect(getByTestId('signup-link')).toBeTruthy();
        // Click signup link
        fireEvent.click(getByTestId('signup-link'));
        // Signup link has triggered all other windows to close
        expect(mockDispatch).toHaveBeenNthCalledWith(1,{ type: "SET_STATE", state: { showLoginWindow: false, showSignUpWindow: false, showAccountWindow: false }});
        // Signup window trigger state changes and displayed
        expect(mockDispatch).toHaveBeenNthCalledWith(2,{ type: "SET_STATE", state: { showSignUpWindow: true }});
    });
});
