import { useState, useRef, useEffect, useContext } from 'react'
import { useNavigate } from 'react-router-dom'
import { UserContext } from '../context/UserContext'
import MoneyUnit from './MoneyUnit'
import '../styles/money.css'

function Money() {
  const [user, setUser] = useContext(UserContext)
  const [tags, setTags] = useState(['income', 'spend', 'invest', 'life', 'play', 'drink', 'food'])
  const [selectedTags, setSelectedTags] = useState([])
  const [moneyList, setMoneyList] = useState([])
  const [addSwitch, setAddSwitch] = useState(false)
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
      if (res.status === 200) {
        e.target.form.reset()
        fetchMoneyList()
      } else {
        res.text().then(err => setErr(err))
      }
    })
  }

  const deleteMoney = e => {
    e.preventDefault()
    let formdata = new FormData(e.target.form)
    let jsondata = Object.fromEntries(formdata.entries())
    console.log(`_id: ${jsondata._id}`)
  }

  return (
    <main>
      <h1>Money</h1>
      <div className='add-icon-container'>
        <div className='add-icon' onClick={e => setAddSwitch(!addSwitch)}></div>
        <form className={`add-container ${addSwitch ? '' : 'display-none'}`}>
          <div>
            <input type='date' name='date' placeholder='Date' /><br />
            <input type='number' step='100' name='amount' placeholder='Amount' /><br />
            <input type='text' name='summary' placeholder='Summary' /><br />
            <input ref={tagInput} type='text' placeholder='Add custom tag' /><br />
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
      <div>
        <div className='err margintop2'>{err}</div>
        {moneyList.map((v, i) => {
          return <MoneyUnit key={i} money={v} funcs={{deleteMoney: deleteMoney}}/>
        })}
      </div>
    </main>
  )
}

export default Money
