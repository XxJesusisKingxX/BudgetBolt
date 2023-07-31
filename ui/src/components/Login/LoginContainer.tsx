import { useContext } from "react";
import Login from "./Login";
import Context from "../../context/Context";
import LoginWindow from "./LoginWindow/LoginWindow";
import SignupWindow from "./SignupWindow/SignupWindow";

const LoginContainer = () => {
    const { showCreateWindow, showLoginWindow, dispatch } = useContext(Context);
    const handleCreateWindow = (open: boolean) => {
        if (open) {
            handleLoginWindow(false)
            dispatch({ type:"SET_STATE", state:{ showCreateWindow: true }});
        } else {
            dispatch({ type:"SET_STATE", state:{ showCreateWindow: false }});
        }
    };
    const handleLoginWindow = (open: boolean) => {
        if (open) {
            dispatch({ type:"SET_STATE", state:{ showLoginWindow: true }});
        } else {
            dispatch({ type:"SET_STATE", state:{ showLoginWindow: false }});
        }
    };
    return (
        <>
            <Login open={() => handleLoginWindow(true)}/>
            {showLoginWindow ? 
            <LoginWindow open={() => handleCreateWindow(true)} close={() => handleLoginWindow(false)}/> 
            : showCreateWindow ? <SignupWindow close={() => handleCreateWindow(false)}/> : ""}
        </>
    );
};

export default LoginContainer;