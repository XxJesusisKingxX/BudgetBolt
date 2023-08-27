import React from 'react';

// Props interface for the PlaidLink component
interface Props {
    plaidFunction: Function; // Function to be called when the button is clicked
    ready: boolean;          // Indicates whether the button is ready to be clicked
}

// PlaidLink component definition
const PlaidLink: React.FC<Props> = ({ plaidFunction, ready }) => {
    return (
        // Button element with conditional disabled state and onClick handler
        <button
            data-testid='plaid-button'      // Data attribute for testing purposes
            className="btn btn--plaid"      // CSS classes for styling
            onClick={() => plaidFunction()} // Click event handler to call the provided function
            disabled={!ready}               // Disabling the button if not ready
        >
            Add Account
        </button>
    );
};

export default PlaidLink;
