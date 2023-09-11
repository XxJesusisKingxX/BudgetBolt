import PlaidLink from '../../Plaid/PlaidLink';

const AccountWindow = () => {
    return (
        <div data-testid='account-window' className='windowcont windowcont--plaid'>
            <h1 className='windowcont__title'>Setup Account</h1>
            <PlaidLink/>
        </div>
    );
};

export default AccountWindow;