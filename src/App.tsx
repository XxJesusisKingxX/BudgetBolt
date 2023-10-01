import { useEffect, useContext, useCallback } from "react";
import Auth from "./pages/Auth"
import Header from "./pages/Header/Header";
import Menu from "./pages/Menu/MenuContainer";
import { EndPoint } from "./constants/endpoints";
import Sideview from "./pages/Sideview";
import Dashboard from "./pages/Dashboard";
import AppContext from "./context/AppContext";
import { LoginProvider } from "./context/LoginContext";
import { ThemeProvider } from "./context/ThemeContext";
import { getCookie } from "./utils/cookie";
import PlaidLink from "./pages/Plaid/PlaidLink";

const App = () => {
  const { profile, dispatch } = useContext(AppContext);

  const generateToken = useCallback(
    async () => {
      const response = await fetch(EndPoint.CREATE_LINK_TOKEN, {
        method: "POST",
      });

      const data = await response.json();

      if (data) {
        dispatch({ type: "SET_STATE", state: { linkToken: data.link_token }});
        localStorage.setItem("link_token", data.link_token);
      }

    },
    [dispatch]
  );

  useEffect(() => {
    const init = async () => {
      if (window.location.href.includes("?oauth_state_id=")) {
        dispatch({ type: "SET_STATE", state: { linkToken: localStorage.getItem("link_token") }});
        return;
      } else {
        generateToken();
      }
    };
    if (getCookie("UID")) init();
  }, [profile, generateToken, dispatch]);

  return (
      <Header>
        <Menu/>
        <LoginProvider>
          <ThemeProvider>
            <Auth/>
            {getCookie("UID") != null ?
            <>
              <Sideview/>
              <Dashboard/>
              <div className="add-account">
                <PlaidLink/>
              </div>
            </>
            :
            null
            }
          </ThemeProvider>
        </LoginProvider>
      </Header>
  );
};

export default App;