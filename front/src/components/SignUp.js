import { useState } from 'react';

function SignUp(props) {
  const [err, setErr] = useState('');

  const signUp = e => {
    e.preventDefault();
    setErr('');
    let formdata = new FormData(e.target.form);
    let jsondata = Object.fromEntries(formdata.entries());

    if (jsondata.username === '' || jsondata.password === '') {
      setErr('Fill out username and password');
      return;
    }

    if (jsondata.password !== jsondata.password_confirm) {
      setErr('Check password confirmation');
      return;
    }

    fetch('/api/sign/up', {
      method: 'post',
      headers: {'Content-type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsondata)
    }).then(res => {
      if (res.status === 200) {
        alert('Sign up succeed');
        props.setIsSignIn(true);
      } else {
        res.text().then(err => setErr(err));
      }
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
      <div className='err'>{err}</div>
    </div>
  );
}

export default SignUp;
