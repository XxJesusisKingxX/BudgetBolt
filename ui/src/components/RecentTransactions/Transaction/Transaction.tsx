import './Transaction.css'
import reel from './icons/reel.png'

interface Props {
    account: string
    transaction: string
    amount: number
    bottom?: React.CSSProperties
}
const Transaction: React.FC<Props> = ({account, transaction, amount, bottom}) => {
    return (
        <div className="transaction-container" style={bottom}>
            <img src={reel} className="transaction-icon"/>
            <span className="account-name">{account}</span>
            <br/>
            <span className="transaction-name">{transaction}........................${amount}</span>
        </div>
    );
};

export default Transaction