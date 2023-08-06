import "./SignupWindow.css";
import { ChangeEventHandler, KeyboardEventHandler, MouseEventHandler } from "react";

interface Props {
    mode: string,
    close: Function
    error: boolean
    isTakenName: boolean
    isInvalidInput: boolean
    isLoading: boolean
    username: string
    password: string
    userChange: ChangeEventHandler<HTMLInputElement>
    userKeyUp: KeyboardEventHandler<HTMLInputElement>
    passChange: ChangeEventHandler<HTMLInputElement>
    passKeyUp: KeyboardEventHandler<HTMLInputElement>
    signup: MouseEventHandler<HTMLDivElement>
    signupOnEnter: KeyboardEventHandler<HTMLInputElement>
};

const SignupWindow: React.FC<Props> = ({ mode, close, error, isTakenName, isInvalidInput, isLoading, username, password, userKeyUp, userChange, passKeyUp, passChange, signup, signupOnEnter }) => {
    const cancel = `/images/${mode}/cancel.png`
    const loading = `/images/${mode}/loading.png`
    return (
        <>
            <div className="create-window-container">
                <h1 className="login-window-title ">Create a Account<img className="create-window-cancel" src={cancel} onClick={() => close()}/></h1>
                {isTakenName ? <div className="login-invalid">Apologies, but the desired username is currently unavailable. Please consider selecting an alternative username.</div> : null}
                {isInvalidInput ? <div className="login-invalid">Enter a valid username and password.</div> : null}
                <input id="username" value={username} onKeyUp={userKeyUp} onChange={userChange} placeholder="User" required/>
                <input id="password" type="password" value={password} onKeyDown={signupOnEnter} onKeyUp={passKeyUp} onChange={passChange} placeholder="Password" required/>
                <br/>
                {error ? <div className="signup-failed">We apologize, but there seems to be an issue on our end. Please try again later.</div> : null}
                {isLoading ? <img src={loading} className="signup-loading"/> : <div onClick={signup} className="signup-acc-button">Submit</div>}
            </div>
        </>
    );
};

export default SignupWindow;