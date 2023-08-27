import { useEffect, useContext, useCallback } from 'react';
import { usePlaidLink } from 'react-plaid-link';
import PlaidLink from './PlaidLink';
import UserContext from '../../context/UserContext';
import AppContext from '../../context/AppContext';
import LoginContext from '../../context/LoginContext';

// PlaidLinkContainer component
const PlaidLinkContainer = () => {
  // Accessing the user's profile, linkToken, and loginDispatch from contexts
  const { profile, userDispatch } = useContext(UserContext);
  const { linkToken } = useContext(AppContext);
  const { loginDispatch } = useContext(LoginContext);

  // Callback function to handle success after linking accounts with Plaid Link
  const onSuccess = useCallback(
    async (public_token: string) => {
      const linkAccounts = async () => {
        // Sending a POST request to set the access token
        await fetch("/api/set_access_token", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
          },
          body: new URLSearchParams({
            public_token: public_token,
            profile: profile,
          }),
        });

        // Sending a POST request to create accounts
        await fetch("/api/accounts/create", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
          },
          body: new URLSearchParams({
            profile: profile,
          }),
        });

        // Updating the login state
        loginDispatch({ type: "SET_STATE", state: { isLogin: true } });
      };
      linkAccounts();
      window.history.pushState("", "", "/");
    },
    [userDispatch, loginDispatch, profile]
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
        <PlaidLink plaidFunction={open} ready={ready}/>
    </>
  );
};

export default PlaidLinkContainer;
