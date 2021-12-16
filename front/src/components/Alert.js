import { useContext } from 'react'
import { AlertContext } from '../context/AlertContext'

function Alert() {
  const [alertMsg, setAlertMsg] = useContext(AlertContext)
  return (
    <>
      {alertMsg &&
        <>
          <div className='overlay' />
          <div className='alert-box'>
            <div className='alert'>
              <div className='alert-msg'>{alertMsg}</div>
              <button onClick={e => setAlertMsg(undefined)}>OK</button>
            </div>
          </div>
        </>
      }
    </>
  )
}

export default Alert
