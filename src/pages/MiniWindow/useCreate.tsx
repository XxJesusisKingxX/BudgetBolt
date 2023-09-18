import { useState } from "react";
import ExpenseComponent from "./components/ExpenseComponent";
import { EndPoint } from "../../constants/endpoints";

// Interface for the shape of a expense
export interface Expense {
    ID?: number
    Name: string   // The name for the expense
    Limit: string  // Limit of the budgeted expense
    Spent: string  // Amount spent in the budgeted expense
}

export const useCreate = () => {
    const [expenses, setExpenses] = useState<Expense[]>([]);

    const getExpenses = async () => {
        const response = await fetch(EndPoint.GET_EXPENSES, {
            method: "GET"
        })

        if (response.ok) {
            const data = await response.json()
            setExpenses(data["expenses"])
        }
    };

    const showExpenses = () => {
        return expenses ? (
            expenses.slice().map((expense) => (
              <ExpenseComponent
                key={expense.ID}
                name={expense.Name}
                limit={expense.Limit}
                spent={expense.Spent}
              />
            ))
          ) : null;
    }

    const addExpenses = async (expense: Expense) => {
        await fetch(EndPoint.CREATE_EXPENSES, {
            method: "POST",
            body: new URLSearchParams({
                name: expense.Name,
                limit: expense.Limit,
                spent: expense.Spent
            })
        })

        getExpenses();
    }

    return {
        getExpenses,
        addExpenses,
        showExpenses
    };
};