<template>
  <v-app>
    <router-view></router-view>
  </v-app>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Roboto&display=swap");
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
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
