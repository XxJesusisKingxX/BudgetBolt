import PlaidLink from "../../Plaid/PlaidLinkContainer";

const AccountWindow = () => {
    return (
        <div className="windowcont windowcont--plaid">
            <h1 className="windowcont__title">Setup Account</h1>
            <PlaidLink/>
        </div>
    );
};

export default AccountWindow;