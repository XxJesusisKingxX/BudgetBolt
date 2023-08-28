import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import AuthContainer from '../..';
import { cleanup, queryByTestId } from '@testing-library/react';

// Global States Intialization
const showAccountWindow = initState.showAccountWindow;

afterEach(() => {
    cleanup();
    initState.showAccountWindow = showAccountWindow;
})

describe("Render AccountWindow:", () => {
    test("hide form", () => {
        const {queryByTestId} = renderWithLoginContext(<AuthContainer/>);
        expect(queryByTestId('account-window')).toBeFalsy();
    });

    test("show form", () => {
        initState.showAccountWindow = true
        const {queryByTestId, getByTestId} = renderWithLoginContext(<AuthContainer/>);
        expect(getByTestId('account-window')).toBeTruthy();
        expect(queryByTestId('signup-window')).toBeFalsy();
        expect(queryByTestId('login-window')).toBeFalsy();
    });

});