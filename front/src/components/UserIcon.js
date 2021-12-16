import { useState, useContext } from 'react'
import { useNavigate } from 'react-router-dom'
import { UserContext } from '../context/UserContext'
import { AlertContext } from '../context/AlertContext'
import '../styles/userIcon.css'

function UserIcon() {
  const [promptMsg, setPromptMsg] = useState(undefined)
  const [user, setUser] = useContext(UserContext)
  const [alertMsg, setAlertMsg] = useContext(AlertContext)
  const navigate = useNavigate()

  const signOut = e => {
    e.preventDefault()
    fetch('/api/sign/out').then(res => {
      if (res.status === 200) {
        setUser(undefined)
        navigate('/', {replace: true})
      } else {
        res.text().then(err => alert(err))
      }
    })
  }

  const deleteAccount = e => {
    e.preventDefault()
    let formdata = new FormData(e.target.form)
    let jsondata = Object.fromEntries(formdata.entries())
    jsondata.username = user.username

    fetch('/api/sign/delete', {
      method: 'delete',
      headers: {'Content-type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsondata)
    }).then(res => {
      if (res.status === 200) {
        setUser(undefined)
        setPromptMsg(undefined)
        navigate('/', {replace: true})
      } else {
        res.text().then(err => setAlertMsg(err))
      }
    })
  }

  const closePrompt = e => {
    e.preventDefault()
    setPromptMsg(undefined)
  }

  return (
    <>
      <details>
        <summary className='user-icon'>
          <h3 className='user-alphabet'>{user.username.substring(0, 1)}</h3>
        </summary>
        <div className='dropdown-menu user-dropdown-menu'>
          <div className='dropdown-item'>Name</div>
          <div className='dropdown-divider'></div>
          <div className='dropdown-item dropdown-action-item' onClick={signOut}>Sign out</div>
          <div className='dropdown-item dropdown-action-item' onClick={e => setPromptMsg('Input your password')}>Delete Account</div>
        </div>
      </details>
      {promptMsg &&
        <>
          <div className='overlay' />
          <div className='prompt-box'>
            <div className='prompt'>
              <form>
                <div className='prompt-msg'>{promptMsg}</div>
                <div className='margintop2'>
                  <input type='password' name='password' placeholder='Password' />
                </div>
                <button onClick={deleteAccount}>Submit</button>
                <button onClick={closePrompt}>Cancel</button>
              </form>
            </div>
          </div>
        </>
      }
    </>
  )
}

export default UserIcon
