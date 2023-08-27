import React from 'react';
import BudgetTrender from '../BudgetTrender/BudgetTrender';
import FinancialsPeek from '../FinancialsPeek/FinancialsPeekContainer';
import HealthIndicator from '../HealthIndicator/HealthIndicatorContainer';
import MiniWindow from '../MiniWindow/MiniWindowContainer';
import Upcoming from '../Upcoming/Upcoming';
import './Overview.css';

// Props interface for the Overview component
interface Props {
    user: string; // User's name
    date: string; // Current date
}

const Overview: React.FC<Props> = ({ user, date }) => {
    return (
        <div className='overview'>
            {/* Header section with user, date, and title */}
            <div className='overview__header'>
                <span className='overview__header__user'>Welcome, {user}</span>
                <span className='overview__header__date'>~ Today is {date} ~</span>
                <span className='overview__header__title'>Budget Overview</span>
            </div>
            {/* Widgets section */}
            <div className='overview__widgets'>
                {/* Top widgets */}
                <div className='overview__widgets__top'>
                    {/* Left side top widgets */}
                    <div className='overview__widgets__top__left'>
                        {/* HealthIndicator component */}
                        <HealthIndicator />
                        {/* Upcoming component */}
                        <Upcoming />
                    </div>
                    {/* Right side top widgets */}
                    <div className='overview__widgets__top__right'>
                        {/* FinancialsPeek component */}
                        <FinancialsPeek />
                        {/* MiniWindow component */}
                        <MiniWindow />
                    </div>
                </div>
                {/* Bottom widget */}
                <div className='overview__widgets__bottom'>
                    {/* BudgetTrender component */}
                    <BudgetTrender />
                </div>
            </div>
        </div>
    );
};

export default Overview;
