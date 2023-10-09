import BillComponent from './components/BillComponent';
import { useContext, useState } from 'react';
import { EndPoint } from '../../constants/endpoints';
import { getDateView } from '../../utils/formatDate';
import AppContext from '../../context/AppContext';
import ThemeContext from '../../context/ThemeContext';
import '../../assets/Loading.css'

// Interface for the shape of a bill
export interface Bills {
  [key: string]: {
    name: string
    total_amount: number
    max_amount: number
    average_amount: number
    due_date: string
    earliest_date_cycle: string
    previous_date_cycle: string
    last_date_cycle: string
    frequency: number
    status: string
    degraded: number
    category: string
  };
}

export const useBill = () => {
    const [bills, setBills] = useState<Bills>({});
    const [isLoading, setLoading] = useState(false);

    const { budgetView } = useContext(AppContext)
    const { mode } = useContext(ThemeContext);

    const getBills = async () => {
        const date = getDateView(new Date(), budgetView)
        const query = `?date=${date}&recurring=${budgetView}`
        setLoading(true)
        try {
            const response = await fetch(EndPoint.GET_TRANSACTIONS + query, {
                method: "GET"
            })
            if (response.ok) {
                const data = await response.json();
                if (data) {
                    setBills(data["bills"]);
                }
                setLoading(false);
            }
        } catch(error) {
            console.log(error);
        }

    };

    const showBills = (loading = isLoading, billList: Bills = bills) => {
        //TODO TEST
        const getCategoryIcon = (category: string) => {
            switch (category) {
                case "Loan Payments":
                    return 'loanpayment'
                case "Rent And Utilities":
                    return 'housing'
                case "Home Improvement":
                    return 'housing'
                case "Government And Non Profit":
                    return 'gov'
                case "Entertainment":
                    return 'entertainment'
                case "Food And Drink":
                    return 'entertainment'
                case "Medical":
                    return 'healthcare'
                case "Personal Care":
                    return 'selfcare'
                case "Transportation":
                    return 'transportation'
                default:
                    return 'misc'
            }
        };

        const handleDaysLeft= (currentDate: Date, date: string) => {
            const targetDate = new Date(date);

            // Calculate the time difference in milliseconds
            const timeDiff = targetDate.getTime() - currentDate.getTime();

            // Convert milliseconds to days
            const daysLeft = Math.ceil(timeDiff / (1000 * 60 * 60 * 24));

            return daysLeft;
        };

        const capitalizeFirstLetterEachWord = (inputString: string) => {
            // Split the input string into words
            const words = inputString.replaceAll("_"," ").split(' ');
          
            // Capitalize the first letter of each word
            const capitalizedWords = words.map((word) => {
              if (word.length === 0) {
                // Skip empty words (e.g., multiple spaces)
                return word;
              }
              // Capitalize the first letter and keep the rest in lowercase
              return word.charAt(0).toUpperCase() + word.slice(1).toLowerCase();
            });
          
            // Join the words back together with spaces
            const resultString = capitalizedWords.join(' ');
          
            return resultString;
        }
        console.log(billList)
        let billCount = 0; // Track amount of bills added
        const maxBillsShown = 9; // max amount of bills to show
        const maxChar = 4; // max amount characters to show for bill name
        const loadingIcon = `/images/${mode}/loading.png`;
        const rows = Object.keys(billList)
        .map((name: any, idx) => {
            // Formatting
            const bill = billList[name]
            const billName = bill.name.length > maxChar ? bill.name.slice(0, maxChar) + "*" : bill.name
            const billShortName = capitalizeFirstLetterEachWord(billName);
            const billAvgAmount = bill.average_amount.toFixed(2);
            const billCategoryName = capitalizeFirstLetterEachWord(bill.category);
            const billCategory = getCategoryIcon(billCategoryName);
            const daysLeft = handleDaysLeft(new Date(), bill.due_date);
            if (daysLeft > 0 && bill.status == "MATURE" && billCount < maxBillsShown) {
                billCount += 1
                return (
                    <BillComponent
                        key={idx}
                        shortName={billShortName}
                        name={bill.name}
                        price={billAvgAmount}
                        dueDate={bill.due_date}
                        categoryName={billCategoryName}
                        category={billCategory}
                    />
                );
            }
        });
        return (
            <>
                {!loading && billList ? (
                  billCount === 0 ? (
                    <span className='nobill'>No bills to show</span>
                  ) : (
                    rows
                  )
                ) : (
                <img className='loading' src={loadingIcon} alt="Loading" />
                )}
            </>
        );
    }

    return {
        showBills,
        getBills,
    };
};