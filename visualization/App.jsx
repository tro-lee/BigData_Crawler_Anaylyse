import React, { useState } from 'react'
import CountryCharts from './modules/CountryCharts'
import SegCharts from './modules/SegCharts'
import SegChartsPro from './modules/SegChartsPro'
function App() {
  const [chart, setChart] = useState(3)
  if (chart === 1) {
    return (
      <>
        <button onClick={() => setChart(2)}>Next Chart</button>
        <SegCharts></SegCharts>
      </>
    )
  }
  else if (chart === 2) {
    return (
      <>
        <button onClick={() => setChart(3)}>Next Chart</button>
        <SegChartsPro></SegChartsPro>
      </>
    )
  }
  else if (chart === 3) {
    return (
      <>
        <button onClick={() => setChart(1)}>Next Chart</button>
        <CountryCharts></CountryCharts>
      </>
    )
  }
  else {
    return (
      <>
        <h1>暂时没有后面的了</h1>
        <button onClick={() => setChart(1)}>Next Chart</button>
      </>
    )
  }
}

export default App

