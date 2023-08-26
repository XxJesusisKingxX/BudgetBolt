import { ChangeEventHandler, KeyboardEventHandler } from "react";

interface Props {
    mode: string
    close: Function
    open: Function
    login: Function
    isAuthError: boolean
    isInvalidPass: boolean
    isInvalidName: boolean
    isNameError: boolean
    showLoading: boolean
    username: string
    password: string
    userChange: ChangeEventHandler<HTMLInputElement>
    userKeyUp: KeyboardEventHandler<HTMLInputElement>
    passChange: ChangeEventHandler<HTMLInputElement>
    passKeyUp: KeyboardEventHandler<HTMLInputElement>
    loginOnEnter: KeyboardEventHandler<HTMLInputElement>
}
const LoginWindow: React.FC<Props> = ({ mode, close, open, isNameError, isAuthError, isInvalidPass, isInvalidName, showLoading, username, password,userKeyUp, userChange, passKeyUp, passChange, login, loginOnEnter}) => {
    const cancel = `/images/${mode}/cancel.png`
    const loading = `/images/${mode}/loading.png`
    return (
        <div className="windowcont windowcont--auth">
            <h1 className="windowcont__title">Sign In<img className="closeicon" src={cancel} onClick={() => close()}/></h1>
            <div className="auth auth--username">
                <input className="auth__input auth__input--roundedinsde" value={username} onKeyUp={userKeyUp} onChange={userChange} placeholder="User" required/>
                {isInvalidName ?
                    <div className="err err--usernameinvalid">
                        Enter a valid username:
                        <br/>
                        - Must start with a letter or an underscore (_)
                        <br/>
                        - Valid characters: 0-9, A-z, (_)
                        <br/>
                    </div>
                    :
                    null
                }
            </div>
            <div className="auth auth--password">
                <input className="auth__input auth__input--roundedinsde" type="password" value={password} onKeyDown={loginOnEnter} onKeyUp={passKeyUp} onChange={passChange} placeholder="Password" required/>
                {isInvalidPass ?
                    <div className="err err--passwordinvalid">
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
                    </div>
                    :
                    null
                }
            </div>
            Create an Account:<span className="link link--signup" onClick={() => open()}>Sign Up</span>
            {isAuthError ?
                <div className="err err--passwordinvalid">Oops! The username or password is incorrect.</div>
                :
                null
            }
            {isNameError ?
                <div className="err err--passwordinvalid">Oops! The username does not exist</div>
                :
                null
            }
            {showLoading ?
                <img src={loading} className="loading loading--login"/>
                :
                <button onClick={() =>login()} className="btn btn--login">Login</button>
            }
        </div>
    );
};

export default LoginWindow;