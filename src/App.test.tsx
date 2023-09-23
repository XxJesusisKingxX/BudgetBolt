import '@testing-library/jest-dom'
import { waitFor } from '@testing-library/react'
import { mockLocalStorage } from './utils/test';
import { setupJestCanvasMock } from 'jest-canvas-mock';
import App from './App';
import { mockDispatch, renderWithAppContext } from './context/mock/AppContext.mock';
import { deleteCookie } from './utils/cookie';
import fetchMock  from 'jest-fetch-mock';
import { EndPoint } from './constants/endpoints';

afterEach(() => {
    deleteCookie("UID");
    window.localStorage.clear;
})
beforeEach(() => {
    setupJestCanvasMock();
    fetchMock.resetMocks();
})

describe("App",() => {
    mockLocalStorage();

    test("generate token with correct params and set token in local storage", async () => {
        // Mock
        fetchMock.enableMocks();
        fetchMock.mockResponse(JSON.stringify({link_token:"token"}), {status: 200})
        // Assign
        document.cookie = "UID=123; path=/";
        // Render
        renderWithAppContext(<App/>)
        // Assert
        await waitFor(() => {
            expect(fetchMock).toHaveBeenCalledWith(
                EndPoint.CREATE_LINK_TOKEN, {
                    method: "POST",
                }
            )
        })
        expect(mockDispatch).toBeCalledWith({"state": {"linkToken": "token"}, "type": "SET_STATE"});
        expect(window.localStorage.getItem('link_token')).toBe("token");
    })
    test("generate token from local storage if already oauth", async () => {
        // Mock
        fetchMock.enableMocks();
        fetchMock.mockResponse(JSON.stringify({link_token:"token"}), {status: 200})
        // Assign
        document.cookie = "UID=123; path=/";
        Object.defineProperty(window, 'location', {
            value: {
            href: 'http://example.com/?oauth_state_id=',
            },
        });
        // Render
        await waitFor(() => {
            renderWithAppContext(<App/>)
        })
        // Assert
        expect(mockDispatch).toBeCalledWith({"state": {"linkToken": "token"}, "type": "SET_STATE"});
    })
})