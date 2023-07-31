import './SignupWindow.css';
import cancel from '../icons/cancel-dark.png'

interface Props {
    close: Function
};

const SignupWindow: React.FC<Props> = ({ close }) => {
    return (
        <>
            <div className="create-window-container">
                <h1 className="create-window-title">Create a Account<img className="create-window-cancel" src={cancel} onClick={() => close()}/></h1>
                <input id="create-username" type="text" placeholder="User"/>
                <input id="create-password" type="text" placeholder="Password"/>
                <br/>
                <button className="create-submit-button">Submit</button>
            </div>
        </>
    );
};

export default SignupWindow;