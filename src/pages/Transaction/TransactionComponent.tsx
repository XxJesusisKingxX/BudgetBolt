import React from 'react';
import './Transaction.css';

// Props interface for the Transaction component
interface Props {
    mode: string;        // Theme mode for determining image paths
    account: string;     // Account information for the transaction
    transaction: string; // Transaction name or description
    amount: number;      // Amount of the transaction
}

const TransactionComponent: React.FC<Props> = ({ mode, account, transaction, amount }) => {
    // Dynamic image path for the reel icon based on the provided mode
    const reel = `/images/${mode}/transactions/reel.png`;

    return (
        <div className='trans'> {/* Container div for the transaction */}
            <img src={reel} className='trans__icon' alt='Transaction Icon' /> {/* Transaction icon */}
            <span className='trans__acc'>{account}</span> {/* Account information */}
            <br />
            <span className='trans__name'>{transaction}...${amount}</span> {/* Transaction details */}
        </div>
    );
};

export default TransactionComponent;
