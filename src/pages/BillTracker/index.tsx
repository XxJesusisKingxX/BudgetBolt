import Bill from './components/Bill';
import '../../assets/Bill.css'

const BillTracker = () => {
    return (
        <div className='billtracker'>
            <span className='billtracker__txt'>Upcoming Bills</span>
            <Bill/>
        </div>
    );
}

export default BillTracker