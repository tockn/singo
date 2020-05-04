<template>
  <div class="media-query">
    <slot />
  </div>
</template>

<script>
export default {
  name: "MediaQuery"
};
</script>

<style scoped>
.media-query {
  width: 640px;
  margin: auto;
}
@media screen and (max-width: 768px) {
  .media-query {
    width: 640px;
  }
}
@media screen and (max-width: 480px) {
  .media-query {
    width: 98%;
  }
}
@media screen and (max-width: 320px) {
  .media-query {
    width: 98%;
  }
}
</style>
