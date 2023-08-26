import { useContext, useEffect, useState } from "react";
import { EndPoint } from "../../constants/endpoints"
import { useAppStateActions } from "../../redux/useUserContextState"
import Transaction from "./Transaction";
import Context from "../../context/Context";

interface Transaction {
    ID: number
    From: string
    Amount: number
    Vendor: string
}

const TransactionContainer = () => {
    const [transactions, setTransactions] = useState<Transaction[]>([]);
    const [isLoading, setIsLoading] = useState(false);
    const { mode, profile, isLogin, isTransactionsRefresh } = useContext(Context);
    const { setLastTransactionsUpdateState, setTransactionsRefreshState } = useAppStateActions();
    const maxPeek = 6;
    const maxChar = 18;
    const everyHour = 3600000;
    
    useEffect(() => {
        const sidebar = document.getElementById("sidebar");
        const isAnimationDone = (done: boolean) => {
            if (done) {
                if (sidebar) {
                    sidebar.onanimationend = null;
                }
            } else {
                if (sidebar) {
                    sidebar.onanimationend = () => {
                        setIsLoading(true)
                    };
                }
            }
        }
        const checkHourlyUpdate = () => {
            setLastTransactionsUpdateState( new Date() )
            setTransactionsRefreshState(!isTransactionsRefresh)
        };
        const retrieveTransactions = async () => {
            try {
                await fetch(EndPoint.CREATE_TRANSACTIONS, {
                    method: "POST",
                    body: new URLSearchParams ({
                        username: profile
                    })
                });
                const baseURL = window.location.href
                const url = new URL(EndPoint.GET_TRANSACTIONS, baseURL);
                url.search = new URLSearchParams(({
                    username: profile
                })).toString();
                const retrieveResponse = await fetch(url, {
                    method: "GET",
                });
                const data = await retrieveResponse.json();
                setTransactions(data["transactions"]);
                isAnimationDone(true);
                setIsLoading(false);
            } catch (error) {
                console.log("failed to retrieve transactions")
            }
        };
        if (isLogin) {
            isAnimationDone(false);
            retrieveTransactions();
            const intervalId = setInterval(checkHourlyUpdate, everyHour);
            return () => clearInterval(intervalId);
        }
    }, [isTransactionsRefresh, isLogin]);

    const loading = `/images/${mode}/loading.png`;
    return (
        <>
            {!isLoading ? transactions.slice(0, maxPeek).map((transaction) => (
                <Transaction
                    key={transaction.ID}
                    account={transaction.From}
                    transaction={transaction.Vendor.length < maxChar ? transaction.Vendor : "Click to see more"}
                    amount={transaction.Amount}
                    mode={mode}
                />
            )) : (
                <img className="loading loading--trans" src={loading}/>
            )}
        </>
    );
};

export default TransactionContainer;