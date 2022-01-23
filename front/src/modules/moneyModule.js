export default {
  getIncomeAndSpending: moneyList => {
    let income = []
    let spending = []
    moneyList.map(money => {
      if (money.amount >= 0) {
        income.push(money)
      } else {
        spending.push(money)
      }
    })
    return [income, spending]
  },
  
  getTotalAmount: moneyList => {
    let total = 0
    moneyList.map(money => total += money.amount)
    return total
  }
}