import { useState } from 'react'
import MoneyForm from './MoneyForm'

function MoneyUnit(props) {
  const [formSwitch, setFormSwitch] = useState(false)

  return (
    <div className='money-unit-container'>
      <div className='money-unit-subcon-0'>
        <div className='money-unit-subcon-1'>
          <div className='col-date'>{props.money.date.split('T')[0]}</div>
          <div className={`col-amount ${props.money.amount < 0 ? 'red':'blue'}`}>{props.money.amount.toLocaleString()}</div>
          <div className='col-summary'>{props.money.summary}</div>
        </div>
        <div className='money-unit-subcon-2 margintop1'>
          {props.money.tags.map((v, i) => {
            return <div key={i} className='tag display-tag'>{v}</div>
          })}
        </div>
      </div>
      <div className='icon-container'>
        <div className='del-icon' onClick={e => props.funcs.deleteMoney(props.money._id)}></div>
        <div className='mod-icon' onClick={e => setFormSwitch(!formSwitch)}></div>
        <button onClick={e => console.log(props.money)}>Print</button>
      </div>
      {formSwitch &&
        <MoneyForm money={props.money} funcs={{setFormSwitch: setFormSwitch}} />
      }
    </div>
  )
}

export default MoneyUnit
