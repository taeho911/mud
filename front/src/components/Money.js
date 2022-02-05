import { useState, useEffect, useContext } from 'react'
import { useNavigate } from 'react-router-dom'
import { UserContext } from '../context/UserContext'
import MoneyPostForm from './MoneyPostForm'
import MoneyUnit from './MoneyUnit'
import MoneyStat from './MoneyStat'
import MoneyAnal from './MoneyAnal'
import '../styles/money.css'
import MoneySetting from './MoneySetting'

function Money() {
  const today = new Date()
  const yearMonthRegex = new RegExp('^[0-9]{4}-(1?[0-2]|0?[1-9])$')

  const [user, setUser] = useContext(UserContext)
  const [yearMonth, setYearMonth] = useState(today.toISOString().slice(0, 7))
  const [moneyList, setMoneyList] = useState([])
  const [addSwitch, setAddSwitch] = useState(false)
  const [screenSwitch, setScreenSwitch] = useState(0)
  const [err, setErr] = useState('')
  const navigate = useNavigate()

  const fetchMoneyListByMonth = month => {
    let splited = month.split('-')
    fetch(`/api/money/get?year=${parseInt(splited[0])}&month=${parseInt(splited[1])}&count=1`).then(res => {
      switch (res.status) {
      case 200:
        res.json().then(data => setMoneyList(data))
        break
      case 401:
        setUser(undefined)
        navigate('/', {replace: true})
        break
      default:
        res.text().then(err => setErr(err))
        break
      }
    })
  }

  useEffect(() => fetchMoneyListByMonth(yearMonth), [yearMonth])

  const deleteMoney = _id => {
    fetch('/api/money/delete', {
      method: 'delete',
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
      body: JSON.stringify({_id: _id})
    }).then(res => {
      switch (res.status) {
        case 200:
          setMoneyList(moneyList.filter((v, i) => v._id !== _id))
          break
        case 401:
          setUser(undefined)
          navigate('/', {replace: true})
          break
        default:
          res.text().then(err => setErr(err))
          break
        }
    })
  }

  const changeYearMonth = e => {
    e.preventDefault()
    setErr('')
    if (yearMonthRegex.test(e.target.value)) {
      setYearMonth(e.target.value)
    }
  }

  return (
    <main>
      <h1>Money</h1>

      <div>
        <div className='margintop2'>
          <div className={`common-icon add-icon ${addSwitch ? 'icon-active' : ''}`}
            onClick={e => setAddSwitch(!addSwitch)}></div>
          <div className={`common-icon list-icon ${screenSwitch === 0 ? 'icon-active' : ''}`}
            onClick={e => setScreenSwitch(0)}></div>
          <div className={`common-icon stat-icon ${screenSwitch === 1 ? 'icon-active' : ''}`}
            onClick={e => setScreenSwitch(1)}></div>
          <div className={`common-icon anal-icon ${screenSwitch === 2 ? 'icon-active' : ''}`}
            onClick={e => setScreenSwitch(2)}></div>
          <div className={`common-icon setting-icon ${screenSwitch === 3 ? 'icon-active' : ''}`}
            onClick={e => setScreenSwitch(3)}></div>
        </div>

        {addSwitch &&
          <MoneyPostForm
            moneyList={moneyList}
            setMoneyList={setMoneyList}
            setErr={setErr}
            yearMonth={yearMonth} />
        }
      </div>

      <div className='err margintop2'>{err}</div>

      <div>
        <input className='money-month' type='month' defaultValue={yearMonth} onChange={changeYearMonth}></input>
      </div>

      <div>
        {screenSwitch === 0 && moneyList.map((v, i) => {
          return <MoneyUnit key={i} money={v}
            moneyList={moneyList}
            deleteMoney={deleteMoney}
            setMoneyList={setMoneyList} />
        })}
        {screenSwitch === 1 &&
          <MoneyStat moneyList={moneyList} />
        }
        {screenSwitch === 2 &&
          <MoneyAnal />
        }
        {screenSwitch === 3 &&
          <MoneySetting />
        }
      </div>
    </main>
  )
}

export default Money
