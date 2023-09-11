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
