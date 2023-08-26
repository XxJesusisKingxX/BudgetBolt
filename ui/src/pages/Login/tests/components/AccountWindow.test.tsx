import '@testing-library/jest-dom'
import { fireEvent, getByTestId, getByText, queryByTestId, waitFor } from '@testing-library/react'
import AccountWindow from '../../AccountWindow';
import { initState, renderWithContext } from '../../../../test-utils';

// Global States Intialization
const isLogin = initState.isLogin;
const profile = initState.profile;

afterEach(() => {
    initState.isLogin = isLogin;
    initState.profile = profile;
})

describe("Render AccountWindow:", () => {
    test("show form", () => {
        const element = renderWithContext(<AccountWindow/>).getByTestId('login-form');
        expect(element).toBeTruthy();
    });

    test("hide form", () => {
        initState.isLogin = true;
        const element = renderWithContext(<Login/>).queryByTestId('login-form');
        expect(element).toBeFalsy();
    });

    test("show button", () => {
        const element = renderWithContext(<Login/>).getByTestId("login-button")
        expect(element).toBeTruthy();
    });
});