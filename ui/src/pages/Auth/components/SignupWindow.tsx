import { ChangeEventHandler, KeyboardEventHandler } from "react";

interface Props {
    mode: string; // The current mode of the login window (e.g., "light" or "dark")
    close: Function; // Function to close the login window
    signup: Function; // Function to open the sign-up window
    serverError: boolean; // Indicates if a server error has occurred
    isTakenName: boolean; // Indicates if the chosen username is already taken
    isInvalidPass: boolean; // Indicates if the entered password is invalid
    isInvalidName: boolean; // Indicates if the entered username is invalid
    showLoading: boolean; // Indicates if loading state should be shown
    username: string; // The current username input value
    password: string; // The current password input value
    userChange: ChangeEventHandler<HTMLInputElement>; // Event handler for changes in the username input
    userKeyUp: KeyboardEventHandler<HTMLInputElement>; // Event handler for key up events in the username input
    passChange: ChangeEventHandler<HTMLInputElement>; // Event handler for changes in the password input
    passKeyUp: KeyboardEventHandler<HTMLInputElement>; // Event handler for key up events in the password input
    signupOnEnter: KeyboardEventHandler<HTMLInputElement>; // Event handler for triggering signup on Enter key press
};

const SignupWindow: React.FC<Props> = ({ mode, close, serverError, isTakenName, isInvalidPass, isInvalidName, showLoading, username, password, userKeyUp, userChange, passKeyUp, passChange, signup, signupOnEnter }) => {
    const cancel = `/images/${mode}/cancel.png`;
    const loading = `/images/${mode}/loading.png`;

    return (
        <div data-testid='signup-window' className="windowcont windowcont--auth">
            <h1 className="windowcont__title">Create an Account<img className="closeicon" src={cancel} onClick={() => close()} alt="Close" /></h1>
            <div className="auth auth--username">
                {/* Display error if the username is already taken */}
                {isTakenName ?
                    <div data-testid='taken-err' className="err err--usernameinvalid">
                        Apologies, but the username is already in use.
                        <br />
                        Please select a different username.
                    </div>
                    :
                    null
                }
                {/* Username input */}
                <input aria-label="username" className="auth__input auth__input--roundedinsde" value={username} onKeyUp={userKeyUp} onChange={userChange} placeholder="Username" required />
                {/* Display error if the username is invalid */}
                {isInvalidName ?
                    <div data-testid='invalid-name' className="err err--usernameinvalid">
                        Enter a valid username:
                        <br />
                        - Must start with a letter or an underscore (_)
                        <br />
                        - Valid characters: 0-9, A-z, (_)
                    </div>
                    :
                    null
                }
            </div>
            <div className="auth auth--password">
                {/* Password input */}
                <input aria-label="password" className="auth__input auth__input--roundedinsde" type="password" value={password} onKeyDown={signupOnEnter} onKeyUp={passKeyUp} onChange={passChange} placeholder="Password" required />
                {/* Display error if the password is invalid */}
                {isInvalidPass ?
                    <div data-testid='invalid-pass' className="err err--passwordinvalid">
                        Enter a valid password:
                        <br />
                        Must have the following:
                        {/* Password requirements */}
                    </div>
                    :
                    null
                }
                {/* Display server error */}
                {serverError ?
                    <div data-testid='server-err' className="err err--servererr">
                        We apologize, but there seems to be an issue.
                        <br />
                        Please try again later.
                    </div>
                    :
                    null
                }
            </div>
            {/* Display loading or signup button */}
            {showLoading ?
                <img data-testid='signup-loading' src={loading} className="loading loading--create" alt="Loading" />
                :
                <button data-testid='signup-button' onClick={() => signup()} className="btn btn--create">Submit</button>}
        </div>
    );
};

export default SignupWindow;