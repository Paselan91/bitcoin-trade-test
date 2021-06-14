<!-- buildで落ちる時用 -->
<!-- <template>
  <div></div>
</template> -->

<template>
  <div class="chart"></div>
</template>

<script src="https://cdn.jsdelivr.net/npm/apexcharts"></script>
<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator"
import axios from "axios"

@Component({})
export default class ComboChart2 extends Vue {
  mounted() {
    const options = {
      chart: {
        height: 350,
        type: "line",
        zoom: {
          enabled: false
        }
      },
      dataLabels: {
        enabled: false
      },
      stroke: {
        curve: "straight"
      },
      series: [
        {
          name: "Desktops",
          data: [10, 41, 35, 51, 49, 62, 69, 91, 148]
        }
      ],
      title: {
        text: "Product Trends by Month",
        align: "left"
      },
      grid: {
        row: {
          colors: ["#f3f3f3", "transparent"], // takes an array which will be repeated on columns
          opacity: 0.5
        }
      },
      xaxis: {
        categories: [
          "Jan",
          "Feb",
          "Mar",
          "Apr",
          "May",
          "Jun",
          "Jul",
          "Aug",
          "Sep"
        ]
      }
    }

    const chart = new ApexCharts(document.querySelector("#chart"), options)
    chart.render()
  }
}
</script>
