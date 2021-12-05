import { useState } from 'react';
import errors from '../errors';

function SignUp() {
  const [err, setErr] = useState(errors.noError);

  const signUp = e => {
    e.preventDefault();
    let formData = new FormData(e.target.form);
    let jsonData = Object.fromEntries(formData.entries());

    if (jsonData.username === '' || jsonData.password === '') {
      setErr(errors.emptyField);
      return
    }

    if (jsonData.password !== jsonData.password_confirm) {
      setErr(errors.wrongPwdConfirm);
      return
    }

    fetch('/api/sign/up', {
      method: 'post',
      headers: {'Content-type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsonData)
    }).then(res => {
      console.log(res);
    });
  }

  return (
    <div>
      <h1>Sign Up</h1>
      <form>
        <input type='text' name='username' placeholder='Username' /><br />
        <input type='password' name='password' placeholder='Password' /><br />
        <input type='password' name='password_confirm' placeholder='Password Confirm' /><br />
        <button onClick={signUp}>Submit</button>
        <button type='reset'>Reset</button>
      </form>
      {err.msg !== '' && <div className='err'>{`<${err.code}> ${err.msg}`}</div>}
    </div>
  );
}

export default SignUp;
