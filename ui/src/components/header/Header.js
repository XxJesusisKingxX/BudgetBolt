import "./Header.css";
import Navigation from "./Navigation/Navigation";

const Header = (props) => {
    return (
        <>
            <div className="title">BUDGETBOLT</div>
            <div className="border_line"></div>
            <Navigation />
            {props.children}
        </>
    );
};
  
export default Header;