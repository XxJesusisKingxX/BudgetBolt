import { useContext, useEffect, useState } from 'react';
import { EndPoint } from '../../constants/endpoints';
import TransactionComponent from './TransactionComponent';
import AppContext from '../../context/AppContext';
import ThemeContext from '../../context/ThemeContext';
import LoginContext from '../../context/LoginContext';
import { getCookie } from '../../utils/cookie';
import { getDateView } from '../../utils/formatDate';

// Transaction interface for a single transaction
interface Transactions {
    ID: string;
    AccountName: string;
    Amount: number;
    Vendor: string;
}

const Transaction = () => {
    const [transactions, setTransactions] = useState<Transactions[] | null>(null);
    const [isLoading, setIsLoading] = useState(false);

    // Accessing isTransactionsRefresh, dispatch, mode, and isLogin from contexts
    const { budgetView, isTransactionsRefresh, dispatch } = useContext(AppContext);
    const { mode } = useContext(ThemeContext);
    const { isLogin } = useContext(LoginContext);

    const maxPeek = 6; // Maximum number of transactions to display
    const maxChar = 18; // Maximum number of characters for transaction name
    const everyHour = 3600000; // Interval for hourly update check

    useEffect(() => {
        // Function to update transactions every hour
        const checkHourlyUpdate = () => {
            dispatch({ type: "SET_STATE", state: { lastTransactionsUpdate: new Date(), isTransactionsRefresh: !isTransactionsRefresh } });
        };

        // Function to retrieve transactions from the server
        const retrieveTransactions = async () => {
            try {
                // Store new transactions for users
                await fetch(EndPoint.CREATE_TRANSACTIONS, {
                    method: "POST",
                });

                // Get transactions for user
                const baseURL = window.location.href;
                const url = new URL(baseURL + EndPoint.GET_TRANSACTIONS);
                url.search = new URLSearchParams({
                    date: getDateView(new Date(), budgetView)
                }).toString();
                const retrieveResponse = await fetch(url, {
                    method: "GET",
                });
                
                if (retrieveResponse.ok) {
                    const data = await retrieveResponse.json();
                    setTransactions(data["transactions"]);
                } else {
                    console.error("failed to retrieve transactions");
                    setIsLoading(false);
                }

                dispatch({ type: "SET_STATE", state: { lastTransactionsUpdate: new Date()}});
                setIsLoading(false);
            } catch (error) {
                console.log("failed to retrieve transactions");
                setIsLoading(false);
            }
        };

        // If logged in, initiate animations and transaction retrieval
        if (getCookie("UID")) {
            setIsLoading(true)
            retrieveTransactions();
            const intervalId = setInterval(checkHourlyUpdate, everyHour);
            return () => clearInterval(intervalId);
        }
    }, [isTransactionsRefresh, budgetView, isLogin, dispatch]);

    const loading = `/images/${mode}/loading.png`;

    return (
        <>
            {!isLoading && transactions ? transactions.slice(0, maxPeek).map((transaction) => (
                <TransactionComponent
                    key={transaction.ID}
                    account={transaction.AccountName}
                    transaction={transaction.Vendor.length < maxChar ? transaction.Vendor : "Click to see more"}
                    amount={transaction.Amount}
                    mode={mode}
                />
            )) : (
                <img className="loading loading--trans" src={loading} alt="Loading" />
            )}
        </>
    );
};

export default Transaction;
