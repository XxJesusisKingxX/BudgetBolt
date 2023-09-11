import { screen } from '@testing-library/react';
import { initAppState, initLoginState, renderAllContext } from '../../../context/mock/Context.mock';
import Sideview from '..';

// Global States Intialization
const lastTransactionsUpdate = initAppState.lastTransactionsUpdate;
afterEach(() => {
    initAppState.lastTransactionsUpdate = lastTransactionsUpdate;
})

describe("Sideview", () => {
    test('renders SideviewComponent with formatted last update time when logged in', () => {
        initLoginState.isLogin = true;
        initAppState.lastTransactionsUpdate = new Date("2023-09-10 10:00:00");
        renderAllContext(<Sideview/>);
        expect(screen.getByText('Last updated: 9/10/2023 10:00:00 AM')).toBeTruthy();
    });
})
