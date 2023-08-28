import './Upcoming.css'
import Bill from '../Bill'

const Upcoming = () => {
    return (
        <div className='upcoming'>
            <span className='upcoming__txt'>Upcoming Bills</span>
            <Bill/>
        </div>
    );
}

export default Upcoming