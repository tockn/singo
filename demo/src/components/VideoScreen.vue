<template>
  <video :ref="refName" autoplay playsinline muted></video>
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from "vue-property-decorator";

@Component
export default class VideoScreen extends Vue {
  @Prop({ required: true })
  public screenId: string;
  @Prop({ required: true })
  public stream: MediaStream;
  @Prop({ required: true })
  public muted: boolean;

  private video: HTMLVideoElement;

  mounted() {
    this.video = this.$refs[this.refName] as HTMLVideoElement;
    this.video.srcObject = this.stream;
    this.video.muted = this.muted;
  }

  get refName(): string {
    return `screen-${this.screenId}`;
  }

  @Watch("stream")
  onChangeStream() {
    this.video.srcObject = this.stream;
  }

  @Watch("muted")
  onChangeMuted() {
    this.video.muted = this.muted;
  }
}
</script>

<style scoped>
video {
  width: 50vw;
  height: 50vh;
}
@media screen and (max-width: 480px) {
  video {
    width: 100vw;
    height: 25vh;
  }
}
</style>
