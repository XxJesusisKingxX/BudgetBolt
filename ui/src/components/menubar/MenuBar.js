import './MenuBar.css'

function MenuBar() {
    return (
        <div className="menubar">
            <div className="top_line"></div>
            <div className="middle_line"></div>
            <div className="bottom_line"></div>
            <div className="hidden_text">Menu</div>
            <div className="hidden_menu">
                <ul><a href="">Budget Overview</a></ul>
                <ul><a href="">Transactions</a></ul>
                <ul><a href="">Expenses</a></ul>
                <ul><a href="">Income</a></ul>
                <ul><a href="">Goals and Savings</a></ul>
                <ul><a href="">Reminders and Alerts</a></ul>
                <ul><a href="">Settings</a></ul>
                <ul><a href="">Help and Support</a></ul>
            </div>
        </div>
    );
}
  
export default MenuBar;
