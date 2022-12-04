StoragePlace
<template>
  <PageLayout>
    <q-scroll-area class="fit">
      <div class="sticky">
        <q-img :src="imgUrl" height="200px" position="50% 100%">
          <div class="absolute-bottom row">
            <h2 class="text-subtitle1 section-title text-md text-uppercase">
              Miesta uloženia
            </h2>
            <q-space />
            <div class="q-gutter-xs">
              <q-btn
                v-if="
                  activeItems.length > 0 && activeItems[0].storagePlaceId !== -1
                "
                round
                flat
                title="Zmazať"
                icon="delete"
                @click="onDelete"
              ></q-btn>
              <q-btn
                v-if="activeItems.length > 0"
                round
                flat
                icon="close"
                title="Zatvoriť"
                @click="activeItems = []"
              ></q-btn>
              <q-btn
                v-if="activeItems[0]?.storagePlaceId !== -1"
                round
                icon="add"
                flat
                title="Vytvoriť"
                @click="
                  activeItems = [
                    {
                      storagePlaceId: -1,
                      title: '',
                      code: '',
                    },
                  ]
                "
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
      <div v-else class="row q-col-gutter-md q-pt-md q-mr-md">
        <div class="col-12 col-sm-7" v-if="!activeItems[0] || $q.screen.gt.xs">
          <q-table
            ref="itemsTableRef"
            dark
            flat
            square
            card-container-class="card-container"
            :rows="storagePlaces"
            :columns="(columns as any)"
            selection="single"
            :selected="activeItems"
            row-key="storagePlaceId"
          >
            <template v-slot:header="props">
              <q-tr :props="props">
                <q-th v-for="col in props.cols" :key="col.name" :props="props">
                  {{ col.label }}
                </q-th>
              </q-tr>
            </template>
            <template v-slot:body="props">
              <q-tr
                class="cursor-pointer"
                :props="props"
                @click="
                  activeItems = props.row === activeItems[0] ? [] : [props.row]
                "
              >
                <q-td v-for="col in props.cols" :key="col.name" :props="props">
                  {{ col.value }}
                </q-td>
              </q-tr>
            </template>
          </q-table>
        </div>
        <div class="col-12 col-sm-5" v-if="!!activeItems[0] || $q.screen.gt.xs">
          <q-card class="sticky-card" v-if="activeItems[0]" flat square>
            <q-card-section>
              <StoragePlaceForm
                :key="activeItems[0]?.storagePlaceId ?? NaN"
                :storagePlaceId="activeItems[0]?.storagePlaceId"
                @submitted="onSubmitted"
              ></StoragePlaceForm>
            </q-card-section>
          </q-card>
          <div
            v-else
            class="text-white section-title text-subtitle1 text-center q-pt-xl"
          >
            Žiadne miesto uloženia nie je vybraté
          </div>
        </div>
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
  z-index: 1;
}
.sticky-card {
  position: sticky;
  top: 80px;
}
.test {
  height: 1200px;
}
.q-table__card {
  background-color: var(--bg-darktransparent);
}
.q-card {
  background-color: var(--bg-semitransparent);
}
</style>
<script lang="ts" setup>
import imgUrl from "@assets/megan-thomas-xMh_ww8HN_Q-unsplash copy.png";
import PageLayout from "@components/common/PageLayout.vue";
import { ref, computed } from "vue";
import StoragePlaceForm from "@storage/StoragePlaceForm.vue";
import { QTable, useQuasar } from "quasar";
import {
  useDeleteStoragePlaceMutation,
  useStoragePlaces,
} from "@storage/StoragePlaceQuery";
import { StoragePlace } from "@api/storagePlace";

const activeItems = ref<StoragePlace[]>([]);
const itemsTableRef = ref<QTable>();

const $q = useQuasar();

const { data: storagePlaces, isLoading } = useStoragePlaces();
const { mutateAsync } = useDeleteStoragePlaceMutation();

const onSubmitted = ([data, updated]: [StoragePlace, boolean]) => {
  activeItems.value = [data];
  if (!updated) {
    itemsTableRef.value?.lastPage();
  }
};

const onDelete = () => {
  if (activeItems.value[0]?.storagePlaceId === undefined) return;
  mutateAsync(activeItems.value[0].storagePlaceId).then((resp) => {
    activeItems.value = [];
  });
};

const columns = computed(() => [
  {
    name: "Názov",
    align: "left",
    label: "Názov",
    field: "title",
    sortable: false,
  },
  {
    name: "Kód",
    align: "center",
    label: "Kód",
    field: "code",
    sortable: false,
  },
]);
</script>
