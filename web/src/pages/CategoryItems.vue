<template>
  <PageLayout>
    <q-scroll-area class="fit">
      <div class="sticky">
        <q-img :src="imgUrl" height="200px" position="50% 100%">
          <div class="absolute-bottom row">
            <h2 class="text-subtitle1 section-title text-md text-uppercase">
              Kategórie
            </h2>
            <q-space />
            <q-btn round outline icon="add"></q-btn>
          </div>
        </q-img>
      </div>
      <div class="row">
        <div class="col-6"><div class="test">test a</div></div>
        <div class="col-6">test b</div>
      </div>
    </q-scroll-area>
  </PageLayout>
</template>
<style scope>
.section-title {
  font-size: large;
  line-height: 0.75rem;
}
.sticky {
  position: sticky;
  top: -127px;
}
.test {
  height: 1200px;
}
</style>
<script lang="ts" setup>
import imgUrl from "@assets/megan-thomas-xMh_ww8HN_Q-unsplash copy.png";
import PageLayout from "@components/common/PageLayout.vue";
</script>
