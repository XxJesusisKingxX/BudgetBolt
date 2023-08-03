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
                console.log("success")
            }
        } catch (error) {
            setIsLoading(false);
            console.error("Error fetching data:", error);
        }
    }
    const handleLogin = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
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
            }
        } catch (error) {
            setIsLoading(false);
            console.error("Error fetching data:", error);
        }
    }
    const handleUserInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setUsername(event.target.value);
    }
    const handlePassInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
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
            />
            ) : (
                null
            )}
        </>
    );
};

export default LoginContainer;