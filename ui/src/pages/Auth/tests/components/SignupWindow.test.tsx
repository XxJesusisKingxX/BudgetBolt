import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';
import { screen } from '@testing-library/react';

// Global States Intialization
const showSignUpWindow = initState.showSignUpWindow;

afterEach(() => {
    initState.showSignUpWindow = showSignUpWindow;
})

describe("Render AccountWindow:", () => {
    test("hide form", () => {
        renderWithLoginContext(<Auth/>);
        expect(screen.queryByRole('heading', {name: "Create an Account Close"})).toBeFalsy();
    });

    test("show form", () => {
        initState.showSignUpWindow = true
        renderWithLoginContext(<Auth/>)
        expect(screen.getByRole('heading', {name: "Create an Account Close"})).toBeTruthy();
        expect(screen.queryByRole('heading', {name: "Sign In Close"})).toBeFalsy();
        expect(screen.queryByRole('heading', {name: "Setup Account"})).toBeFalsy();
    });

});