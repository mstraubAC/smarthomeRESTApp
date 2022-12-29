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
  name: 'ChartElectricEnergyMoneyDaily',
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
      type: Array,
      default: () => [Title]
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
        plugins: {
          title: {
            display: true,
            text: "Täglicher Energie relatierter Einnahmen-/Ausgabenfluss (letzten 30 Tage)",
          },
        },
        scales: {
            x: {
                stacked: true,
                title: {
                  display: true,
                  text: 'Tag',
                }
            },
            y: {
                beginAtZero: true,
                stacked: true,
                title: {
                  display: true,
                  text: 'EUR',
                }
            },
        },
        interaction: {
            intersect: false,
        }
      },
      loaded: false,
    }
  },
  methods: {
    async getElectricEnergyConsumptionMonthlyMoneyFlowDaily() {
        this.loaded = false;
        this.labels = [ ];
        this.datasets = [ ];
        
        //
        const today = new Date()
        const todayStr = today.toISOString().slice(0, 10);
        var startDate = new Date(today);
        startDate.setMonth(startDate.getMonth() - 1);
        const startDateStr = startDate.toISOString().slice(0, 10);
        const url = process.env.VUE_APP_SMARTHOME_API + "/v1/aggregates/electricconsumption/moneyflow/daily?startDate=" + startDateStr + "&endDate=" + todayStr;
        const res = await fetch(url);
        const rawData = await res.json();

        //
        this.chartData.datasets.push({
            label:'Strombezugskosten',
            backgroundColor: 'rgb(255, 99, 132)',
            // stack: 'VNB',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Einspeiseverguetung',
            backgroundColor: 'rgb(75, 192, 192)',
            // stack: 'VNB',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Umsatzsteuer auf Eigenverbrauch', 
            backgroundColor: 'rgb(201, 203, 207)',
            // stack: 'Electric energy',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Einsparung Bezugskosten', 
            backgroundColor: 'rgb(201, 203, 207)',
            // stack: 'Electric energy',
            order: 0,
            data: [],
            });
        rawData.forEach(element => {
            var logdate = new Date(Date.parse(element.logdate));
            var logdateStr = logdate.getFullYear() + "-" + String(logdate.getMonth()+1).padStart(2, '0') + "-" + String(logdate.getDate()).padStart(2, '0');
            this.chartData.labels.push(logdateStr);
            this.chartData.datasets[0].data.push(element.utilitiesBoughInclVat);
            this.chartData.datasets[1].data.push(element.utilitiesSoldInclVat);
            this.chartData.datasets[2].data.push(element.vatToPayForDirectConsumption);
            this.chartData.datasets[3].data.push(-1. * element.savedUtilitiesBuy);
        });
        this.loaded = true;
    }
  },
  mounted() {
    this.getElectricEnergyConsumptionMonthlyMoneyFlowDaily()
  }
}
</script>