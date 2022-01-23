export default {
  getIncomeAndSpending: moneyList => {
    let income = 0
    let spending = 0
    moneyList.map(money => {
      if (money.amount >= 0) {
        income += money.amount
      } else {
        spending += money.amount
      }
    })
    return [income, spending]
  },

  getTotalAmount: moneyList => {
    let total = 0
    moneyList.map(money => total += money.amount)
    return total
  },

  getTotalAmountOfEachTag: moneyList => {
    let totals = new Map()
    moneyList.map(money => {
      if (money.tags.length == 0) {
        let key = 'extra'
        if (totals.has(key)) {
          totals.set(key, totals.get(key) + money.amount)
        } else {
          totals.set(key, money.amount)
        }
      } else {
        money.tags.map(tag => {
          if (totals.has(tag)) {
            totals.set(tag, totals.get(tag) + money.amount)
          } else {
            totals.set(tag, money.amount)
          }
        })
      }
    })
    return [Array.from(totals.keys()), Array.from(totals.values())]
  }
}