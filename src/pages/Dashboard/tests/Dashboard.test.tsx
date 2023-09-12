import { render, screen } from '@testing-library/react';
import { mockLocalStorage } from '../../../utils/test';
import Dashboard from '..';

afterEach(() => {
  window.localStorage.clear();
})

describe('Dashboard Component', () => {
  mockLocalStorage();
  test('renders user name and date', () => {
    window.localStorage.setItem('profile', "John");
    render(<Dashboard/>);
    expect(screen.getByText('Welcome, JOHN')).toBeTruthy();
    expect(screen.getByTestId('dashboard-date').textContent).toMatch(/[A-Za-z]{3} \d{1,2}(st|nd|rd|th), \d{4}/);
  });
});
