<template>
  <div class="left-column">
    <div class="header">
      <h1>Zasobar</h1>
    </div>
    <div>Menu {{ loggedIn }}</div>
  </div>
</template>

<style scoped>
.header {
  width: 100%;
  margin: 0;
  display: flex;
  justify-content: center;
  background-color: var(--el-color-primary-light-7);
  height: 60px;
}
.header h1 {
  vertical-align: middle;
}
</style>

<script lang="ts" setup>
import { useAuthenticationStore } from "@/auth/authentication";
import { storeToRefs } from "pinia";

const { loggedIn } = storeToRefs(useAuthenticationStore());
</script>
