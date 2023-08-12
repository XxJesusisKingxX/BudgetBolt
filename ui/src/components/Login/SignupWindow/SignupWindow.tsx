import { ChangeEventHandler, KeyboardEventHandler } from "react";

interface Props {
    mode: string,
    close: Function
    signup: Function
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
    signupOnEnter: KeyboardEventHandler<HTMLInputElement>
};

const SignupWindow: React.FC<Props> = ({ mode, close, error, isTakenName, isInvalidPass, isInvalidName, showLoading, username, password, userKeyUp, userChange, passKeyUp, passChange, signup, signupOnEnter }) => {
    const cancel = `/images/${mode}/cancel.png`
    const loading = `/images/${mode}/loading.png`
    return (
        <div className="windowcont windowcont--auth">
            <h1 className="windowcont__title">Create a Account<img className="closeicon" src={cancel} onClick={() => close()}/></h1>
            <div className="auth auth--username">
                {isTakenName ?
                    <div className="err err--usernameinvalid">
                        Apologies, but the username is already in use.
                        <br/>
                        Please select a different username.
                        <br/>
                    </div>
                    :
                    null
                }
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
                <input className="auth__input auth__input--roundedinsde" type="password" value={password} onKeyDown={signupOnEnter} onKeyUp={passKeyUp} onChange={passChange} placeholder="Password" required/>
                {isInvalidPass ?
                    <div className="err err--passwordinvalid">
                        Enter a valid password:
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
                {error ?
                    <div className="err err--servererr">
                        We apologize, but there seems to be an issue.
                        <br/>
                        Please try again later.
                    </div>
                    :
                    null
                }
            </div>
            {showLoading ?
                <img src={loading} className="loading loading--create"/>
                :
                <button onClick={() => signup()} className="btn btn--create">Submit</button>}
        </div>
    );
};

export default SignupWindow;