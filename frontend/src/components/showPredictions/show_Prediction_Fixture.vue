<!--Component to display the fixture for each player in the leagues tab-->
<template>
  <div class="pageHeader">
    <h1 class="testFixturesTitle"></h1>
  </div>
  <div class="matchdayButtons">
    <button type="button" class="matchdayButton" @click="showmatchday = 1">
      Matchday 1
    </button>
    <button type="button" class="matchdayButton" @click="showmatchday = 2">
      Matchday 2
    </button>
    <button type="button" class="matchdayButton" @click="showmatchday = 3">
      Matchday 3
    </button>
  </div>
  <div class="predictionsDisplay" v-if="scoresLoaded">
    <ul class="predictionsList">
      <li
        v-for="fixture in fixtures.slice(
          (showmatchday - 1) * 16,
          16 * showmatchday
        )"
        :key="fixture.id"
        class="item"
      >
        <!--Using Show_matchPrediction to display the match predictions-->
        <Show_matchPrediction
          class="predictionDisplay"
          :Fixture="fixture"
          :Prediction="predictions.get(fixture.id)"
        />
      </li>
    </ul>
  </div>
</template>

<script>
import Show_matchPrediction from "./show_matchPrediction.vue";
import axios from "axios";

// get request fixtures, for å prøve post request kommenter vekk koden under
export default {
  name: "FixtureData",
  data() {
    return {
      //Array with the fixtures
      fixtures: [],
      //Map with the predictions
      predictions: new Map(),
      scoresLoaded: false,
      showmatchday: 1,
    };
  },
  //Setting the props
  props: {
    userid: {
      required: true,
    },
  },
  async created() {
    try {
      //Getting the fixtures from backend with a get request
      const response = await axios.get(this.backendUrl + "/otp/fixtures");
      this.fixtures = response.data;
      //Setting predbuf to the predictions
      let predBuf = await this.getPredictions();
      //Looping through the predictions
      for (let i in predBuf.predictions) {
        this.predictions.set(predBuf.predictions[i].matchid, {
          home: predBuf.predictions[i].homescore,
          away: predBuf.predictions[i].awayscore,
          points: predBuf.predictions[i].points,
        });
      }
      //Setting scoresloaded
      this.scoresLoaded = true;
    } catch (e) {
      console.error(e);
    }
  },

  methods: {
    //Function to get predictions from backend with a post requst
    async getPredictions() {
      const data = {
        userid: parseInt(this.userid),
      };
      const predJSON = JSON.stringify(data);

      const response = await axios.post(
        this.backendUrl + "/otp/getpredictions",
        predJSON
      );
      return response.data;
    },
  },
  components: { Show_matchPrediction },
};
</script>

<style scoped>
.item {
  color: white;
  width: 100%;
  display: flex;
  justify-content: center;
}

.predictionsDisplay {
  display: flex;
  flex-direction: column;
}

.predictionsList {
  display: grid;
  grid-template-columns: 1fr 1fr;
  place-items: center;
  width: 100%;
  margin-top: 20px;
  margin: 5px;
  padding: 0px 0px 0px 0px;
}

.pageHeader {
  display: flex;
  justify-content: center;
}

.matchdayButtons {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  width: 100%;
}

.matchdayButton {
  font-size: 26px;
  border-radius: 12px;
  color: black;
  border: 2px black;
  width: 15%;
}

.matchdayButton:hover {
  cursor: pointer;
  background-color: rgb(47, 62, 163);
  box-shadow: 0 12px 16px 0 rgba(0, 0, 0, 0.24),
    0 17px 50px 0 rgba(0, 0, 0, 0.19);
  color: white;
}

@media (max-width: 850px) {
  .predictionsList {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 750px) {
  .matchdayButtons {
    display: flex;
    flex-direction: column;
    width: 100%;
    align-items: center;
  }
  .matchdayButton {
    width: 50%;
    margin-bottom: 10px;
  }
}
</style>
