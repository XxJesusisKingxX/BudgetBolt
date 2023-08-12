import { useContext, useEffect, useState } from "react";
import { validatePass, validateUser } from "../../utils/validator";
import { AuthType } from "../../enums/auth"
import Auth from "./Auth";
import Context from "../../context/Context";
import LoginWindow from "./LoginWindow/LoginWindow";
import SignupWindow from "./SignupWindow/SignupWindow";
import AccountWindow from "./AccountWindow/AccountWindow";
import { useAppStateActions } from "../../redux/redux";
import { EndPoint } from "../../enums/endpoints";

const AuthContainer = () => {
    const { isLoading, isLogin, mode } = useContext(Context);
    const { setLoadingState, setLoginState, setProfileState } = useAppStateActions();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [showLogin, setShowLogin] = useState(true);
    const [showCreateWindow, setShowCreateWindow] = useState(false);
    const [showLoginWindow, setShowLoginWindow] = useState(false);
    const [showUsernameError, setShowUsernameError] = useState(false);
    const [showPasswordError, setShowPasswordError] = useState(false);
    const [showNameError, setShowNameError] = useState(false);
    const [showServerError, setShowServerError] = useState(false);
    const [showAuthError, setShowAuthError] = useState(false);
    const [showAccountWindow, setShowAccountWindow] = useState(false);

    useEffect(() => {
        const validate = () => {
            if (!validateUser(username)) {
                setShowUsernameError(true);
            } else {
                setShowUsernameError(false);
            }
            if (!validatePass(password)) {
                setShowPasswordError(true);
            } else {
                setShowPasswordError(false);
            }
        };
        validate();
    }, [username, password]);
    // ********* Handlers *************
    const handleAuthentication = (authType: number) => {
        if (!showUsernameError && !showPasswordError && username != "" && password != "") {
            setLoadingState(true);
            const startAuth = async () => {
                try {
                    const response = await fetch(authType == AuthType.Login
                    ? EndPoint.GET_PROFILE
                    : EndPoint.CREATE_PROFILE, {
                        method: "POST",
                        body: new URLSearchParams ({
                            username: username,
                            password: password
                        })
                    });
                    if (response.ok) {
                        setProfileState(username.toLocaleLowerCase())
                        if (authType == AuthType.SignUp) {
                            setShowAccountWindow(true);
                            setShowLogin(false);
                        } else if (authType == AuthType.Login) {
                            setLoginState(true);
                            clearFields();
                        }
                    } else if (response.status == 409) {
                        setLoadingState(false);
                        setShowNameError(true);
                    } else if (response.status == 401) {
                        setLoadingState(false);
                        setShowAuthError(true);
                    } else if (response.status == 404) {
                        setLoadingState(false);
                        setShowNameError(true);
                    } else {
                        setLoadingState(false);
                        setShowServerError(true);
                    }
                } catch (error) {
                    setLoadingState(false);
                    setShowServerError(true);
                }
            }
            startAuth()
        }
    };
    const handleAuthOnEnter = (event: React.KeyboardEvent<HTMLInputElement>, authType: number) => {
        if (event.key === "Enter") {
            handleAuthentication(authType);
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
    // ********* Utils *************
    const clearFields = () => {
        setUsername("");
        setPassword("");
        setLoadingState(false);
        setShowServerError(false);
        setShowUsernameError(false);
        setShowPasswordError(false);
        setShowAuthError(false);
        setShowLoginWindow(false);
        setShowCreateWindow(false);
        setShowAccountWindow(false);
        setShowNameError(false);
    }
    const logout = () => {
        clearFields();
        localStorage.clear();
        setLoginState(false);
    }
    const login = () => {
        clearFields();
        setShowLoginWindow(true);
    }
    return (
        <>
            {!isLogin && !showLoginWindow ? (
            <Auth
                authType={() => {
                    login();
                }}
                authName="Login"
            />
            ) : (isLogin ? (
            <Auth
                authType={() => {
                    logout();
                }}
                authName="Logout"
            />
            ) : (
                null
            ))}

            {showLoginWindow ? (
            <LoginWindow
                isInvalidName={showUsernameError}
                isInvalidPass={showPasswordError}
                isAuthError={showAuthError}
                isNameError={showNameError}
                showLoading={isLoading}
                loginOnEnter={(event) => handleAuthOnEnter(event, AuthType.Login)}
                login={() => handleAuthentication(AuthType.Login)}
                userKeyUp={handleUserKeyUp}
                userChange={handleUserChange}
                username={username}
                passKeyUp={handlePassKeyUp}
                passChange={handlePassChange}
                password={password}
                open={() => {
                    clearFields();
                    setShowLoginWindow(false);
                    setShowCreateWindow(true);
                }}
                close={() => {
                    setShowLogin(true);
                    clearFields();
                }}
                mode={mode}
            />
            ) : (
                null
            )}

            {!showAccountWindow && showCreateWindow ? (
            <SignupWindow
                isTakenName={showNameError}
                isInvalidName={showUsernameError}
                isInvalidPass={showPasswordError}
                showLoading={isLoading}
                signupOnEnter={(event) => handleAuthOnEnter(event, AuthType.SignUp)}
                signup={() => handleAuthentication(AuthType.SignUp)}
                userKeyUp={handleUserKeyUp}
                userChange={handleUserChange}
                username={username}
                passKeyUp={handlePassKeyUp}
                passChange={handlePassChange}
                password={password}
                close={() => {
                    setShowLogin(true);
                    clearFields();
                }}
                error={showServerError}
                mode={mode}
            />
            ) : (
                null
            )}

            {!isLogin && showAccountWindow ? (
            <AccountWindow/>
            ) : (
                null
            )}
        </>
    );
};

export default AuthContainer;