import { useState } from 'react'
import MoneyForm from './MoneyForm'

function MoneyList(props) {
  const [displayFlag, setDisplayFlag] = useState(false)

  return (
    <>
      <hr />
      <div className='money-unit-container'>
        <div className='money-unit-subcon-0'>
          <div className='money-unit-subcon-1'>
            <div className='col-date'>{props.money.date.split('T')[0]}</div>
            <div className='col-amount'>{props.money.amount.toLocaleString()}</div>
            <div className='col-summary'>{props.money.summary}</div>
          </div>
          <div className='money-unit-subcon-2 margintop2'>
            {props.money.tags.map((v, i) => {
              return <div key={i} className='tag display-tag'>{v}</div>
            })}
          </div>
        </div>
        <div className='icon-container'>
          <div className='del-icon' onClick={props.funcs.deleteMoney}></div>
          <div className='mod-icon' onClick={e => setDisplayFlag(!displayFlag)}></div>
        </div>
        {displayFlag &&
          <MoneyForm
            date={props.money.date}
            amount={props.money.amount}
            summary={props.money.summary}
            tags={props.money.tags}
            selectedTags={props.money.tags}
            funcs={{setDisplayFlag: setDisplayFlag}}
          />
        }
      </div>
    </>
  )
}

export default MoneyList
