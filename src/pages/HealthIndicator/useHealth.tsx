import { useContext } from "react";
import { Health } from "../../constants/style";
import AppContext from "../../context/AppContext";

export const useHealth = () => {
    const { totalExpenses, totalIncome, dispatch } = useContext(AppContext)

    const healthLevel = 1 - totalExpenses / totalIncome;
    const calculateHealth = (lvl: number = healthLevel) => {
        if (lvl >= 1) {
            dispatch({ type:'SET_STATE', state: { health: Health.LOW }})
        } else if (lvl >= 0.5) {
            dispatch({ type:'SET_STATE', state: { health: Health.HIGH }})
        } else if ( lvl >= 0.2) {
            dispatch({ type:'SET_STATE', state: { health: Health.MEDIUM }})
        } else if (lvl < 0.2) {
            dispatch({ type:'SET_STATE', state: { health: Health.LOW }})
        } else {
            dispatch({ type:'SET_STATE', state: { health: Health.NONE }})
        }
    }

    return {
        calculateHealth
    };
};