import { useState, useRef } from 'react'

function MoneyForm(props) {
  const [tags, setTags] = useState(props.tags)
  const [selectedTags, setSelectedTags] = useState(props.selectedTags)
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

  const postMoney = e => {
    console.log('postMoney')
  }

  return (
    <form className='add-container'>
      <div>
        <input type='date' name='date' placeholder='Date' defaultValue={props.date} /><br />
        <input type='number' step='100' name='amount' placeholder='Amount' defaultValue={props.number} /><br />
        <input type='text' name='summary' placeholder='Summary' defaultValue={props.summary} /><br />
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
  )
}

export default MoneyForm
