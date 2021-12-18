import { useState, useRef } from 'react'
import '../styles/money.css'

function Money() {
  const [tags, setTags] = useState(['income', 'spend', 'invest', 'life', 'play', 'drink', 'food'])
  const [selectedTags, setSelectedTags] = useState([])
  const tagInput = useRef(undefined)

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
          <button>Submit</button>
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
        </div>
      </form>
    </main>
  )
}

export default Money
