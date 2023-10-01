import { useContext, useState } from "react";
import ThemeContext from "../../context/ThemeContext";

export const useChart= () => {

    const { mode } = useContext(ThemeContext)

    const populateChart = () => {
        return (
            <div>HEELo</div>
        );
    };

    return {
        populateChart
    };
};