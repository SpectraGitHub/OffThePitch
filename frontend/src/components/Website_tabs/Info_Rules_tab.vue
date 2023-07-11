<!--Component which displays the FAQs, Rules and the point calculation for the user. The pointcalculation and the example of the a point calculation are displayed in tables-->
<template>
  <div class="front">
    <Top_Bar />
  </div>
  <div class="main">
    <div class="buttons">
      <button
        @click="(questions = !questions), (rules = false), (pointcalc = false)"
      >
        FAQs
      </button>
      <button
        @click="(rules = !rules), (pointcalc = false), (questions = false)"
      >
        Rules
      </button>
      <button
        @click="(pointcalc = !pointcalc), (questions = false), (rules = false)"
      >
        Point calculations
      </button>
    </div>
    <div v-if="questions" class="FAQs">
      <h3>
        Forgot your password? For now it's not possible to reset your password.
      </h3>
      <h3>
        What is 1x2? 1 is defined as home victory, x is defined as draw and 2 is
        defined as away victory.
      </h3>
      <h3>
        Can not set predictions? If the first matchday has started, you should
        not be able to predict the results for matchday 1 or set predictition
        for trophy winners or topscorers.
      </h3>
      <h3>
        Can not see other players predictions? You can not see other peoples
        predictions before a matchday has started.
      </h3>
      <h3>
        What is a matchday? A matchday is the period all 16 teams in the world
        cup play their first, second or third match.
      </h3>
      <h3>
        Do I need to set all predictions before matchday 1? No you only need to
        set the predictions for matchday 1, the topscorers and the trophy
        winners. Before matchday 2 starts, you need to set the predictions for
        matchday 2, the same with matchday 3 and the playoffs.
      </h3>
      <h3>
        When is the deadline before a matchday? The deadline is the kickoff time
        for the first match in the matchday.
      </h3>
      <h3>
        I am logged in, but the pages has no content. That means that your token
        has expired, log in again and it will work
      </h3>
      <h3>
        How do I get points? It's two tables under the point calculations
        section that shows you how to get points.
      </h3>
      <h3>
        My topscorers is out of world cup, can I change them? No, you can not
        change topscorers if the get injured or if they are out of the
        tournament. You should choose topscorers that you think will get the
        longest in the tournament.
      </h3>
      <h3>
        How do I join a league? All players are automatically joined to the
        overall league, but you can also join other league. Then you need to get
        a league code from someone in the league you want to join.
      </h3>
      <h3>
        How do I create a league? Go to the create new league section, type in a
        league name and invite people to your league with the league code.
      </h3>
    </div>
    <div v-if="rules" class="Rules">
      <h3>This tournament prediction game is for world cup.</h3>
      <h3>
        The goal is to get the maximum points possible to win over your friends.
      </h3>
      <h3>
        Create different leagues and set predictions to get the maximum out of
        the prediction game.
      </h3>
      <h3>
        Points will be calculated automatically and you can see how the points
        are calculated on the points calculation section.
      </h3>
      <h3>The most important rule, have fun!</h3>
    </div>
    <div v-if="pointcalc" class="PointCalc">
      <button
        class="PointsTable"
        @click="(PointsTable = !PointsTable), (ExampleTable = false)"
      >
        Points table 🡇
      </button>
      <table v-if="PointsTable" class="points">
        <tr>
          <th>Event</th>
          <th>Points</th>
        </tr>
        <tr class="middleHeader">
          Top scorers
        </tr>
        <tr>
          <td>Selected top scorers score a goal</td>
          <td>2 point</td>
        </tr>
        <tr>
          <td>Your selected top scorer is the top scorer of the tournament</td>
          <td>10 point</td>
        </tr>
        <tr class="middleHeader">
          Trophy winners
        </tr>
        <tr>
          <td>Correct trophy winner(s)</td>
          <td>12 points</td>
        </tr>
        <tr>
          <td>Correct trophy winner, but wrong trophy</td>
          <td>5 points</td>
        </tr>
        <tr class="middleHeader">
          Match results
        </tr>
        <tr>
          <td>Correct home goals</td>
          <td>1 point</td>
        </tr>
        <tr>
          <td>Correct away goals</td>
          <td>1 point</td>
        </tr>
        <tr>
          <td>Correct goal difference</td>
          <td>2 points</td>
        </tr>
        <tr>
          <td>Correct result (1x2)</td>
          <td>2 point</td>
        </tr>
        <tr>
          <td>Guessed more than 3.5 goals in a match and if that is correct</td>
          <td>1 point</td>
        </tr>
        <tr>
          <td>Guessed more than 5.5 goals in a match and if that is correct</td>
          <td>2 point</td>
        </tr>
        <tr>
          <td>Guessed more than 7.5 goals in a match and if that is correct</td>
          <td>4 point</td>
        </tr>
      </table>
      <button
        @click="(ExampleTable = !ExampleTable), (PointsTable = false)"
        class="ExamplesTable"
      >
        Example calculation of points 🡇
      </button>
      <table v-if="ExampleTable" class="example">
        <tr>
          <th>Guessed result</th>
          <th>Actual result</th>
          <th>Points</th>
          <th>Explanation</th>
        </tr>
        <tr>
          <td>2-0</td>
          <td>1-0</td>
          <td>2 + 1 = 3</td>
          <td>Correct 1x2 + correct away goals</td>
        </tr>
        <tr>
          <td>4-1</td>
          <td>3-0</td>
          <td>2 + 2 = 4</td>
          <td>Correct goal difference + correct 1x2</td>
        </tr>
        <tr>
          <td>9-1</td>
          <td>4-4</td>
          <td>4</td>
          <td>Correct more than 7.5 goals</td>
        </tr>
        <tr>
          <td>8-0</td>
          <td>0-9</td>
          <td>4</td>
          <td>Correct more than 7.5 goals</td>
        </tr>
        <tr>
          <td>1-1</td>
          <td>4-4</td>
          <td>2 + 2 = 4</td>
          <td>Correct goal difference + correct 1x2</td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import Top_Bar from "../Header/Header.vue";

