import { useEffect, useContext, useCallback } from 'react';
import { usePlaidLink } from 'react-plaid-link';
import AppContext from '../../context/AppContext';
import LoginContext from '../../context/LoginContext';
import { EndPoint } from '../../constants/endpoints';

// PlaidLinkContainer component
const PlaidLink = () => {
  // Accessing the user's profile, linkToken, and loginDispatch from contexts
  const { profile } = useContext(AppContext);
  const { linkToken } = useContext(AppContext);
  const { loginDispatch } = useContext(LoginContext);

  // Callback function to handle success after linking accounts with Plaid Link
  const onSuccess = useCallback(
    async (public_token: string) => {
      const linkAccounts = async () => {
        // Step 1. Create plaid access token
        const getToken = await fetch(EndPoint.CREATE_ACCESS_TOKEN, {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
          },
          body: new URLSearchParams({
            public_token: public_token,
            profile: profile,
          }),
        });

        // Step 2. Create all accounts
        const getAccounts = await fetch(EndPoint.CREATE_ACCOUNTS, {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
          },
          body: new URLSearchParams({
            profile: profile,
          }),
        });

        // Step 3. Create transactions for users
        const createTrans = await fetch(EndPoint.CREATE_TRANSACTIONS, {
          method: "POST",
          body: new URLSearchParams({
              profile: profile,
          }),
        });

        if (getAccounts.ok && getToken.ok && createTrans.ok) {
          loginDispatch({ type: "SET_STATE", state: { isLogin: true } });
        } else {
          console.error("ERROR: Accounts#%d & Token#%d", getAccounts.status, getToken.status)
        }
      };
      linkAccounts();
      window.history.pushState("", "", "/");
    },
    [loginDispatch, profile]
  );

  let isOauth = false;
  // Configuration for Plaid Link
  const config: Parameters<typeof usePlaidLink>[0] = {
    token: linkToken!,
    onSuccess,
  };

  if (window.location.href.includes("?oauth_state_id=")) {
    config.receivedRedirectUri = window.location.href;
    isOauth = true;
  }

  // Using the Plaid Link hook
  const { open, ready } = usePlaidLink(config);

  // Automatically open Plaid Link if isOauth is true and ready
  useEffect(() => {
    if (isOauth && ready) {
      open();
    }
  }, [ready, open, isOauth]);

  return (
    <>
        <button
            className="btn btn--plaid"
            onClick={() => open()}     // Click event handler to call the provided function
            disabled={!ready}          // Disabling the button if not ready
        >
            Add Account
        </button>
    </>
  );
};

export default PlaidLink;
