import './Home.css';
import Sideview from '../../components/RecentTransactions/Sideview';
import Overview from '../../components/Overview/OverviewContainer';

const Home = () => {
    return (
        <>
            <Sideview/>
            <Overview/>
        </>
    );
};
  
export default Home;