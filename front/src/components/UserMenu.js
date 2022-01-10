import { useState, useContext } from 'react'
import { useNavigate } from 'react-router-dom'
import { UserContext } from '../context/UserContext'
import { AlertContext } from '../context/AlertContext'

function UserMenu() {
  const [promptMsg, setPromptMsg] = useState(undefined)
  const [user, setUser] = useContext(UserContext)
  const [alertMsg, setAlertMsg] = useContext(AlertContext)
  const navigate = useNavigate()

  const signOut = e => {
    e.preventDefault()
    fetch('/api/sign/out').then(res => {
      switch (res.status) {
      case 200:
      case 401:
        setUser(undefined)
        navigate('/', {replace: true})
        break
      default:
        res.text().then(err => setAlertMsg(err))
        break
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
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsondata)
    }).then(res => {
      switch (res.status) {
      case 200:
      case 401:
        setUser(undefined)
        setPromptMsg(undefined)
        navigate('/', {replace: true})
        break
      default:
        res.text().then(err => setAlertMsg(err))
        break
      }
    })
  }

  const closePrompt = e => {
    e.preventDefault()
    setPromptMsg(undefined)
  }

  return (
    <>
      <div className='dropdown-item'>{user.username}</div>
      <div className='dropdown-divider'></div>
      <div className='dropdown-action-item' onClick={signOut}>Sign out</div>
      <div className='dropdown-action-item' onClick={e => setPromptMsg('Input your password')}>Delete Account</div>
      <div className='dropdown-divider'></div>
      
      {promptMsg && <>
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
      </>}
    </>
  )
}

export default UserMenu
