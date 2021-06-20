<!-- buildで落ちる時用 -->
<!--
<template>
  <div>combo chart</div>
</template>
-->
<!--
<template>
  <div class="chart-box">
    <div id="chart-candlestick">
      <apexchart
        height="290"
        :options="chartOptions"
        :series="series"
      ></apexchart>
    </div>
    <div id="chart">
      <apexchart
        type="line"
        height="350"
        :options="chartOptionsCombo"
        :series="seriesCombo"
      ></apexchart>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator"

@Component({})
export default class ComboChart extends Vue {
  seriesData = [
    {
      x: new Date(1538879400000),
      y: [6604.44, 6604.44, 6600, 6603.5]
    },
    {
      x: new Date(1538881200000),
      y: [6603.5, 6603.99, 6597.5, 6603.86]
    },
    {
      x: new Date(1538883000000),
      y: [6603.85, 6605, 6600, 6604.07]
    },
    {
      x: new Date(1538884800000),
      y: [6604.98, 6606, 6604.07, 6606]
    }
  ]

  public seriesCombo = [
    {
      name: "TEAM A",
      type: "column",
      data: [23, 11, 22, 27, 13, 22, 37, 21, 44, 22, 30]
    },
    {
      name: "TEAM B",
      type: "area",
      data: [44, 55, 41, 67, 22, 43, 21, 41, 56, 27, 43]
    },
    {
      name: "TEAM C",
      type: "line",
      data: [30, 25, 36, 30, 45, 35, 64, 52, 59, 36, 39]
    }
  ]

  seriesData2 = [
    {
      x: new Date(1538879400000),
      y: 6600
    },
    {
      x: new Date(1538881200000),
      y: 6600
    },
    {
      x: new Date(1538883000000),
      y: 6600
    },
    {
      x: new Date(1538884800000),
      y: 6621
    }
  ]

  seriesData3 = [
    {
      x: new Date(1538879400000),
      y: 6621
    },
    {
      x: new Date(1538881200000),
      y: 6663
    },
    {
      x: new Date(1538883000000),
      y: 6600
    },
    {
      x: new Date(1538884800000),
      y: 6640
    }
  ]

  public series = [
    {
      name: "Website Blog3",
      type: "line",
      data: this.seriesData3,
      options: {
        colors: ["blue"]
      }
    },
    {
      name: "Website Blog",
      type: "candlestick",
      data: this.seriesData
    }
  ]

  public chartOptions: any = {
    chart: {
      type: "candlestick",
      height: 290,
      id: "candles",
      toolbar: {
        autoSelected: "pan",
        show: false
      },
      zoom: {
        enabled: false
      }
    },
    plotOptions: {
      candlestick: {
        colors: {
          upward: "#3C90EB",
          downward: "#DF7D46"
        }
      }
    },
    xaxis: {
      type: "datetime"
    }
  }

  public chartOptionsCombo: any = {
    chart: {
      height: 350,
      type: "line",
      stacked: false
    },
    stroke: {
      width: [0, 2, 5],
      curve: "smooth"
    },
    plotOptions: {
      bar: {
        columnWidth: "50%"
      }
    },

    fill: {
      opacity: [0.85, 0.25, 1],
      gradient: {
        inverseColors: false,
        shade: "light",
        type: "vertical",
        opacityFrom: 0.85,
        opacityTo: 0.55,
        stops: [0, 100, 100, 100]
      }
    },
    labels: [
      "01/01/2003",
      "02/01/2003",
      "03/01/2003",
      "04/01/2003",
      "05/01/2003",
      "06/01/2003",
      "07/01/2003",
      "08/01/2003",
      "09/01/2003",
      "10/01/2003",
      "11/01/2003"
    ],
    markers: {
      size: 0
    },
    xaxis: {
      type: "datetime"
    },
    yaxis: {
      title: {
        text: "Points"
      },
      min: 0
    },
    tooltip: {
      shared: true,
      intersect: false,
      y: {
        formatter(y) {
          if (typeof y !== "undefined") {
            return y.toFixed(0) + " points"
          }
          return y
        }
      }
    }
  }
}
</script>
-->
