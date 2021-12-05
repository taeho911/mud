import { useState } from 'react';
import errors from '../errors';

function SignIn() {
  const [err, setErr] = useState(errors.noError);

  const signIn = e => {
    e.preventDefault();
    let formData = new FormData(e.target.form);
    let jsonData = Object.fromEntries(formData.entries());

    fetch('/api/sign/in', {
      method: 'post',
      headers: {'Content-type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsonData)
    }).then(res => {
      console.log(res);
    });
  }

  return (
    <div>
      <h1>Sign In</h1>
      <form>
        <input type='text' name='username' placeholder='Username' /><br />
        <input type='password' name='password' placeholder='Password' /><br />
        <button onClick={signIn}>Submit</button>
        <button type='reset'>Reset</button>
      </form>
      {err.msg != '' && <div className='err'>{`<${err.code}> ${err.msg}`}</div>}
    </div>
  );
}

export default SignIn;
