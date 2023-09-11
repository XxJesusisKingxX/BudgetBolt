import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';
import { screen } from '@testing-library/react';

// Global States Intialization
const showLoginWindow = initState.showLoginWindow;

afterEach(() => {
    initState.showLoginWindow = showLoginWindow;
})

describe("Render AccountWindow:", () => {
    test("hide form", () => {
        renderWithLoginContext(<Auth/>);
        expect(screen.queryByRole('heading', {name: "Sign In"})).toBeFalsy();
    });

    test("show form only", () => {
        initState.showLoginWindow = true
        renderWithLoginContext(<Auth/>);
        expect(screen.getByRole('heading', {name: "Sign In Close"})).toBeTruthy();
        expect(screen.queryByRole('heading', {name: "Setup Account"})).toBeFalsy();
        expect(screen.queryByRole('heading', {name: "Create an Account"})).toBeFalsy();
    });
});