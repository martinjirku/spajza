<template><div class="container"></div></template>

<style scoped>
.container {
  background: rgb(198, 226, 255);
  background: linear-gradient(
    162deg,
    rgba(198, 226, 255, 1) 0%,
    rgba(198, 239, 255, 1) 67%,
    rgba(0, 187, 255, 1) 100%
  );
  width: 100%;
  height: 100%;
}
</style>

<script lang="ts" setup>
import { ref } from "vue";

const activeIndex2 = ref("1");
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
</script>
