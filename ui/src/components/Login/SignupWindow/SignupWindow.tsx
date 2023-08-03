import './SignupWindow.css';
import cancel from '../icons/cancel-dark.png'
import loading from '../../../images/loading.png'
import { ChangeEventHandler, FormEventHandler } from 'react';

interface Props {
    close: Function
    error: boolean
    isLoading: boolean
    username: string
    password: string
    userChange: ChangeEventHandler<HTMLInputElement>
    passChange: ChangeEventHandler<HTMLInputElement>
    signup: FormEventHandler<HTMLFormElement>
};

const SignupWindow: React.FC<Props> = ({ close, error, isLoading, username, password, userChange, passChange, signup, }) => {
    return (
        <>
            <div className="create-window-container">
                <h1 className="create-window-title">Create a Account<img className="create-window-cancel" src={cancel} onClick={() => close()}/></h1>
                <form onSubmit={signup}>
                    <input id="create-username" value={username} onChange={userChange} placeholder="User" required/>
                    <input id="create-password" type="password" value={password} onChange={passChange} placeholder="Password" required/>
                    <br/>
                    {!isLoading ? <button type="submit" className="create-submit-button">Submit</button> : <img src={loading} className="signup-loading"/>}
                    {error ? <div>{error}</div> : null}
                </form>
            </div>
        </>
    );
};

export default SignupWindow;