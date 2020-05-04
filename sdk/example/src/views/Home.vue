<template>
  <div class="home">
    <media-query>
      <div class="join-field">
        <v-text-field
          v-model="roomName"
          class="room-name"
          placeholder="Room Name"
          outlined
        />
        <v-btn class="join-button" @click="join">Join</v-btn>
      </div>
    </media-query>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MediaQuery from "@/components/MediaQuery.vue";
@Component({
  components: {MediaQuery}
})
export default class Home extends Vue {
  public roomName = '';
  public join() {
    this.$router.push(`/rooms/${this.roomName}`)
  }
}
</script>

<style>
  .home {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    margin: auto;
    height: 3.2rem;
  }
  .join-field {
    display: flex;
    justify-content: space-around;
  }
  .join-button {
    height: 56px !important;
  }
</style>
