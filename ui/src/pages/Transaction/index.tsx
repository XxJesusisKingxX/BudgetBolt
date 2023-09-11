import { useContext, useEffect, useState } from 'react';
import { EndPoint } from '../../constants/endpoints';
import Transaction from './TransactionComponent';
import AppContext from '../../context/AppContext';
import ThemeContext from '../../context/ThemeContext';
import LoginContext from '../../context/LoginContext';
import TransactionComponent from './TransactionComponent';

// Transaction interface for a single transaction
interface Transaction {
    ID: number;
    From: string;
    Amount: number;
    Vendor: string;
}

// TransactionContainer component
const TransactionContainer = () => {
    const [transactions, setTransactions] = useState<Transaction[]>([]);
    const [isLoading, setIsLoading] = useState(false);

    // Accessing profile, isTransactionsRefresh, dispatch, mode, and isLogin from contexts
    const { profile, isTransactionsRefresh, dispatch } = useContext(AppContext);
    const { mode } = useContext(ThemeContext);
    const { isLogin } = useContext(LoginContext);

    const maxPeek = 6; // Maximum number of transactions to display
    const maxChar = 18; // Maximum number of characters for transaction name
    const everyHour = 3600000; // Interval for hourly update check

    useEffect(() => {
        const sidebar = document.getElementById("sidebar");

        // Function to update transactions every hour
        const checkHourlyUpdate = () => {
            dispatch({ type: "SET_STATE", state: { lastTransactionsUpdate: new Date(), isTransactionsRefresh: !isTransactionsRefresh } });
        };

        // Function to retrieve transactions from the server
        const retrieveTransactions = async () => {
            try {
                await fetch(EndPoint.CREATE_TRANSACTIONS, {
                    method: "POST",
                    body: new URLSearchParams({
                        username: profile,
                    }),
                });

                const baseURL = window.location.href;
                const url = new URL(EndPoint.GET_TRANSACTIONS, baseURL);
                url.search = new URLSearchParams({
                    username: profile,
                }).toString();
                const retrieveResponse = await fetch(url, {
                    method: "GET",
                });
                
                if (retrieveResponse.ok) {
                    const data = await retrieveResponse.json();
                    setTransactions(data["transactions"]);
                } else {
                    console.error("failed to retrieve transactions");
                }

                if (sidebar) sidebar.onanimationend = null
                setIsLoading(false);
            } catch (error) {
                console.log("failed to retrieve transactions");
            }
        };

        // If logged in, initiate animations and transaction retrieval
        if (isLogin) {
            if (sidebar) sidebar.onanimationend = () => { setIsLoading(true) }; //set loading after animation of sidebar
            retrieveTransactions();
            const intervalId = setInterval(checkHourlyUpdate, everyHour);
            return () => clearInterval(intervalId);
        }
    }, [isTransactionsRefresh, isLogin]);

    const loading = `/images/${mode}/loading.png`;

    return (
        <>
            {!isLoading && transactions.length !== 0 ? transactions.slice(0, maxPeek).map((transaction) => (
                <TransactionComponent
                    key={transaction.ID}
                    account={transaction.From}
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

export default TransactionContainer;
