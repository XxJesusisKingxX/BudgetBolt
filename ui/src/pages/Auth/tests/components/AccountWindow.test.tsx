import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';
import { screen } from '@testing-library/react';

// Global States Intialization
const showAccountWindow = initState.showAccountWindow;

afterEach(() => {
    initState.showAccountWindow = showAccountWindow;
})

describe("Render AccountWindow:", () => {
    test("hide form", () => {
        renderWithLoginContext(<Auth/>);
        expect(screen.queryByRole('heading', {name: "Setup Account"})).toBeFalsy();
    });

    test("show form", () => {
        initState.showAccountWindow = true
        renderWithLoginContext(<Auth/>);
        expect(screen.queryByRole('heading', {name: "Setup Account"})).toBeTruthy();
        expect(screen.queryByRole('heading', {name: "Create an Account"})).toBeFalsy();
        expect(screen.queryByRole('heading', {name: "Sign In"})).toBeFalsy();
    });

});