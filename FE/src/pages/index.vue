<!-- buildで落ちる時用 -->
<!-- 
<template>
  <div>past candle</div>
</template>
-->
<template>
  <div>
    <v-card class="mx-auto my-12" max-width="1024">
      <v-container fluid>
        <v-row>
          <v-col :cols="3">
            <v-select
              v-model="selectedPeriods"
              :items="periods"
              label="Periods"
              item-text="label"
              item-value="value"
              outlined
            ></v-select>
          </v-col>
          <v-col :cols="3">
            <v-select
              v-model="selectedTerms"
              :items="terms"
              label="Term"
              item-text="label"
              item-value="value"
              outlined
            ></v-select>
          </v-col>
          <v-col :cols="3">
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
                  v-bind="attrs"
                  outlined
                  @blur="date = parseDate(dateFormatted)"
                  v-on="on"
                ></v-text-field>
              </template>
              <v-date-picker v-model="date" no-title></v-date-picker>
            </v-menu>
          </v-col>
          <v-col :cols="3">
            <v-menu :offset-y="true" :close-on-content-click="false">
              <template #activator="{ on, attrs }">
                <v-btn
                  v-bind="attrs"
                  outlined
                  width="200"
                  height="55"
                  v-on="on"
                >
                  Indicators
                </v-btn>
              </template>
              <v-card width="600">
                <v-list :height="60">
                  <v-list-item>
                    <v-col :cols="3">
                      <v-checkbox v-model="isSmas" :label="`SMAS`"></v-checkbox>
                    </v-col>
                    <v-col v-if="isSmas" :cols="3">
                      <v-text-field
                        v-model="sma1"
                        label="SMA1"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-col v-if="isSmas" :cols="3">
                      <v-text-field
                        v-model="sma2"
                        label="SMA2"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-col v-if="isSmas" :cols="3">
                      <v-text-field
                        v-model="sma3"
                        label="SMA3"
                        required
                      ></v-text-field>
                    </v-col>
                  </v-list-item>
                </v-list>
                <v-list :height="60">
                  <v-list-item>
                    <v-col :cols="3">
                      <v-checkbox v-model="isEmas" :label="`EMAS`"></v-checkbox>
                    </v-col>
                    <v-col v-if="isEmas" :cols="3">
                      <v-text-field
                        v-model="ema1"
                        label="EMA1"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-col v-if="isEmas" :cols="3">
                      <v-text-field
                        v-model="ema2"
                        label="EMA2"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-col v-if="isEmas" :cols="3">
                      <v-text-field
                        v-model="ema3"
                        label="EMA3"
                        required
                      ></v-text-field>
                    </v-col>
                  </v-list-item>
                </v-list>
                <v-list :height="60">
                  <v-list-item>
                    <v-col :col="3">
                      <v-checkbox
                        v-model="isBBands"
                        :label="`Bolinger Bands`"
                      ></v-checkbox>
                    </v-col>
                    <v-col v-if="isBBands" :col="3">
                      <v-text-field
                        v-model="bbandsN"
                        label="N"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-col v-if="isBBands" :col="3">
                      <v-text-field
                        v-model="bbandsK"
                        label="K"
                        required
                      ></v-text-field>
                    </v-col>
                  </v-list-item>
                </v-list>
                <v-list :height="60">
                  <v-list-item>
                    <v-col>
                      <v-checkbox
                        v-model="isIchimokuCloud"
                        :label="`Ichimoku Cloud`"
                      ></v-checkbox>
                    </v-col>
                  </v-list-item>
                </v-list>
                <v-list :height="60">
                  <v-list-item>
                    <v-col>
                      <v-checkbox
                        v-model="isVolume"
                        :label="`Volume`"
                      ></v-checkbox>
                    </v-col>
                  </v-list-item>
                </v-list>
                <v-list :height="60">
                  <v-list-item>
                    <v-col :col="3">
                      <v-checkbox v-model="isRsi" :label="`RSI`"></v-checkbox>
                    </v-col>
                    <v-col v-if="isRsi" :col="3">
                      <v-text-field
                        v-model="rsi"
                        label="Period"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-spacer></v-spacer>
                    <v-spacer></v-spacer>
                  </v-list-item>
                </v-list>
                <v-list>
                  <v-list-item>
                    <v-col :col="3">
                      <v-checkbox v-model="isMacd" :label="`MACD`"></v-checkbox>
                    </v-col>
                    <v-col v-if="isMacd" :col="3">
                      <v-text-field
                        v-model="macd1"
                        label="Period 1"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-col v-if="isMacd" :col="3">
                      <v-text-field
                        v-model="macd2"
                        label="Period 2"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-col v-if="isMacd" :col="3">
                      <v-text-field
                        v-model="macd3"
                        label="Period 3"
                        required
                      ></v-text-field>
                    </v-col>
                  </v-list-item>
                </v-list>
              </v-card>
            </v-menu>
          </v-col>
        </v-row>
      </v-container>
      <v-container>
        <client-only>
          <v-row
            v-if="isLoading"
            justify="center"
            align-content="center"
            class="text-center no-chart-area"
          >
            <v-progress-circular :size="70" indeterminate></v-progress-circular>
          </v-row>
          <div v-else>
            <div id="chart-box">
              <div id="chart-candlestick">
                <apexchart
                  type="candlestick"
                  height="350"
                  :options="candlestickChartOptions"
                  :series="candlestickSeries"
                ></apexchart>
              </div>
              <div v-if="isVolume" id="chart-bar">
                <apexchart
                  type="bar"
                  height="350"
                  :options="volumeChartOptions"
                  :series="volumeChartSeries"
                ></apexchart>
              </div>
              <div v-if="isRsi" id="chart-line">
                <apexchart
                  type="line"
                  height="350"
                  :options="rsiLineChartOptions"
                  :series="rsiLineSeries"
                ></apexchart>
              </div>
              <div v-if="isLine" id="chart">
                <apexchart
                  type="line"
                  height="350"
                  :options="lineChartOptions"
                  :series="lineSeries"
                ></apexchart>
              </div>
              <div v-if="isMacd" id="chart">
                <apexchart
                  type="line"
                  height="350"
                  :options="macdChartOptions"
                  :series="macdSeries"
                >
                </apexchart>
              </div>
            </div>
          </div>
        </client-only>
        <div class="text-center">
          <v-btn color="primary" rounded @click="fetchPastCandle"
            >Get Data</v-btn
          >
        </div>
      </v-container>
    </v-card>
    <!-- <v-container>
      <v-btn class="my-3 text-center" @click="checkParamerters"
        >Check Qeury Parameters</v-btn
      >
    </v-container> -->
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "nuxt-property-decorator"
import axios from "axios"

