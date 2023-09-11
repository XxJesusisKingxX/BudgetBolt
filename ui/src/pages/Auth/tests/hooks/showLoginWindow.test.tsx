import '@testing-library/jest-dom'
import { screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { mockLoginDispatch, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';

describe("Show LoginWindow:", () => {
    test("click login button", async () => {
        renderWithLoginContext(<Auth/>);
        // Login button exists
        expect(screen.getByRole('button', { name: "Login" })).toBeTruthy();
        // Click login button
        userEvent.click(screen.getByRole('button', { name: "Login" }));
        // Assertion
        expect(mockLoginDispatch).toHaveBeenNthCalledWith(1, { type: "SET_STATE", state: { showLoginWindow: true } });
    });
});
