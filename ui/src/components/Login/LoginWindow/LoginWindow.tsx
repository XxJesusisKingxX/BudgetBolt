import './LoginWindow.css';
import cancel from '../icons/cancel-dark.png'
import loading from '../../../images/loading.png'
import { ChangeEventHandler, MouseEventHandler } from 'react';

interface Props {
    close: Function
    open: Function
    error: boolean
    isLoading: boolean
    username: string
    password: string
    userChange: ChangeEventHandler<HTMLInputElement>
    passChange: ChangeEventHandler<HTMLInputElement>
    login: MouseEventHandler<HTMLDivElement>
}
const LoginWindow: React.FC<Props> = ({ close, open, error, isLoading, username, password, userChange, passChange, login }) => {
    return (
        <>
            <div className="login-window-container">
                <h1 className="login-window-title">Sign In<img className="login-window-cancel" src={cancel} onClick={() => close()}/></h1>
                <input id="username" value={username} onChange={userChange} name="user"placeholder="User" required/>
                <input id="password" type="password" value={password} onChange={passChange} name="pass" placeholder="Password" required/>
                <div className="login-window-create">Create an Account: <span className="login-window-signup" onClick={() => open()}>Sign Up</span></div>
                {error ? <div className="login-error">Login Failed</div> : null}
                {!isLoading ? <div onClick={login} className="login-acc-button"><span>Login</span></div> : <img src={loading} className="login-loading"/>}
            </div>
        </>
    );
};

export default LoginWindow;