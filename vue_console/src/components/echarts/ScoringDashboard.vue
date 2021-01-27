<!-- 
    Echarts 评分仪表盘组件
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
    name: "ScoringDashboard",
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
        Scorename:{
            type: String,
            default: "成绩评定"
        },
        Score:{
            type: Number,
            default: 50
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
            this.chart = echarts.init(this.$el,'macarons')
            this.chart.setOption({
                series: [{
                    type: 'gauge',
                    startAngle: 180,
                    endAngle: 0,
                    min: 0,
                    max: 100,
                    splitNumber: 8,
                    axisLine: {
                        lineStyle: {
                            width: 6,
                            color: [
                                [0.25, '#FF6E76'],
                                [0.5, '#FDDD60'],
                                [0.75, '#58D9F9'],
                                [1, '#7CFFB2 ']
                            ]
                        }
                    },
                    pointer: {
                        icon: 'path://M12.8,0.7l12,40.1H0.7L12.8,0.7z',
                        length: '80%',
                        width: 5,
                        offsetCenter: [0, '20%'],
                        itemStyle: {
                            color: 'auto'
                        }
                    },
                    axisTick: {
                        length: 12,
                        lineStyle: {
                            color: 'auto',
                            width: 2
                        }
                    },
                    splitLine: {
                        length: 20,
                        lineStyle: {
                            color: 'auto',
                            width: 5
                        }
                    },
                    axisLabel: {
                        color: '#464646',
                        fontSize: 20,
                        distance: -60,
                        formatter: function (value) {
                            if (value === 87.5) {
                                return '优';
                            }
                            else if (value === 62.5) {
                                return '中';
                            }
                            else if (value === 37.5) {
                                return '良';
                            }
                            else if (value === 12.5) {
                                return '差 ';
                            }
                        }
                    },
                    title: {
                        offsetCenter: [0, '-40%'],
                        fontSize: 30
                    },
                    detail: {
                        fontSize: 40,
                        offsetCenter: [0, '20%'],
                        valueAnimation: true,
                        formatter: function (value) {
                            return Math.round((value)) + '%';
                        },
                        color: 'auto'
                    },
                    data: [{
                        value: 100 - this.Score,
                        name: this.Scorename
                    }]
                }]
            })
        }
    }
}
</script>

<style>

</style>