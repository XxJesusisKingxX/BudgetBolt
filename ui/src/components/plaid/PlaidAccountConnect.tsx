import React, { useEffect, useContext, useCallback } from "react";
import { usePlaidLink } from "react-plaid-link";
import "./PlaidAccountConnect.css"
import Context from "../../Context";

const Link = () => {
  const { linkToken, dispatch } = useContext(Context);

  const onSuccess = useCallback(
    (public_token: string) => {
      // If the access_token is needed, send public_token to server
      const exchangePublicTokenForAccessToken = async () => {
        const response = await fetch("/api/set_access_token", {
          method: "POST",
          headers: {
            "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
          },
          body: `public_token=${public_token}`,
        });
        if (!response.ok) {
          dispatch({
            type: "SET_STATE",
            state: {
              itemId: `no item_id retrieved`,
              accessToken: `no access_token retrieved`,
              isItemAccess: false,
            },
          });
          return;
        }
        const data = await response.json();
        dispatch({
          type: "SET_STATE",
          state: {
            itemId: data.item_id,
            accessToken: data.access_token,
            isItemAccess: true,
          },
        });
      };
      exchangePublicTokenForAccessToken();
      dispatch({ type: "SET_STATE", state: { linkSuccess: true } });
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
    <button className="plaid_addbutton" onClick={() => open()} disabled={!ready}>
      Add Account
    </button>
  );
};

export default Link;
