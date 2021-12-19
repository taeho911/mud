function MoneyList(props) {
  return (
    <div className='money-unit-container'>
      <div className='col-date'>{props.money.date.split('T')[0]}</div>
      <div className='col-amount'>{props.money.amount}</div>
      <div className='col-tags'>
        {props.money.tags.map((v, i) => {
          return (
            <div key={i} className='tag display-tag'>{v}</div>
          )
        })}
      </div>
      <div className='col-summary'>{props.money.summary}</div>
    </div>
  )
}

export default MoneyList
