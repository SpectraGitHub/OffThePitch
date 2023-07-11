<!--Component which displays the fixture on the fixture tab with teh use of the FixtureData from the Fixture.vue-->
<template>
  <Top_Bar />
  <div class="main">
    <FixtureData />
  </div>
</template>

<script>
import Top_Bar from "../Header/Header.vue";
import FixtureData from "../matchFixture/Fixture.vue";
export default {
  name: "Fixtures_Screen",
  components: {
    Top_Bar,
    FixtureData,
  },
};
</script>

<style scoped>
.main {
  width: auto;
  height: auto;
  background: white;
  border: solid 2px #fff;
}
</style>
