import Home_Screen from "./components/Website_tabs/Home_tab.vue";
import Leagues_Screen from "./components/Website_tabs/League_tab.vue";
import Fixtures_Screen from "./components/Website_tabs/Fixture_tab.vue";
import Prediction_Screen from "./components/Website_tabs/Prediction_tab.vue";
import Info_Rules_Screen from "./components/Website_tabs/Info_Rules_tab.vue";
import LogIn_Screen from "./components/Website_tabs/logIn_tab.vue";
import CreateNew_Screen from "./components/Website_tabs/createNewUser_tab.vue";
import joinLeague from "./components/Leagues/joinLeague.vue";
import createNewLeague from "./components/Leagues/createNewLeague.vue";
import playerX_predictions from "./components/showPredictions/playerX_predictions.vue";
import { createRouter, createWebHistory } from "vue-router";
/*
 * Setting up the routes for all the different pages
 */
const routes = [
  {
    name: "Home",
    component: Home_Screen,
    path: "/",
  },
  {
    name: "Leagues",
    component: Leagues_Screen,
    path: "/leagues",
  },
  {
    name: "FixtureData",
    component: Fixtures_Screen,
    path: "/fixtures",
  },
  {
    name: "My predictions",
    component: Prediction_Screen,
    path: "/my_predictions",
  },
  {
    name: "Info/rules",
    component: Info_Rules_Screen,
    path: "/info_rules",
  },
  {
    name: "Log in",
    component: LogIn_Screen,
    path: "/login",
  },
  {
    name: "Create new user",
    component: CreateNew_Screen,
    path: "/newuser",
  },
  {
    name: "Join league",
    component: joinLeague,
    path: "/joinleague",
  },
  {
    name: "Create league",
    component: createNewLeague,
    path: "/createleague",
  },
  {
    name: "Player X predictions",
    component: playerX_predictions,
    path: "/predictions/:id",
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

//Route to login if the user is not logged in
router.beforeEach((to, from, next) => {
  if (!sessionStorage.getItem("currentState")) {
    if (to.path == "/my_predictions") {
      alert("Login to enter this page!");
      router.push({ path: "/login" });
    }

    if (to.path == "/leagues") {
      alert("Login to enter this page!");
      router.push({ path: "/login" });
    }

    next();
  } else if (sessionStorage.getItem("currentState")) {
    if (to.path == "/login") {
      router.push({ path: "/" });
    }

    if (to.path == "/newuser") {
      router.push({ path: "/" });
    }

    next();
  }
});

export default router;
