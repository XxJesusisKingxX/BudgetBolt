import { useEffect, useContext, useCallback } from "react";
import Auth from "./pages/Auth"
import Header from "./pages/Header/Header";
import Menu from "./pages/Menu/MenuContainer";
import { EndPoint } from "./constants/endpoints";
import Sideview from "./pages/Sideview";
import Dashboard from "./pages/Dashboard";
import AppContext, { AppProvider } from "./context/AppContext";
import { LoginProvider } from "./context/LoginContext";
import { ThemeProvider } from "./context/ThemeContext";

const App = () => {
  const { profile, dispatch } = useContext(AppContext);

  const generateToken = useCallback(
    async () => {
      const baseURL = window.location.href;
      const url = new URL(EndPoint.CREATE_LINK_TOKEN, baseURL)
      const response = await fetch(url, {
        method: "POST",
        body: new URLSearchParams ({
          username: profile
        })
      });
      if (!response.ok) {
        dispatch({ type: "SET_STATE", state: { linkToken: '' }});
        return;
      }
      const data = await response.json();
      if (data) {
        if (data.error != null) {
          dispatch({ type: "SET_STATE", state: { linkToken: '' }});
          return;
        }
        dispatch({ type: "SET_STATE", state: { linkToken: data.link_token }});
        localStorage.setItem("link_token", data.link_token);
      }
    },
    [dispatch, profile]
  );

  useEffect(() => {
    const init = async () => {
      if (window.location.href.includes("?oauth_state_id=")) {
        dispatch({ type: "SET_STATE", state: { linkToken: localStorage.getItem("link_token") }});
        return;
      }
      generateToken();
    };
    init();
  }, [dispatch, profile, generateToken]);

  return (
      <Header>
        <AppProvider>
          <ThemeProvider>
            <Menu/>
            <LoginProvider>
              <Auth/>
            </LoginProvider>
            <Sideview/>
            <Dashboard/>
          </ThemeProvider>
        </AppProvider>
      </Header>
  );
};

export default App;
