<template>
  <v-app>
    <v-row>
      <v-col>
        <v-card width="600px">
          <ChartElectricEnergyMoneyYearly/>
        </v-card>
      </v-col>
      <v-col>
        <v-card width="600px">
          <ChartElectricEnergyMoneyMonthly/>
        </v-card>
      </v-col>
      <v-col>
        <v-card width="600px">
          <ChartElectricEnergyMoneyDaily/>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-card width="600px">
          <ChartElectricEnergyFlowPVYearly/>
        </v-card>        
      </v-col>
      <v-col>
        <v-card width="600px">
          <ChartElectricEnergyFlowPVMonthly/>
        </v-card>
      </v-col>
      <v-col>
        <v-card width="600px">
          <ChartElectricEnergyFlowPVDaily/>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-card width="600px">
          <ChartElectricEnergyFlowYearly/>
        </v-card>
      </v-col>
      <v-col>
        <v-card width="600px">
          <ChartElectricEnergyFlowMonthly/>
        </v-card>
      </v-col>
      <v-col>
        <v-card width="600px">
          <ChartElectricEnergyFlowDaily/>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-card width="600px">
          <ChartHeatpumpYearly/>
        </v-card>
      </v-col>
      <v-col>
        <v-card width="600px">
          <ChartHeatpumpMonthly/>
        </v-card>
      </v-col>
      <v-col>
        <v-card width="600px">
          <ChartHeatpumpDaily/>
        </v-card>
      </v-col>
    </v-row>>
  </v-app>
</template>

<script>
import ChartHeatpumpDaily from './ChartHeatpumpDaily.vue'
import ChartHeatpumpMonthly from './ChartHeatpumpMonthly.vue'
import ChartHeatpumpYearly from './ChartHeatpumpYearly.vue'

import ChartElectricEnergyMoneyDaily from './ChartElectricEnergyMoneyDaily.vue'
import ChartElectricEnergyMoneyMonthly from './ChartElectricEnergyMoneyMonthly.vue'
import ChartElectricEnergyMoneyYearly from './ChartElectricEnergyMoneyYearly.vue'

import ChartElectricEnergyFlowDaily from './ChartElectricEnergyFlowDaily.vue'
import ChartElectricEnergyFlowMonthly from './ChartElectricEnergyFlowMonthly.vue'
import ChartElectricEnergyFlowYearly from './ChartElectricEnergyFlowYearly.vue'

import ChartElectricEnergyFlowPVDaily from './ChartElectricEnergyFlowPVDaily.vue'
import ChartElectricEnergyFlowPVMonthly from './ChartElectricEnergyFlowPVMonthly.vue'
import ChartElectricEnergyFlowPVYearly from './ChartElectricEnergyFlowPVYearly.vue'

// this.$ga.disable();

export default {
  name: 'App',
  components: { 
      ChartHeatpumpDaily, ChartHeatpumpMonthly, ChartHeatpumpYearly,
      ChartElectricEnergyMoneyDaily, ChartElectricEnergyMoneyMonthly, ChartElectricEnergyMoneyYearly, 
      ChartElectricEnergyFlowDaily, ChartElectricEnergyFlowMonthly, ChartElectricEnergyFlowYearly, 
      ChartElectricEnergyFlowPVDaily, ChartElectricEnergyFlowPVMonthly, ChartElectricEnergyFlowPVYearly
  },
  data () {
    return {
      showPassword: false,
    }
  },
}
</script>