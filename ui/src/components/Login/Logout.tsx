import "./LoginLogout.css";

interface Props {
    logout: Function
};

const Logout: React.FC<Props> = ({ logout }) => {
    return (
        <>
            <button className="login-button" onClick={() => logout()}>Logout</button>
        </>
    );
};

export default Logout;