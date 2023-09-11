import '@testing-library/jest-dom'
import { waitFor } from '@testing-library/react'
import { mockLocalStorage, mockingFetch } from './utils/test';
import App from './App';
import { mockDispatch, renderAllContext } from './context/mock/Context.mock';

describe("App",() => {
    mockLocalStorage();
    test("generate token", async () => {
        const mockFetch = mockingFetch(200, {link_token:"token"})
        await waitFor(() => {
            renderAllContext(<App/>)
        })
        expect(mockDispatch).toBeCalledWith({"state": {"linkToken": "token"}, "type": "SET_STATE"});
        // Assert successful login
        expect(window.localStorage.getItem('link_token')).not.toBeNull();
        mockFetch.mockRestore()
    })
    test("generate token fetch failed", async () => {
        const mockFetch = mockingFetch(500, {})
        await waitFor(() => {
            renderAllContext(<App/>)
        })
        // Assert successful login
        expect(mockDispatch).toBeCalledWith({"state": {"linkToken": ""}, "type": "SET_STATE"});
        mockFetch.mockRestore()
    })
    test("generate token fetch has error", async () => {
        const mockFetch = mockingFetch(200, {error:"failed"})
        await waitFor(() => {
            renderAllContext(<App/>)
        })
        // Assert successful login
        expect(mockDispatch).toBeCalledWith({"state": {"linkToken": ""}, "type": "SET_STATE"});
        mockFetch.mockRestore()
    })
    test("generate token from local storage if already oauth", async () => {
        const mockFetch = mockingFetch(200, {link_token:"token"})
        // Mock window
        Object.defineProperty(window, 'location', {
            value: {
            href: 'http://example.com/?oauth_state_id=',
            },
        });
        await waitFor(() => {
            renderAllContext(<App/>)
        })
        // Assert successful login
        expect(mockDispatch).toBeCalledWith({"state": {"linkToken": "token"}, "type": "SET_STATE"});
        mockFetch.mockRestore()
    })
})