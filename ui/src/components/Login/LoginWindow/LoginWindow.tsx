import './LoginWindow.css';
import cancel from '../icons/cancel-dark.png'

interface Props {
    close: Function
    open: Function
}
const LoginWindow: React.FC<Props> = ({ close, open }) => {
    return (
        <>
            <div className="login-window-container">
                <h1 className="login-window-title">Sign In<img className="login-window-cancel" src={cancel} onClick={() => close()}/></h1>
                <input id="username" type="text" placeholder="User"/>
                <input id="password" type="text" placeholder="Password"/>
                <div className="login-window-create">Create an Account: <span className="login-window-signup" onClick={() => open()}>Sign Up</span></div>
            </div>
        </>
    );
};

export default LoginWindow;