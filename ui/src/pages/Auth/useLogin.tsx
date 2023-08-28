import { useContext, useState } from 'react';
import { AuthType } from '../../constants/auth';
import UserContext from '../../context/UserContext';
import { EndPoint } from '../../constants/endpoints';
import LoginContext from '../../context/LoginContext';

// Custom hook for handling user authentication (login and signup)
export const useLogin = (username: string, password: string, valid: boolean) => {
    // Accessing user and login context
    const { userDispatch } = useContext(UserContext);
    const { loginDispatch } = useContext(LoginContext);

    // State variables to manage various error states and loading state
    const [serverErr, setServerErr] = useState(false); // Indicates a server error
    const [nameErr, setNameErr] = useState(false); // Indicates a non-existent username
    const [authErr, setAuthErr] = useState(false); // Indicates an authentication error (wrong credentials)
    const [takenNameErr, setTakenNameErr] = useState(false); // Indicates that the username is already taken
    const [showLoading, setShowLoading] = useState(false); // Indicates loading state

    // Function to clear all error states
    const clearErrors = () => {
        setServerErr(false);
        setAuthErr(false);
        setNameErr(false);
        setTakenNameErr(false);
    };

    // Function to handle authentication process (login or signup)
    const handleAuth = async (authType: number) => {
        if (valid) {
            setShowLoading(true);
            try {
                let response = null;
                if (authType === AuthType.Login) {
                    // Construct URL for login
                    const baseURL = window.location.href;
                    const GET_PROFILE_URL = new URL(EndPoint.GET_PROFILE, baseURL);
                    GET_PROFILE_URL.search = new URLSearchParams({
                        username: username,
                        password: password
                    }).toString();

                    // Fetch login endpoint
                    response = await fetch(GET_PROFILE_URL, {
                        method: "GET"
                    });
                } else if (authType === AuthType.SignUp) {
                    // Fetch signup endpoint
                    response = await fetch(EndPoint.CREATE_PROFILE, {
                        method: "POST",
                        body: new URLSearchParams({
                            username: username,
                            password: password
                        })
                    });
                }

                // Process response
                if (response?.ok) {
                    setShowLoading(false);
                    const lowercaseUsername = username.toLocaleLowerCase();
                    // Update user and login context upon successful authentication
                    userDispatch({ type: "SET_STATE", state: { profile: lowercaseUsername } });
                    loginDispatch({ type: "SET_STATE", state: { isLogin: true } });
                    if (authType == AuthType.SignUp) {
                        loginDispatch({ type: "SET_STATE", state: { showAccountWindow: true } });
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
                console.log(error);
                setShowLoading(false);
                setServerErr(true);
            }
        }
    };

    // Return the authentication-related states and functions
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