<template>
  <div class="home" ref="home">
    <media-query class="middle">
      <div class="join-field">
        <v-text-field
          v-model="roomName"
          class="room-name"
          placeholder="Room Name"
          outlined
          @keydown.enter="join"
        />
        <v-btn class="join-button" @click="join">Join</v-btn>
      </div>
    </media-query>
    <video-menu />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import MediaQuery from "@/components/MediaQuery.vue";
import VideoMenu from "@/components/VideoMenu.vue";
import { disableBodyScroll } from 'body-scroll-lock';
@Component({
  components: { VideoMenu, MediaQuery }
})
export default class Home extends Vue {
  public roomName = "";
  private ref: Element;

  mounted () {
    this.ref = this.$refs.home as Element;
    disableBodyScroll(this.ref);
  }

  public join() {
    if (this.roomName === "") {
      alert("error: Room Name is empty");
      return;
    }
    this.$router.push(`/rooms/${this.roomName}`);
  }
}
</script>

<style>
.middle {
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
