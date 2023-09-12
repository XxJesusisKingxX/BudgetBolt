import { useContext, useEffect, useState } from 'react';
import { validatePass, validateUser } from '../../utils/validator';
import { AuthType } from '../../constants/auth'
import LoginWindow from './components/LoginWindow';
import SignupWindow from './components/SignupWindow';
import AccountWindow from './components/AccountWindow';
import { useLogin } from './useLogin';
import LoginContext from '../../context/LoginContext';
import "../../assets/auth/styles/Auth.css";
import ThemeContext from '../../context/ThemeContext';
import AppContext from '../../context/AppContext';

const Auth = () => {
    // Get theme mode from context
    const { mode } = useContext(ThemeContext);
    const { dispatch } = useContext(AppContext);

    // Get login-related states from contexts
    const { showAccountWindow, showSignUpWindow, showLoginWindow, isLogin, loginDispatch } = useContext(LoginContext);

    // State for username and password inputs
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    // States to handle showing errors
    const [showUsernameError, setShowUsernameError] = useState(false);
    const [showPasswordError, setShowPasswordError] = useState(false);

    // Calculate if the form is valid
    const valid = !showUsernameError && !showPasswordError && username !== '' && password !== '';

    // Use custom hook for login-related functionality
    const { serverErr, nameErr, takenNameErr, authErr, showLoading, clearErrors, handleAuth } = useLogin(username, password, valid);

    // Validate username and password on input change
    useEffect(() => {
        const validate = () => {
            if (!validateUser(username) && username) {
                setShowUsernameError(true);
            } else {
                setShowUsernameError(false);
            }
            if (!validatePass(password) && password) {
                setShowPasswordError(true);
            } else {
                setShowPasswordError(false);
            }
        };
        validate();
    }, [username, password]);

    // Handle authentication on Enter key press
    const handleAuthOnEnter = (event: React.KeyboardEvent<HTMLInputElement>, authType: number) => {
        if (event.key === "Enter") {  // Changed "SET_STATE" to "Enter"
            handleAuth(authType);
        }
    }

    // Event handlers for input changes
    const handleUserChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setUsername(event.target.value);
    }
    const handleUserKeyUp = (event: React.KeyboardEvent<HTMLInputElement>) => {
        setUsername(event.currentTarget.value);
    }
    const handlePassChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
    }
    const handlePassKeyUp = (event: React.KeyboardEvent<HTMLInputElement>) => {
        setPassword(event.currentTarget.value);
    }

    // Close authentication windows
    const handleWindowClose = () => {
        setUsername('');
        setPassword('');
        clearErrors();
        loginDispatch({ type: "SET_STATE", state: { showLoginWindow: false, showSignUpWindow: false, showAccountWindow: false }});
    }

    // Show signup window
    const handleCreateWindow = () => {
        handleWindowClose();
        loginDispatch({ type: "SET_STATE", state: { showSignUpWindow: true }});
    }

    // Show login window
    const handleLogin = () => {
        loginDispatch({ type: "SET_STATE", state: { showLoginWindow: true }});
    }

    // Handle logout
    const handleLogout = () => {
        localStorage.clear();
        dispatch({ type: "SET_STATE", state: { profile: '' }});
        loginDispatch({ type: "SET_STATE", state: { isLogin: false }});
        handleWindowClose();
    }

    return (
        <>
            {/* Render login or logout button based on login status */}
            {!isLogin ? (
                <button data-testid='auth-button' className="btn btn--auth" onClick={() => handleLogin()}>Login</button>
            ) : (
                <button data-testid='auth-button' className="btn btn--auth" onClick={() => handleLogout()}>Logout</button>
            )}
            {/* Render login window */}
            {showLoginWindow ? (
                <LoginWindow
                    // Props related to errors and input values
                    isInvalidName={showUsernameError}
                    isInvalidPass={showPasswordError}
                    isAuthError={authErr}
                    isNameError={nameErr}
                    serverError={serverErr}
                    showLoading={showLoading}
                    username={username}
                    password={password}
                    // Event handlers
                    loginOnEnter={(event) => handleAuthOnEnter(event, AuthType.Login)}
                    login={() => handleAuth(AuthType.Login)}
                    userKeyUp={handleUserKeyUp}
                    userChange={handleUserChange}
                    passKeyUp={handlePassKeyUp}
                    passChange={handlePassChange}
                    openSignUp={handleCreateWindow}
                    close={handleWindowClose}
                    mode={mode}
                />
            ) : (
                null
            )}
            {/* Render signup window */}
            {showSignUpWindow ? (
                <SignupWindow
                    // Props related to errors and input values
                    isTakenName={takenNameErr}
                    isInvalidName={showUsernameError}
                    isInvalidPass={showPasswordError}
                    showLoading={showLoading}
                    serverError={serverErr}
                    username={username}
                    password={password}
                    // Event handlers
                    signupOnEnter={(event) => handleAuthOnEnter(event, AuthType.SignUp)}
                    signup={() => handleAuth(AuthType.SignUp)}
                    userKeyUp={handleUserKeyUp}
                    userChange={handleUserChange}
                    passKeyUp={handlePassKeyUp}
                    passChange={handlePassChange}
                    close={handleWindowClose}
                    mode={mode}
                />
            ) : (
                null
            )}
            {/* Render account window */}
            {showAccountWindow ? (
                <AccountWindow/>
            ) : (
                null
            )}
        </>
    );
};

export default Auth;
