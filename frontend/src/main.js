import { createApp } from "vue";
import App from "./App.vue";
import router from "./routers";

import "./assets/main.css";

let app = createApp(App);
app.config.globalProperties.backendUrl = "http://10.212.170.213:8080"; //USED FOR BUILD
//app.config.globalProperties.backendUrl = "http://localhost:8080";

/*
 *   Calculates where a prediction has gotten its points from
 */
const GOALDIFFERENCE = 2;
const HOMEGOALS = 1;
const AWAYGOALS = 1;
const HOMETIEAWAY = 2;
const ABOVESEVENPOINTFIVE = 4;
const ABOVEFIVEPOINTFIVE = 2;
const ABOVETHREEPOINTFIVE = 1;

app.config.globalProperties.calculatePoints = function (result, prediction) {
  let points = new Map();

  let resultDiff = result["home-team"].score - result["away-team"].score;
  let predDiff = prediction.home - prediction.away;
  //Setting points for each category
  if (resultDiff == predDiff) {
    points.set("Goal difference", GOALDIFFERENCE);
  }

  if (result["home-team"].score == prediction.home) {
    points.set("Home goals", HOMEGOALS);
  }

  if (result["away-team"].score == prediction.away) {
    points.set("Away goals", AWAYGOALS);
  }

  if (result["away-team"].score > result["home-team"].score) {
    if (prediction.away > prediction.home) {
      points.set("1X2", HOMETIEAWAY);
    }
  } else if (result["away-team"].score < result["home-team"].score) {
    if (prediction.away < prediction.home) {
      points.set("1X2", HOMETIEAWAY);
    }
  } else {
    if (prediction.away == prediction.home) {
      points.set("1X2", HOMETIEAWAY);
    }
  }

  let totalGoals = result["away-team"].score + result["home-team"].score;
  let predictedGoals = prediction.home + prediction.away;

  //Checking how many goals socred in order to give the correct amount of points
  if (totalGoals > 7 && predictedGoals > 7) {
    points.set("Above 7.5 goals", ABOVESEVENPOINTFIVE);
  } else if (totalGoals > 5 && predictedGoals > 5) {
    points.set("Above 5.5 goals", ABOVEFIVEPOINTFIVE);
  } else if (totalGoals > 3 && predictedGoals > 3) {
    points.set("Above 3.5 goals", ABOVETHREEPOINTFIVE);
  }

  return points;
};

app.use(router).mount("#app");
