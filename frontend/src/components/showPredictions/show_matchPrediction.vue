<template>
  <div class="matchDisplay">
    <!--Checking the status of the match, displaying points if the match is done-->
    <div v-if="Fixture.status.short == 'FT'">
      {{ this.Prediction.points + "p" }}
    </div>
    <div class="fixtureDisplay">
      <h2 class="teamNamesDisplay">
        <div class="homeTeamName">
          <!--Using the FlagLogo component to display the home team logo-->
          <FlagLogo :TeamName="Fixture['home-team'].name" />
          {{ Fixture["home-team"].name }}
        </div>
        <div class="finalResult">
          <!--Displaying the score if the status of the match is full time-->
          <div v-if="Fixture.status.short == 'FT'">
            {{ Fixture["home-team"].score }} - {{ Fixture["away-team"].score }}
          </div>
          <!--A message if the match is not done-->
          <div v-if="Fixture.status.short != 'FT'">
            <h4 class="matchStatus">Match not finished</h4>
          </div>
        </div>
        <div class="awayTeamName">
          {{ Fixture["away-team"].name }}
          <!--Using the FlagLogo component to display the away team logo-->
          <FlagLogo :TeamName="Fixture['away-team'].name" />
        </div>
      </h2>
    </div>
    <div class="predictionDisplay" v-if="predictionExist">
      <div class="predictions">
        <h4 v-if="Fixture.status.round_started">Predicted score:</h4>
        <div v-if="Fixture.status.round_started">
          {{ this.Prediction.home }} - {{ this.Prediction.away }}
        </div>
      </div>
      <div class="scrollListPos">
        <div class="pointcalcScrollList" v-if="Fixture.status.short == 'FT'">
          <!--Loops through points and setting the points the each players leagues page. This outputs the reasons why the player has gotten point and how many points-->
          <div class="pointscalc" v-for="item in this.points" :key="item">
            <div class="points">{{ item[0] + ": " + item[1] + "p" }}</div>
          </div>
        </div>
      </div>
    </div>
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
      predictionExist: false,
      points: undefined,
    };
  },
  //Setting the props
  props: {
    Fixture: {
      required: true,
    },
    Prediction: {
      type: Object,
      required: false,
    },
  },
  async created() {
    try {
      if (this.Prediction) {
        this.predictionExist = true;
        this.points = this.calculatePoints(this.Fixture, this.Prediction);
      }
    } catch (e) {
      console.error(e);
    }
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
  background-color: rgb(179, 177, 175);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  width: 100%;
  height: 170px;
  margin: 7px;
}
.fixtureDisplay {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
}
.teamNamesDisplay {
  display: flex;
  width: 100%;
  flex-direction: row;
  justify-content: space-between;
  font-family: "Franklin Gothic Medium", "Arial Narrow", Arial, sans-serif;
  font-size: 15px;
  color: black;
}

.homeTeamName {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
}

.awayTeamName {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  align-items: center;
}

.finalResult {
  display: flex;
  align-items: center;
  font-size: x-large;
  position: absolute;
  width: 100%;
  flex-direction: row;
  justify-content: center;
  align-items: center;
}

.predictionDisplay {
  color: black;
  font-size: small;
  width: 100%;
  height: 200px;
}

.predictions {
  font-size: medium;
  display: flex;
  width: 100%;
  justify-content: center;
  gap: 20px;
}

.pointscalc {
  display: flex;
  width: 100%;
  justify-content: center;
  gap: 20px;
}

.points {
  display: flex;
  align-items: center;
}

.pointcalcScrollList {
  width: 200px;
  height: 60px;
  overflow: auto;
}

.scrollListPos {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  justify-content: center;
}

@media (max-width: 1350px) {
  .homeTeamName {
    flex-direction: column;
    justify-content: center;
  }
  .awayTeamName {
    flex-direction: column;
    align-items: flex-end;
  }

  .awayTeamName > * {
    order: -1;
  }
}

@media (max-width: 850px) {
  .finalResult {
    font-size: large;
  }
}
</style>
