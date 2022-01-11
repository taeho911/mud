import { useState, useRef, useEffect, useContext } from 'react'
import { useNavigate } from 'react-router-dom'
import { UserContext } from '../context/UserContext'
import MoneyUnit from './MoneyUnit'
import '../styles/money.css'

function Money() {
  let today = new Date()
  const monthList = ['All']

  const [user, setUser] = useContext(UserContext)
  const [tags, setTags] = useState(['income', 'spend', 'invest', 'life', 'play', 'drink', 'food'])
  const [selectedTags, setSelectedTags] = useState([])
  const [moneyList, setMoneyList] = useState([])
  const [addSwitch, setAddSwitch] = useState(false)
  const [statSwitch, setStatSwitch] = useState(false)
  const [monthSwitch, setMonthSwitch] = useState()
  const [err, setErr] = useState('')
  const tagInput = useRef(undefined)
  const navigate = useNavigate()


  const fetchMoneyList = () => {
    fetch('/api/money/get').then(res => {
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

  const fetchMoneyListByMonth = (year, month) => {
    fetch(`/api/money/get?year=${year}&month=${month}`).then(res => {
      console.log(res.status)
    })
  }

  useEffect(fetchMoneyList, [])

  const addTags = e => {
    e.preventDefault()
    let tag = tagInput.current.value
    if (tag.length > 0 && !tags.includes(tag)) {
      tags.push(tag)
      setTags([...tags])

      selectedTags.push(tag)
      setSelectedTags([...selectedTags])
      tagInput.current.value = ''
    }
  }

  const updateSelectedTags = tag => {
    if (selectedTags.includes(tag)) {
      setSelectedTags(selectedTags.filter(item => {return item !== tag}))
    } else {
      selectedTags.push(tag)
      setSelectedTags([...selectedTags])
    }
  }

  const postMoney = e => {
    e.preventDefault()
    let formdata = new FormData(e.target.form)
    let jsondata = Object.fromEntries(formdata.entries())

    // 'yyyy-mm-dd' 형식의 날짜는 Go에서 unmarshal시 제대로 파싱되지 않기 때문에 아래의 코드를 추가
    jsondata.date = new Date(jsondata.date)
    // type='number'는 string을 반환하기 때문에 명시적으로 숫자형으로 파싱해주는 작업 필요
    jsondata.amount = parseFloat(jsondata.amount)
    jsondata.tags = selectedTags

    fetch('/api/money/post', {
      method: 'post',
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsondata)
    }).then(res => {
      switch(res.status) {
        case 200:
          e.target.form.reset()
          res.json().then(data => {
            moneyList.push(data)
            setMoneyList([...moneyList.sort((a, b) => {
              if (a.date < b.date) return 1
              if (a.date > b.date) return -1
              return 0
            })])
          })
          break
        default:
          res.text().then(err => setErr(err))
          break
      }
    })
  }

  const deleteMoney = _id => {
    fetch('/api/money/delete', {
      method: 'delete',
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
      body: JSON.stringify({_id: _id})
    }).then(res => {
      switch (res.status) {
        case 200:
          setMoneyList(moneyList.filter((v, i) => v._id != _id))
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

  return (
    <main>
      <h1>Money</h1>

      <div>
        <div className='margintop2'>
          <div className={`add-icon ${addSwitch ? 'add-icon-active' : ''}`}
            onClick={e => setAddSwitch(!addSwitch)}></div>
          <div className={`stat-icon ${statSwitch ? 'stat-icon-active' : ''}`}
            onClick={e => setStatSwitch(!statSwitch)}></div>
        </div>

        <form className={`add-container ${addSwitch ? '' : 'display-none'}`}>
          <div>
            <input type='date' name='date' placeholder='Date' /><br />
            <input type='number' step='100' name='amount' placeholder='Amount' /><br />
            <input type='text' name='summary' placeholder='Summary' /><br />
            <input ref={tagInput} type='text' placeholder='Add custom tag' /><br />
          </div>
          <div>
            <button onClick={addTags}>Add Tag</button>
            <button onClick={postMoney}>Submit</button>
          </div>
          <div>
            {tags.map((tag, i) => {
              return <span key={i} 
                className={`tag ${selectedTags.includes(tag) ? 'selected-tag' : ''}`}
                onClick={e => updateSelectedTags(tag)}
                onDoubleClick={e => setTags(tags.filter(item => {return item !== tag}))}>{tag}</span>
            })}
          </div>
        </form>
      </div>

      <div className='err margintop2'>{err}</div>

      <div>
        {["All", 1,2,3,4,5,6,7,8,9,10,11,12].map(month => {
          return <button className='btn-month' onClick={e => console.log(month)}>{month}</button>
        })}
      </div>

      {/* https://www.npmjs.com/package/react-chartjs-2 */}

      <div>
        {!statSwitch && moneyList.map((v, i) => {
          return <MoneyUnit key={i} money={v}
            moneyList={moneyList}
            deleteMoney={deleteMoney}
            setMoneyList={setMoneyList} />
        })}
      </div>
    </main>
  )
}

export default Money
