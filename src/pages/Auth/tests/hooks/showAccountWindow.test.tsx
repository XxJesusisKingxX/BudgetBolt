import '@testing-library/jest-dom'
import { screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event';
import { initState, renderWithLoginContext } from '../../../../context/mock/LoginContext.mock';
import Auth from '../..';

// Global States Intialization
const showAccountWindow = initState.showAccountWindow;

afterEach(() => {
    initState.showAccountWindow = showAccountWindow;
})

describe("Show LoginWindow:", () => {
    test("click login button", async () => {
        initState.showAccountWindow = true;
        renderWithLoginContext(<Auth/>);
        // Plaid button exists
        expect(screen.getByRole('button', { name: "Add Account" })).toBeTruthy();
        // Click plaid button
        userEvent.click(screen.getByRole('button', { name: "Add Account" }));
        // No need to test mock trigger because a succesful sign up will test it
        // No need to test plaid button since already tested here
    });
});
