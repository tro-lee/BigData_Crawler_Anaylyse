import * as echarts from 'echarts';
import { useEffect } from 'react';
import segTitle from '../../result/seg_content_result.json';

export default function () {
  useEffect(() => {
    const nameArr = []
    const countArr = []
    Reflect.ownKeys(segTitle).forEach(key => {
      nameArr.push(key)
      countArr.push(segTitle[key])
    })

    genEcharts(nameArr, countArr)
  }, [])

  return (
    <div className="seg-title" style={{ width: '100vw', height: '100vh' }}></div>
  )
}

function genEcharts(x, y) {
  var chartDom = document.querySelector('.seg-title');
  var myChart = echarts.init(chartDom);
  var option;

  option = {
    title: {
      text: '新闻词频统计, 当前数量:' + x.length,
      left: 10
    },
    toolbox: {
      feature: {
        dataZoom: {
          yAxisIndex: false
        },
        saveAsImage: {
          pixelRatio: 2
        }
      }
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      }
    },
    grid: {
      bottom: 90
    },
    dataZoom: [
      {
        type: 'inside'
      },
      {
        type: 'slider'
      }
    ],
    xAxis: {
      data: x,
      silent: false,
      splitArea: {
        show: false
      }
    },
    yAxis: {
      splitArea: {
        show: false
      }
    },
    series: [
      {
        type: 'bar',
        data: y,
        // Set `large` for large data amount
        large: true
      }
    ]
  };


  option && myChart.setOption(option);
}