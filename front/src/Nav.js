import React, { useContext, useEffect } from 'react'
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom'
import { UserContext } from './context/UserContext'
import App from './App'
import UserMenu from './components/UserMenu'
import Sign from './components/Sign'
import Account from './components/Account'
import Money from './components/Money'
import './styles/nav.css'

function Nav() {
  const [user, setUser] = useContext(UserContext)

  // SPA의 유저 세션 유지를 위한 코드
  useEffect(() => {
    fetch('/api/sign/confirm').then(res => {
      if (res.status === 200) res.json().then(data => setUser(data))
    })
  }, [])

  return (
    <BrowserRouter>
      <details className='menu'>
        <summary className='dropdown-summary'>Menu</summary>
        <div className='dropdown-menu'>
          {user &&
            <UserMenu />
          }
          <Link to='/'>Home</Link>
          {!user && 
            <Link to='sign'>Sign</Link>
          }
          {user && <>
            <Link to='account'>Account</Link>
            <Link to='money'>Money</Link>
          </>}
        </div>
      </details>
      <Routes>
        <Route path='/' element={<App />} />
        {!user && 
          <Route path='sign' element={<Sign />} />
        }
        {user && <>
          <Route path='account' element={<Account />} />
          <Route path='money' element={<Money />} />
        </>}
      </Routes>
    </BrowserRouter>
  )
}

export default Nav
