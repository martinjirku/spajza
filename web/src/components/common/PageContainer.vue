<template>
  <div class="page-container">
    <slot></slot>
  </div>
</template>
<script lang="ts" setup></script>
<style>
.page-container {
  background: linear-gradient(
    113.04deg,
    #e68a8a -0.64%,
    #e6a15c 37.48%,
    #e08053 68.8%,
    #ebeb8f 99.61%
  );
  height: 100vh;
  width: 100vw;
}
</style>
