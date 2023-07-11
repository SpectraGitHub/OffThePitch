<!--Component to display where the user can type in their predictions -->
<template>
  <div class="matchDisplay">
    <h4 class="fixtureDateDisplay">{{ Fixture.date }} {{ Fixture.time }}</h4>
    <h2 class="teamNamesDisplay">
      <div class="HomeTeamName">
        <!--Using the FlagLogo component to display the home team logo-->
        <FlagLogo :TeamName="Fixture['home-team'].name" />
        {{ Fixture["home-team"].name }}
      </div>
      <div class="AwayTeamName">
        {{ Fixture["away-team"].name }}
        <!--Using the FlagLogo component to display the away team logo-->
        <FlagLogo :TeamName="Fixture['away-team'].name" />
      </div>
      <!--Checking if the round has started-->
      <div class="prediction" v-if="Fixture.status.round_started != true">
        <input
          @change="savePrediction"
          v-model="home"
          type="number"
          id="homeScore"
          name="homeScore"
          min="0"
          max="20"
        />
        -
        <input
          @change="savePrediction"
          v-model="away"
          type="number"
          id="awayScore"
          name="awayScore"
          min="0"
          max="20"
        />
      </div>
      <!--Checking if round has started-->
      <div class="prediction" v-if="Fixture.status.round_started == true">
        <h3>Matchday started</h3>
      </div>
    </h2>
  </div>
</template>

<script scoped>
import FlagLogo from "../Flag/FlagLogo.vue";
export default {
  name: "prediction-display",
  emits: ["change"],
  data() {
    return {
      image: FlagLogo,
      homeScore: 0,
      awayScore: 0,
      home: undefined,
      away: undefined,
    };
  },
  //Props so that the component could be used in another component and send a variable back to this one
  props: {
    Fixture: {
      required: true,
    },
    SetScores: {
      required: false,
    },
  },
  //Mointing the saved scores on the screen if the user have set any previous scores
  mounted() {
    if (this.SetScores) {
      if (this.SetScores.get(this.Fixture.id)) {
        this.home = this.SetScores.get(this.Fixture.id).home;
        this.away = this.SetScores.get(this.Fixture.id).away;
      }
    }
  },
  methods: {
    //Function to save the home and away scores for each matchid
    savePrediction: function () {
      const info = {
        matchid: this.Fixture.id,
        homescore: this.home,
        awayscore: this.away,
      };
      if (info.homescore != undefined && info.awayscore != undefined) {
        this.$emit("savePred", info);
      }
    },
  },
  components: { FlagLogo },
};
</script>

<style scoped>
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
.matchDisplay {
  border-style: solid;
  border-width: 5px;
  border-color: blue;
  background-color: antiquewhite;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100px;
  margin: 7px;
}
.teamNamesDisplay {
  display: flex;
  width: 100%;
  flex-direction: row;
  justify-content: space-between;
  font-family: "Franklin Gothic Medium", "Arial Narrow", Arial, sans-serif;
  font-size: clamp(15px, 1.5vw, 40px);
  color: black;
}

.HomeTeamName {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
}

.AwayTeamName {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
}

.fixtureDateDisplay {
  color: black;
  display: flex;
  justify-content: center;
}

.prediction {
  display: flex;
  position: absolute;
  font-size: xx-large;
  width: 100%;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}
.prediction > input {
  width: 10%;
  margin: 10px;
  border-radius: 20px;
  height: 30px;
  text-align: center;
  background-color: rgb(200, 200, 200);
}

.prediction > h3 {
  font-size: clamp(15px, 1.5vw, 40px);
}

@media (max-width: 1350px) {
  .HomeTeamName {
    flex-direction: column;
    justify-content: center;
  }
  .AwayTeamName {
    flex-direction: column;
    align-items: flex-end;
  }

  .AwayTeamName > * {
    order: -1;
  }
  .matchDisplay {
    height: 150px;
  }
}
</style>
