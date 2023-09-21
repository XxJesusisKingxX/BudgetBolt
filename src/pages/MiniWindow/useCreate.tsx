import { useContext, useState } from "react";
import ExpenseComponent from "./components/ExpenseComponent";
import { EndPoint } from "../../constants/endpoints";
import ThemeContext from "../../context/ThemeContext";

// Interface for the shape of a expense
export interface Expense {
    ID: string
    Name: string   // The name for the expense
    Limit: string  // Limit of the budgeted expense
    Spent: string  // Amount spent in the budgeted expense
}

export const useCreate = () => {
    const [expenses, setExpenses] = useState<Expense[]>([]);
    const [isLoading, setLoading] = useState(false);

    const { mode } = useContext(ThemeContext)

    const updateExpense = async (id: string, limit: string) => {
        if (id && limit) {
            setLoading(true)
            const response = await fetch(EndPoint.UPDATE_EXPENSES, {
                method: "POST",
                body: new URLSearchParams({
                    id: id,
                    limit: limit,
                })
            })
    
            if (response.ok) {
                setLoading(false)
                getExpenses()
            }
        }
        else {
            console.error("ERROR: empty limit")
        }
    };

    const getExpenses = async () => {
        setLoading(true)
        const response = await fetch(EndPoint.GET_EXPENSES, {
            method: "GET"
        })

        if (response.ok) {
            const data = await response.json()
            setExpenses(data["expenses"])
            setLoading(false)
        }
    };

    const showExpenses = (loading = isLoading, expensesList = expenses) => {
        const loadingIcon = `/images/${mode}/loading.png`;
        return !loading && expensesList ? (
            expensesList.slice().map((expense : Expense) => (
              <ExpenseComponent
                key={expense.ID}
                update={updateExpense}
                id={expense.ID}
                name={expense.Name}
                limit={expense.Limit}
                spent={expense.Spent}
              />
            ))
          ) : <img className='miniwindow__budget__view__loading' src={loadingIcon} alt="Loading" />;
    }

    const addExpenses = async (expense: Expense) => {
        const response = await fetch(EndPoint.CREATE_EXPENSES, {
            method: "POST",
            body: new URLSearchParams({
                name: expense.Name,
                limit: expense.Limit,
                spent: expense.Spent
            })
        })
        if (response.ok) getExpenses();
    }

    return {
        getExpenses,
        addExpenses,
        showExpenses,
        isLoading
    };
};