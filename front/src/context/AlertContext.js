import { useState, createContext } from 'react'

export const AlertContext = createContext()

export const AlertProvider = props => {
  const [alertMsg, setAlertMsg] = useState(undefined)
  return (
    <AlertContext.Provider value={[alertMsg, setAlertMsg]}>
      {props.children}
    </AlertContext.Provider>
  )
}
