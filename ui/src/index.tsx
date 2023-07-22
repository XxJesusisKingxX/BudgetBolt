import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import './AppConfig.css';
import { QuickstartProvider } from "./context/Context";

ReactDOM.render(
  <React.StrictMode>
    <QuickstartProvider>
      <App />
    </QuickstartProvider>
  </React.StrictMode>,
  document.getElementById("root")
);

