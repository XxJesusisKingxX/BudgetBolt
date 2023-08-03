import './LoginWindow.css';
import cancel from '../icons/cancel-dark.png'
import loading from '../../../images/loading.png'
import { ChangeEventHandler, FormEventHandler } from 'react';

interface Props {
    close: Function
    open: Function
    isLoading: boolean
    username: string
    password: string
    userChange: ChangeEventHandler<HTMLInputElement>
    passChange: ChangeEventHandler<HTMLInputElement>
    login: FormEventHandler<HTMLFormElement>
}
const LoginWindow: React.FC<Props> = ({ close, open, isLoading, username, password, userChange, passChange, login }) => {
    return (
        <>
            <div className="login-window-container">
                <h1 className="login-window-title">Sign In<img className="login-window-cancel" src={cancel} onClick={() => close()}/></h1>
                <form onSubmit={login}>
                    <input id="username" value={username} onChange={userChange} name="user"placeholder="User" required/>
                    <input id="password" type="password" value={password} onChange={passChange} name="pass" placeholder="Password" required/>
                    <br/>
                    {!isLoading ? <button type="submit" className="login-submit-button">Login</button> : <img src={loading} className="login-loading"/>}
                    <div className="login-window-create">Create an Account: <span className="login-window-signup" onClick={() => open()}>Sign Up</span></div>
                </form>
            </div>
        </>
    );
};

export default LoginWindow;