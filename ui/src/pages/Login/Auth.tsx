import "./Auth.css";

interface Props {
    authType: Function
    authName: string
};

const Auth: React.FC<Props> = ({ authType, authName }) => {
    return (
        <>
            <button className="btn btn--auth" onClick={() => authType()}>{authName}</button>
        </>
    );
};

export default Auth;