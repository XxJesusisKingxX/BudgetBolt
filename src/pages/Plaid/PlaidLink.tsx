import { useEffect, useContext, useCallback } from 'react';
import { usePlaidLink } from 'react-plaid-link';
import AppContext from '../../context/AppContext';
import LoginContext from '../../context/LoginContext';
import { EndPoint } from '../../constants/endpoints';

// PlaidLinkContainer component
const PlaidLink = () => {
  // Accessing the user's profile, linkToken, and loginDispatch from contexts
  const { linkToken } = useContext(AppContext);
  const { loginDispatch } = useContext(LoginContext);

  // Callback function to handle success after linking accounts with Plaid Link
  const onSuccess = useCallback(
    async (public_token: string) => {
      const linkAccounts = async () => {
        // Step 1. Create plaid access token
        const getToken = await fetch(EndPoint.CREATE_ACCESS_TOKEN, {
          method: "POST",
          body: new URLSearchParams({
            public_token: public_token,
          }),
        });

        // Step 2. Create all accounts
        const getAccounts = await fetch(EndPoint.CREATE_ACCOUNTS, {
          method: "POST",
        });

        // Step 3. Create transactions for users
        const createTrans = await fetch(EndPoint.CREATE_TRANSACTIONS, {
          method: "POST",
        });

        if (getAccounts.ok && getToken.ok && createTrans.ok) {
          loginDispatch({ type: "SET_STATE", state: { isLogin: true } });
        } else {
          console.error("ERROR: Accounts#%d & Token#%d & Transactions#%d", getAccounts.status, getToken.status, createTrans.status)
        }
      };
      linkAccounts();
      window.history.pushState("", "", "/");
    },
    [loginDispatch]
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
