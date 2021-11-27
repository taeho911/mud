import React from 'react';
import ReactDOM from 'react-dom';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import App from './App';
import Sign from './components/Sign';
import Account from './components/Account';
import './styles/global.css'
import './styles/index.css';

ReactDOM.render(
  <div className='background'>
    <BrowserRouter>
      <nav className='nav'>
        <div className='nav-container'>
          <Link to='/'>Home</Link>
          <Link to='sign'>Sign</Link>
          <Link to='account'>Account</Link>
        </div>
      </nav>
      <div className='wrapper'>
        <Routes>
          <Route path='/' element={<App />} />
          <Route path='sign' element={<Sign />} />
          <Route path='account' element={<Account />} />
        </Routes>
      </div>
    </BrowserRouter>
  </div>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
