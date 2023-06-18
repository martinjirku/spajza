<template>
  <div v-if="!isLoading" class="row fit q-col-gutter-sm q-pt-md">
    <div
      class="col-12 col-sm-6 col-md-4 col-lg-3"
      v-for="(i, index) in items ?? []"
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
              @save="(value: string) => emit('update:title', i.storageItemId ?? 0, value)"
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
              @save="(value: number) => emit('update:location', i.storageItemId ?? 0, value)"
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
</template>

<script lang="ts" setup>
import { StorageItem } from "@api/storage";
import { StoragePlace } from "@api/storagePlace";
import { Unit } from "@api/unit";
import { createStoragePlaceOptions } from "@storage/StoragaPlace";
import { QPopupProxy } from "quasar";
import { computed, ref } from "vue";

const consumptPopup = ref<QPopupProxy[]>();

interface Props {
  isLoading?: boolean;
  items?: StorageItem[];
  storagePlaces?: StoragePlace[];
  units?: Unit[];
}
const props = defineProps<Props>();
interface Emits {
  (event: "update:title", storageItemId: number, value: string): void;
  (event: "update:location", storageItemId: number, value: number): void;
}

const storagePlaceOptions = computed(() =>
  createStoragePlaceOptions(props.storagePlaces)
);

const emit = defineEmits<Emits>();
</script>
