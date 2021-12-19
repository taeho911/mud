function MoneyList(props) {

  return (
    <>
      <hr />
      <div className='money-unit-container'>
        <div>
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
        <div className='del-icon-container'>
          <div className='del-icon'></div>
        </div>
      </div>
    </>
  )
}

export default MoneyList
