<template>
  <page-container v-if="auth.authStatus === 'checking'">
    <q-inner-loading
      :showing="true"
      label="Please wait..."
      label-class="text-teal"
      label-style="font-size: 1.1em"
    />
  </page-container>
  <router-view v-if="auth.authStatus === 'loaded'"></router-view>
</template>

<script lang="ts" setup>
import PageContainer from "./components/common/PageContainer.vue";
import { useAuthenticationStore } from "./auth/authentication";
const auth = useAuthenticationStore();
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Roboto&display=swap");
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

:root {
  --bg-gradient: linear-gradient(
    113.04deg,
    #e68a8a -0.64%,
    #e6a15c 37.48%,
    #e08053 68.8%,
    #ebeb8f 99.61%
  );
}

body {
  padding: 0;
  margin: 0;
}
.header {
  padding: 0;
}
.el-card {
  --el-card-bg-color: rgba(255, 255, 255, 0.8);
}
a {
  text-decoration: none;
}
a:hover {
  text-decoration: underline;
}
</style>
