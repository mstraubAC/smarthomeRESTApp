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
  name: 'ChartElectricEnergyFlowPVDaily',
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
            text: "Täglicher PV Energiefluss (letzten 30 Tage)",
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
                  text: 'Energie / kWh',
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
    async getElectricEnergyConsumptionMonthlyFlowPVDaily() {
        this.loaded = false;
        this.labels = [ ];
        this.datasets = [ ];
        
        //
        const today = new Date()
        const todayStr = today.toISOString().slice(0, 10);
        var startDate = new Date(today);
        startDate.setMonth(startDate.getMonth() - 1);
        const startDateStr = startDate.toISOString().slice(0, 10);
        const url = process.env.VUE_APP_SMARTHOME_API + "/v1/aggregates/electricconsumption/flow/daily?startDate=" + startDateStr + "&endDate=" + todayStr;
        const res = await fetch(url);
        const rawData = await res.json();

        this.chartData.datasets.push({
            label:'PV', 
            backgroundColor: 'rgb(234, 250, 24)',
            stack: 'Erzeuger',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Einspeisung',
            backgroundColor: 'rgb(75, 192, 192)',
            stack: 'Senken',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Eigenverbrauch', 
            backgroundColor: 'rgb(201, 203, 207)',
            stack: 'Senken',
            order: 0,
            data: [],
            });

        rawData.forEach(element => {
            // console.log(element);
            var logdate = new Date(Date.parse(element.logdate));
            var logdateStr = logdate.getFullYear() + "-" + String(logdate.getMonth()+1).padStart(2, '0') + "-" + String(logdate.getDate()).padStart(2, '0');
            this.chartData.labels.push(logdateStr);
            this.chartData.datasets[0].data.push(element.pvGeneration);
            this.chartData.datasets[1].data.push(element.electricGridFeedIn);
            var eigenVerbrauch = element.pvGeneration - element.electricGridFeedIn
            if (eigenVerbrauch < 0) {
              this.chartData.datasets[2].data.push(0);
            }
            else {
              this.chartData.datasets[2].data.push(eigenVerbrauch);
            }
        });
        this.loaded = true;
    }
  },
  mounted() {
    this.getElectricEnergyConsumptionMonthlyFlowPVDaily()
  }
}
</script>