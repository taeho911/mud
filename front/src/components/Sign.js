import React from 'react';
import SignIn from './SignIn';
import SignUp from './SignUp';

function Sign() {
  const [isSignIn, setIsSignIn] = React.useState(true);

  return (
    <div className='container'>
      <div>
        <button onClick={e => setIsSignIn(true)}>Sign in</button>
        <button onClick={e => setIsSignIn(false)}>Sign up</button>
      </div>
      {isSignIn
        ? <SignIn />
        : <SignUp />
      }
    </div>
  );
}

export default Sign;
