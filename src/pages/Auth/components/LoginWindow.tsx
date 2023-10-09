import { ChangeEventHandler, KeyboardEventHandler } from 'react';
import '../../../assets/Button.css'
import '../../../assets/Auth.css'
import '../../../assets/Error.css'
import '../../../assets/Loading.css'

interface Props {
    mode: string;                                         // The current mode of the login window (e.g., 'light' or 'dark')
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
        <div data-testid='login-window' className='window'>
            <span className='window__title'>Sign In<img className='closeicon' src={cancel} onClick={() => close()} alt='Close' /></span>
            {/* Username input */}
            <input aria-label='username' className='window__form' value={username} onKeyUp={userKeyUp} onChange={userChange} placeholder='Username' disabled={showLoading} required />
            {/* Display error if the username is invalid */}
            {isInvalidName ?
                <div data-testid='invalid-name' className='err err--login-signup'>
                    Enter a valid username:
                    <br />
                    - Must start with a letter or an underscore (_)
                    <br />
                    - Valid characters: 0-9, A-z, (_)
                </div>
                :
                null
            }
            {/* Password input */}
            <input aria-label='password' className='window__form' type='password' value={password} onKeyDown={loginOnEnter} onKeyUp={passKeyUp} onChange={passChange} placeholder='Password' disabled={showLoading} required />
            {/* Display error if the password is invalid */}
            {isInvalidPass ?
                <div data-testid='invalid-pass' className='err err--login-signup'>
                    Password is invalid.
                    <br />
                    - At least one lowercase letter
                    <br />
                    - At least one uppercase letter
                    <br />
                    - At least one digit
                    <br />
                    - At least one special character (!@#$%^&*)
                    <br />
                    - Minimum 8 characters
                </div>
                :
                null
            }
            {/* Link to open the sign-up window */}
            <span className='window__subtitle'>Create an Account</span><span data-testid='signup-link' className={`window__link ${showLoading ? 'window__link--disable' : ''}`} onClick={() => openSignUp()}>Sign Up</span>
            {/* Display authentication error */}
            {isAuthError ?
                <div data-testid='auth-err' className='err err--login-signup'>Oops! The username or password is incorrect.</div>
                :
                null
            }
            {/* Display username error */}
            {isNameError ?
                <div data-testid='name-err' className='err err--login-signup'>Oops! The username does not exist</div>
                :
                null
            }
            {/* Display server error */}
            {serverError ?
                <div data-testid='server-err' className='err err--login-signup'>
                    We apologize, but there seems to be an issue.
                    <br />
                    Please try again later.
                </div>
                :
                null
            }
            {/* Display loading or login button */}
            <div className='window__button'>
                {showLoading ?
                    <img data-testid='login-loading' src={loading} className='loading loading--small' alt='Loading' />
                    :
                    <button data-testid='login-button' onClick={() => login()} className='btn btn--login'>Login</button>
                }
            </div>
        </div>
    );
};

export default LoginWindow;