// Format date in a MMM DD YYYY with month placeholder being first 3 letters of month
export function formatOverviewDate(currentDate) {
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
    switch (month) {
        case 1:
            month = "Jan"
            break;
        case 2:
            month = "Feb"
            break;
        case 3:
            month = "Mar"
            break;
        case 4:
            month = "Apr"
            break;
        case 5:
            month = "May"
            break;
        case 6:
            month = "Jun"
            break;
        case 7:
            month = "Jul"
            break;
        case 8:
            month = "Aug"
            break;
        case 9:
            month = "Sep"
            break;
        case 10:
            month = "Oct"
            break;
        case 11:
            month = "Nov"
            break;
        case 12:
            month = "Dec"
            break;
        default:
            month = "N/A"
            break;
    };
    const date = `${month} ${day}, ${year}`;
    return date;
};