@Component({})
export default class GenericChart extends Vue {
  isLoading: Boolean = false

  isLine: Boolean = false

  isSmas: Boolean = false
  sma1: number = 7
  sma2: number = 14
  sma3: number = 21

  isEmas: Boolean = false
  ema1: number = 7
  ema2: number = 14
  ema3: number = 21

  isBBands: Boolean = false
  bbandsN: number = 20
  bbandsK: number = 2

  isIchimokuCloud: Boolean = false

  isVolume: Boolean = false

  isRsi: Boolean = false
  rsi: number = 14

  isMacd: Boolean = false
  macd1: number = 12
  macd2: number = 26
  macd3: number = 9

  url: string = "/api/v1/candles"
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

  terms: Array<object> = [
    { label: "After", value: "after" },
    { label: "Before", value: "before" }
  ]

  indicators: Array<object> = [
    { id: 1, name: "SMA" },
    { id: 2, name: "EMA" },
    { id: 3, name: "MACD" },
    { id: 4, name: "Bolinger Bands" },
    { id: 5, name: "Ichimoku Cloud" }
  ]

  selectedPeriods = this.periods[3].value
  selectedTerms = this.terms[0].value

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
      height: 350,
      id: "candles",
      toolbar: {
        autoSelected: "pan",
        show: false
      },
      zoom: {
        enabled: false
      }
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

  public volumeChartSeries = [
    {
      name: "volume",
      data: []
    }
  ]

