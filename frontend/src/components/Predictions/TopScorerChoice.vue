<!--Component to display yhe topscorer choices-->
<template>
  <div class="players">
    <h4>Which team does your topscorer play for?</h4>
    <select v-model="selectedTeam">
      <!--Looping through AllPossiblePlayers-->
      <option
        v-for="team in this.AllPossiblePlayers"
        :key="team.team_id"
        :value="{ players: team.players, team_id: team.team_id }"
      >
        {{ team.team_name }}
      </option>
    </select>
    <div class="playerSelect" v-if="selectedTeam != undefined">
      <h4>Choose topscorer</h4>
      <select v-model="selectedPlayer" @change="savePrediction">
        <option
          v-for="player in this.selectedTeam.players"
          :value="{ player_id: player.id }"
          :key="player.id"
        >
          {{ player["first-name"] }} {{ player["last-name"] }}
        </option>
      </select>
    </div>
  </div>
</template>

<script scoped>
export default {
  name: "TopScorerChoice",
  emits: ["saveSelectedTopscorers"],
  data() {
    return {
      selectedTeam: undefined,
      selectedPlayer: undefined,
    };
  },
  //setting props
  props: {
    AllPossiblePlayers: {
      required: true,
    },

    Index: undefined,
  },
  methods: {
    //Function to save the topscorers
    savePrediction: function () {
      const info = {
        team_id: this.selectedTeam.team_id,
        player_id: this.selectedPlayer.player_id,
      };
      if (
        info.team_id != undefined &&
        info.player_id != undefined &&
        this.Index != undefined
      ) {
        this.$emit("saveSelectedTopscorers", { p: info, i: this.Index });
      }
    },
  },
};
</script>

<style scoped>
.players {
  display: flex;
  flex-direction: column;
  align-items: center;
  border: 100px;
  border-color: black;
  outline-style: solid;
  height: 100px;
  width: 100%;
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}
</style>
