<template><div class="container"></div></template>

<style scoped>
.container {
  background: rgb(198, 226, 255);
  background: var(--bg-gradient);
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
