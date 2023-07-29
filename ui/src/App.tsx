import React, { useEffect, useContext, useCallback } from "react";

import PlaidLink from "./components/Plaid/PlaidLinkContainer";
import Context from "./context/Context";
import Header from "./components/Header/Header";
import Home from "./pages/Home/Home";

const App = () => {
  const { dispatch } = useContext(Context);

  const generateToken = useCallback(
    async () => {
      const response = await fetch("/api/create_link_token", {
        method: "POST",
      });
      if (!response.ok) {
        dispatch({ type: "SET_STATE", state: { linkToken: null } });
        return;
      }
      const data = await response.json();
      if (data) {
        if (data.error != null) {
          dispatch({
            type: "SET_STATE",
            state: {
              linkToken: null
            },
          });
          return;
        }
        dispatch({ type: "SET_STATE", state: { linkToken: data.link_token } });
      }
      // Save the link_token to be used later in the Oauth flow.
      localStorage.setItem("link_token", data.link_token);
    },
    [dispatch]
  );

  useEffect(() => {
    const init = async () => {
      if (window.location.href.includes("?oauth_state_id=")) {
        dispatch({
          type: "SET_STATE",
          state: {
            linkToken: localStorage.getItem("link_token"),
          },
        });
        return;
      }
      generateToken();
    };
    init();
  }, [dispatch, generateToken]);

  return (
    <>
      <Header>
        <Home/>
        <PlaidLink/>
      </Header>
    </>
  );
};

export default App;
