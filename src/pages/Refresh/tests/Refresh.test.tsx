import { render, screen, waitFor } from '@testing-library/react';
import Refresh from '..';
import userEvent from '@testing-library/user-event';
import { mockDispatch, renderAllContext } from '../../../context/mock/Context.mock';
import RefreshComponent from '../RefreshComponent';

describe("Refresh", () => {
  test("renders", () => {
    render(<Refresh/>).container;
    expect(screen.getByRole('img', { name: "Refresh" })).toBeTruthy();
  });
  test("trigger refresh", async () => {
    renderAllContext(<Refresh/>);
    // Trigger refresh button
    userEvent.click(screen.getByRole('img', { name: "Refresh" }));
    // Assertion
    expect(mockDispatch).toBeCalledWith({"state": {"isTransactionsRefresh": true, "lastTransactionsUpdate": expect.any(Date)}, "type": "SET_STATE"});
  });
  test("trigger classname switch alt", () => {
    const props = {
      mode: "light",
      isRefresh: true,
      refresh: () => {}
    }
    renderAllContext(<RefreshComponent {...props}/>).container;
    // Assertion
    expect(screen.getByRole('img', { name: "Refresh" }).classList[0]).toContain('sidebar__refresh')
    expect(screen.getByRole('img', { name: "Refresh" }).classList[1]).toContain('sidebar__refresh--load')

  });
  test("trigger classname switch regular", () => {
    const props = {
      mode: "light",
      isRefresh: false,
      refresh: () => {}
    }
    renderAllContext(<RefreshComponent {...props}/>).container;
    // Assertion
    expect(screen.getByRole('img', { name: "Refresh" }).classList[0]).toContain('sidebar__refresh')
    expect(screen.getByRole('img', { name: "Refresh" }).classList[1]).toContain('sidebar__refresh--loadalt')
  });
});
