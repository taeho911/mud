import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js'
import { Doughnut } from 'react-chartjs-2'
import moneyModule from "../modules/moneyModule"

ChartJS.register(ArcElement, Tooltip, Legend)

function MoneyStat(props) {
  const [income, spending] = moneyModule.getIncomeAndSpending(props.moneyList)

  const incomeSpendingData = {
    labels: ['Income', 'Spending'],
    datasets: [
      {
        data: [moneyModule.getTotalAmount(income), moneyModule.getTotalAmount(spending)],
        backgroundColor: ['#1a73e8', '#ff3e00'],
        borderColor: ['blue', 'red'],
        borderWidth: 1
      }
    ]
  }

  return (
    <>
      <div>
        <Doughnut data={incomeSpendingData} />
      </div>
    </>
  )
}

export default MoneyStat
