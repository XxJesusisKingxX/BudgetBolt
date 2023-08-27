import { FC } from "react";

// Props interface for the Refresh component
interface Props {
    mode: string;       // The mode for determining image path
    isRefresh: boolean; // Indicates whether the refresh button is in a "refreshing" state
    refresh: Function;  // Function to be called when the button is clicked
}

// Refresh component definition
const Refresh: FC<Props> = ({ mode, isRefresh, refresh }) => {
    // Dynamic image path based on the provided mode
    const refreshIcon = `/images/${mode}/refresh.png`;

    return (
        // Image element for the refresh button
        <img
            className={isRefresh ? "sidebar__refresh sidebar__refresh--load" : "sidebar__refresh sidebar__refresh--loadalt"}
            onClick={() => refresh()} // Click event handler to call the provided refresh function
            src={refreshIcon}         // Source of the refresh icon image
            alt="Refresh"             // Alternative text for accessibility
        />
    );
};

export default Refresh;
