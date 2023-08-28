import { useEffect, useState } from 'react';
import Menu from './Menu';

const MenuContainer = () => {
    // State to control whether the dropdown menu is shown
    const [showDropdown, setShowDropdown] = useState(false);

    // Event handler for mouse over event
    const handleMouseOver = () => {
        setShowDropdown(true);
    };

    // Event handler for mouse out event
    const handleMouseOut = () => {
        setShowDropdown(false);
    };

    useEffect(() => {
        // Get the menu area element by its ID
        const menuArea = document.getElementById('menu');

        // Add a 'mouseover' event listener to the menu area
        if (menuArea) {
            menuArea.addEventListener('mouseover', handleMouseOver);
        }

        // Cleanup: Remove the event listener when the component unmounts
        return () => {
            if (menuArea) {
                menuArea.removeEventListener('mouseover', handleMouseOver);
            }
        };
    }, []);

    return (
        // Render the Menu component and pass event handlers and state as props
        <Menu
            onMouseOver={handleMouseOver}
            onMouseOut={handleMouseOut}
            showDropdown={showDropdown}
        />
    );
};

export default MenuContainer;
