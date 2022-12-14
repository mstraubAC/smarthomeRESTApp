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
  name: 'ChartHeatpumpYearly',
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
            text: "Jährlicher elektrischer Energiebedarf und thermische Energiebereitstellung Wärmepumpe",
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
    async getHeatpumpYearly() {
        this.loaded = false;
        this.labels = [ ];
        this.datasets = [ ];
        
        //
        const url = process.env.VUE_APP_SMARTHOME_API + "/v1/aggregates/heatpump/yearly";
        const res = await fetch(url);
        const rawData = await res.json();

        //
        this.chartData.datasets.push({
            label:'Thermische Energie',
            backgroundColor: 'rgb(255, 99, 132)',
            stack: 'Thermisch',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Elektrische Energie Waermepumpe',
            backgroundColor: 'rgb(75, 192, 192)',
            stack: 'Electrisch',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Electrische Energie inkl. Steuerung', 
            backgroundColor: 'rgb(201, 203, 207)',
            stack: 'Electrisch',
            order: 0,
            data: [],
            });
        rawData.forEach(element => {
            var logdate = new Date(Date.parse(element.logdate));
            var logdateStr = logdate.getFullYear();
            this.chartData.labels.push(logdateStr);
            this.chartData.datasets[0].data.push(element.totalThermalEnergyInkWh);
            this.chartData.datasets[1].data.push(element.totalElectricEnergyHeatSourceOnlyInkWh);
            var electricEnergyControl = 0
            if (element.totalElectricEnergyIncludingControlInkWh > 0) {
              electricEnergyControl = element.totalElectricEnergyIncludingControlInkWh - element.totalElectricEnergyHeatSourceOnlyInkWh
            }
            this.chartData.datasets[2].data.push(electricEnergyControl);
        });
        this.loaded = true;
    }
  },
  mounted() {
    this.getHeatpumpYearly()
  }
}
</script>