import '@testing-library/jest-dom'
import { screen, waitFor } from '@testing-library/react';
import fetchMock  from 'jest-fetch-mock';
import FinancialsSnapshot from '../index';
import { renderWithAppContext, mockDispatch, initAppState } from '../../../context/mock/AppContext.mock';

const totalIncome = initAppState.totalIncome;
const totalExpenses = initAppState.totalExpenses;

afterEach(() => {
    initAppState.totalIncome = totalIncome;
    initAppState.totalExpenses = totalExpenses;
})

beforeEach(() => {
    fetchMock.resetMocks();
});

describe("FinancialsSnapshot", () => {
    test('renders income, expenses, savings, chart percentage, and trigger upsertIncome', async () => {
        // Set context
        initAppState.totalIncome = 1000;
        initAppState.totalExpenses = 550;

        // Mock
        fetchMock.enableMocks();
        fetchMock.mockResponseOnce(JSON.stringify({}), { status: 200 });
        fetchMock.mockResponseOnce(JSON.stringify({"incomes":[{"income_amount":800},{"income_amount":1000}]}), { status: 200 });

        renderWithAppContext(<FinancialsSnapshot/>);

        // Assert
        expect(screen.getByRole('heading', {name: "Total Income: $1000.00"})).toBeInTheDocument();
        expect(screen.getByRole('heading', {name: "Total Expenses: $550.00"})).toBeInTheDocument();
        expect(screen.getByRole('heading', {name: "Total Savings: $450.00"})).toBeInTheDocument();
        expect(screen.getByText("45%")).toBeInTheDocument();
        expect(screen.getByText("55%")).toBeInTheDocument();
        await waitFor(() => {
            expect(mockDispatch).toBeCalledWith({"state": {"totalIncome": 1800}, "type": "SET_STATE"});
        })
    });
    test('does not render chart percentage', () => {
        // Set context
        initAppState.totalIncome = 0;
        initAppState.totalExpenses = 0;

        renderWithAppContext(<FinancialsSnapshot/>);

        // Assert
        expect(screen.queryByTestId('financials-snapshot-top-percent')).toBeFalsy();
        expect(screen.queryByTestId('financials-snapshot-btm-percent')).toBeFalsy();
    });
    test('render chart percentage over', () => {
        // Set context
        initAppState.totalIncome = 1000;
        initAppState.totalExpenses = 1100;

        renderWithAppContext(<FinancialsSnapshot/>);

        // Assert
        expect(screen.getByText("0%")).toBeInTheDocument();
        expect(screen.getByText("110%")).toBeInTheDocument();
    });
});