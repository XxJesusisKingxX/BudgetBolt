import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';
import { screen } from '@testing-library/react';

// Global States Intialization
const showSignUpWindow = initState.showSignUpWindow;

afterEach(() => {
    initState.showSignUpWindow = showSignUpWindow;
})

describe("Render SignUpWindow:", () => {
    test("hide form", () => {
        renderWithLoginContext(<Auth/>);
        expect(screen.queryByText("Create an Account")).toBeFalsy();
    });

    test("show form", () => {
        initState.showSignUpWindow = true
        renderWithLoginContext(<Auth/>)
        expect(screen.getByText("Create an Account")).toBeTruthy();
        expect(screen.queryByText("Sign In")).toBeFalsy();
        expect(screen.queryByText("Setup Account")).toBeFalsy();
    });

});