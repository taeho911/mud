import React, { useContext, useEffect } from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import { UserContext } from './context/UserContext';
import App from './App';
import Sign from './components/Sign';
import Account from './components/Account';
import Money from './components/Money';
import Test from './components/Test';

function Nav() {
  const [user, setUser] = useContext(UserContext);

  // SPA의 유저 세션 유지를 위한 코드
  useEffect(() => {
    fetch('/api/sign/confirm').then(res => {
      if (res.status === 200) res.json().then(data => setUser(data));
    });
  }, []);
  
  return (
    <BrowserRouter>
      <nav className='nav'>
        <div className='nav-container'>
          <Link to='/'>Mud</Link>
          {!user && 
            <Link to='sign'>Sign</Link>
          }
          {user &&
            <React.Fragment>
              <Link to='account'>Account</Link>
              <Link to='money'>Money</Link>
            </React.Fragment>
          }
          <Link to='test'>Test</Link>
        </div>
      </nav>
      <div className='wrapper'>
        <Routes>
          <Route path='/' element={<App />} />
          {!user && 
            <Route path='sign' element={<Sign />} />
          }
          {user &&
            <React.Fragment>
              <Route path='account' element={<Account />} />
              <Route path='money' element={<Money />} />
            </React.Fragment>
          }
          <Route path='test' element={<Test />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default Nav;
