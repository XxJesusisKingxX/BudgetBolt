import '@testing-library/jest-dom'
import { render, screen } from '@testing-library/react';
import FinancialsSnapshot from '../FinancialsSnapshotComponent';

describe("Render Charts", () => {
    test('renders FinancialsSnapshot component with income, expenses, savings, and trend', () => {
        const props = {
            income: '5000',
            expenses: '3000',
            savings: '2000',
            level: '3',
            percentage: 0
        };
        render(<FinancialsSnapshot {...props} />);
        // Check if income, expenses, savings, and trend are rendered
        expect(screen.getByRole('heading', {name: `Total Income: $${props.income}`})).toBeInTheDocument();
        expect(screen.getByRole('heading', {name: `Total Expenses: $${props.expenses}`})).toBeInTheDocument();
        expect(screen.getByRole('heading', {name: `Total Savings: $${props.savings}`})).toBeInTheDocument();
    });
});
