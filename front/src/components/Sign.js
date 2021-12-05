import { useState } from 'react';
import SignIn from './SignIn';
import SignUp from './SignUp';
import '../styles/sign.css';

function Sign() {
  const [isSignIn, setIsSignIn] = useState(true);

  return (
    <main className='sign-container'>
      <div className='sign-tab-box'>
        <span
          className={isSignIn ? 'sign-tab-active' : 'sign-tab'}
          onClick={e => setIsSignIn(true)}>Sign in</span>
        <span
          className={isSignIn ? 'sign-tab' : 'sign-tab-active'}
          onClick={e => setIsSignIn(false)}>Sign up</span>
      </div>
      {isSignIn
        ? <SignIn />
        : <SignUp />
      }
    </main>
  );
}

export default Sign;
