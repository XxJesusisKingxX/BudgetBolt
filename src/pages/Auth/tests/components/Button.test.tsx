import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';
import { screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { mockLocalStorage } from '../../../../utils/test';
import { renderAllContext, mockLoginDispatch, mockDispatch, initLoginState } from '../../../../context/mock/Context.mock';

// Global States Intialization
const isLogin = initState.isLogin;

beforeEach(() => {
    window.localStorage.clear();
})

afterEach(() => {
    initState.isLogin = isLogin;
})

describe("Render Button:", () => {
    mockLocalStorage();

    test("show 'Login'", () => {
        renderWithLoginContext(<Auth/>)
        expect(screen.getByRole('button', { name: "Login" })).toBeTruthy();
    });

    test("show 'Logout'", () => {
        initState.isLogin = true
        renderWithLoginContext(<Auth/>)
        expect(screen.getByRole('button', { name: "Logout" })).toBeTruthy();
    });

    test("click 'Logout'", () => {
        initLoginState.isLogin = true
        renderAllContext(<Auth/>)
        userEvent.click(screen.getByRole('button', { name: "Logout" }))
        // Assert
        expect(mockLoginDispatch).toBeCalledWith({"state": {"isLogin": false}, "type": "SET_STATE"});
        expect(mockLoginDispatch).toBeCalledWith({ type: "SET_STATE", state: { showLoginWindow: false, showSignUpWindow: false, showAccountWindow: false }})
        expect(mockDispatch).toBeCalledWith({ type: "SET_STATE", state: { profile: '' }})
        // Assert local storage
        const items = window.localStorage.getAll()
        expect(Object.keys(items).length).toBe(0);
    });
});