<!-- buildで落ちる時用 -->
<!-- <template>
  <div></div>
</template> -->

<template>
  <div>
    <v-container class="grey lighten-5 text-center">
      <h1>Past Candle</h1>
      <v-row :align="align" no-gutters>
        <v-col :cols="5">
          <v-container>
            <h3 class="text-center">Choose Periods</h3>
            <v-select
              v-model="selectedPeriods"
              :items="periods"
              label="Periods"
              item-text="label"
              item-value="value"
              solo
              return-object
              width="100px"
            ></v-select>
          </v-container>
        </v-col>
        <v-spacer></v-spacer>
        <v-col :cols="7">
          <h3 class="text-center">Choose Date</h3>
          <v-row>
            <v-col cols="5">
              <v-container fluid>
                <v-radio-group v-model="beforeAfter" mandatory col>
                  <v-radio label="After" value="after"></v-radio>
                  <v-radio label="Before" value="before"></v-radio>
                </v-radio-group>
              </v-container>
            </v-col>
            <v-col cols="7">
              <v-container>
                <v-row>
                  <v-col cols="12" lg="6">
                    <v-menu
                      :close-on-content-click="false"
                      transition="scale-transition"
                      offset-y
                      max-width="290px"
                      min-width="auto"
                    >
                      <template #activator="{ on, attrs }">
                        <v-text-field
                          v-model="computedDateFormatted"
                          label="Date"
                          persistent-hint
                          prepend-icon="mdi-calendar"
                          v-bind="attrs"
                          @blur="date = parseDate(dateFormatted)"
                          v-on="on"
                        ></v-text-field>
                      </template>
                      <v-date-picker v-model="date" no-title></v-date-picker>
                    </v-menu>
                  </v-col>
                </v-row>
              </v-container>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="2">
          <v-container class="px-0" fluid>
            <v-checkbox v-model="isSmas" :label="`SMAS`"></v-checkbox>
            <div v-if="isSmas">
              <v-text-field v-model="sma1" label="SMA1" required></v-text-field>
              <v-text-field v-model="sma2" label="SMA2" required></v-text-field>
              <v-text-field v-model="sma3" label="SMA3" required></v-text-field>
            </div>
          </v-container>
        </v-col>
        <v-col cols="2">
          <h3>You selected</h3>
          <p>
            {{ selectedPeriods.label }} {{ beforeAfter }}
            {{ date }}
          </p>
        </v-col>
        <v-col cols="2">
          <v-btn class="my-3" @click="fetchPastCandle">Get Data</v-btn>
        </v-col>
        <v-col cols="2">
          <v-btn class="my-3" @click="healthCheck">health Check</v-btn>
        </v-col>
      </v-row>
    </v-container>
    <v-container>
      <client-only>
        <div v-if="isLoading" class="text-center">
          <v-progress-circular :size="50" indeterminate></v-progress-circular>
        </div>
        <div v-else>
          <div id="chart">
            <apexchart
              type="candlestick"
              height="350"
              :options="candlestickChartOptions"
              :series="candlestickSeries"
            ></apexchart>
          </div>
          <div id="chart">
            <apexchart
              type="line"
              height="350"
              :options="lineChartOptions"
              :series="lineSeries"
            ></apexchart>
          </div>
        </div>
      </client-only>
    </v-container>

    <v-container class="my-5">
      <v-row justify="center">
        <v-btn to="/" nuxt elevation="3" width="10%">
          Back
        </v-btn>
      </v-row>
    </v-container>

    <v-container>
      <v-btn class="my-3 text-center" @click="checkParamerters"
        >Check Qeury Parameters</v-btn
      >
    </v-container>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "nuxt-property-decorator"
import axios from "axios"

@Component({})
export default class GenericChart extends Vue {
  isLoading: Boolean = false
  isSmas: Boolean = false
  sma1: number = 7
  sma2: number = 14
  sma3: number = 21

  url: string = "api/v1/candles"
  beforeAfter: string = "after"
  periods: Array<object> = [
    { label: "1m", value: 60 },
    { label: "3m", value: 180 },
    { label: "5m", value: 300 },
    { label: "15m", value: 900 },
    { label: "30m", value: 1800 },
    { label: "1h", value: 3600 },
    { label: "2h", value: 7200 },
    { label: "4h", value: 14400 },
    { label: "6h", value: 21600 },
    { label: "12h", value: 43200 },
    { label: "1d", value: 86400 },
    { label: "3d", value: 259200 },
    { label: "1w", value: 604800 }
  ]

  selectedPeriods = { label: "15m", value: 900 }

  today = new Date()
  date = new Date().toISOString().substr(0, 10)
  dateFormatted = this.formatDate(new Date().toISOString().substr(0, 10))
  get computedDateFormatted() {
    return this.formatDate(this.date)
  }

  formatDate(date: any) {
    if (!date) return null

    const [year, month, day] = date.split("-")
    return `${month}/${day}/${year}`
  }

  parseDate(date: any) {
    if (!date) return null

    const [month, day, year] = date.split("/")
    return `${year}-${month.padStart(2, "0")}-${day.padStart(2, "0")}`
  }

