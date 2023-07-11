<!--Component to display the prediction tab with the help of Prediction_Fixture component, Prediction_TopScorers component and Prediction_TrophyWinners component-->
<template>
  <Top_Bar />
  <div class="buttons">
    <button class="buttonChoices" type="button" @click="display = 1">
      Fixtures
    </button>
    <button class="buttonChoices" type="button" @click="display = 2">
      Top Scorers
    </button>
    <button class="buttonChoices" type="button" @click="display = 3">
      Trophy
    </button>
  </div>
  <div v-if="display == 1">
    <Prediction_Fixture :userid="id" />
  </div>
  <div v-if="display == 2">
    <!--Checks if the round has started-->
    <Prediction_TopScorers
      :roundOneStarted="this.fixtures[0].status.round_started"
    />
  </div>

  <div v-if="display == 3">
    <!--Checks if the round has started-->
    <Prediction_TrophyWinners
      :userid="id"
      :roundOneStarted="this.fixtures[0].status.round_started"
    />
  </div>
</template>

<script>
import Top_Bar from "../Header/Header.vue";
import axios from "axios";
import Prediction_Fixture from "../Predictions/Prediction_Fixture.vue";
import Prediction_TopScorers from "../Predictions/Prediction_TopScorers.vue";
import Prediction_TrophyWinners from "../Predictions/Prediction_TrophyWinners.vue";
export default {
  components: {
    Prediction_Fixture,
    Top_Bar,
    Prediction_TopScorers,
    Prediction_TrophyWinners,
  },
  data() {
    return {
      display: undefined,
      //Array with the fixtures
      fixtures: [],
      fixturesloaded: false,
    };
  },
  async created() {
    try {
      //Getting the fixtures from backend on create
      const response = await axios.get(this.backendUrl + "/otp/fixtures");
      this.fixtures = response.data;
      this.display = 1;
    } catch (e) {
      console.error(e);
    }
  },
};
</script>
<style scoped>
.buttons {
  display: flex;
  justify-content: space-around;
  text-align: center;
}
.buttonChoices {
  margin-top: 30px;
  font-size: 26px;
  border-radius: 12px;
  color: black;
  border: 2px black;
  width: 15%;
}

.buttonChoices:hover {
  cursor: pointer;
}

@media (max-width: 700px) {
  .buttonChoices {
    width: 30%;
  }
}
</style>
