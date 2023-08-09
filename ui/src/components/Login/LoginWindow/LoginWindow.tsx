import "./LoginWindow.css";
import { ChangeEventHandler, KeyboardEventHandler, MouseEventHandler } from "react";

interface Props {
    mode: string
    close: Function
    open: Function
    isAuthError: boolean
    isInvalidPass: boolean
    isInvalidName: boolean
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
const LoginWindow: React.FC<Props> = ({ mode, close, open, isAuthError, isInvalidPass, isInvalidName, showLoading, username, password,userKeyUp, userChange, passKeyUp, passChange, login, loginOnEnter}) => {
    const cancel = `/images/${mode}/cancel.png`
    const loading = `/images/${mode}/loading.png`
    return (
        <>
            <div className="login-window-container">
                <h1 className="login-window-title">Sign In<img className="login-window-cancel" src={cancel} onClick={() => close()}/></h1>
                <div id="username">
                    <input value={username} onKeyUp={userKeyUp} onChange={userChange} name="user"placeholder="User" required/>
                    {isInvalidName ?
                    <span>
                        Enter a valid username:
                        <br/>
                        - Must start with a letter or an underscore (_)
                        <br/>
                        - Valid characters: 0-9, A-z, (_)
                        <br/>
                    </span> : null}
                </div>
                <div id="password">
                    <input type="password" value={password} onKeyDown={loginOnEnter} onKeyUp={passKeyUp} onChange={passChange} name="pass" placeholder="Password" required/>
                    {isInvalidPass ?
                    <span>
                        Password is invalid.
                        <br/>
                        Must have the followings:
                        <br/>
                        - At least one digit (0-9)
                        <br/>
                        - At least one lowercase letter (a-z)
                        <br/>
                        - At least one uppercase letter (A-Z)
                        <br/>
                        - At least one special character: !@#$%^&*
                        <br/>
                        - At least 8 characters long
                    </span> : null}
                </div>
                <div className="login-window-create">Create an Account: <span className="login-window-signup" onClick={() => open()}>Sign Up</span></div>
                {isAuthError ?
                    <span className="auth-error">Oops! The username or password is incorrect.</span> : null}
                {showLoading ? <img src={loading} className="login-loading"/> : <div onClick={login} className="login-acc-button"><span>Login</span></div>}
            </div>
        </>
    );
};

export default LoginWindow;