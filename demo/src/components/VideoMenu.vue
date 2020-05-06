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
    <transition name="top-left">
      <div v-show="opened">
        <v-btn
          class="mx-2 my-2"
          fab
          dark
          color="cyan"
          @click="muteButtonClicked"
        >
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
          <font-awesome-icon
            v-show="!videoOn"
            icon="video-slash"
            class="icon"
          />
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
    </transition>
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
  private viewPortWidth: number;
  private viewPortHeight: number;

  mounted() {
    document.ontouchmove = this.touchMove;
    document.onmousemove = this.mouseMove;
    document.ontouchend = this.endDragging;
    document.onmouseup = this.endDragging;
    this.ref = this.$refs.buttons as Element;
    const vel = this.$refs.openButton as Vue;
    this.openButtonRef = vel.$el;
    this.viewPortWidth = window.innerWidth;
    this.viewPortHeight = window.innerHeight;
    this.setDefaultPosition();
    window.onorientationchange = () => {
      this.viewPortWidth = window.innerHeight;
      this.viewPortHeight = window.innerWidth;
      this.setDefaultPosition();
    };
  }
  setDefaultPosition() {
    this.updateButtonsPosition(0, 99999);
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

  // getにすると、this.viewPortWidthを書き換えても発火しないので関数にする
  calculateLeftMax(): number {
    return this.viewPortWidth - this.openButtonRef.clientWidth;
  }
  calculateTopMax(): number {
    return (
      this.viewPortHeight - this.refHeight - this.openButtonRef.clientHeight / 2
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
    if (left > this.calculateLeftMax()) left = this.calculateLeftMax();
    if (left < this.leftMin) left = this.leftMin;
    if (top > this.calculateTopMax()) top = this.calculateTopMax();
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
.top-left-enter-active,
.top-left-leave-active {
  transform: translate(0px, 0px);
  will-change: opacity;
  transition: transform 225ms cubic-bezier(0, 0, 0.2, 1) 0ms,
    opacity 225ms cubic-bezier(0.4, 0, 0.2, 1) 0ms;
}

.top-left-enter,
.top-left-leave-to {
  opacity: 0;
  transform: translateY(0) translateX(-30%);
}
</style>
