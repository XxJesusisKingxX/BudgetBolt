import { useContext, useState } from "react";
import Login from "./Login";
import Context from "../../context/Context";
import LoginWindow from "./LoginWindow/LoginWindow";
import SignupWindow from "./SignupWindow/SignupWindow";

const LoginContainer = () => {
    const { isLogin, profile, dispatch } = useContext(Context);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [showLogin, setShowLogin] = useState(true);
    const [showCreateWindow, setShowCreateWindow] = useState(false);
    const [showLoginWindow, setShowLoginWindow] = useState(false);
    const [showLoading, setShowLoading] = useState(false);
    const [showValidationError, setShowValidationError] = useState(false);
    const [showServerError, setShowServerError] = useState(false);

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
                    if (response.status == 200) {
                        clearFields();
                        dispatch({ type:"SET_STATE", state:{ isLogin: true }});
                        dispatch({ type:"SET_STATE", state:{ profile: username.toLocaleUpperCase() }});
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
    }
    return (
        <>
            {!isLogin ? (
            <Login
                open={() => {
                    setShowLoginWindow(true);
                    setShowLogin(false);
                }}
            />
            ) : (
                null
            )}
            {showLoginWindow ? (
            <LoginWindow
                isInvalidInput={showValidationError}
                isLoading={showLoading}
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
            />
            ) : (
                null
            )}
            {showCreateWindow ? (
            <SignupWindow
                isInvalidInput={showValidationError}
                isLoading={showLoading}
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
            />
            ) : (
                null
            )}
        </>
    );
};

export default LoginContainer;