<template>
  <PageLayout>
    <q-scroll-area class="fit">
      <div class="sticky">
        <q-img :src="imgUrl" height="200px" position="50% 100%">
          <div class="absolute-bottom row">
            <h2 class="text-subtitle1 section-title text-md text-uppercase">
              Špajza
            </h2>
            <q-space />
            <div v-if="createNew" class="q-gutter-xs">
              <q-btn
                round
                flat
                icon="close"
                title="Zatvoriť"
                @click="createNew = false"
              ></q-btn>
            </div>
            <div v-else class="q-gutter-xs">
              <q-btn
                round
                flat
                icon="add"
                title="Pridať"
                @click="createNew = true"
              ></q-btn>
            </div>
          </div>
        </q-img>
      </div>
      <div v-if="isLoading" class="fit row q-pt-xl">
        <q-space></q-space>
        <div class="items-center">
          <q-spinner-ball class="self-center" size="80px" color="grey-1" />
        </div>
        <q-space></q-space>
      </div>
      <div v-if="createNew" class="row fit q-col-gutter-sm q-pt-md q-mb-sm">
        <div class="col-12 shadow-4">
          <q-card flat square>
            <q-card-section>
              <StorageItemForm />
            </q-card-section>
          </q-card>
        </div>
      </div>
      <div v-if="!isLoading" class="row fit q-col-gutter-sm q-pt-md">
        <div
          class="col-12 col-sm-6 col-md-4 col-lg-3"
          v-for="i in items?.items"
        >
          <q-card flat square>
            <q-card-section>
              <div class="text-h5">
                {{ i.title }}
                <q-popup-edit v-model="i.title" buttons v-slot="scope">
                  <q-input
                    v-model="scope.value"
                    dense
                    title="Názov"
                    autofocus
                    @keyup.enter="scope.set"
                  />
                </q-popup-edit>
              </div>
            </q-card-section>
            <q-separator />
            <q-card-section>
              <div>
                {{
                  storagePlaces?.find(
                    (s) => s.storagePlaceId === i.storagePlaceId
                  )?.title
                }}
              </div>
              <div>
                {{ i.currentAmount }}
                {{ units?.find((u) => u.name === i.unit)?.symbol }}
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
    </q-scroll-area>
  </PageLayout>
</template>
<script lang="ts" setup>
import imgUrl from "@assets/megan-thomas-xMh_ww8HN_Q-unsplash copy.png";
import { useUnits } from "@categories/UnitQuery";
import PageLayout from "@components/common/PageLayout.vue";
import { useStoryPlaces } from "@storage/StoragePlaceQuery";
import { useStorageItems } from "@storage/StorageQuery";
import { computed, ref } from "vue";
import StorageItemForm from "@storage/StorageItemForm.vue";

const createNew = ref(false);

const { data: items, isLoading: isLoadingStorageItems } = useStorageItems();
const { data: units, isLoading: isLoadingUnits } = useUnits();
const { data: storagePlaces, isLoading: isLoadingStoragePlaces } =
  useStoryPlaces();

const isLoading = computed(() => {
  return (
    isLoadingStorageItems.value ||
    isLoadingUnits.value ||
    isLoadingStoragePlaces.value
  );
});
</script>
