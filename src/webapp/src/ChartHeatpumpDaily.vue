<template>
  <Bar v-if="loaded"
    :chart-options="chartOptions"
    :chart-data="chartData"
    :chart-id="chartId"
    :dataset-id-key="datasetIdKey"
    :plugins="plugins"
    :css-classes="cssClasses"
    :styles="styles"
    :width="width"
    :height="height"
  />
</template>

<script>
import { Bar } from 'vue-chartjs/legacy'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)

export default {
  name: 'ChartHeatpumpDaily',
  components: { Bar },
  props: {
    chartId: {
      type: String,
      default: '0'
    },
    datasetIdKey: {
      type: String,
      default: 'label'
    },
    width: {
      type: Number,
      default: 320
    },
    height: {
      type: Number,
      default: 240
    },
    cssClasses: {
      default: '',
      type: String
    },
    styles: {
      type: Object,
      default: () => {}
    },
    plugins: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      chartData: {
        labels: [ ],
        datasets: [ ]
      },
      chartOptions: {
        responsive: true,
        scales: {
            x: {
                stacked: true,
            },
            y: {
                beginAtZero: true,
                stacked: true,
            },
            // y2: {
            //     type: 'linear',
            //     position: 'right'
            // }
        },
        interaction: {
            intersect: false,
        }
      },
      loaded: false,
    }
  },
  methods: {
    async getHeatpumpDaily() {
        this.loaded = false;
        this.labels = [ ];
        this.datasets = [ ];
        
        //
        const today = new Date()
        const todayStr = today.toISOString().slice(0, 10);
        var startDate = new Date(today);
        startDate.setMonth(startDate.getMonth() - 1);
        const startDateStr = startDate.toISOString().slice(0, 10);
        const res = await fetch("http://localhost:7777/v1/aggregates/heatpump/daily?startDate=" + startDateStr + "&endDate=" + todayStr);
        const rawData = await res.json();

        //
        this.chartData.datasets.push({
            label:'Thermal Energy in kWh',
            backgroundColor: 'rgb(255, 99, 132)',
            stack: 'Thermal energy',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Electric Energy Heatpump in kWh',
            backgroundColor: 'rgb(75, 192, 192)',
            stack: 'Electric energy',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Electric Energy Control in kWh', 
            backgroundColor: 'rgb(201, 203, 207)',
            stack: 'Electric energy',
            order: 0,
            data: [],
            });
        // this.chartData.datasets.push({
        //     label:'Daily Work Ratio', 
        //     // backgroundColor: 'rgb(201, 203, 207)',
        //     stack: 'Worknumbers',
        //     // type: 'line',
        //     yAxisID: 'y2',
        //     data: [],
        //     });
        rawData.forEach(element => {
            // console.log(element)
            var logdate = new Date(Date.parse(element.logdate));
            var logdateStr = logdate.getFullYear() + "-" + String(logdate.getMonth()+1).padStart(2, '0') + "-" + String(logdate.getDay()+1).padStart(2, '0');
            this.chartData.labels.push(logdateStr);
            this.chartData.datasets[0].data.push(element.totalThermalEnergyInkWh);
            this.chartData.datasets[1].data.push(element.totalElectricEnergyHeatSourceOnlyInkWh);
            this.chartData.datasets[2].data.push(element.totalElectricEnergyIncludingControlInkWh - element.totalElectricEnergyHeatSourceOnlyInkWh);
            // this.chartData.datasets[3].data.push(element.dailyWorkCoefficientHeatSourceOnly);
        });
        this.loaded = true;
    }
  },
  mounted() {
    this.getHeatpumpDaily()
  }
}
</script>