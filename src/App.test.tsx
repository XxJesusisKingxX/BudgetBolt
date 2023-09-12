import '@testing-library/jest-dom'
import { waitFor } from '@testing-library/react'
import { mockLocalStorage, mockingFetch } from './utils/test';
import App from './App';
import { initAppState, mockDispatch, renderWithAppContext } from './context/mock/AppContext.mock';

const profile = initAppState.profile;
afterEach(() => {
    initAppState.profile = profile;
})

describe("App",() => {
    mockLocalStorage();
    test("generate token", async () => {
        const mockFetch = mockingFetch(200, {link_token:"token"})
        initAppState.profile = "testing"
        renderWithAppContext(<App/>)
        await waitFor(() => {
            expect(mockDispatch). toBeCalledWith({"state": {"linkToken": "token"}, "type": "SET_STATE"});
        })
        // Assert successful login
        expect(window.localStorage.getItem('link_token')).not.toBeNull();
        mockFetch.mockRestore()
    })
    test("generate token from local storage if already oauth", async () => {
        const mockFetch = mockingFetch(200, {link_token:"token"})
        initAppState.profile = "testing"
        // Mock window
        Object.defineProperty(window, 'location', {
            value: {
            href: 'http://example.com/?oauth_state_id=',
            },
        });
        await waitFor(() => {
            renderWithAppContext(<App/>)
        })
        // Assert successful login
        expect(mockDispatch).toBeCalledWith({"state": {"linkToken": "token"}, "type": "SET_STATE"});
        mockFetch.mockRestore()
    })
})