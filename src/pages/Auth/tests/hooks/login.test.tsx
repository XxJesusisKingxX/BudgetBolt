import '@testing-library/jest-dom'
import { fireEvent, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event';
import Auth from '../..';
import { initLoginState, mockDispatch, mockLoginDispatch, renderAllContext } from '../../../../context/mock/Context.mock';
import { mockingFetch } from '../../../../utils/test';


// Global States Intialization
const showLoginWindow = initLoginState.showLoginWindow;
afterEach(() => {
    initLoginState.showLoginWindow = showLoginWindow;

})

describe("Login",() => {
    const loginWorkFlow = (username: string = "test", password: string = "Password1!", loading = true) => {
        initLoginState.showLoginWindow = true
        renderAllContext(<Auth/>);
        // Fill login form
        userEvent.type(screen.getByLabelText('username'), username)
        userEvent.type(screen.getByLabelText('password'), password)
        // Verfiy loading is not shown
        expect(screen.queryByTestId('login-loading')).toBeFalsy();
        // Attempt login
        userEvent.click(screen.getByTestId('login-button'));
        // Verfiy loading is shown
        if (loading) {
            expect(screen.getByTestId('login-loading')).toBeTruthy();
        }
    }

    test("successfully login with 'enter' pressed", async () => {
        const mockFetch = mockingFetch(200);
        initLoginState.showLoginWindow = true
        renderAllContext(<Auth/>)
        userEvent.type(screen.getByLabelText('username'), 'test')
        userEvent.type(screen.getByLabelText('password'), 'Password1!')
        // Assert user input changes
        expect(screen.getByLabelText('username')).toHaveValue('test')
        expect(screen.getByLabelText('password')).toHaveValue('Password1!')
        // Press enter keydown
        fireEvent.keyDown(screen.getByLabelText('password'), { key: 'Enter', code: 'Enter' })
        // Assert successful login
        await waitFor(()=> {
            expect(mockDispatch).toHaveBeenCalledWith({type: "SET_STATE", state: { profile: "test" }});
        })
        expect(mockLoginDispatch).toHaveBeenCalledWith({type: "SET_STATE", state: { isLogin: true }});
        expect(mockLoginDispatch).not.toHaveBeenCalledWith({type: "SET_STATE", state: { showAccountWindow: true }});
        mockFetch.mockRestore();
    })
    test("successfully login", async () => {
        const mockFetch = mockingFetch(200);
        loginWorkFlow();
        // Login success now verify: setup account window does not shows, login is set, and profile name is updated
        await waitFor(()=> {
            expect(mockDispatch).toHaveBeenCalledWith({type: "SET_STATE", state: { profile: "test" }});
        })
        expect(mockLoginDispatch).toHaveBeenCalledWith({type: "SET_STATE", state: { isLogin: true }});
        expect(mockLoginDispatch).not.toHaveBeenCalledWith({type: "SET_STATE", state: { showAccountWindow: true }});
        mockFetch.mockRestore();
    })
    test("username input invalid", async () => {
        const mockFetch = mockingFetch(200);
        // Perform the same steps to login as user
        loginWorkFlow("2323","",false);
        // Assertion
        await waitFor(()=> {
            expect(screen.getByTestId('invalid-name')).toBeTruthy();
        })
        mockFetch.mockRestore();
    })
    test("password input invalid", async () => {
        const mockFetch = mockingFetch(200);
        // Perform the same steps to login as user
        loginWorkFlow("Test","Per",false);
        // Assertion
        await waitFor(()=> {
            expect(screen.getByTestId('invalid-pass')).toBeTruthy();
        })
        mockFetch.mockRestore();
    })

    describe("Errors", () => {
        test("failed with name not existing error", async () => {
            const mockFetch = mockingFetch(404);
            // Perform the same steps to login as user
            loginWorkFlow();
            // Assertion
            await waitFor(()=> {
                expect(screen.getByTestId('name-err')).toBeTruthy();
            })
            mockFetch.mockRestore();
        })
        test("failed with authentication failed error", async () => {
            const mockFetch = mockingFetch(401);
            // Perform the same steps to login as user
            loginWorkFlow();
            // Assertion
            await waitFor(()=> {
                expect(screen.getByTestId('auth-err')).toBeTruthy();
            })
            mockFetch.mockRestore();
        })
        test("failed with internal server error", async () => {
            const mockFetch = mockingFetch(500);
            // Perform the same steps to login as user
            loginWorkFlow();
            // Assertion
            await waitFor(()=> {
                expect(screen.getByTestId('server-err')).toBeTruthy();
            })
            mockFetch.mockRestore();
        })
    })
})