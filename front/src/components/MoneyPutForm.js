import { useState, useRef } from 'react'

function MoneyPutForm(props) {
  const [tags, setTags] = useState(props.money.tags)
  const [selectedTags, setSelectedTags] = useState(props.money.tags)
  const [err, setErr] = useState('')
  const tagInput = useRef(undefined)

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

  const replaceMoneyListElem = (moneyList, data) => {
    let i = moneyList.findIndex(money => money._id === data._id)
    moneyList[i] = data
    return moneyList
  }

  const putMoney = e => {
    e.preventDefault()
    let formdata = new FormData(e.target.form)
    let jsondata = Object.fromEntries(formdata.entries())

    jsondata.date = new Date(jsondata.date)
    jsondata.amount = parseFloat(jsondata.amount)
    jsondata.tags = selectedTags
    jsondata._id = props.money._id
    jsondata.username = props.money.username
    jsondata.maketime = props.money.maketime

    fetch('/api/money/put', {
      method: 'put',
      headers: {'Content-Type': 'application/json;charset=UTF-8'},
      body: JSON.stringify(jsondata)
    }).then(res => {
      switch(res.status) {
        case 200:
          props.setFormSwitch(false)
          res.json().then(data => props.setMoneyList([...replaceMoneyListElem(props.moneyList, data)]))
          break
        default:
          res.text().then(err => setErr(err))
          break
      }
    })
  }

  return (
    <>
      <div className='overlay' onClick={e => props.setFormSwitch(false)}></div>
      <form className='add-container mod-container'>
        <div>
          <input type='date' name='date' placeholder='Date' defaultValue={props.money.date.split('T')[0]} /><br />
          <input type='number' step='100' name='amount' placeholder='Amount' defaultValue={props.money.amount} /><br />
          <input type='text' name='summary' placeholder='Summary' defaultValue={props.money.summary} /><br />
          <input ref={tagInput} type='text' placeholder='Add custom tag' /><br />
        </div>
        <div>
          <button onClick={addTags}>Add Tag</button>
          <button onClick={putMoney}>Submit</button>
          <button onClick={e => props.setFormSwitch(false)}>X</button>
        </div>
        <div>
          {tags.map((tag, i) => {
            return <span key={i} 
              className={`tag ${selectedTags.includes(tag) ? 'selected-tag' : ''}`}
              onClick={e => updateSelectedTags(tag)}
              onDoubleClick={e => setTags(tags.filter(item => {return item !== tag}))}>{tag}</span>
          })}
        </div>
        {err.length > 0 &&
          <div className='err margintop2'>{err}</div>
        }
      </form>
    </>
  )
}

export default MoneyPutForm
