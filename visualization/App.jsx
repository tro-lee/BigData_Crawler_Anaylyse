import React, { useState } from 'react'
import CountryCharts from './modules/CountryCharts'
import CountryChartsPro from './modules/CountryChartsPro'
import CountryChartsProTree from './modules/CountryChartsProTree'
import SegCharts from './modules/SegCharts'
import SegChartsPro from './modules/SegChartsPro'

const Charts = [
  <SegCharts></SegCharts>,
  <SegChartsPro></SegChartsPro>,
  <CountryCharts></CountryCharts>,
  <CountryChartsPro></CountryChartsPro>,
  <CountryChartsProTree></CountryChartsProTree>
]

function App() {
  const [chart, setChart] = useState(0)
  const next = () => {
    setChart(chart => (chart + 1) % Charts.length)
  }

  function NextButton() {
    return (<button onClick={next}>Next Chart</button>)
  }

  const CurChart = Charts[chart]

  return (
    <>
      <NextButton></NextButton>
      {CurChart}
    </>
  )

}

export default App

