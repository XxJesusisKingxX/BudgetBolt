import React from 'react';
import BudgetTrender from '../Charts/Chart';
import FinancialsSnapshot from '../FinancialsSnapshot';
import HealthIndicator from '../HealthIndicator';
import BudgetWindow from '../BudgetWindow';
import BillTracker from '../BillTracker';
import './Dashboard.css';

interface Props {
    user: string; // User's name
    date: string; // Current date
}

const DashboardComponent: React.FC<Props> = ({ user, date }) => {
    return (
        <div className='dashboard'>
            {/* Header section with user, date, and title */}
            <div className='dashboard__header'>
                <span className='dashboard__header__user'>Welcome, {user}</span>
                <span data-testid='dashboard-date' className='dashboard__header__date'>~ Today is {date} ~</span>
                <span className='dashboard__header__title'>Budget Dashboard</span>
            </div>
            {/* Widgets section */}
            <div className='dashboard__widgets'>
                {/* Top widgets */}
                <div className='dashboard__widgets__top'>
                    {/* Left side top widgets */}
                    <div className='dashboard__widgets__top__left'>
                        {/* HealthIndicator component */}
                        <HealthIndicator />
                        {/* BillTracker component */}
                        <BillTracker />
                    </div>
                    {/* Right side top widgets */}
                    <div className='dashboard__widgets__top__right'>
                        {/* FinancialsSnapshot component */}
                        <FinancialsSnapshot />
                        {/* BudgetWindow component */}
                        <BudgetWindow />
                    </div>
                </div>
                {/* Bottom widget */}
                <div className='dashboard__widgets__bottom'>
                    {/* BudgetTrender component */}
                    {/* <BudgetTrender /> */}
                </div>
            </div>
        </div>
    );
};

export default DashboardComponent;
