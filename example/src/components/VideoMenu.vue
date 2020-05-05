<template>
  <div class="buttons" ref="buttons" :style="buttonsStyle">
    <v-btn
      class="mx-2"
      fab
      dark
      large
      color="cyan"
      ref="openButton"
      @mousedown="startDragging"
      @touchstart="startDragging"
      @click="openClicked"
    >
      <font-awesome-icon icon="bars" class="icon" />
    </v-btn>
    <div v-show="opened">
      <v-btn class="mx-2 my-2" fab dark color="cyan" @click="muteButtonClicked">
        <font-awesome-icon v-show="!muted" icon="microphone" class="icon" />
        <font-awesome-icon
          v-show="muted"
          icon="microphone-slash"
          class="icon"
        />
      </v-btn>
      <v-btn
        class="mx-2 my-2"
        fab
        dark
        color="cyan"
        @click="videoButtonClicked"
      >
        <font-awesome-icon v-show="videoOn" icon="video" class="icon" />
        <font-awesome-icon v-show="!videoOn" icon="video-slash" class="icon" />
      </v-btn>
      <v-btn
        class="mx-2 my-2"
        fab
        dark
        color="cyan"
        @click="leaveButtonClicked"
      >
        <font-awesome-icon icon="sign-out-alt" class="icon" />
      </v-btn>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator";

@Component
export default class VideoMenu extends Vue {
  @Prop()
  public muted: boolean;
  @Prop()
  public videoOn: boolean;
  @Prop()
  onMuteStatusChanged: (muted: boolean) => Function;
  @Prop()
  onVideoStatusChanged: (status: boolean) => Function;
  @Prop()
  onLeaveClicked: () => Function;

  private dragging = false;
  private buttonsStyle = "";
  private ref: Element;
  private openButtonRef: Element;
  private opened = true;
  private refHeight = 0;
  private onMouseMoved = false;

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
    if (this.onMouseMoved) {
      this.onMouseMoved = false;
      return;
    }
    this.opened = !this.opened;
  }
  muteButtonClicked() {
    this.onMuteStatusChanged(!this.muted);
  }
  videoButtonClicked() {
    this.onVideoStatusChanged(!this.videoOn);
  }
  leaveButtonClicked() {
    this.onLeaveClicked();
  }

  private startDragging() {
    this.dragging = true;
  }
  private endDragging() {
    this.dragging = false;
  }

  get leftMax(): number {
    return window.innerWidth - this.openButtonRef.clientWidth;
  }
  get topMax(): number {
    return (
      window.innerHeight - this.refHeight - this.openButtonRef.clientHeight / 2
    );
  }
  get leftMin(): number {
    return 0;
  }
  get topMin(): number {
    return 0;
  }
  private mouseMove(e: MouseEvent) {
    if (!this.dragging) return;
    this.onMouseMoved = true;
    this.updateButtonsPosition(e.clientX, e.clientY);
  }
  private touchMove(e: TouchEvent) {
    if (!this.dragging) return;
    this.updateButtonsPosition(e.touches[0].clientX, e.touches[0].clientY);
  }
  private updateButtonsPosition(mouseX: number, mouseY: number) {
    this.ref = this.$refs.buttons as Element;
    this.refHeight = this.ref.clientHeight;
    let top = mouseY - this.openButtonRef.clientHeight / 2;
    let left = mouseX - this.openButtonRef.clientWidth / 2;
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
  z-index: 123456;
}
.icon {
  font-size: 24px;
}
</style>
