import "./LoginWindow.css";
import { ChangeEventHandler, KeyboardEventHandler, MouseEventHandler } from "react";

interface Props {
    mode: string
    close: Function
    open: Function
    error: boolean
    isInvalidInput: boolean
    showLoading: boolean
    username: string
    password: string
    userChange: ChangeEventHandler<HTMLInputElement>
    userKeyUp: KeyboardEventHandler<HTMLInputElement>
    passChange: ChangeEventHandler<HTMLInputElement>
    passKeyUp: KeyboardEventHandler<HTMLInputElement>
    login: MouseEventHandler<HTMLDivElement>
    loginOnEnter: KeyboardEventHandler<HTMLInputElement>
}
const LoginWindow: React.FC<Props> = ({ mode, close, open, error, isInvalidInput, showLoading, username, password,userKeyUp, userChange, passKeyUp, passChange, login, loginOnEnter}) => {
    const cancel = `/images/${mode}/cancel.png`
    const loading = `/images/${mode}/loading.png`
    return (
        <>
            <div className="login-window-container">
                <h1 className="login-window-title">Sign In<img className="login-window-cancel" src={cancel} onClick={() => close()}/></h1>
                {isInvalidInput ?
                <div className="login-invalid">
                    Enter a valid username and password.
                    <br/>
                    - Must start with a letter or an underscore (_)
                    <br/>
                    - Valid characters: 0-9, A-z, (_).
                    <br/>
                </div> : null}
                <input id="username" value={username} onKeyUp={userKeyUp} onChange={userChange} name="user"placeholder="User" required/>
                <input id="password" type="password" value={password} onKeyDown={loginOnEnter} onKeyUp={passKeyUp} onChange={passChange} name="pass" placeholder="Password" required/>
                <div className="login-window-create">Create an Account: <span className="login-window-signup" onClick={() => open()}>Sign Up</span></div>
                {error ? <div className="login-failed">Oops! The username or password you entered is incorrect. Please double check and try again.</div> : null}
                {showLoading ? <img src={loading} className="login-loading"/> : <div onClick={login} className="login-acc-button"><span>Login</span></div>}
            </div>
        </>
    );
};

export default LoginWindow;