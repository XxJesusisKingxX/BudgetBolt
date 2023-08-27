import { useContext, useState } from "react";
import { AuthType } from "../../constants/auth";
import UserContext from "../../context/UserContext";
import { EndPoint } from "../../constants/endpoints";
import LoginContext from "../../context/LoginContext";

export const useLogin = (username: string, password: string, valid: boolean) => {
    const { userDispatch} = useContext(UserContext);
    const { loginDispatch } = useContext(LoginContext);
    const [serverErr, setServerErr] = useState(false);
    const [nameErr, setNameErr] = useState(false);
    const [authErr, setAuthErr] = useState(false);
    const [takenNameErr, setTakenNameErr] = useState(false);
    const [showLoading, setShowLoading] = useState(false);

    const clearErrors = () => {
        setServerErr(false)
        setAuthErr(false)
        setNameErr(false)
        setTakenNameErr(false)
    }

    const handleAuth = async (authType: number) => {
        if (valid) {
            setShowLoading(true);
            try {
                let response = null
                if (authType === AuthType.Login) {
                    const baseURL = window.location.href;
                    const GET_PROFILE_URL = new URL(EndPoint.GET_PROFILE, baseURL);
                    GET_PROFILE_URL.search = new URLSearchParams(({
                        username: username,
                        password: password
                    })).toString();
                    response = await fetch(GET_PROFILE_URL, {
                            method: "GET"
                        }
                    );
                } else if (authType === AuthType.SignUp) {
                    response = await fetch(EndPoint.CREATE_PROFILE, {
                            method: "POST",
                            body: new URLSearchParams ({
                                username: username,
                                password: password
                            })
                        }
                    );
                }
                if (response?.ok) {
                    setShowLoading(false)
                    const lowercaseUsername = username.toLocaleLowerCase()
                    userDispatch({ type: "SET_STATE", state: { profile: lowercaseUsername }});
                    userDispatch({ type: "SET_STATE", state: { isLogin: true }});
                    if (authType == AuthType.SignUp) {
                        loginDispatch({ type: "SET_STATE", state: { showAccountWindow: true }});
                    }
                } else if (response?.status == 409) {
                    setShowLoading(false);
                    setTakenNameErr(true);
                } else if (response?.status == 401) {
                    setShowLoading(false);
                    setAuthErr(true);
                } else if (response?.status == 404) {
                    setShowLoading(false);
                    setNameErr(true);
                } else {
                    setShowLoading(false);
                    setServerErr(true);
                }
            } catch (error) {
                console.log(error)
                setShowLoading(false);
                setServerErr(true);
            }
        }
    }
    
    return {
        serverErr,
        authErr,
        nameErr,
        takenNameErr,
        showLoading,
        clearErrors,
        handleAuth
    };
};