import "./AccountWindow.css";
import PlaidLink from "../../Plaid/PlaidLinkContainer";

interface Props {
    mode: string
    close: Function
}
const AccountWindow: React.FC<Props> = ({ mode, close }) => {
    const cancel = `/images/${mode}/cancel.png`
    return (
        <>
            <div className="account-window-container">
                <img className="account-window-cancel" src={cancel} onClick={() => close()}/>
                <h1 className="account-window-title">Set Up Your Account</h1>
                <span className="account-link-button"><PlaidLink/></span>
            </div>
        </>
    );
};

export default AccountWindow;