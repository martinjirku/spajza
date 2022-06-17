<template>
  <div class="page-container">
    <slot></slot>
  </div>
</template>
<script lang="ts" setup></script>
<style>
.page-container {
  background: var(--bg-gradient);
  height: 100vh;
  width: 100vw;
}
</style>
