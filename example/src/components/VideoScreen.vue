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

  mounted() {
    const el = this.$refs[this.refName] as HTMLVideoElement;
    el.srcObject = this.stream;
  }

  get refName(): string {
    return `screen-${this.screenId}`;
  }

  @Watch("stream")
  onChangeStream() {
    const el = this.$refs[this.refName] as HTMLVideoElement;
    el.srcObject = this.stream;
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
