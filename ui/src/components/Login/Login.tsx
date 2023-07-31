import './Login.css';

interface Props {
    open: Function
};

const Login: React.FC<Props> = ({ open }) => {
    return (
        <>
            <button className="login-button" onClick={() => open()}>Login | Signup</button>
        </>
    );
};

export default Login;