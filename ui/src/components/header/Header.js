import './Header.css'

function Header(props) {
    return (
        <>
            <div className="title">BUDGETBOLT</div>
            <div className="border_line"></div>
            {props.children}
        </>
    );
}
  
export default Header;