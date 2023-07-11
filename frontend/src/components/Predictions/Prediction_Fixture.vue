<!--This component is similar to the Fixture component, but it's for the prediciton page where the user can type in predictions-->
<template>
  <div class="pageHeader">
    <h1 class="testFixturesTitle">Make your predictions!</h1>
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
    <button type="button" class="matchdayButton" @click="showmatchday = 4">
      Playoffs
    </button>
  </div>
  <div class="predictionsDisplay">
    <!--If the playoffs button is not pressed-->
    <ul v-if="showmatchday != 4" class="predictionsList">
      <!--Looping through the fixtures from (x-1)*16 -->
      <li
        v-for="fixture in fixtures.slice(
          (showmatchday - 1) * 16,
          16 * showmatchday
        )"
        :key="fixture.id"
        class="item"
      >
        <!--Using the PredictionDisplay compontent and sending in fixture and scores. The compontent should only display of scoresLoaded is true-->
        <PredictionDisplay
          @savePred="savePrediction"
          class="predictionDisplay"
          :Fixture="fixture"
          :SetScores="scores"
          v-if="scoresLoaded"
        />
      </li>
    </ul>

    <!-- Playoff Tabs below-->
    <div v-if="showmatchday === 4">
      <h1 class="playoffTitle">Round of 16</h1>
      <!--Checking the length-->
      <ul v-if="fixtures.length >= 56" class="predictionsList">
        <!--Looping through the round of 16 matches-->
        <li
          v-for="fixture in fixtures.slice(48, 56)"
          :key="fixture.id"
          class="item"
        >
          <!--Using the PredictionDisplay compontent and sending in fixture and scores. The compontent should only display of scoresLoaded is true-->
          <PredictionDisplay
            @savePred="savePrediction"
            class="predictionDisplay"
            :Fixture="fixture"
            :SetScores="scores"
            v-if="scoresLoaded"
          />
        </li>
      </ul>

      <h1 class="playoffTitle">Quarter Finals</h1>
      <!--Checking the length of fixtures-->
      <ul v-if="fixtures.length >= 60" class="predictionsList">
        <!--Looping through the quarter finals matches-->
        <li
          class="item"
          v-for="fixture in fixtures.slice(56, 60)"
          :key="fixture.id"
        >
          <!--Using the PredictionDisplay compontent and sending in fixture and scores. The compontent should only display of scoresLoaded is true-->
          <PredictionDisplay
            @savePred="savePrediction"
            class="predictionDisplay"
            :Fixture="fixture"
            :SetScores="scores"
            v-if="scoresLoaded"
          />
        </li>
      </ul>

      <h1 class="playoffTitle">Semi Finals</h1>
      <!--Checking the fixture length-->
      <ul v-if="fixtures.length >= 62" class="predictionsList">
        <!--Looping through the semi finals matches-->
        <li
          class="item"
          v-for="fixture in fixtures.slice(60, 62)"
          :key="fixture.id"
        >
          <!--Using the PredictionDisplay compontent and sending in fixture and scores. The compontent should only display of scoresLoaded is true-->
          <PredictionDisplay
            @savePred="savePrediction"
            class="predictionDisplay"
            :Fixture="fixture"
            :SetScores="scores"
            v-if="scoresLoaded"
          />
        </li>
      </ul>

      <h1 class="playoffTitle">Third Place Finals</h1>
      <!--Checking the length to display the third place finals-->
      <ul v-if="fixtures.length >= 63" class="predictionsList">
        <li
          class="item"
          v-for="fixture in fixtures.slice(62, 63)"
          :key="fixture.id"
        >
          <!--Using the PredictionDisplay compontent and sending in fixture and scores. The compontent should only display of scoresLoaded is true-->
          <PredictionDisplay
            @savePred="savePrediction"
            class="predictionDisplay"
            :Fixture="fixture"
            :SetScores="scores"
            v-if="scoresLoaded"
          />
        </li>
      </ul>

      <h1 class="playoffTitle">Finals</h1>
      <!--Checking for the finals-->
      <ul v-if="fixtures.length >= 64" class="predictionsList">
        <li
          class="item"
          v-for="fixture in fixtures.slice(63, 64)"
          :key="fixture.id"
        >
          <!--Using the PredictionDisplay compontent and sending in fixture and scores. The compontent should only display of scoresLoaded is true-->
          <PredictionDisplay
            @savePred="savePrediction"
            class="predictionDisplay"
            :Fixture="fixture"
            :SetScores="scores"
            v-if="scoresLoaded"
          />
        </li>
      </ul>
    </div>

    <div class="saveButtonDiv">
      <button class="savePredictionsButton" @click="submitPredictions()">
        Save Predictions
      </button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import PredictionDisplay from "./PredictionDisplay.vue";

