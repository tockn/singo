<template>
  <v-app class="app">
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        singo
      </div>
    </v-app-bar>

    <v-content>
      <router-view />
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "App",

  components: {},

  data: () => ({
    //
  })
});
</script>

<style>
  .app {
  }
</style>
