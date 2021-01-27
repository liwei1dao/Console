<template>
  <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import resize from '@/mixins/resize'

export default {
    name:"PieChart",
    mixins: [resize],
    props: {
        className: {
        type: String,
        default: 'chart'
        },
        data :{
            type: Object,
            default:function(){
                return {
                    title:"demo",
                    chartdata:[
                        {
                            value: 320, 
                            name: 'Industries'
                        },
                        {
                            value: 240, 
                            name: 'Technology'
                        }
                    ]
                }
            }
        },
        width: {
        type: String,
        default: '100%'
        },
        height: {
        type: String,
        default: '300px'
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
        var legenddata = []
        for (let i = 0; i < this.data.chartdata.length; ++i) {
            legenddata.push(this.data.chartdata[i].name)
        }

        this.chart.setOption({
            tooltip: {
                trigger: 'item',
                formatter: '{a} <br/>{b} : {c} ({d}%)'
            },
            legend: {
                left: 'center',
                bottom: '10',
                data: legenddata
            },
            series: [
                {
                    name: this.data.title,
                    type: 'pie',
                    roseType: 'radius',
                    radius: [15, 95],
                    center: ['50%', '38%'],
                    data: this.data.chartdata,
                    animationEasing: 'cubicInOut',
                    animationDuration: 2600
                }
            ]
        })
        }
    }
}
</script>


<style>

</style>