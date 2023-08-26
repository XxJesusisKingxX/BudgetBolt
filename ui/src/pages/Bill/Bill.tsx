import { FC } from "react"
import "./Bill.css"
import { ModeType } from "../../constants/style"

interface Props {
    name: string
    price: string
    daysLeft: number
    dueDate: string
    category: string
    spacing: number
    mode: string
    icon?: string
}
const Bill: FC<Props> = ({ icon, mode, name, price, daysLeft, dueDate, category, spacing }) => {
    const pic = `/images/${mode}/bills/${icon}.png`;
    return (
        <div className="bill">
            {/* TODO remove if statmemnt/optional for icon when can handling */}
            {icon ? <img className="bill__icon" src={pic}/> : ""}{name}: ${price} ~ {daysLeft} days
            <div style={{top:`${7+spacing}rem`}} className="bill__fullview">
                <div className="bill__fullview__cont"/>
                <div className="bill__fullview__txt">
                    <div className="bill__fullview__txt__title">
                        {name}
                    </div>
                    <div className="bill__fullview__txt__cols">
                        <div className="bill__fullview__txt__cols__lcol">
                            Due Date:
                            <br/>
                            {dueDate}
                        </div>
                        <div className="bill__fullview__txt__cols__rcol">
                            Category:
                            <br/>
                            {category}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Bill

// Full Name:
// Displaying the full name of the person or entity associated with the bill can provide quick identification.

// Bill Due Date:
// Include the due date of the bill to let users know when it needs to be paid.

// Bill Amount:
// Show the total amount of the bill so users can immediately see how much they owe.

// Bill Type or Category:
// Categorize the bill (e.g., utility, rent, subscription) so users can understand what the bill is for.

// Bill Status:
// Indicate whether the bill has been paid or is pending.

// Payment History:
// If applicable, show a summary of past payments made toward this bill.

// Bill Description:
// Provide a brief description of the bill to remind users of its purpose.

// Contact Information:
// Include contact details (e.g., phone number, email) for the bill issuer in case users have questions.

// Options for Action:
// Offer buttons or links for users to take actions like making a payment or viewing more details.

// Additional Notes or Comments:
// Allow users to add and view any notes or comments related to the bill.