export default {
  name: "Info_Rules_Screen",
  data: function () {
    return {
      pointcalc: false,
      questions: false,
      rules: false,
      PointsTable: false,
      ExampleTable: false,
    };
  },
  components: {
    Top_Bar,
  },
  methods: {
    onClick() {
      location.href = "/";
    },
  },
};
</script>

<style scoped>
.main {
  width: auto;
  height: auto;
  padding-bottom: 0;
  background: white;
}

.PointsTable {
  top: 20px;
  font-size: xx-large;
}

.PointsTable:hover {
  cursor: pointer;
}

.ExamplesTable {
  top: 30px;
  font-size: xx-large;
}
.ExamplesTable:hover {
  cursor: pointer;
}

.PointCalc > button {
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 100%;
  border: 0px;
  padding: 0px 0px 0px 25px;
  height: 40px;
  background: white;
}

.middleHeader {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: x-large;
}

.points {
  font-family: Helvetica, sans-serif;
  border-collapse: collapse;
  width: 70%;
  margin: 30px;
}

.points td,
.points th {
  border: 1px solid #ddd;
  padding: 8px;
}

.points tr:nth-child(even) {
  background-color: #f2f2f2;
}

.points tr:hover {
  background-color: #ddd;
}

.points th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: rgb(47, 62, 163);
  color: white;
}

.example {
  font-family: Helvetica, sans-serif;
  border-collapse: collapse;
  width: 70%;
  margin: 30px;
}

.example td,
.example th {
  border: 1px solid #ddd;
  padding: 8px;
}

.example tr:nth-child(even) {
  background-color: #f2f2f2;
}

.example tr:hover {
  background-color: #ddd;
}

.example th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: rgb(47, 62, 163);
  color: white;
}

.buttons {
  width: 40%;
  display: flex;
  flex-direction: row;
  padding: 0px 0px 0px 25px;
  margin-top: 30px;
}

.buttons > button {
  font-size: 30px;
  height: 70px;
  width: 100%;
  margin: 0px;
  border-radius: 10px;
}

.buttons > button:hover {
  cursor: pointer;
  background-color: rgb(82, 82, 82);
  color: white;
}

.FAQs > h3 {
  border: solid 2px black;
  margin: 10px;
}

.Rules > h3 {
  border: solid 2px black;
  margin: 10px;
}

@media (max-width: 500px) {
  .buttons {
    display: flex;
    flex-direction: column;
    width: 80%;
  }
  .buttons > button {
    margin-bottom: 10px;
  }
  .ExamplesTable {
    margin-top: 50px;
    margin-bottom: 50px;
  }
}
</style>
