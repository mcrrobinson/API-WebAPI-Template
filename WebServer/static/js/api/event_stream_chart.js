var valueChartCounter;
chartColor = "#FFFFFF";

var ctx = document.getElementById('bigDashboardChart').getContext("2d");

var gradientStroke = ctx.createLinearGradient(500, 0, 100, 0);
gradientStroke.addColorStop(0, '#80b6f4');
gradientStroke.addColorStop(1, chartColor);

var gradientFill = ctx.createLinearGradient(0, 200, 0, 50);
gradientFill.addColorStop(0, "rgba(128, 182, 244, 0)");
gradientFill.addColorStop(1, "rgba(255, 255, 255, 0.24)");

var numsamples = 60;

var myChartData = {
  labels: [],
  datasets: [{
    label: "",
    borderColor: chartColor,
    pointBorderColor: chartColor,
    pointBackgroundColor: "#1e3d60",
    pointHoverBackgroundColor: "#1e3d60",
    pointHoverBorderColor: chartColor,
    pointBorderWidth: 1,
    pointHoverRadius: 7,
    pointHoverBorderWidth: 2,
    pointRadius: 0,
    pointHitRadius: 10,
    data: [],
    fill: true,
    backgroundColor: gradientFill,
    borderWidth: 2,
    data: []
  }]
}

var myChartOptions = {
  layout: {
    padding: {
      left: 20,
      right: 20,
      top: 0,
      bottom: 0
    }
  },
  showLines: true,
  animation: { duration: 500, easing: 'linear' },
  maintainAspectRatio: false,
  tooltips: false,
  legend: {
    position: "bottom",
    fillStyle: "#FFF",
    display: false
  },
  scales: {
    yAxes: [{
      ticks: {
        fontColor: "rgba(255,255,255,0.4)",
        fontStyle: "bold",
        beginAtZero: true,
        maxTicksLimit: 5,
        padding: 10,
        min: 0
      },
      gridLines: {
        drawTicks: true,
        drawBorder: false,
        display: true,
        color: "rgba(255,255,255,0.1)",
        zeroLineColor: "transparent"
      }

    }],
    xAxes: [{
      gridLines: {
        zeroLineColor: "transparent",
        display: false,

      },
      ticks: {
        padding: 10,
        fontColor: "rgba(255,255,255,0.4)",
        fontStyle: "bold"
      }
    }]
  }
}


for (var i = 0; i < numsamples; i++) {
  myChartData.labels.push('');
  myChartData.datasets[0].data.push(null);
}

var myChart = new Chart(ctx, {
  type: 'line',
  data: myChartData,
  options: myChartOptions
});

setInterval(function randomdata() {
  myChart.data.datasets[0].data.shift();
  myChart.data.labels.shift();

  var ts = new Date();
  var seconds = ts.getSeconds();
  label = ts.getHours() + ":" + ts.getMinutes() + ":" + seconds;
    
  myChart.data.datasets[0].data.push(valueChartCounter);
  myChart.data.labels.push(label);
  valueChartCounter = 0;
  myChart.update();

}, 1000);