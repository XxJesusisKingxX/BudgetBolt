import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';
import { cleanup } from '@testing-library/react';

// Global States Intialization
const showSignUpWindow = initState.showSignUpWindow;

afterEach(() => {
    cleanup();
    initState.showSignUpWindow = showSignUpWindow;
})

describe("Render AccountWindow:", () => {
    test("hide form", () => {
        const {queryByTestId} = renderWithLoginContext(<Auth/>);
        expect(queryByTestId('signup-window')).toBeFalsy();
    });

    test("show form", () => {
        initState.showSignUpWindow = true
        const {getByTestId, queryByTestId} = renderWithLoginContext(<Auth/>)
        expect(getByTestId('signup-window')).toBeTruthy();
        expect(queryByTestId('account-window')).toBeFalsy();
        expect(queryByTestId('login-window')).toBeFalsy();
    });

});