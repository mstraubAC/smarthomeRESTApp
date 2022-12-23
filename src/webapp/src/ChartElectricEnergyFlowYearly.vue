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
  name: 'ChartElectricEnergyFlowYearly',
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
            text: "Jährlicher Energiefluss",
          },
        },
        scales: {
            x: {
                stacked: true,
                title: {
                  display: true,
                  text: 'Jahr',
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
    async getElectricEnergyConsumptionMonthlyFlowYearly() {
        this.loaded = false;
        this.labels = [ ];
        this.datasets = [ ];
        
        //
        const res = await fetch("http://localhost:7777/v1/aggregates/electricconsumption/flow/yearly");
        const rawData = await res.json();

        //
        this.chartData.datasets.push({
            label:'Strombezug',
            backgroundColor: 'rgb(255, 99, 132)',
            stack: 'VNB',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Einspeisung',
            backgroundColor: 'rgb(75, 192, 192)',
            stack: 'VNB',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Heizung', 
            backgroundColor: 'rgb(255, 0, 0)',
            stack: 'Verbraucher',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'IT', 
            backgroundColor: 'rgb(251, 0, 255)',
            stack: 'Verbraucher',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Wallbox', 
            backgroundColor: 'rgb(239, 240, 0)',
            stack: 'Verbraucher',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Haushalt', 
            backgroundColor: 'rgb(201, 203, 207)',
            stack: 'Verbraucher',
            order: 0,
            data: [],
            });
        rawData.forEach(element => {
            var logdate = new Date(Date.parse(element.logdate));
            var logdateStr = logdate.getFullYear();
            this.chartData.labels.push(logdateStr);
            this.chartData.datasets[0].data.push(element.electricGridConsumption);
            this.chartData.datasets[1].data.push(-1. * element.electricGridFeedIn);
            this.chartData.datasets[2].data.push(element.heatingConsumption);
            this.chartData.datasets[3].data.push(element.itConsumption);
            this.chartData.datasets[4].data.push(element.wallboxConsumption);
            var eigenVerbrauch = element.pvGeneration - element.electricGridFeedIn
            if (eigenVerbrauch < 0) {
              this.chartData.datasets[5].data.push(0);
            }
            else {
              this.chartData.datasets[5].data.push(element.electricGridConsumption + eigenVerbrauch - element.heatingConsumption - element.itConsumption - element.wallboxConsumption);
            }
        });
        this.loaded = true;
    }
  },
  mounted() {
    this.getElectricEnergyConsumptionMonthlyFlowYearly()
  }
}
</script>