import { useContext, useState } from "react";
import { EndPoint } from "../../constants/endpoints";
import AppContext from "../../context/AppContext";
import { getDateView } from "../../utils/formatDate";

export const useIncome = () => {
    const [incomes, setIncomes] = useState(null);

    const { budgetView, dispatch } = useContext(AppContext);

    const upsertIncome = async () => {
        try{
            console.log(budgetView)
            const response = await fetch(EndPoint.UPSERT_INCOMES, {
                method: "POST",
                body: new URLSearchParams({
                    date: getDateView(new Date(), budgetView),
                })
            })
    
            if (response.ok) {
                const inc = await getIncomes();
                storeIncomeTotal(inc)
            }
        } catch(error) {
            console.log(error);
        }
    };

    const getIncomes = async () => {
        try {
            const response = await fetch(EndPoint.GET_INCOMES, {
                method: "GET"
            })
            if (response.ok) {
                const data = await response.json();
                if (data) {
                    setIncomes(data["incomes"]);
                    return data["incomes"]
                }
            }
        } catch(error) {
            console.log(error);
        }

    };

    const storeIncomeTotal = (incomesList: any) => {
        let total: number = 0.00;
        incomesList.slice().map((income: any) => {
            total += income.income_amount > 0 ? income.income_amount : 0
        });
        dispatch({ type:'SET_STATE', state: { totalIncome: total }})
    }

    return {
        upsertIncome,
    };
};

// test