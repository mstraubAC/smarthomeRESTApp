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
  name: 'ChartElectricEnergyFlowPVMonthly',
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
            text: "Jährlicher PV Energiefluss",
          },
        },
        scales: {
            x: {
                stacked: true,
                title: {
                  display: true,
                  text: 'Monat',
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
    async getElectricEnergyConsumptionMonthlyFlowPVMonthly() {
        this.loaded = false;
        this.labels = [ ];
        this.datasets = [ ];
        
        //
        const res = await fetch("http://localhost:7777/v1/aggregates/electricconsumption/flow/monthly");
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
            var logdateStr = logdate.getFullYear() + "-" + String(logdate.getMonth()+1).padStart(2, '0');
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
    this.getElectricEnergyConsumptionMonthlyFlowPVMonthly()
  }
}
</script>