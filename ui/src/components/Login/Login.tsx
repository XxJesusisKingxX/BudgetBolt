import "./LoginLogout.css";

interface Props {
    login: Function
};

const Login: React.FC<Props> = ({ login }) => {
    return (
        <>
            <button className="login-button" onClick={() => login()}>Login</button>
        </>
    );
};

export default Login;