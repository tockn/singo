<template>
  <div
    class="buttons"
    ref="buttons"
    :style="buttonsStyle"
  >
    <v-btn
      class="mx-2"
      fab dark large
      color="cyan"
      ref="openButton"
      @mousedown="startDragging"
      @touchstart="startDragging"
      @click="openClicked"
    >
      <font-awesome-icon icon="bars" class="icon"/>
    </v-btn>
    <div v-show="opened">
      <v-btn class="mx-2 my-2" fab dark color="cyan">
        <font-awesome-icon icon="microphone" class="icon"/>
      </v-btn>
      <v-btn class="mx-2 my-2" fab dark color="cyan">
        <font-awesome-icon icon="video" class="icon"/>
      </v-btn>
      <v-btn class="mx-2 my-2" fab dark color="cyan">
        <font-awesome-icon icon="sign-out-alt" class="icon"/>
      </v-btn>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class VideoMenu extends Vue {
  private dragging = false;
  private buttonsStyle = "";
  private ref: Element;
  private openButtonRef: Element;
  private opened = false;
  private refHeight = 0;

  mounted() {
    document.ontouchmove = this.touchMove;
    document.onmousemove = this.mouseMove;
    document.ontouchend = this.endDragging;
    document.onmouseup = this.endDragging;
    this.ref = this.$refs.buttons as Element;
    const vel = this.$refs.openButton as Vue;
    this.openButtonRef = vel.$el;
  }
  openClicked() {
    this.opened = !this.opened;
  }
  private startDragging(e: TouchEvent | MouseEvent) {
    this.dragging = true;
  }
  private endDragging() {
    this.dragging = false;
  }

  get leftMax(): number {
    return window.innerWidth - this.openButtonRef.clientWidth;
  }
  get topMax(): number {
    return window.innerHeight - this.refHeight - this.openButtonRef.clientHeight / 2;
  }
  get leftMin(): number {
    return 0;
  }
  get topMin(): number {
    return 0;
  }
  private mouseMove(e: MouseEvent) {
    if (!this.dragging) return;
    this.ref = this.$refs.buttons as Element;
    this.refHeight = this.ref.clientHeight;
    let top = e.clientY - this.openButtonRef.clientHeight / 2;
    let left = e.clientX - this.openButtonRef.clientWidth / 2;
    if (left > this.leftMax) left = this.leftMax;
    if (left < this.leftMin) left = this.leftMin;
    if (top > this.topMax) top = this.topMax;
    if (top < this.topMin) top = this.topMin;
    this.buttonsStyle = `top: ${top}px; left: ${left}px`;
  }
  private touchMove(e: TouchEvent) {
    if (!this.dragging) return;
    this.ref = this.$refs.buttons as Element;
    this.refHeight = this.ref.clientHeight;
    let top = e.touches[0].clientY - this.openButtonRef.clientHeight / 2;
    let left = e.touches[0].clientX - this.openButtonRef.clientWidth / 2;
    if (left > this.leftMax) left = this.leftMax;
    if (left < this.leftMin) left = this.leftMin;
    if (top > this.topMax) top = this.topMax;
    if (top < this.topMin) top = this.topMin;
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
.icon {
    font-size: 24px;
  }
</style>
