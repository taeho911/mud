import { useState, useRef, useEffect, useContext } from 'react'
import { UserContext } from '../context/UserContext'
import MoneyUnit from './MoneyUnit'
import '../styles/money.css'

function Money() {
  const [user, setUser] = useContext(UserContext)
  const [tags, setTags] = useState(['income', 'spend', 'invest', 'life', 'play', 'drink', 'food'])
  const [selectedTags, setSelectedTags] = useState([])
  const [moneyList, setMoneyList] = useState([])
  const [err, setErr] = useState('')
  const tagInput = useRef(undefined)

  const fetchMoneyList = () => {
    fetch('/api/money').then(res => {
      if (res.status === 200) {
        res.json().then(data => setMoneyList(data))
      } else {
        res.text().then(err => setErr(err))
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
    jsondata.tags = selectedTags

    fetch('/api/money/post', {
      method: 'post',
      headers: {'Content-type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsondata)
    }).then(res => {
      if (res.status === 200) {
        fetchMoneyList()
      } else {
        res.text().then(err => setErr(err))
      }
    })
  }

  return (
    <main>
      <h1>Money</h1>
      <form>
        <div>
          <input type='date' name='date' placeholder='Date' />
          <input type='number' step='100' name='amount' placeholder='Amount' />
          <input type='text' name='summary' placeholder='Summary' />
          <input ref={tagInput} className='tag-input' type='text' placeholder='Add custom tag' />
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
        <div>
          <div className='err'>{err}</div>
          {moneyList.map((v, i) => {
            return <MoneyUnit key={i} money={v} />
          })}
        </div>
      </form>
    </main>
  )
}

export default Money
