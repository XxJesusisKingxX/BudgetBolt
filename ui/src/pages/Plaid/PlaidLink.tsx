interface Props {
  plaidFunction: Function
  ready: boolean
}

const PlaidLink: React.FC<Props> = ({ plaidFunction, ready }) => {
  return (
    <button data-testid='plaid-button' className="btn btn--plaid" onClick={() => plaidFunction()} disabled={!ready}>Add Account</button>
  );
};

export default PlaidLink;