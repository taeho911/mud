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
    let formData = new FormData(e.target.form)
    let jsonData = Object.fromEntries(formData.entries())

    fetch('/api/sign/in', {
      method: 'post',
      headers: {'Content-type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsonData)
    }).then(res => {
      if (res.status === 200) {
        res.json().then(data => {
          setUser(data)
          navigate('/', {replace: true})
        })
      } else {
        res.text().then(err => setErr(err))
      }
    })
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
      <div className='err'>{err}</div>
    </div>
  )
}

export default SignIn
