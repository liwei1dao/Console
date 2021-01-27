<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import resize from '@/mixins/resize'

export default {
  mixins: [resize],
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '350px'
    },
    autoResize: {
      type: Boolean,
      default: true
    },
    data: {
      type: Object,
      default: function(){
          return{
            xAxis:{
                data:[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23]
            },
            yAxis:{
                max:100
            },
            series:[
                {
                    name:"demo",
                    color:"#FF005A",
                    data:[0,0,0,0,0],
                }
            ]
          }
      }
    }
  },
  data() {
    return {
      chart: null
    }
  },
  watch: {
    chartData: {
      deep: true,
      handler(val) {
        this.setOptions(val)
      }
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart()
    })
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
      this.chart = echarts.init(this.$el, 'macarons')
      this.setOptions()
    },
    setOptions() {
        var legenddata = []
        var seriesdata = []
        for (let i = 0; i < this.data.series.length; ++i) {
            legenddata.push(this.data.series[i].name)
            seriesdata.push({
                name: this.data.series[i].name, itemStyle: {
                    normal: {
                        color: this.data.series[i].color,
                        lineStyle: {
                            color: this.data.series[i].color,
                            width: 2
                        }
                    }
                },
                smooth: true,
                type: 'line',
                data: this.data.series[i].data,
                animationDuration: 2800,
                animationEasing: 'cubicInOut'
            })
        }
        this.chart.setOption({
            xAxis: {
                data: this.data.xAxis.data,
                boundaryGap: false,
                axisTick: {
                    show: false
                }
            },
            grid: {
                left: 10,
                right: 10,
                bottom: 20,
                top: 30,
                containLabel: true
            },
            tooltip: {
                trigger: 'axis',
                axisPointer: {
                    type: 'cross'
                },
                padding: [5, 10]
            },
            yAxis: {
                max : this.data.yAxis.max,
                axisTick: {
                    show: true,
                    length :10,  
                }
            },
            legend: {
                data: legenddata
            },
            series: seriesdata
        })
    }
  }
}
</script>
