import { useEffect, useContext, useCallback } from "react";
import Auth from "./pages/Login/AuthContainer"
import Context from "./context/Context";
import Header from "./pages/Header/Header";
import Menu from "./pages/Menu/MenuContainer";
import { EndPoint } from "./constants/endpoints";
import { useAppStateActions } from "./redux/useUserContextState";
import Sideview from "./pages/Sideview/SideviewContainer";
import Overview from "./pages/Overview/OverviewContainer";

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
        localStorage.setItem("link_token", data.link_token);
      }
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
      <Header>
        <Menu/>
        <Auth/>
        <Sideview/>
        <Overview/>
      </Header>
  );
};

export default App;
