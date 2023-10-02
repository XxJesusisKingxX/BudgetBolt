import { useContext, useState } from "react";
import ExpenseComponent from "./components/ExpenseComponent";
import { EndPoint } from "../../constants/endpoints";
import ThemeContext from "../../context/ThemeContext";
import { getDateView } from "../../utils/formatDate";
import { BudgetView } from "../../constants/view";
import AppContext from "../../context/AppContext";

// Interface for the shape of a expense
export interface Expense {
    ID: string
    Name: string   // The name for the expense
    Limit: string  // Limit of the budgeted expense
    Spent: string  // Amount spent in the budgeted expense
}

export const useExpense = () => {
    const [expenses, setExpenses] = useState<Expense[]>([]);
    const [isLoading, setLoading] = useState(false);

    const { mode } = useContext(ThemeContext);
    const { dispatch } = useContext(AppContext);

    const updateAllExpenses = async (currentView: BudgetView = BudgetView.MONTHLY) => {
        try {
            setLoading(true)
            const response = await fetch(EndPoint.UPDATE_ALL_EXPENSES, {
                method: "POST",
                body: new URLSearchParams({
                    date: getDateView(new Date(), currentView),
                })
            })
    
            if (response.ok) {
                setLoading(false);
                const exp = await getExpenses()
                storeExpenseTotal(exp);
            }
        } catch (error) {
            console.log(error);
        }
    };

    const updateExpense = async (id: string, limit: string) => {
        
        if (id && limit) {
            try {
                setLoading(true)
                const response = await fetch(EndPoint.UPDATE_EXPENSES, {
                    method: "POST",
                    body: new URLSearchParams({
                        id: id,
                        limit: limit,
                    })
                })
        
                if (response.ok) {
                    setLoading(false);
                    getExpenses();
                }
            } catch (error) {
                console.log(error);
            }
        }
        else {
            console.error("ERROR: empty limit and/or id");
        }
    };

    const getExpenses = async () => {
        setLoading(true)
        try {
            const response = await fetch(EndPoint.GET_EXPENSES, {
                method: "GET"
            })
            if (response.ok) {
                const data = await response.json();
                if (data) {
                    setExpenses(data["expenses"]);
                    setLoading(false)
                    return data["expenses"]
                }
                setLoading(false);
            }
        } catch(error) {
            console.log(error);
        }

    };

    const showExpenses = (loading = isLoading, expensesList = expenses) => {
        const loadingIcon = `/images/${mode}/loading.png`;

        const rows = expensesList.slice().map((expense: any) => (
            <ExpenseComponent
            key={expense.expense_id}
            update={updateExpense}
            id={expense.expense_id}
            name={expense.expense_name}
            limit={expense.expense_limit}
            spent={expense.expense_spent}
            />
        ));
        return (
        !loading && expensesList ?
        rows 
        : 
        <img className='miniwindow__budget__view__loading' src={loadingIcon} alt="Loading" />
        );

    }

    const addExpenses = async (expense: Expense) => {
        try {
            const response = await fetch(EndPoint.CREATE_EXPENSES, {
                method: "POST",
                body: new URLSearchParams({
                    name: expense.Name,
                    limit: expense.Limit,
                    spent: expense.Spent,
                }),
            });
            if (response.ok) getExpenses();
        } catch (error) {
            console.log(error);
        }
        
    }

    const storeExpenseTotal = (expensesList: any) => {
        let total: number = 0.00;
        expensesList.slice().map((expense: any) => {
            total += expense.expense_spent > 0 ? expense.expense_spent : 0
            return null;
        });
        dispatch({ type:'SET_STATE', state: { totalExpenses: total }})
    }

    return {
        getExpenses,
        addExpenses,
        showExpenses,
        updateExpense,
        updateAllExpenses,
        isLoading
    };
};