const PREDICTIONS = [];

// get request fixtures, for å prøve post request kommenter vekk koden under
export default {
  name: "FixtureData",
  data() {
    return {
      //Array with all the matches
      fixtures: [],
      //Scores is returned as a map
      scores: new Map(),
      scoresLoaded: false,
      //Automatically set to show the first matchday on the prediction page
      showmatchday: 1,
      //Empty struct of predictions which gets updated with information from backend
      predictions: {
        matchid: "",
        homescore: "",
        awayscore: "",
      },
      //Array with the saved predictions for this user
      savedPredictions: [],
    };
  },
  async created() {
    try {
      //Get request from backend to get the fixtures
      const response = await axios.get(this.backendUrl + "/otp/fixtures");
      //Setting the fixtures variable
      this.fixtures = response.data;
      //Getting the scores and setting it to the score buffer variable
      let scoreBuf = await this.getScores();
      //Looping through the scoreBuf and setting the home and away scores for the different matchIds
      for (let i in scoreBuf) {
        this.scores.set(scoreBuf[i].matchid, {
          home: scoreBuf[i].homescore,
          away: scoreBuf[i].awayscore,
        });
      }
      //Setting scores loaded to through after all the scores is loaded
      this.scoresLoaded = true;
    } catch (e) {
      console.error(e);
    }
  },

  methods: {
    //Function that gets all the users predicitons
    async getScores() {
      try {
        //Setting the headers for the get request
        axios.defaults.headers.get["token"] = this.getToken();

        //Get request from backend
        const response = await axios.get(
          this.backendUrl + "/otp/getpredictions"
        );
        return response.data["predictions"];
      } catch (e) {
        console.error(e);
        //Redirect to login, clearing session storage and outputting a message to the user when the session has expired.
        location.href = "/login";
        sessionStorage.clear();
        alert("Session expired, please login again");
      }
    },
    //Getting the current session token
    getToken() {
      return sessionStorage.getItem("token");
    },
    //Saving the prediction
    savePrediction(pred) {
      //Looping through the predictions array
      for (let i = 0; i < PREDICTIONS.length; i++) {
        if (PREDICTIONS[i].matchid == pred.matchid) {
          PREDICTIONS.splice(i, 1);
        }
      }
      //Pusing the pred to the predicitons
      PREDICTIONS.push(pred);
    },
    //Submits the predictions when the button is pressed
    submitPredictions() {
      const data = {
        predictions: PREDICTIONS,
      };
      //Checking if there are set any predictions
      if (PREDICTIONS.length == 0) {
        alert("Please set some predictions before submitting");
        return;
      }
      //Looping through PREDICTIONS to check for negative numbers
      for (let i = 0; i < PREDICTIONS.length; i++) {
        if (PREDICTIONS[i].homescore < 0 || PREDICTIONS[i].awayscore < 0) {
          alert("Please do not enter negative numbers.");
          return;
        }
      }
      //Setting the headers for the post request
      axios.defaults.headers.post["token"] = this.getToken();

      const scoreJSON = JSON.stringify(data);
      try {
        //Post request to the backend
        axios
          .post(this.backendUrl + "/otp/savepredictions", scoreJSON)
          .then(function () {}.bind(this), alert("Predictions saved"));
      } catch (e) {
        console.error(e);
        //Relocate to login, clearing session storage and outputting a message to the user if the session has expired
        location.href = "/login";
        sessionStorage.clear();
        alert("Session expired, please login again");
      }
    },
  },

  components: { PredictionDisplay },
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

.playoffTitle {
  width: 100%;
  text-align: center;
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

.saveButtonDiv {
  position: relative;
  display: flex;
  height: 75px;
  justify-content: center;
}

.savePredictionsButton {
  text-align: center;
  font-size: x-large;
  width: 50%;
  border-radius: 20px;
}

.savePredictionsButton:hover {
  cursor: pointer;
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

@media (min-width: 850px) {
  .matchdayButtons {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
    width: 100%;
  }
}
</style>
