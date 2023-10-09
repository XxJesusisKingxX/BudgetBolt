import PlaidLink from '../../Plaid/PlaidLink';

const AccountWindow = () => {
    return (
        <div data-testid='account-window' className='window window__button'>
            <span className='window__title'>Setup Account</span>
            <PlaidLink/>
        </div>
    );
};

export default AccountWindow;