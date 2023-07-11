<!--Component to show score prediction, topscorer prediction and trophy winners for the selected player in the league-->
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
    <!--Using the Show_Prediciton_Fixture component to show the predictions-->
    <Show_Prediction_Fixture :userid="id" />
  </div>
  <div v-if="display == 2">
    <!--Checks if the round has started, then it uses Show_Prediction_TopScorers to display the predicted topscorers if the round has started-->
    <div v-if="this.fixtures[0].status.round_started">
      <Show_Prediction_TopScorers :userid="id" />
    </div>
    <div v-if="!this.fixtures[0].status.round_started">
      <h2 class="roundnotstarted">
        You can not see others top scorers before the first matchday has started
      </h2>
    </div>
  </div>
  <div v-if="display == 3">
    <!--Checks if the round has started, then it uses Show_Prediction_TrophyWinners to display the predicted trophywinners if the round has started-->
    <div v-if="this.fixtures[0].status.round_started">
      <Show_Prediction_TrophyWinners :userid="id" />
    </div>
    <div v-if="!this.fixtures[0].status.round_started">
      <h2 class="roundnotstarted">
        You can not see others trophy winners before the first matchday has
        started
      </h2>
    </div>
  </div>
</template>

<script>
import Top_Bar from "../Header/Header.vue";
import axios from "axios";
import Show_Prediction_Fixture from "./show_Prediction_Fixture.vue";
import Show_Prediction_TopScorers from "./show_Prediction_TopScorers.vue";
import Show_Prediction_TrophyWinners from "./show_Prediction_TrophyWinners.vue";
export default {
  components: {
    Top_Bar,
    Show_Prediction_Fixture,
    Show_Prediction_TrophyWinners,
    Show_Prediction_TopScorers,
  },
  data() {
    return {
      //Array with the fixtures
      fixtures: [],
      display: undefined,
    };
  },
  //Using id as props to get the users id when the name is pressed
  props: {
    id: {
      required: true,
    },
  },
  async created() {
    try {
      //getting the fixtures from backend on create
      this.display = 1;
      const response = await axios.get(this.backendUrl + "/otp/fixtures");
      this.fixtures = response.data;
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

.roundnotstarted {
  font-size: x-large;
  margin: 20px;
}

@media (max-width: 700px) {
  .buttonChoices {
    width: 30%;
  }
  .roundnotstarted {
    font-size: medium;
  }
}
</style>