  @Watch("data")
  watchDate() {
    this.dateFormatted = this.formatDate(this.date)
  }

  public candlestickSeries = [
    {
      data: []
    }
  ]

  public candlestickChartOptions: any = {
    chart: {
      type: "candlestick",
      height: 350
    },
    title: {
      text: "CandleStick Chart",
      align: "left"
    },
    xaxis: {
      type: "datetime",
      minHeight: undefined,
      maxHeight: undefined
    },
    yaxis: {
      tooltip: {
        enabled: true
      },
      min: undefined,
      max: undefined
    }
  }

  public lineSeries = [
    {
      name: "SMA1",
      data: []
    },
    {
      name: "SMA2",
      data: []
    },
    {
      name: "SMA3",
      data: []
    }
  ]

  public lineChartOptions: any = {
    chart: {
      height: 350,
      type: "line",
      zoom: {
        enabled: true
      }
    },
    dataLabels: {
      enabled: false
    },
    stroke: {
      curve: "straight",
      width: 1
    },
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
      categories: []
    },
    yaxis: {
      min: 3500000,
      max: 3700000
    }
  }

  // mounted() {
  //   const yesterday = new Date(this.today)
  //   yesterday.setDate(yesterday.getDate() - 1)
  //   this.date = yesterday.toISOString().substr(0, 10)
  //   this.fetchPastCandle()
  // }

  async healthCheck() {
    await axios
      .get(this.url)
      .then(res => {
        if (res.status === 200) {
          // console.log("response")
          // console.log(res.data)
          console.log("res")
          console.log(res)
        }
      })
      .catch(e => {
        alert("ヘルスチェックに失敗しました")
        console.error("error!!!")
        console.error(e)
        return false
      })
  }

  async fetchPastCandle() {
    this.isLoading = true

    // FIXME: 変数名かえたい
    const unixtimestamp = Math.floor(Date.parse(this.date) / 1000)

    const requestData = {
      periods: this.selectedPeriods.value,
      beforeAfter: this.beforeAfter,
      time: unixtimestamp,
      smas: this.isSmas ? "1" : "", // FIXME:
      ...(this.isSmas
        ? {
            sma1: this.sma1,
            sma2: this.sma2,
            sma3: this.sma3
          }
        : {})
      // ...(this.isSmas
      //   ? {
      //       smas: {
      //         sma1: this.sma1,
      //         sma2: this.sma2,
      //         sma3: this.sma3
      //       }
      //     }
      //   : {})
    }
    console.log("requestData")
    console.log(requestData)

    // FIXME: anyに型をつける
    let data: Array<any>

    await axios
      .get(this.url, { params: requestData })
      .then(res => {
        if (res.status === 200) {
          console.log("response")
          console.log(res.data)
          // data = res.data.candles.result[this.selectedPeriods.value]
          data = res.data.candles

          /**
           *  Response
           *  https://docs.cryptowat.ch/rest-api/markets/ohlc
           *
           *  [
           *      CloseTime,
           *      OpenPrice,
           *      HighPrice,
           *      LowPrice,
           *      ClosePrice,
           *      Volume,
           *      QuoteVolume
           *  ]
           **/
          const chartData: any = data.map(value => ({
            x: new Date(value.close_time * 1000),
            y: [value.open, value.high, value.low, value.close]
          }))

          this.candlestickSeries = [
            {
              data: chartData
            }
          ]

          const maxYaxis = Math.max(
            ...chartData.map((value: any) => value.y[1])
          )
          const minYaxis = Math.min(
            ...chartData.map((value: any) => value.y[2])
          )
          const maxXaxis = Math.max(...chartData.map((value: any) => value.x))
          const minXaxis = Math.min(...chartData.map((value: any) => value.x))

          this.candlestickChartOptions.yaxis.max = maxYaxis
          this.candlestickChartOptions.yaxis.min = minYaxis
          this.candlestickChartOptions.xaxis.maxHeight = maxXaxis
          this.candlestickChartOptions.xaxis.minHeight = minXaxis

          if (res.data.smas) {
            const sma1: any = res.data.smas[0].values
            const sma2: any = res.data.smas[1].values
            const sma3: any = res.data.smas[2].values

            this.lineSeries[0].data = [...sma1]
            this.lineSeries[1].data = [...sma2]
            this.lineSeries[2].data = [...sma3]

            this.lineChartOptions.xaxis.categories = chartData.map(function(
              item: any
            ) {
              return item.x
            })
            this.lineChartOptions.yaxis.max = maxYaxis
            this.lineChartOptions.yaxis.min = minYaxis
            console.log("this.lineSeries.data")
            console.log(this.lineSeries[0].data)
          }
        }
      })
      .catch(e => {
        alert("Apiの取得に失敗しました")
        console.error(e)
        return false
      })
    this.isLoading = false
  }

  // TODO: デバッグ用
  // checkParamerters() {
  //   console.log("requestData")
  //   console.log(requestData)
  // }
  // @Prop({ required: true, default: "bar" }) type: string
}
</script>
