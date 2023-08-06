import React, { useEffect, useContext, useCallback } from "react";
import { usePlaidLink } from "react-plaid-link";
import "./PlaidLink.css"
import Context from "../../context/Context";
import PlaidLink from "./PlaidLink";

const PlaidLinkContainer = () => {
  const { linkToken, profile, dispatch } = useContext(Context);

  const onSuccess = useCallback(
    (public_token: string) => {
      const linkAccounts = async () => {
        await fetch("/api/set_access_token", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
          },
          body: new URLSearchParams({
            public_token: public_token,
            profile: profile
          }),
        });
        await fetch("/api/accounts/create", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
          },
          body: new URLSearchParams({
            profile: profile
          }),
        });
        dispatch({ type: "SET_STATE", state: { isLogin: true } });
      };
      linkAccounts();
      window.history.pushState("", "", "/");
    },
    [dispatch]
  );

  let isOauth = false;
  const config: Parameters<typeof usePlaidLink>[0] = {
    token: linkToken!,
    onSuccess,
  };

  if (window.location.href.includes("?oauth_state_id=")) {
    // TODO: figure out how to delete this ts-ignore
    // @ts-ignore
    config.receivedRedirectUri = window.location.href;
    isOauth = true;
  }

  const { open, ready } = usePlaidLink(config);

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