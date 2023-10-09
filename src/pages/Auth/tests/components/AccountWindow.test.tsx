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
        expect(screen.queryByText("Setup Account")).toBeFalsy();
    });

    test("show form", () => {
        initState.showAccountWindow = true
        renderWithLoginContext(<Auth/>);
        expect(screen.getByText("Setup Account")).toBeTruthy();
        expect(screen.queryByText("Sign In")).toBeFalsy();
        expect(screen.queryByText("Create an Account")).toBeFalsy();
    });

});