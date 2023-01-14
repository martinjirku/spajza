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

      <div v-if="!isLoading" class="row fit q-col-gutter-sm q-pt-md">
        <div
          class="col-12 col-sm-6 col-md-4 col-lg-3"
          v-for="(i, index) in items?.items"
        >
          <q-card flat square>
            <q-card-section>
              <div class="text-h5">
                <span class="text-dark cursor-pointer">
                  {{ i.title }}
                </span>
                <q-popup-edit
                  v-bind:model-value="i.title"
                  buttons
                  v-slot="scope"
                  @save="(value: string) => updateTitle({storageItemId: i.storageItemId ?? 0, value})"
                >
                  <q-input
                    v-model="scope.value"
                    dense
                    title="Názov"
                    autofocus
                    @keyup.enter="scope.set"
                  />
                </q-popup-edit>
              </div>
              <div class="row inline no-wrap items-center">
                <q-icon name="place" class="q-mr-sm text-accent" size="xs" />
                <span class="text-subtitle2 text-primary cursor-pointer">
                  {{
                    storagePlaces?.find(
                      (s) => s.storagePlaceId === i.storagePlaceId
                    )?.title
                  }}
                </span>
                <q-popup-edit
                  v-bind:model-value="i.storagePlaceId"
                  buttons
                  v-slot="scope"
                  @save="(value: number) => updateLocation({storageItemId: i.storageItemId ?? 0, value})"
                >
                  <q-select
                    label="Miesto uloženia"
                    v-model="scope.value"
                    :options="storagePlaceOptions"
                    v-bind="scope.value"
                    map-options
                    emit-value
                  >
                    <template v-slot:no-option>
                      <q-item>
                        <q-item-section class="text-grey">
                          Neexistujú žiadne miesta
                        </q-item-section>
                      </q-item>
                    </template>
                  </q-select>
                </q-popup-edit>
              </div>
            </q-card-section>
            <q-separator />
            <q-card-section>
              <div class="row">
                <div class="col-6 col-md-4">
                  <q-popup-proxy>
                    <q-list bordered separator>
                      <q-item v-for="p in i.consumptions">{{
                        `-${p.amount} ${
                          units?.find((u) => u.name === p.unit)?.symbol ?? ""
                        }`
                      }}</q-item>
                    </q-list>
                  </q-popup-proxy>
                  <div class="text-subtitle2">Váha</div>
                  <div class="text-h5 text-weight-bold">
                    {{ i.currentAmount }}
                    <span class="text-weight-regular">
                      {{ units?.find((u) => u.name === i.unit)?.symbol }}
                    </span>
                  </div>
                </div>
                <div class="col-6 col-md-8">
                  <q-btn class="fit" outline flat>
                    <q-popup-proxy ref="consumptPopup">
                      <StorageItemConsumptForm
                        :id="i.storageItemId ?? 0"
                        :default-unit="i.unit ?? ''"
                        :default-value="i.currentAmount ?? 0"
                        :close="consumptPopup?.[index]?.hide"
                      />
                    </q-popup-proxy>
                    <q-icon left size="1.5em" name="remove_circle_outline" />
                    Upotrebiť
                  </q-btn>
                </div>
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
import { useStoragePlaces } from "@storage/StoragePlaceQuery";
import {
  useStorageItems,
  useUpdateStorageItemTitleMutation,
  useUpdateStorageItemLocationMutation,
} from "@storage/StorageQuery";
import StorageItemConsumptForm from "@storage/StorageItemConsumptForm.vue";
import { computed, ref, watch } from "vue";
import StorageItemForm from "@storage/StorageItemForm.vue";
import { createStoragePlaceOptions } from "@storage/StoragaPlace";
import { QPopupProxy } from "quasar";

type FormType = "closed" | "create-new" | "upload-barcode";

const formState = ref<FormType>("closed");
const consumptPopup = ref<QPopupProxy[]>();

const { data: items, isLoading: isLoadingStorageItems } = useStorageItems();
const { data: units, isLoading: isLoadingUnits } = useUnits();
const { mutateAsync: updateTitle } = useUpdateStorageItemTitleMutation();
const { mutateAsync: updateLocation } = useUpdateStorageItemLocationMutation();
const { data: storagePlaces, isLoading: isLoadingStoragePlaces } =
  useStoragePlaces();

const storagePlaceOptions = computed(() =>
  createStoragePlaceOptions(storagePlaces.value)
);

const isLoading = computed(() => {
  return (
    isLoadingStorageItems.value ||
    isLoadingUnits.value ||
    isLoadingStoragePlaces.value
  );
});
</script>
