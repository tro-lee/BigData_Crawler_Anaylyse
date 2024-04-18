import * as echarts from 'echarts';
import { useEffect } from 'react';
import segTitle from '../../result/seg_title_result.json';

export default function () {
  useEffect(() => {
    let roseResult = []
    Reflect.ownKeys(segTitle).forEach(key => {
      if (segTitle[key] < 100) return
      roseResult.push({ value: segTitle[key], name: key })
    })

    roseResult = roseResult.sort((a, b) => a.value - b.value)
    roseResult = roseResult.slice(0, 100)


    genEcharts(roseResult)
  }, [])

  return (
    <div className="seg-title" style={{ width: '100vw', height: '100vh' }}></div>
  )
}

function genEcharts(result) {
  var chartDom = document.querySelector('.seg-title');
  var myChart = echarts.init(chartDom);
  var option;

  option = {
    title: {
      text: '2023-2024年新闻',
      subtext: '热门词频',
      left: 'center'
    },
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b} : {c} ({d}%)'
    },
    toolbox: {
      show: true,
      feature: {
        mark: { show: true },
        dataView: { show: true, readOnly: false },
        restore: { show: true },
        saveAsImage: { show: true }
      }
    },
    series: [
      {
        name: 'Radius Mode',
        type: 'pie',
        radius: [20, 300],
        center: ['25%', '50%'],
        roseType: 'radius',
        itemStyle: {
          borderRadius: 5
        },
        label: {
          show: false
        },
        emphasis: {
          label: {
            show: true
          }
        },
        data: result
      },
      {
        name: 'Area Mode',
        type: 'pie',
        radius: [20, 300],
        center: ['75%', '50%'],
        roseType: 'area',
        itemStyle: {
          borderRadius: 5
        },
        data: result
      }
    ]
  };

  option && myChart.setOption(option);
}