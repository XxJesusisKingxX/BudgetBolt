import { useContext, useEffect, useState } from "react";
import { validatePass, validateUser } from "../../utils/validator";
import { AuthType } from "../../constants/auth"
import Auth from "./Auth";
import Context from "../../context/Context";
import LoginWindow from "./LoginWindow/LoginWindow";
import SignupWindow from "./SignupWindow/SignupWindow";
import AccountWindow from "./AccountWindow";
import { useLogin } from "./useLogin";

const AuthContainer = () => {
    const {profile, isLogin, mode, dispatch} = useContext(Context);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [showUsernameError, setShowUsernameError] = useState(false);
    const [showPasswordError, setShowPasswordError] = useState(false);
    const valid = !showUsernameError && !showPasswordError && username != "" && password != ""
    const {serverErr, nameErr, takenNameErr, authErr, showLoading, showAccountWindow, login} = useLogin(username, password, valid);
    const [showCreateWindow, setShowCreateWindow] = useState(false);
    const [showLoginWindow, setShowLoginWindow] = useState(false);

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
        if (event.key === "Enter") {
            login(authType);
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
        setUsername("");
        setPassword("");
        setShowLoginWindow(false);
        setShowCreateWindow(false);
    }
    const handleLoginWindow = () => {
        handleWindowClose();
        setShowLoginWindow(false);
        setShowCreateWindow(true);
    }

    return (
        <>
            {!isLogin && !showLoginWindow ? (
            <Auth
                authType={() => {
                    setShowLoginWindow(true);
                }}
                authName="Login"
            />
            ) : (isLogin ? (
            <Auth
                authType={() => {
                    localStorage.clear();
                    dispatch({ type: "SET_STATE", state: { isLogin: false }});
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
                isAuthError={authErr}
                isNameError={nameErr}
                showLoading={showLoading}
                loginOnEnter={(event) => handleAuthOnEnter(event, AuthType.Login)}
                login={() => login(AuthType.Login)}
                userKeyUp={handleUserKeyUp}
                userChange={handleUserChange}
                username={username}
                passKeyUp={handlePassKeyUp}
                passChange={handlePassChange}
                password={password}
                open={handleLoginWindow}
                close={handleWindowClose}
                mode={mode}
            />
            ) : (
                null
            )}

            {!showAccountWindow && showCreateWindow ? (
            <SignupWindow
                isTakenName={takenNameErr}
                isInvalidName={showUsernameError}
                isInvalidPass={showPasswordError}
                showLoading={showLoading}
                signupOnEnter={(event) => handleAuthOnEnter(event, AuthType.SignUp)}
                signup={() => login(AuthType.SignUp)}
                userKeyUp={handleUserKeyUp}
                userChange={handleUserChange}
                username={username}
                passKeyUp={handlePassKeyUp}
                passChange={handlePassChange}
                password={password}
                close={handleWindowClose}
                error={serverErr}
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