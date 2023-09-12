import '@testing-library/jest-dom'
import { waitFor } from '@testing-library/react'
import { mockLocalStorage, mockingFetch } from './utils/test';
import App from './App';
import { mockDispatch, renderWithAppContext } from './context/mock/AppContext.mock';
import { deleteCookie } from './utils/cookie';

afterEach(() => {
    deleteCookie("UID")
})

describe("App",() => {
    mockLocalStorage();
    test("generate token", async () => {
        const mockFetch = mockingFetch(200, {link_token:"token"})
        document.cookie = "UID=123; path=/";
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
        document.cookie = "UID=123; path=/";
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