<!--Component which displays all the selected top scorers-->
<template>
  <div class="selectedTopScorersBox">
    <h2 class="">Selected top scorers:</h2>
    <div class="topScorers">
      <!--Looping through the topscorers and outputting the topscorer player name, the flag to the team and how many goals-->
      <h3 v-for="topscorer in this.topscorers" :key="topscorer.player_id">
        {{ topscorer.player_name }}
        <FlagLogo :TeamName="topscorer.team_name" /> ({{ topscorer.goals }})
      </h3>
    </div>
  </div>
</template>

<script>
import FlagLogo from "../Flag/FlagLogo.vue";

export default {
  name: "SelectedTopScorers",
  components: { FlagLogo },
  //Props so this component can be used in another component
  props: {
    players: {
      required: true,
    },
  },
  data: function () {
    return {
      topscorers: this.players,
    };
  },
};
</script>

<style scoped>
.selectedTopScorersBox > h2 {
  display: flex;
  justify-content: center;
  margin-bottom: 10px;
}

.topScorers > h3 {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: row;
}

.topScorers {
  margin-top: 30px;
}

.selectedTopScorersBox {
  border: 5px;
  border-color: black;
  border-style: solid;
  width: 100%;
}
</style>
