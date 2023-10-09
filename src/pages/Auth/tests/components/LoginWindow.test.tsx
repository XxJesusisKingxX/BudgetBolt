import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';
import { screen } from '@testing-library/react';

// Global States Intialization
const showLoginWindow = initState.showLoginWindow;

afterEach(() => {
    initState.showLoginWindow = showLoginWindow;
})

describe("Render LoginWindow:", () => {
    test("hide form", () => {
        renderWithLoginContext(<Auth/>);
        expect(screen.queryByText("Sign In")).toBeFalsy();
    });

    test("show form only", () => {
        initState.showLoginWindow = true
        renderWithLoginContext(<Auth/>);
        expect(screen.getByText("Sign In")).toBeTruthy();
        expect(screen.queryByText("Setup Account")).toBeFalsy();
    });
});