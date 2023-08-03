import { useContext, useEffect, useState } from 'react';
import { updateTransactionsEveryHour } from '../../../utils/UpdateTimer'
import loadingSpinner from './icons/loading.png';
import Transaction from './Transaction';
import Context from '../../../context/Context';
import './Transaction.css'

interface Transaction {
    ID: number
    From: string
    Amount: number
    Vendor: string
}
const TransactionContainer = () => {
    const [transactions, setTransactions] = useState<Transaction[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const {isTransactionsRefresh, lastTransactionsUpdate, dispatch} = useContext(Context);
    const maxPeek = 6
    const maxChar = 18
    const everyHour = 3600000

    useEffect(() => {
        const checkHourlyUpdate = () => {
            const isPastHour = updateTransactionsEveryHour(lastTransactionsUpdate)
            if (isPastHour) {
                dispatch({ type: "SET_STATE" , state: { lastTransactionsUpdate: new Date(), isTransactionsRefresh: !isTransactionsRefresh }})
            }
        };
        const retrieveTransactions = async () => {
            try {
                await fetch("/api/transactions/create", {
                    method: "POST",
                    body: new URLSearchParams ({
                        username: "test_user"
                    })
                });
                const retrieveResponse = await fetch("/api/transactions/get?username=test_user", {
                    method: "GET",
                });
                const data = await retrieveResponse.json();
                setTransactions(data["transactions"]);
                setIsLoading(false);
            } catch (error) {
                
            }
        };
        // retrieveTransactions();
        const intervalId = setInterval(checkHourlyUpdate,  everyHour);
        return () => clearInterval(intervalId);
    }, [isTransactionsRefresh]);
    return (
        <>
            {!isLoading ? transactions.slice(0, maxPeek).map((transaction) => (
                <Transaction key={transaction.ID} bottom={{marginBottom:'-35px'}} account={transaction.From} transaction={transaction.Vendor.length < maxChar ? transaction.Vendor : "Click to see more"} amount={transaction.Amount}/>
            )) : (
                <img className="transaction-loading" src={loadingSpinner}/>
            )}
        </>
    );
};

export default TransactionContainer;