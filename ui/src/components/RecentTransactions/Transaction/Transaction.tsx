import "./Transaction.css";

interface Props {
    mode: string
    account: string
    transaction: string
    amount: number
};

const Transaction: React.FC<Props> = ({mode, account, transaction, amount}) => {
    const reel = `/images/${mode}/transactions/reel.png`;
    return (
        <div className="trans">
            <img src={reel} className="trans__icon"/>
            <span className="trans__acc">{account}</span>
            <br/>
            <span className="trans__name">{transaction}...${amount}</span>
        </div>
    );
};

export default Transaction;