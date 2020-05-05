<template>
  <div
    class="buttons"
    ref="buttons"
    @mousedown="startDragging"
    @mouseup="endDragging"
    :style="buttonsStyle"
  >
    <v-btn class="mx-2" fab dark large color="cyan">
      <v-icon dark>mdi-pencil</v-icon>
    </v-btn>
    <v-btn class="mx-2" fab dark large color="cyan">
      <v-icon dark>mdi-pencil</v-icon>
    </v-btn>
    <v-btn class="mx-2" fab dark large color="cyan">
      <v-icon dark>mdi-pencil</v-icon>
    </v-btn>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class VideoMenu extends Vue {
  private dragging = false;
  private buttonsStyle = "";
  private ref: Element;

  mounted() {
    document.onmousemove = this.mouseMove;
    this.ref = this.$refs.buttons as Element;
  }
  private startDragging() {
    this.dragging = true;
  }
  private endDragging() {
    this.dragging = false;
  }
  private mouseMove(e: MouseEvent) {
    if (!this.dragging) return;
    const top = e.clientY - this.ref.clientHeight / 2;
    const left = e.clientX - this.ref.clientWidth / 2;
    this.buttonsStyle = `top: ${top}px; left: ${left}px`;
  }
}
</script>

<style scoped>
.buttons {
  position: absolute;
  top: 0;
  left: 0;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
</style>
