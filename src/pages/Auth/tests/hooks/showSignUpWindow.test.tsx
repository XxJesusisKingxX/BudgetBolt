import '@testing-library/jest-dom'
import { screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { initState, mockLoginDispatch, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';

// Global States Intialization
const showLoginWindow = initState.showLoginWindow;

afterEach(() => {
    initState.showLoginWindow = showLoginWindow;
})

describe("Show SignupWindow:", () => {
    test("click signup link", async () => {
        initState.showLoginWindow = true;
        renderWithLoginContext(<Auth/>);
        // Signup link button exists
        expect(screen.getByTestId('signup-link')).toBeTruthy();
        // Click signup link
        userEvent.click(screen.getByText("Sign Up"));
        // Signup link has triggered all other windows to close
        expect(mockLoginDispatch).toBeCalledWith({ type: "SET_STATE", state: { showLoginWindow: false, showSignUpWindow: false, showAccountWindow: false }});
        // Signup window trigger state changes and displayed
        expect(mockLoginDispatch).toBeCalledWith({ type: "SET_STATE", state: { showLoginWindow: false }});
        expect(mockLoginDispatch).toBeCalledWith({ type: "SET_STATE", state: { showSignUpWindow: false }});
        expect(mockLoginDispatch).toBeCalledWith({ type: "SET_STATE", state: { showSignUpWindow: true }});
    });
});
