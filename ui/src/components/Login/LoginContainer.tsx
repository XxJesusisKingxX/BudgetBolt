import { useContext, useState } from "react";
import Login from "./Login";
import Context from "../../context/Context";
import LoginWindow from "./LoginWindow/LoginWindow";
import SignupWindow from "./SignupWindow/SignupWindow";

const LoginContainer = () => {
    const { isLogin, dispatch } = useContext(Context);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [showCreateWindow, setShowCreateWindow] = useState(false);
    const [showLoginWindow, setShowLoginWindow] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [showLoginError, setShowLoginError] = useState(false);
    const [showSignUpError, setShowSignUpError] = useState(false);

    const showLogin = (show: boolean) => {
        if (show) {
            dispatch({ type:"SET_STATE", state:{ isLogin: false }});
        } else {
            dispatch({ type:"SET_STATE", state:{ isLogin: true }});
        }
    }
    const handleSignUp = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        setIsLoading(true);
        try {
            const response = await fetch("/api/profile/create", {
                method: "POST",
                body: new URLSearchParams ({
                    username: username,
                    password: password
                })
            });
            if (response.status == 200) {
                setIsLoading(false);
                setShowCreateWindow(false);
            }
        } catch (error) {
            setIsLoading(false);
            console.error("Error fetching data:", error);
        }
    }

    const handleLogin = async () => {
        setShowLoginError(false);
        setIsLoading(true);
        try {
            const response = await fetch("/api/profile/get", {
                method: "POST",
                body: new URLSearchParams ({
                    username: username,
                    password: password
                })
            });
            if (response.status == 200) {
                dispatch({ type:"SET_STATE", state:{ isLogin: true }});
                setIsLoading(false);
                setShowLoginWindow(false);
                let data = await response.json()
                console.log(data["id"])
            } else {
                setIsLoading(false);
                setShowLoginError(true);
            }
        } catch (error) {
            setIsLoading(false);
            setShowLoginError(true);
        }
    }
    const handleUserInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setUsername(event.target.value);
        let maxChar = 25
        const validStart = new RegExp("^[a-zA-Z_][a-zA-Z0-9_]$"); // make sure starts with _ or alpha minimum and the follwing can be numebr, alpha , or underscore
        const isUnder = username.length <= maxChar ? true : false;
        if (validStart.test(username) && isUnder) {
            console.log("valid")
        } else {
            console.log("invalid")
        }
    }
    const handlePassInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
        //TODO FIX
        // const complexityRegex = new RegExp("^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$")
        // if (complexityRegex.test(password)) {
        //     console.log("valid")
        // } else {
        //     console.log("invalid")
        // }
    }
    const clearFields = () => {
        setUsername("");
        setPassword("");
    }
    return (
        <>
            {!isLogin ? (
            <Login 
                open={() => {
                    setShowLoginWindow(true);
                    showLogin(false);
                }}
                
            />
            ) : (
                null
            )}
            {showLoginWindow ? (
            <LoginWindow
                isLoading={isLoading}
                login={handleLogin}
                userChange={handleUserInput}
                username={username}
                passChange={handlePassInput}
                password={password}
                open={() => {
                    setShowLoginWindow(false);
                    setShowCreateWindow(true);
                }}
                close={() => {
                    setShowLoginWindow(false);
                    setShowCreateWindow(false);
                    clearFields();
                    showLogin(true);
                }}
                error={showLoginError}
            />
            ) : (
                null
            )}
            {showCreateWindow ? (
            <SignupWindow
                isLoading={isLoading}
                signup={handleSignUp}
                userChange={handleUserInput}
                username={username}
                passChange={handlePassInput}
                password={password}
                close={() => {
                    setShowCreateWindow(false);
                    clearFields();
                    showLogin(true);
                }}
                error={showSignUpError}
            />
            ) : (
                null
            )}
        </>
    );
};

export default LoginContainer;