  // FIXME: Add each data after Get res
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
    },
    {
      name: "EMA1",
      data: []
    },
    {
      name: "EMA2",
      data: []
    },
    {
      name: "EMA3",
      data: []
    },
    {
      name: "BBands Up",
      data: []
    },
    {
      name: "BBands Middle",
      data: []
    },
    {
      name: "BBands Low",
      data: []
    },
    {
      name: "Ichimoku Tenkan",
      data: []
    },
    {
      name: "Ichimoku Kijun",
      data: []
    },
    {
      name: "Ichimoku SenkouA",
      data: []
    },
    {
      name: "Ichimoku SenkouB",
      data: []
    },
    {
      name: "Ichimoku Chikou",
      data: []
    }
  ]

  public rsiLineSeries = [
    {
      name: "RSI",
      data: []
    }
  ]

  public rsiLineChartOptions: any = {
    chart: {
      height: 200,
      type: "line",
      zoom: {
        enabled: true
      }
    },
    annotations: {
      yaxis: [
        {
          y: 30,
          borderColor: "#00E396",
          label: {
            borderColor: "#00E396",
            style: {
              color: "#fff",
              background: "#00E396"
            },
            text: "30%"
          }
        },
        {
          y: 70,
          borderColor: "#00E396",
          label: {
            borderColor: "#00E396",
            style: {
              color: "#fff",
              background: "#00E396"
            },
            text: "70%"
          }
        }
      ]
    },
    dataLabels: {
      enabled: false
    },
    stroke: {
      curve: "straight",
      width: 1
    },
    title: {
      text: "RSI Chart",
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
      min: 0,
      max: 100
    }
  }

  public volumeChartOptions: any = {
    chart: {
      height: 350,
      type: "bar",
      brush: {
        enabled: true,
        target: "candles"
      },
      selection: {
        enabled: true,
        xaxis: {
          min: undefined,
          max: undefined
        },
        fill: {
          color: "#ccc",
          opacity: 0.4
        },
        stroke: {
          color: "#0D47A1"
        }
      }
    },
    dataLabels: {
      enabled: false
    },
    plotOptions: {
      bar: {
        columnWidth: "80%",
        colors: {
          ranges: [
            {
              from: -1000,
              to: 0,
              color: "#F15B46"
            },
            {
              from: 1,
              to: 10000,
              color: "#FEB019"
            }
          ]
        }
      }
    },
    stroke: {
      width: 0
    }
    // xaxis: {
    //   type: "datetime",
    //   axisBorder: {
    //     offsetX: 13
    //   }
    // },
    // yaxis: {
    //   labels: {
    //     show: false
    //   }
    // }
  }

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
      text: "Line Chart",
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
      min: NaN,
      max: NaN
    }
  }

  public macdSeries: any = []

  public macdChartOptions: any = {
    chart: {
      height: 350,
      type: "line"
    },
    stroke: {
      curve: "straight",
      width: 1
    },
    title: {
      text: "MACD",
      align: "left"
    }
    // xaxis: {
    //   type: "datetime"
    // }
  }

  // TODO: uncommentout after app developed
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
      periods: this.selectedPeriods,
      beforeAfter: this.selectedTerms,
      time: unixtimestamp,
      smas: this.isSmas ? "1" : "", // FIXME: 1
      ...(this.isSmas
        ? {
            sma1: this.sma1,
            sma2: this.sma2,
            sma3: this.sma3
          }
        : {}),
      emas: this.isEmas ? "1" : "", // FIXME: 1
      ...(this.isSmas
        ? {
            ema1: this.ema1,
            ema2: this.ema2,
            ema3: this.ema3
          }
        : {}),
      bbands: this.isBBands ? "1" : "", // FIXME: 1
      ...(this.isBBands
        ? {
            bbandsN: this.bbandsN,
            bbandsK: this.bbandsK
          }
        : {}),
      ichimoku: this.isIchimokuCloud ? "1" : "", // FIXME: 1
      rsi: this.isRsi ? "1" : "", // FIXME: 1
      ...(this.isRsi
        ? {
            rsiPeriod: this.rsi
          }
        : {}),
      macd: this.isMacd ? "1" : "", // FIXME: 1
      ...(this.isMacd
        ? {
            macd1: this.macd1,
            macd2: this.macd2,
            macd3: this.macd3
          }
        : {})
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

          if (this.isVolume) {
            const volumes: any = data.map(function(value) {
              return value.volume
            })
            this.volumeChartSeries[0].data = volumes
          }

          if (res.data.smas) {
            const sma1: any = res.data.smas[0].values
            const sma2: any = res.data.smas[1].values
            const sma3: any = res.data.smas[2].values

            this.lineSeries[0].data = [...sma1]
            this.lineSeries[1].data = [...sma2]
            this.lineSeries[2].data = [...sma3]

            // FIXME: Caculate and get data from BE
            const smaYMax = Math.max(...[...sma1, ...sma2, ...sma3])
            const smaYMin = Math.min(...[...sma1, ...sma2, ...sma3])
            this.lineChartOptions.yaxis.max =
              smaYMax > maxYaxis ? smaYMax : maxYaxis
            this.lineChartOptions.yaxis.min =
              smaYMin < minYaxis ? smaYMin : maxYaxis

            this.lineChartOptions.xaxis.categories = chartData.map(function(
              item: any
            ) {
              return item.x
            })
          }

          if (res.data.emas) {
            const ema1: any = res.data.emas[0].values
            const ema2: any = res.data.emas[1].values
            const ema3: any = res.data.emas[2].values

            this.lineSeries[3].data = [...ema1]
            this.lineSeries[4].data = [...ema2]
            this.lineSeries[5].data = [...ema3]

            // FIXME: Caculate and get data from BE
            const emaYMax = Math.max(...[...ema1, ...ema2, ...ema3])
            const emaYMin = Math.min(...[...ema1, ...ema2, ...ema3])
            this.lineChartOptions.yaxis.max =
              emaYMax > maxYaxis ? emaYMax : maxYaxis
            this.lineChartOptions.yaxis.min =
              emaYMin < minYaxis ? emaYMin : maxYaxis

            if (this.lineChartOptions.xaxis.categories.length === 0) {
              this.lineChartOptions.xaxis.categories = chartData.map(function(
                item: any
              ) {
                return item.x
              })
            }
          }

          if (res.data.bbands) {
            const up: any = res.data.bbands.up
            const mid: any = res.data.bbands.mid
            const down: any = res.data.bbands.down

            this.lineSeries[6].data = [...up]
            this.lineSeries[7].data = [...mid]
            this.lineSeries[8].data = [...down]

            // FIXME: Caculate and get data from BE
            const bbandsYMax = Math.max(...up)
            const bbandsYMin = Math.min(...down)
            this.lineChartOptions.yaxis.max =
              bbandsYMax > maxYaxis ? bbandsYMax : maxYaxis
            this.lineChartOptions.yaxis.min =
              bbandsYMin < minYaxis ? bbandsYMin : maxYaxis

            if (this.lineChartOptions.xaxis.categories.length === 0) {
              this.lineChartOptions.xaxis.categories = chartData.map(function(
                item: any
              ) {
                return item.x
              })
            }
          }

          if (res.data.ichimoku) {
            const tenkan: any = res.data.ichimoku.tenkan
            const kijun: any = res.data.ichimoku.kijun
            const senkoua: any = res.data.ichimoku.senkoua
            const senkoub: any = res.data.ichimoku.senkoub
            const chikou: any = res.data.ichimoku.chikou

            this.lineSeries[9].data = [...tenkan]
            this.lineSeries[10].data = [...kijun]
            this.lineSeries[11].data = [...senkoua]
            this.lineSeries[12].data = [...senkoub]
            this.lineSeries[13].data = [...chikou]

            // FIXME: Caculate and get data from BE
            const ichimokuYMax = Math.max(
              ...[...tenkan, ...kijun, ...senkoua, ...senkoub, ...chikou]
            )
            const ichimokuYMin = Math.min(
              ...[...tenkan, ...kijun, ...senkoua, ...senkoub, ...chikou]
            )
            this.lineChartOptions.yaxis.max =
              ichimokuYMax > maxYaxis ? ichimokuYMax : maxYaxis
            this.lineChartOptions.yaxis.min =
              ichimokuYMin < minYaxis ? ichimokuYMin : maxYaxis

            if (this.lineChartOptions.xaxis.categories.length === 0) {
              this.lineChartOptions.xaxis.categories = chartData.map(function(
                item: any
              ) {
                return item.x
              })
            }
          }

          if (res.data.rsi) {
            const rsi: any = res.data.rsi.values

            this.rsiLineSeries[0].data = [...rsi]

            // FIXME: Caculate and get data from BE
            this.rsiLineChartOptions.xaxis.categories = chartData.map(function(
              item: any
            ) {
              return item.x
            })
          }

          if (res.data.macd) {
            // Macd         []float64 `json:"macd,omitempty"`
            // MacdSignal   []float64 `json:"macd_signal,omitempty"`
            // MacdHist     []float64 `json:"macd_hist,omitempty"`
            const macd: any = res.data.macd.macd
            const macdSignal: any = res.data.macd.macd_signal
            const macdHist: any = res.data.macd.macd_hist

            this.macdSeries = [
              {
                name: "MACD",
                type: "line",
                data: [...macd]
              },
              {
                name: "MACD Signal",
                type: "line",
                data: [...macdSignal]
              },
              {
                name: "MACD Histogram",
                type: "column",
                data: [...macdHist]
              }
            ]

            // FIXME: Caculate and get data from BE
            const macdYMax = Math.max(...[...macd, ...macdSignal])
            const macdYMin = Math.min(...[...macd, ...macdSignal])
            console.log(macdYMax)
            console.log(macdYMin)
            // this.macdChartOptions.yaxis.max = macdYMax
            // this.macdChartOptions.yaxis.min = macdYMin
            // this.macdChartOptions.xaxis.categories = chartData.map(function(
            //   item: any
            // ) {
            //   return item.x
            // })

            console.log("this.macdSeries")
            console.log(this.macdSeries)
          }

          this.isLine =
            this.isSmas || this.isEmas || this.isBBands || this.isIchimokuCloud
        }
      })
      .catch(e => {
        alert("Apiの取得に失敗しました")
        console.error(e)
        return false
      })
    this.isLoading = false
  }
}
</script>

<style lang="sass" scoped>
.no-chart-area
  height: 350px
</style>
