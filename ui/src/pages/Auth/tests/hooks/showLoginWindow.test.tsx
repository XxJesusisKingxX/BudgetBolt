import '@testing-library/jest-dom'
import { mockDispatch, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import AuthContainer from '../..';
import { fireEvent, cleanup } from '@testing-library/react';

afterEach(() => {
    cleanup();
})

describe("Show LoginWindow:", () => {
    test("click login button", async () => {
        const { getByTestId } = renderWithLoginContext(<AuthContainer/>);
        // Login button exists
        expect(getByTestId("auth-button")).toBeTruthy();
        // Click login button
        fireEvent.click(getByTestId("auth-button"));
        // Login window trigger state changes and displayed
        expect(mockDispatch).toHaveBeenNthCalledWith(1, { type: "SET_STATE", state: { showLoginWindow: true } });
    });
});
