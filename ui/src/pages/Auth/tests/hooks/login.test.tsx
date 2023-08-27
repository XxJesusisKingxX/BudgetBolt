import '@testing-library/jest-dom'
import { cleanup, fireEvent, getByTestId, queryByTestId, waitFor } from '@testing-library/react'
import { loginFormFill, mockingFetch } from '../test-utils';
import AuthContainer from '../..';
import UserContext from '../../../../context/UserContext'
import { render } from '@testing-library/react'
import LoginContext from '../../../../context/LoginContext';

// Create mock Login Context
const mockLoginDispatch = jest.fn();
const initLoginState = {
    showAccountWindow: false,
    showLoginWindow: false,
    showSignUpWindow: false,
    isLogin: false,
    loginDispatch: mockLoginDispatch
};

// Create mock User Context
const mockUserDispatch = jest.fn();
const initUserState = {
    profile: "",
    userDispatch: mockUserDispatch
};

// Create mock User Context wrapped around Login Context
const renderContext = () => {
  const element = render(
    <UserContext.Provider value={initUserState}>
        <LoginContext.Provider value={initLoginState}>
            <AuthContainer/>
        </LoginContext.Provider>
    </UserContext.Provider>
  );
  return element;
}

// Global States Intialization
const showLoginWindow = initLoginState.showLoginWindow;
afterEach(() => {
    initLoginState.showLoginWindow = showLoginWindow;
    cleanup();
})

describe("Login",() => {
    const loginWorkFlow = (username: string = "test", password: string = "Password1!", loading = true) => {
        initLoginState.showLoginWindow = true
        const element = renderContext().container;
        // Fill login form
        loginFormFill(username, password, element);
        // Verfiy loading is not shown
        expect(queryByTestId(element,'login-loading')).toBeFalsy();
        // Attempt login
        fireEvent.click(getByTestId(element, 'login-button'));
        // Verify loading is shown
        if (loading) {
            expect(getByTestId(element,'login-loading')).toBeTruthy();
        }
        return element
    }

    test("successfully login", async () => {
        const mockFetch = mockingFetch(200);
        loginWorkFlow();
        // Login success now verify: setup account window does not shows, login is set, and profile name is updated
        await waitFor(()=> {
            expect(mockUserDispatch).toHaveBeenCalledWith({type: "SET_STATE", state: { profile: "test" }});
            expect(mockLoginDispatch).toHaveBeenCalledWith({type: "SET_STATE", state: { isLogin: true }});
            expect(mockLoginDispatch).not.toHaveBeenCalledWith({type: "SET_STATE", state: { showAccountWindow: true }});
        })
        // Cleanup
        mockFetch.mockRestore();
    })
    test("username input invalid", async () => {
        const mockFetch = mockingFetch(200);
        const element = loginWorkFlow("2323","",false);
        // Username validation error is shown
        await waitFor(()=> {
            expect(getByTestId(element, 'invalid-name')).toBeTruthy();
        })
        // Cleanup
        mockFetch.mockRestore();
    })
    test("password input invalid", async () => {
        const mockFetch = mockingFetch(200);
        const element = loginWorkFlow("Test","Per",false);
        // Password validation error is shown
        await waitFor(()=> {
            expect(getByTestId(element, 'invalid-pass')).toBeTruthy();
        })
        // Cleanup
        mockFetch.mockRestore();
    })

    describe("Errors", () => {
        test("failed with name not existing error", async () => {
            const mockFetch = mockingFetch(404);
            const element = loginWorkFlow();
            // Login failed and appropiate error shown
            await waitFor(()=> {
                expect(getByTestId(element, 'name-err')).toBeTruthy();
            })
            // Cleanup
            mockFetch.mockRestore();
        })
        test("failed with authentication failed error", async () => {
            const mockFetch = mockingFetch(401);
            const element = loginWorkFlow();
            // Login failed and appropiate error shown
            await waitFor(()=> {
                expect(getByTestId(element, 'auth-err')).toBeTruthy();
            })
            // Cleanup
            mockFetch.mockRestore();
        })
        test("failed with internal server error", async () => {
            const mockFetch = mockingFetch(500);
            const element = loginWorkFlow();
            // Login failed and appropiate error shown
            await waitFor(()=> {
                expect(getByTestId(element, 'server-err')).toBeTruthy();
            })
            // Cleanup
            mockFetch.mockRestore();
        })
    })
})