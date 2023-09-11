import { ChangeEventHandler, KeyboardEventHandler } from "react";

interface Props {
    mode: string;                                         // The current mode of the login window (e.g., "light" or "dark")
    close: Function;                                      // Function to close the login window
    openSignUp: Function;                                 // Function to open the sign-up window
    login: Function;                                      // Function to handle the login process
    serverError: boolean;                                 // Indicates if a server error has occurred
    isAuthError: boolean;                                 // Indicates if there's an authentication error (wrong username/password)
    isInvalidPass: boolean;                               // Indicates if the entered password is invalid
    isInvalidName: boolean;                               // Indicates if the entered username is invalid
    isNameError: boolean;                                 // Indicates if the username does not exist (for authentication)
    showLoading: boolean;                                 // Indicates if loading state should be shown
    username: string;                                     // The current username input value
    password: string;                                     // The current password input value
    userChange: ChangeEventHandler<HTMLInputElement>;     // Event handler for changes in the username input
    userKeyUp: KeyboardEventHandler<HTMLInputElement>;    // Event handler for key up events in the username input
    passChange: ChangeEventHandler<HTMLInputElement>;     // Event handler for changes in the password input
    passKeyUp: KeyboardEventHandler<HTMLInputElement>;    // Event handler for key up events in the password input
    loginOnEnter: KeyboardEventHandler<HTMLInputElement>; // Event handler for triggering login on Enter key press
}

const LoginWindow: React.FC<Props> = ({ mode, close, openSignUp, serverError, isNameError, isAuthError, isInvalidPass, isInvalidName, showLoading, username, password, userKeyUp, userChange, passKeyUp, passChange, login, loginOnEnter }) => {
    const cancel = `/images/${mode}/cancel.png`;
    const loading = `/images/${mode}/loading.png`;

    return (
        <div data-testid='login-window' className="windowcont windowcont--auth">
            <h1 className="windowcont__title">Sign In<img className="closeicon" src={cancel} onClick={() => close()} alt="Close" /></h1>
            <div className="auth auth--username">
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
                <input aria-label="password" className="auth__input auth__input--roundedinsde" type="password" value={password} onKeyDown={loginOnEnter} onKeyUp={passKeyUp} onChange={passChange} placeholder="Password" required />
                {/* Display error if the password is invalid */}
                {isInvalidPass ?
                    <p data-testid='invalid-pass' className="err err--passwordinvalid">
                        Password is invalid.
                        {/* Password requirements */}
                    </p>
                    :
                    null
                }
            </div>
            {/* Link to open the sign-up window */}
            Create an Account:<span data-testid='signup-link' className="link link--signup" onClick={() => openSignUp()}>Sign Up</span>
            {/* Display authentication error */}
            {isAuthError ?
                <div data-testid='auth-err' className="err err--passwordinvalid">Oops! The username or password is incorrect.</div>
                :
                null
            }
            {/* Display username error */}
            {isNameError ?
                <div data-testid='name-err' className="err err--passwordinvalid">Oops! The username does not exist</div>
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
            {/* Display loading or login button */}
            {showLoading ?
                <img data-testid='login-loading' src={loading} className="loading loading--login" alt="Loading" />
                :
                <button data-testid='login-button' onClick={() => login()} className="btn btn--login">Login</button>
            }
        </div>
    );
};

export default LoginWindow;