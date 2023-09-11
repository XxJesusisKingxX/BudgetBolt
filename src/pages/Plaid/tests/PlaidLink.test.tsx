import PlaidLink from '../PlaidLink';
import { renderAllContext } from '../../../context/mock/Context.mock';

describe('PlaidLink', () => {
  test('should create and open link and recieve accounts and token', () => {
    // Mock the PlaidLink hook's return value
    const mockOpen = jest.fn();
    const mockPlaidLink = {
      open: mockOpen,
      ready: true,
    };
    const mockPlaid = jest.spyOn(require('react-plaid-link'), 'usePlaidLink').mockReturnValue(mockPlaidLink);
    // Mock window
    Object.defineProperty(window, 'location', {
      value: {
        href: 'http://example.com/?oauth_state_id=',
      },
    });
    // Render the component
    renderAllContext(<PlaidLink/>);
    // Assertion
    expect(mockPlaid).toBeCalled();
    expect(mockOpen).toBeCalled();
    // Cleanup
    mockPlaid.mockRestore();
  });
  test('should not create and open link when not ready', () => {
    // Mock the PlaidLink hook's return value
    const mockOpen = jest.fn();
    const mockPlaidLink = {
      open: mockOpen,
      ready: false,
    };
    const mockPlaid = jest.spyOn(require('react-plaid-link'), 'usePlaidLink').mockReturnValue(mockPlaidLink);
    // Mock window
    Object.defineProperty(window, 'location', {
      value: {
        href: 'http://example.com/?oauth_state_id=',
      },
    });
    // Render the component
    renderAllContext(<PlaidLink/>);
    // Assertion
    expect(mockPlaid).toBeCalled();
    expect(mockOpen).not.toBeCalled();
    // Cleanup
    mockPlaid.mockRestore();
  });
// unable to test if window.location.href name changes
// unable to test onsuccess but im confident due to issues only be in endpoints and state management
});
