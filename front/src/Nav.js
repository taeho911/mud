import React, { useContext } from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import { UserContext } from './context/UserContext';
import App from './App';
import Sign from './components/Sign';
import Account from './components/Account';
import Money from './components/Money';
import Test from './components/Test';

function Nav() {
  const [signed, setSigned] = useContext(UserContext);
  
  return (
    <BrowserRouter>
      <nav className='nav'>
        <div className='nav-container'>
          <Link to='/'>Mud</Link>
          {!signed && 
            <Link to='sign'>Sign</Link>
          }
          {signed &&
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
          {!signed && 
            <Route path='sign' element={<Sign />} />
          }
          {signed &&
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
