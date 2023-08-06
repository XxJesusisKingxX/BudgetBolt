import "./Transaction.css";

interface Props {
    mode: string
    account: string
    transaction: string
    amount: number
    bottom?: React.CSSProperties
};

const Transaction: React.FC<Props> = ({mode, account, transaction, amount, bottom}) => {
    const reel = `/images/${mode}/reel.png`;
    return (
        <div className="transaction-container" style={bottom}>
            <img src={reel} className="transaction-icon"/>
            <span className="account-name">{account}</span>
            <br/>
            <span className="transaction-name">{transaction}...${amount}</span>
        </div>
    );
};

export default Transaction;