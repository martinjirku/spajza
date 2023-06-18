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
            <div v-if="formState !== 'closed'" class="q-gutter-xs">
              <q-btn
                round
                flat
                icon="close"
                title="Zatvoriť"
                @click="formState = 'closed'"
              ></q-btn>
            </div>
            <div v-else class="q-gutter-xs">
              <q-btn
                round
                flat
                icon="sym_o_barcode"
                title="Zatvoriť"
                @click="formState = 'upload-barcode'"
              />
              <q-btn
                round
                flat
                icon="add"
                title="Pridať"
                @click="formState = 'create-new'"
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
      <div
        v-if="formState !== 'closed'"
        class="row fit q-col-gutter-sm q-pt-md q-mb-sm"
      >
        <div class="col-12 shadow-4">
          <q-card flat square>
            <q-card-section>
              <StorageItemForm
                :barcode-preload="formState === 'upload-barcode'"
              />
            </q-card-section>
          </q-card>
        </div>
      </div>
      <div v-if="formState === 'upload-barcode'" class="row fit q-col" />

      <card-list
        :is-loading="isLoading"
        :items="itemsData?.items"
        :units="unitsData"
        :storage-places="storagePlaces"
        @update:title="(id, value) => updateTitle({ storageItemId: id, value })"
        @update:location="
          (id, value) => updateLocation({ storageItemId: id, value })
        "
      />
    </q-scroll-area>
  </PageLayout>
</template>
<script lang="ts" setup>
import imgUrl from "@assets/megan-thomas-xMh_ww8HN_Q-unsplash copy.png";
import { useUnits } from "@categories/UnitQuery";
import PageLayout from "@components/common/PageLayout.vue";
import { useStoragePlaces } from "@storage/StoragePlaceQuery";
import {
  useStorageItems,
  useUpdateStorageItemTitleMutation,
  useUpdateStorageItemLocationMutation,
} from "@storage/StorageQuery";
import { computed, ref } from "vue";
import StorageItemForm from "@storage/StorageItemForm.vue";
import CardList from "./CardList.vue";

type FormType = "closed" | "create-new" | "upload-barcode";

const formState = ref<FormType>("closed");

const { data: itemsData, isLoading: isLoadingStorageItems } = useStorageItems();
const { data: unitsData, isLoading: isLoadingUnits } = useUnits();
const { mutateAsync: updateTitle } = useUpdateStorageItemTitleMutation();
const { mutateAsync: updateLocation } = useUpdateStorageItemLocationMutation();
const { data: storagePlaces, isLoading: isLoadingStoragePlaces } =
  useStoragePlaces();

const isLoading = computed(() => {
  return (
    isLoadingStorageItems.value ||
    isLoadingUnits.value ||
    isLoadingStoragePlaces.value
  );
});
</script>
