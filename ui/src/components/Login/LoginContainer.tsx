import { useContext, useState } from "react";
import Login from "./Login";
import Logout from "./Logout";
import Context from "../../context/Context";
import LoginWindow from "./LoginWindow/LoginWindow";
import SignupWindow from "./SignupWindow/SignupWindow";
import AccountWindow from "./AccountWindow/AccountWindow";

const LoginContainer = () => {
    const { isLogin, mode, dispatch } = useContext(Context);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [showLogin, setShowLogin] = useState(true);
    const [showCreateWindow, setShowCreateWindow] = useState(false);
    const [showLoginWindow, setShowLoginWindow] = useState(false);
    const [showLoading, setShowLoading] = useState(false);
    const [showValidationError, setShowValidationError] = useState(false);
    const [showNameError, setShowNameError] = useState(false);
    const [showServerError, setShowServerError] = useState(false);
    const [showAccountWindow, setShowAccountWindow] = useState(false);

    const AuthType = {
        SignUp: 1,
        Login: 2,
    };
    
    const handleAuthentication = (authType: number) => {
        if (!validateUser(username) || !validatePass(password)) {
            setShowServerError(false)
            setShowValidationError(true);
        } else {
            setShowValidationError(false);
            setShowServerError(false);
            setShowLoading(true);
            const startAuth = async () => {
                try {
                    const response = await fetch(authType == AuthType.Login
                    ? "/api/profile/get"
                    : "/api/profile/create", {
                        method: "POST",
                        body: new URLSearchParams ({
                            username: username,
                            password: password
                        })
                    });
                    if (response.ok) {
                        dispatch({ type:"SET_STATE", state:{ profile: username.toLocaleLowerCase() }});
                        if (authType == AuthType.SignUp) {
                            setShowAccountWindow(true)
                            setShowLogin(false);
                        } else if (authType == AuthType.Login) {
                            dispatch({ type:"SET_STATE", state:{ isLogin: true }});
                            clearFields();
                        }
                    } else if (response.status == 409){
                        setShowLoading(false);
                        setShowNameError(true);
                    } else {
                        setShowLoading(false);
                        setShowServerError(true);
                    }
                } catch (error) {
                    setShowLoading(false);
                    setShowServerError(true);
                }
            }
            startAuth()
        }
    }
    const validateUser = (username: string) => {
        let maxChar = 25
        const validStart = new RegExp(`^_?[a-zA-Z][a-zA-Z0-9_]{1,${maxChar}}$`); // make sure starts with _ (if _ must have letter follow) or letter minimum and the following can be a number, letter , or underscore
        const isUnder = username.length <= maxChar ? true : false;
        if (validStart.test(username) && isUnder) {
            return true
        } else {
            return false
        }
    }
    const validatePass = (password: string) => {
        let maxChar = 8
        const complexityRegex = new RegExp(`^(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*]).{${maxChar},}$`)
        if (complexityRegex.test(password)) {
            return true
        } else {
            return false
        }
    }
    const handleAuthOnEnter = (event: React.KeyboardEvent<HTMLInputElement>, authType: number) => {
        if (event.key === "Enter") {
            handleAuthentication(authType)
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
    const clearFields = () => {
        setUsername("");
        setPassword("");
        setShowLoading(false)
        setShowServerError(false)
        setShowValidationError(false)
        setShowLoginWindow(false);
        setShowCreateWindow(false);
        setShowAccountWindow(false);
        setShowNameError(false);
    }
    const logout = () => {
        clearFields();
        localStorage.clear();
        dispatch({ type:"SET_STATE", state:{ isLogin: false }});
    }
    const login = () => {
        clearFields();
        setShowLoginWindow(true);
        setShowLogin(true);
    }
    return (
        <>
            {!isLogin && !showLoginWindow ? (
            <Login
                login={() => {
                    login();
                }}
            />
            ) : (isLogin ? (
            <Logout
                logout={() => {
                    logout();
                }}
            />
            ) : (
                null
            ))}

            {showLoginWindow ? (
            <LoginWindow
                isInvalidInput={showValidationError}
                isLoading={showLoading}
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
                error={showServerError}
                mode={mode}
            />
            ) : (
                null
            )}

            {showCreateWindow ? (
            <SignupWindow
                isInvalidInput={showValidationError}
                isLoading={showLoading}
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
                isTakenName={showNameError}
                error={showServerError}
                mode={mode}
            />
            ) : (
                null
            )}

            {!isLogin && showAccountWindow ? (
            <AccountWindow
                close={() => {
                    clearFields();
                }}
                mode={mode}
            />
            ) : (
                null
            )}
        </>
    );
};

export default LoginContainer;

// Password:
//     At least one digit (0-9).
//     At least one lowercase letter (a-z).
//     At least one uppercase letter (A-Z).
//     At least one special character from the set: !@#$%^&*.
//     At least 8 characters long.