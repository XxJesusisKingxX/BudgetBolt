import './AppConfig.css';
import Header from './components/header/Header'
import MenuBar from './components/menubar/MenuBar';
import Home from './pages/home/Home'

function App() {
  return (
    <>
      <Header>
        <MenuBar></MenuBar>
        <Home></Home>
      </Header>
    </>
  );
}

export default App;