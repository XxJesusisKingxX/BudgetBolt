interface Props {
    action: Function
    name: string
};

const AuthButton: React.FC<Props> = ({ action, name }) => {
    return (
        <>
            <button data-testid='auth-button' className="btn btn--auth" onClick={() => action()}>{name}</button>
        </>
    );
};

export default AuthButton;