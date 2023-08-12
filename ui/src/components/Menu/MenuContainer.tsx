import { useEffect, useState } from "react";
import Menu from "./Menu";

const MenuContainer = () => {
    const [showDropdown, setShowDropdown] = useState(false);
    const handleMouseOver = () => {
        setShowDropdown(true);
    };
    const handleMouseOut = () => {
        setShowDropdown(false);
    };
    useEffect(() => {
      const menuArea = document.getElementById('menu');
      if (menuArea ) {
        menuArea.addEventListener('mouseover', handleMouseOver);
      }
      return () => {
        if (menuArea) {
            menuArea.removeEventListener('mouseover', handleMouseOver);
        }
      };
    }, []);
    return (
        <Menu
            onMouseOver={handleMouseOver}
            onMouseOut={handleMouseOut}
            showDropdown={showDropdown}
        />
    );
};

export default MenuContainer;
