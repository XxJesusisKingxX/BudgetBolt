interface Props {
    action: Function // any action for the button to perform
    name: string     // the name of the button to be displayed
};

const AuthButton: React.FC<Props> = ({ action, name }) => {
    return (
        <>
            <button data-testid='auth-button' className="btn btn--auth" onClick={() => action()}>{name}</button>
        </>
    );
};

export default AuthButton;