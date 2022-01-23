import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  ArcElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';
import { Doughnut, Bar } from 'react-chartjs-2'
import moneyModule from "../modules/moneyModule"

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  ArcElement,
  Title,
  Tooltip,
  Legend
)

function MoneyStat(props) {
  const [income, spending] = moneyModule.getIncomeAndSpending(props.moneyList)
  const [eachTagKeys, eachTagValues] = moneyModule.getTotalAmountOfEachTag(props.moneyList)
  const colorSet = [
    'rgba(255, 99, 132, 0.5)',
    'rgba(54, 162, 235, 0.5)',
    'rgba(255, 206, 86, 0.5)',
    'rgba(75, 192, 192, 0.5)',
    'rgba(153, 102, 255, 0.5)',
    'rgba(255, 159, 64, 0.5)'
  ]
  
  const getColorArray = len => {
    let arr = []
    let loop = parseInt(len / colorSet.length)
    let remain = len % colorSet.length
    for (let i = 0; i < loop; i++) {
      arr = arr.concat(colorSet)
    }
    arr = arr.concat(colorSet.slice(0, remain))
    return arr
  }

  const doughnutOptions = {
    plugins: {
      title: {
        display: true,
        text: 'Income and Spending',
        font: {size: 25}
      }
    }
  }

  const horizontalBarOptions = {
    indexAxis: 'y',
    elements: {
      bar: {borderWidth: 2},
    },
    responsive: true,
    plugins: {
      legend: {display: false},
      title: {
        display: true,
        text: 'Each Tags Stat',
        font: {size: 25}
      }
    }
  }

  const incomeSpendingData = {
    labels: ['Income', 'Spending'],
    datasets: [
      {
        data: [income, spending],
        backgroundColor: ['rgba(53, 162, 235, 0.5)', 'rgba(255, 99, 132, 0.5)'],
        borderWidth: 5
      }
    ]
  }

  const eachTagData = {
    labels: eachTagKeys,
    datasets: [
      {
        data: eachTagValues,
        backgroundColor: getColorArray(eachTagValues.length)
      }
    ]
  }

  return (
    <>
      <div>
        <hr className='margintop4 marginbottom4'></hr>
        <Doughnut data={incomeSpendingData} options={doughnutOptions} />
        <hr className='margintop4 marginbottom4'></hr>
        <Bar data={eachTagData} options={horizontalBarOptions} />
        <hr className='margintop4 marginbottom4'></hr>
        <button onClick={e => console.log(moneyModule.getTotalAmountOfEachTag(props.moneyList))}>getTotalAmountOfEachTag</button>
      </div>
    </>
  )
}

export default MoneyStat
