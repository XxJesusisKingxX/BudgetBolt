import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import AuthContainer from '../..';
import { cleanup, getByTestId, queryByTestId } from '@testing-library/react';

// Global States Intialization
const showLoginWindow = initState.showLoginWindow;

afterEach(() => {
    cleanup();
    initState.showLoginWindow = showLoginWindow;
})

describe("Render AccountWindow:", () => {
    test("hide form", () => {
        const {queryByTestId} = renderWithLoginContext(<AuthContainer/>);
        expect(queryByTestId('login-window')).toBeFalsy();
    });

    test("show form only", () => {
        initState.showLoginWindow = true
        const {queryByTestId, getByTestId} = renderWithLoginContext(<AuthContainer/>);
        expect(getByTestId('login-window')).toBeTruthy();
        expect(queryByTestId('signup-window')).toBeFalsy();
        expect(queryByTestId('account-window')).toBeFalsy();
    });
});