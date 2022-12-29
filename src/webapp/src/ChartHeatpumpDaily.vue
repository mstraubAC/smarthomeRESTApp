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
            text: "Täglicher elektrischer Energiebedarf und thermische Energiebereitstellung Wärmepumpe (letzten 30 Tage)",
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
        const url = process.env.VUE_APP_SMARTHOME_API + "/v1/aggregates/heatpump/daily?startDate=" + startDateStr + "&endDate=" + todayStr;
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
            label:'Elektrische Energie Wärmepumpe',
            backgroundColor: 'rgb(75, 192, 192)',
            stack: 'Electrisch',
            order: 0,
            data: [],
            });
        this.chartData.datasets.push({
            label:'Elektrische Energie Steuerung und Pumpen', 
            backgroundColor: 'rgb(201, 203, 207)',
            stack: 'Electrisch',
            order: 0,
            data: [],
            });
        rawData.forEach(element => {
            var logdate = new Date(Date.parse(element.logdate));
            var logdateStr = logdate.getFullYear() + "-" + String(logdate.getMonth()+1).padStart(2, '0') + "-" + String(logdate.getDate()).padStart(2, '0');
            this.chartData.labels.push(logdateStr);
            this.chartData.datasets[0].data.push(element.totalThermalEnergyInkWh);
            this.chartData.datasets[1].data.push(element.totalElectricEnergyHeatSourceOnlyInkWh);
            this.chartData.datasets[2].data.push(element.totalElectricEnergyIncludingControlInkWh - element.totalElectricEnergyHeatSourceOnlyInkWh);
        });
        this.loaded = true;
    }
  },
  mounted() {
    this.getHeatpumpDaily()
  }
}
</script>