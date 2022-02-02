import { useState } from 'react'
import MoneyPutForm from './MoneyPutForm'

function MoneyUnit(props) {
  const [formSwitch, setFormSwitch] = useState(false)

  return (
    <div className='money-unit-container'>
      <div className='money-unit-subcon-0'>
        <div className='money-unit-subcon-1'>
          <div className='col-date'>{props.money.date.split('T')[0]}</div>
          <div className={`col-amount ${props.money.amount < 0 ? 'red':'blue'}`}>{props.money.amount.toLocaleString()}</div>
        </div>
        <div className='col-summary'>{props.money.summary}</div>
        <div className='money-unit-subcon-2'>
          {props.money.tags.map((v, i) => {
            return <div key={i} className='tag display-tag'>{v}</div>
          })}
        </div>
      </div>
      <div className='icon-container'>
        <div className='del-icon' onClick={e => props.deleteMoney(props.money._id)}></div>
        <div className='mod-icon' onClick={e => setFormSwitch(!formSwitch)}></div>
      </div>
      {formSwitch &&
        <MoneyPutForm money={props.money}
          moneyList={props.moneyList}
          setFormSwitch={setFormSwitch}
          setMoneyList={props.setMoneyList} />
      }
    </div>
  )
}

export default MoneyUnit
