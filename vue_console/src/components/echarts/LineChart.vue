<!-- 
    Echarts 线性图标
    Demo:https://echarts.apache.org/examples/zh/editor.html?c=gauge-grade
 -->
<template>
  <v-card :width="width" :height="height"  ref="main">
  </v-card>
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import resize from '@/mixins/resize'

export default {
    name: "LineChart",
    mixins: [resize],
    props: {
        width: {
            type: [String,Number],
            default: '100%'
        },
        height: {
            type: [String,Number],
            default: '100%'
        },
        keys :{
            type: Array,
            default: () => ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        },
        values:{
            type: Array,
            default: () => [150, 230, 224, 218, 135, 147, 260]
        }
    },
    data() {
        return {
            chart: null
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
            this.chart.setOption({
                xAxis: {
                    type: 'category',
                    data: this.keys
                },
                yAxis: {
                    type: 'value'
                },
                series: [{
                    data: this.values,
                    type: 'line',
                    smooth: false               //false折线 true曲线
                }]
            })
        }
    }
}
</script>
