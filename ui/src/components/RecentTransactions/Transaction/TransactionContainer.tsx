import { useEffect, useState } from 'react';
import Transaction from './Transaction';

interface Transaction {
    From: string
    Amount: number
    Vendor: string
}
const TransactionContainer = () => {
    const [transactions, setTransactions] = useState<Transaction[]>([]);
    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch("/api/transactions/get?username=test_user", {
                    method: "GET",
                });
                const data = await response.json();
                setTransactions(data["transactions"]);
            } catch (error) {
                console.error("Error fetching data:", error);
            }
        };
        fetchData();
    }, []);
    console.log(transactions)
    const maxPeek = 6
    const maxChar = 18
    return (
        <>
            {transactions.slice(0, maxPeek).map((transaction) => (
                <Transaction bottom={{marginBottom:'-35px'}} account={transaction.From} transaction={transaction.Vendor.length < maxChar ? transaction.Vendor : "Click to see more"} amount={transaction.Amount}/>
            ))}
        </>
    );
};

export default TransactionContainer;