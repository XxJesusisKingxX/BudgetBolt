import { useContext, useEffect, useState } from 'react';
import { validatePass, validateUser } from '../../utils/validator';
import { AuthType } from '../../constants/auth'
import AuthButton from './components/AuthButton';
import LoginWindow from './components/LoginWindow';
import SignupWindow from './components/SignupWindow';
import AccountWindow from './components/AccountWindow';
import { useLogin } from './useLogin';
import UserContext from '../../context/UserContext';
import LoginContext from '../../context/LoginContext';
import "../../assets/auth/styles/Auth.css";
import ThemeContext from '../../context/ThemeContext';

const Auth = () => {
    const {mode} = useContext(ThemeContext);
    const {isLogin, userDispatch} = useContext(UserContext);
    const {showAccountWindow, showSignUpWindow, showLoginWindow, loginDispatch} = useContext(LoginContext);

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [showUsernameError, setShowUsernameError] = useState(false);
    const [showPasswordError, setShowPasswordError] = useState(false);

    const valid = !showUsernameError && !showPasswordError && username != '' && password != '';
    const {serverErr, nameErr, takenNameErr, authErr, showLoading, clearErrors, handleAuth} = useLogin(username, password, valid);

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
   
    const handleAuthOnEnter = (event: React.KeyboardEvent<HTMLInputElement>, authType: number) => {
        if (event.key === "SET_STATE") {
            handleAuth(authType);
        }
    }
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
    const handleWindowClose = () => {
        setUsername('');
        setPassword('');
        clearErrors();
        loginDispatch({ type: "SET_STATE", state: { showLoginWindow: false, showSignUpWindow: false, showAccountWindow: false }});
    }
    const handleCreateWindow = () => {
        handleWindowClose();
        loginDispatch({ type: "SET_STATE", state: { showSignUpWindow: true }});
    }
    const handleLogin = () => {
        loginDispatch({ type: "SET_STATE", state: { showLoginWindow: true }});
    }
    const handleLogout = () => {
        localStorage.clear();
        userDispatch({ type: "SET_STATE", state: { isLogin: false }});
    }

    return (
        <>
            {!isLogin ? (
            <AuthButton
                action={handleLogin}
                name="Login"
            />
            ) : (isLogin ? (
            <AuthButton
                action={handleLogout}
                name="Logout"
            />
            ) : (
                null
            ))}
            {showLoginWindow ? (
            <LoginWindow
                isInvalidName={showUsernameError}
                isInvalidPass={showPasswordError}
                isAuthError={authErr}
                isNameError={nameErr}
                serverError={serverErr}
                showLoading={showLoading}
                loginOnEnter={(event) => handleAuthOnEnter(event, AuthType.Login)}
                login={() => handleAuth(AuthType.Login)}
                userKeyUp={handleUserKeyUp}
                userChange={handleUserChange}
                username={username}
                passKeyUp={handlePassKeyUp}
                passChange={handlePassChange}
                password={password}
                openSignUp={handleCreateWindow}
                close={handleWindowClose}
                mode={mode}
            />
            ) : (
                null
            )}

            {showSignUpWindow ? (
            <SignupWindow
                isTakenName={takenNameErr}
                isInvalidName={showUsernameError}
                isInvalidPass={showPasswordError}
                showLoading={showLoading}
                signupOnEnter={(event) => handleAuthOnEnter(event, AuthType.SignUp)}
                signup={() => handleAuth(AuthType.SignUp)}
                userKeyUp={handleUserKeyUp}
                userChange={handleUserChange}
                username={username}
                passKeyUp={handlePassKeyUp}
                passChange={handlePassChange}
                password={password}
                close={handleWindowClose}
                serverError={serverErr}
                mode={mode}
            />
            ) : (
                null
            )}

            {showAccountWindow ? (
            <AccountWindow/>
            ) : (
                null
            )}
        </>
    );
};

export default Auth;