import '@testing-library/jest-dom'
import { initState, renderWithLoginContext, mockLoginDispatch } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';
import { screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { mockLocalStorage } from '../../../../utils/test';

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
        initState.isLogin = true
        renderWithLoginContext(<Auth/>)
        userEvent.click(screen.getByRole('button', { name: "Logout" }))
        // Assert
        expect(mockLoginDispatch).toBeCalledWith({"state": {"isLogin": false}, "type": "SET_STATE"});
        // Assert local storage
        const items = window.localStorage.getAll()
        expect(Object.keys(items).length).toBe(0);
    });
});