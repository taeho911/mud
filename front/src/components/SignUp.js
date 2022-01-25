import { useState, useContext } from 'react'
import { AlertContext } from '../context/AlertContext'

function SignUp(props) {
  const [err, setErr] = useState('')
  const [alertMsg, setAlertMsg] = useContext(AlertContext)

  const signUp = e => {
    e.preventDefault()
    setErr('')
    let formdata = new FormData(e.target.form)
    let jsondata = Object.fromEntries(formdata.entries())

    if (jsondata.username === '' || jsondata.password === '') {
      setErr('Fill out username and password')
      return
    }

    if (jsondata.password !== jsondata.password_confirm) {
      setErr('Check password confirmation')
      return
    }

    fetch('/api/sign/up', {
      method: 'post',
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsondata)
    }).then(res => {
      switch (res.status) {
      case 200:
        setAlertMsg('Sign up succeed')
        props.setIsSignIn(true)
        break
      default:
        res.text().then(err => setErr(err))
        break
      }
    })
  }

  return (
    <div>
      <h1>Sign Up</h1>
      <form className='margintop2'>
        <input type='text' name='username' placeholder='Username' /><br />
        <input type='password' name='password' placeholder='Password' /><br />
        <input type='password' name='password_confirm' placeholder='Password Confirm' /><br />
        <button onClick={signUp}>Submit</button>
      </form>
      <div className='err'>{err}</div>
    </div>
  )
}

export default SignUp
