import '@testing-library/jest-dom'
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';
import { screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { mockLocalStorage } from '../../../../utils/test';
import { renderAllContext, mockLoginDispatch, mockDispatch } from '../../../../context/mock/Context.mock';
import { deleteCookie } from '../../../../utils/cookie';

afterEach(() => {
    deleteCookie("UID");
})

describe("Render Button:", () => {
    mockLocalStorage();
    
    test("show 'Login'", () => {
        renderWithLoginContext(<Auth/>)
        expect(screen.getByRole('button', { name: "Login" })).toBeTruthy();
    });

    test("show 'Logout'", async () => {
        document.cookie = "UID=123; path=/";
        renderWithLoginContext(<Auth/>)
        expect(screen.getByRole('button', { name: "Logout" })).toBeTruthy();
    });

    test("click 'Logout'", () => {
        // Setup env
        document.cookie = "UID=123; path=/";
        const items = window.localStorage.getAll()
        // Test logout click
        renderAllContext(<Auth/>)
        userEvent.click(screen.getByRole('button', { name: "Logout" }))
        // Assert
        expect(mockLoginDispatch).toBeCalledWith({ type: "SET_STATE", state: { showLoginWindow: false, showSignUpWindow: false, showAccountWindow: false }})
        expect(mockDispatch).toBeCalledWith({ type: "SET_STATE", state: { profile: '' }})
        expect(Object.keys(items).length).toBe(0);
        expect(document.cookie).toBe('');
    });
});