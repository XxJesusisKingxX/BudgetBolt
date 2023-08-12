import "./Header.css";
import { FC, ReactNode } from "react";

interface Props {
    children?: ReactNode;
}

const Navigation: FC<Props> = (props: Props) => {
    return (
        <>
            <header className="nav nav--border">
                <span className="nav__logo nav__logo--default">BUDGETBOLT</span>
            </header>
            {props.children}
        </>
    );
};

export default Navigation;
