import Bill from './components/Bill';
import '../../assets/view/styles/billtracker/View.css'

const BillTracker = () => {
    return (
        <div className='billtracker'>
            <span className='billtracker__txt'>Upcoming Bills</span>
            <Bill/>
        </div>
    );
}

export default BillTracker