import { render, screen } from '@testing-library/react';
import DashboardComponent from '../DashboardComponent';

describe('Dashboard Component', () => {
  test('renders user name and date', () => {
    render(<DashboardComponent user="John" date="2023-09-10" />);
    expect(screen.getByText('Welcome, John')).toBeTruthy();
    expect(screen.getByText('~ Today is 2023-09-10 ~')).toBeTruthy();
  });
});
