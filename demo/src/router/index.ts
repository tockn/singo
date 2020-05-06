import Router, { RouteConfig } from "vue-router";
import Home from "../views/Home.vue";
import Room from "../views/Room.vue";
import Vue from "vue";

Vue.use(Router);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/rooms/:id",
    name: "Room",
    component: Room
  }
];

const router = new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
