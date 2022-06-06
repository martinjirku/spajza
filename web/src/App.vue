<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import TopHeader from "./components/common/TopHeader.vue";
import LeftColumn from "./components/common/LeftColumn.vue";
</script>

<template>
  <v-app>
    <router-view></router-view>
  </v-app>
</template>
<!-- <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <LeftColumn></LeftColumn>
      </el-aside>
      <el-container>
        <el-header class="header">
          <TopHeader></TopHeader>
        </el-header>
        <el-main>Main</el-main>
      </el-container>
    </el-container>
  </div> -->

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
