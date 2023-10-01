import { BudgetView } from "../constants/view";

/**
 * Format date in a MM DD YYYY with month placeholder being first 3 letters of month
 *
 * @param {Date} currentDate - the current local time date object
 * @returns {boolean} Format date as such: `Jan 03, 2023`
 */
export function formatOverviewDate(currentDate: Date) {
    const year = currentDate.getFullYear().toString();
    let month = currentDate.getMonth() + 1; // Months are zero-indexed, so add 1
    let day = currentDate.getDate().toString();

    // Parse day ordinal
    if (day.endsWith("1") && !day.endsWith("11")) { 
        day += "st"
    }
    else if (day.endsWith("2") && !day.endsWith("12")) { 
        day += "nd"
    }
    else if (day.endsWith("3") && !day.endsWith("13")) { 
        day += "rd"
    }
    else {
        day += "th"
    };
    
    // Parse month for abbreviation
    let monthStr
    switch (month) {
        case 1:
            monthStr = "Jan"
            break;
        case 2:
            monthStr = "Feb"
            break;
        case 3:
            monthStr = "Mar"
            break;
        case 4:
            monthStr = "Apr"
            break;
        case 5:
            monthStr = "May"
            break;
        case 6:
            monthStr = "Jun"
            break;
        case 7:
            monthStr = "Jul"
            break;
        case 8:
            monthStr = "Aug"
            break;
        case 9:
            monthStr = "Sep"
            break;
        case 10:
            monthStr = "Oct"
            break;
        case 11:
            monthStr = "Nov"
            break;
        case 12:
            monthStr = "Dec"
            break;
        default:
            monthStr = "N/A"
            break;
    };
    const date = `${monthStr} ${day}, ${year}`;
    return date;
};

export const getDateView = (currentDate: Date, view: BudgetView) => {
    let year;
    let month;
    let day;
    let formattedDate;

    const monthlyView = () => {
        year = currentDate.getFullYear();
        month = String(currentDate.getMonth() + 1).padStart(2, '0'); // it's zero-based so no need to backup 1
        formattedDate = `${year}-${month}-01`;
        return formattedDate;
    }
    const weeklyView = () => {
        const dayOfWeek = currentDate.getDay()
            const newDate = new Date(currentDate);
            newDate.setDate(currentDate.getDate() - dayOfWeek); //subtract day of week to get back to beginning of week
            year = newDate.getFullYear();
            month = String(newDate.getMonth() + 1).padStart(2, '0'); // it's zero-based so no need to backup 1
            day = String(newDate.getDate()).padStart(2, '0');
            formattedDate = `${year}-${month}-${day}`;
            return formattedDate;
    }
    const yearlyView = () => {
        year = currentDate.getFullYear();
            formattedDate = `${year}-01-01`;
            return formattedDate;
    }
    switch (view) {
        case BudgetView.WEEKLY:
            return weeklyView();
        case BudgetView.MONTHLY:
            return monthlyView();
        case BudgetView.YEARLY:
            return yearlyView();
        default:
            return monthlyView();
    }
}