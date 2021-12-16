import React from 'react'
import ReactDOM from 'react-dom'
import reportWebVitals from './reportWebVitals'
import { UserProvider } from './context/UserContext'
import { AlertProvider } from './context/AlertContext'
import Nav from './Nav'
import Alert from './components/Alert'
import './styles/global.css'
import './styles/index.css'

ReactDOM.render(
  <div className='background'>
    <AlertProvider>
      <UserProvider>
        <Nav />
        <Alert />
      </UserProvider>
    </AlertProvider>
  </div>,
  document.getElementById('root')
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
