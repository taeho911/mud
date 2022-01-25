import { useState, useContext } from 'react'
import { useNavigate } from 'react-router-dom'
import { UserContext } from '../context/UserContext'

function SignIn() {
  const [user, setUser] = useContext(UserContext)
  const [err, setErr] = useState('')
  const navigate = useNavigate()

  const signIn = e => {
    e.preventDefault()
    setErr('')
    let formdata = new FormData(e.target.form)
    let jsondata = Object.fromEntries(formdata.entries())

    fetch('/api/sign/in', {
      method: 'post',
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsondata)
    }).then(res => {
      switch (res.status) {
      case 200:
        res.json().then(data => {
          setUser(data)
          navigate('/', {replace: true})
        })
        break
      default:
        res.text().then(err => setErr(err))
        break
      }
    })
  }

  return (
    <div>
      <h1>Sign In</h1>
      <form className='margintop2'>
        <input type='text' name='username' placeholder='Username' /><br />
        <input type='password' name='password' placeholder='Password' /><br />
        <button onClick={signIn}>Submit</button>
      </form>
      <div className='err'>{err}</div>
    </div>
  )
}

export default SignIn
