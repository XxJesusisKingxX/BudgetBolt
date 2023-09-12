import '@testing-library/jest-dom'
import { screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event';
import { initLoginState, mockDispatch, mockLoginDispatch, renderAllContext } from '../../../../context/mock/Context.mock';
import Auth from '../..';
import { mockingFetch } from '../../../../utils/test';


// Global States Intialization
const showSignUpWindow = initLoginState.showSignUpWindow;
afterEach(() => {
    initLoginState.showSignUpWindow = showSignUpWindow;
})

describe("Signup",() => {
    const signupWorkflow = (username: string = "test", password: string = "Password1!", loading = true) => {
        initLoginState.showSignUpWindow = true
        renderAllContext(<Auth/>);
        // Fill login form
        userEvent.type(screen.getByLabelText('username'), username)
        userEvent.type(screen.getByLabelText('password'), password)
        // Verfiy loading is not shown
        expect(screen.queryByTestId('signup-loading')).toBeFalsy();
        // Attempt signup
        userEvent.click(screen.getByRole('button', { name: 'Submit' }));
        // Verify loading is shown
        if (loading) {
            expect(screen.getByTestId('signup-loading')).toBeTruthy();
        }
    }

    test("successfully signup", async () => {
        const mockFetch = mockingFetch(200);
        signupWorkflow();
        // Signup success now verify: setup account window does show, login is set, and profile name is updated
        await waitFor(()=> {
            expect(mockDispatch).toHaveBeenCalledWith({type: "SET_STATE", state: { profile: "test" }});
            expect(mockLoginDispatch).toBeCalledWith({type: "SET_STATE", state: { isLogin: true }});
            expect(mockLoginDispatch).toBeCalledWith({type: "SET_STATE", state: { showAccountWindow: true }});
            expect(mockLoginDispatch).toBeCalledWith({type: "SET_STATE", state: { showLoginWindow: false }});
            expect(mockLoginDispatch).toBeCalledWith({type: "SET_STATE", state: { showSignUpWindow: false }});
        })
        // Cleanup
        mockFetch.mockRestore();
    })
    test("username input invalid", async () => {
        const mockFetch = mockingFetch(200);
        signupWorkflow("2323","",false);
        // Username validation error is shown
        await waitFor(()=> {
            expect(screen.getByTestId('invalid-name')).toBeTruthy();
        })
        // Cleanup
        mockFetch.mockRestore();
    })
    test("password input invalid", async () => {
        const mockFetch = mockingFetch(200);
        signupWorkflow("Test","Per",false);
        // Password validation error is shown
        await waitFor(()=> {
            expect(screen.getByTestId('invalid-pass')).toBeTruthy();
        })
        // Cleanup
        mockFetch.mockRestore();
    })

    describe("Errors", () => {
        test("failed with name taken already error", async () => {
            const mockFetch = mockingFetch(409);
            signupWorkflow();
            // Login failed and appropiate error shown
            await waitFor(()=> {
                expect(screen.getByTestId('taken-err')).toBeTruthy();
            })
            // Cleanup
            mockFetch.mockRestore();
        })
        test("failed with internal server error", async () => {
            const mockFetch = mockingFetch(500);
            signupWorkflow();
            // Login failed and appropiate error shown
            await waitFor(()=> {
                expect(screen.getByTestId('server-err')).toBeTruthy();
            })
            // Cleanup
            mockFetch.mockRestore();
        })
    })
})