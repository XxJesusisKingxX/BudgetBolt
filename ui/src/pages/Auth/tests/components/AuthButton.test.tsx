import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/UserContext.mock';
import AuthContainer from '../..';
import { cleanup } from '@testing-library/react';

// Global States Intialization
const isLogin = initState.isLogin;

afterEach(() => {
    cleanup();
    initState.isLogin = isLogin;
})

describe("Render AuthButton:", () => {
    test("show 'Login'", () => {
        const {getByText} = renderWithLoginContext(<AuthContainer/>)
        expect(getByText("Login")).toBeTruthy();
    });

    test("show 'Logout'", () => {
        initState.isLogin = true
        const {getByText} = renderWithLoginContext(<AuthContainer/>)
        expect(getByText("Logout")).toBeTruthy();
    });
});