import { useEffect, useContext, useCallback } from "react";
import Auth from "./components/Login/AuthContainer"
import Context from "./context/Context";
import Header from "./components/Header/Header";
import Home from "./pages/Home/Home";
import Menu from "./components/Menu/MenuContainer";
import { EndPoint } from "./enums/endpoints";
import { useAppStateActions } from "./redux/redux";

const App = () => {
  const { profile, dispatch } = useContext(Context);
  const { setLinkTokenState } = useAppStateActions();

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
        setLinkTokenState("")
        return;
      }
      const data = await response.json();
      if (data) {
        if (data.error != null) {
          setLinkTokenState("")
          return;
        }
        setLinkTokenState(data.link_token);
      }
      localStorage.setItem("link_token", data.link_token);
    },
    [dispatch, profile]
  );

  useEffect(() => {
    const init = async () => {
      if (window.location.href.includes("?oauth_state_id=")) {
        setLinkTokenState(localStorage.getItem("link_token"));
        return;
      }
      generateToken();
    };
    init();
  }, [dispatch, profile, generateToken]);

  return (
    <>
      <Header>
        <Menu/>
        <Auth/>
        <Home/>
      </Header>
    </>
  );
};

export default App;
