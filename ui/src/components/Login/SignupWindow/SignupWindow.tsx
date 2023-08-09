import "./SignupWindow.css";
import { ChangeEventHandler, KeyboardEventHandler, MouseEventHandler } from "react";

interface Props {
    mode: string,
    close: Function
    error: boolean
    isTakenName: boolean
    isInvalidPass: boolean
    isInvalidName: boolean
    showLoading: boolean
    username: string
    password: string
    userChange: ChangeEventHandler<HTMLInputElement>
    userKeyUp: KeyboardEventHandler<HTMLInputElement>
    passChange: ChangeEventHandler<HTMLInputElement>
    passKeyUp: KeyboardEventHandler<HTMLInputElement>
    signup: MouseEventHandler<HTMLDivElement>
    signupOnEnter: KeyboardEventHandler<HTMLInputElement>
};

const SignupWindow: React.FC<Props> = ({ mode, close, error, isTakenName, isInvalidPass, isInvalidName, showLoading, username, password, userKeyUp, userChange, passKeyUp, passChange, signup, signupOnEnter }) => {
    const cancel = `/images/${mode}/cancel.png`
    const loading = `/images/${mode}/loading.png`
    return (
        <>
            <div className="create-window-container">
                <h1 className="login-window-title ">Create a Account<img className="create-window-cancel" src={cancel} onClick={() => close()}/></h1>
                <div id="username">
                    {isTakenName ?
                    <span>
                        Apologies, but the username is already in use.
                        <br/>
                        Please select a different username.
                        <br/>
                    </span> : null}
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
                    <input type="password" value={password} onKeyDown={signupOnEnter} onKeyUp={passKeyUp} onChange={passChange} name="pass" placeholder="Password" required/>
                    {isInvalidPass ?
                    <span>
                        Oops! The username or password is incorrect.
                        <br/>
                        Please double check and try again.
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
                    {error ?
                    <span>
                        We apologize, but there seems to be an issue on our end. Please try again later.
                    </span> : null}
                </div>
                {showLoading ? <img src={loading} className="signup-loading"/> : <div onClick={signup} className="signup-acc-button">Submit</div>}
            </div>
        </>
    );
};

export default SignupWindow;