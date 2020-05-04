import VueRouter from "vue-router";

declare module "vue/types/vue" {
  interface Vue {
    $router: VueRouter;
  }
}
