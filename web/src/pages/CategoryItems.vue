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
            <div class="q-gutter-xs">
              <q-btn
                v-if="
                  activeCategories.length > 0 && activeCategories[0].id !== -1
                "
                round
                flat
                title="Zmazať"
                icon="delete"
                @click="onDelete"
              ></q-btn>
              <q-btn
                v-if="activeCategories.length > 0"
                round
                flat
                icon="close"
                title="Zatvoriť"
                @click="activeCategories = []"
              ></q-btn>
              <q-btn
                v-if="activeCategories[0]?.id !== -1"
                round
                icon="add"
                flat
                title="Vytvoriť"
                @click="
                  activeCategories = [
                    { id: -1, path: '', defaultUnit: 'kilogram', title: '' },
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
        <div
          class="col-12 col-sm-7"
          v-if="!activeCategories[0] || $q.screen.gt.xs"
        >
          <q-table
            ref="categoryTableRef"
            dark
            flat
            square
            card-container-class="card-container"
            :rows="categories"
            :columns="(columns as any)"
            selection="single"
            :selected="activeCategories"
            row-key="id"
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
                  activeCategories =
                    props.row === activeCategories[0] ? [] : [props.row]
                "
              >
                <q-td v-for="col in props.cols" :key="col.name" :props="props">
                  {{ col.value }}
                </q-td>
              </q-tr>
            </template>
          </q-table>
        </div>
        <div
          class="col-12 col-sm-5"
          v-if="!!activeCategories[0] || $q.screen.gt.xs"
        >
          <q-card class="sticky-card" v-if="activeCategories[0]" flat square>
            <q-card-section>
              <CategoryForm
                :key="activeCategories[0]?.id ?? NaN"
                :categoryId="activeCategories[0]?.id"
                @submitted="onSubmitted"
              ></CategoryForm>
            </q-card-section>
          </q-card>
          <div
            v-else
            class="text-white section-title text-subtitle1 text-center q-pt-xl"
          >
            Žiadna kategória nie je vybratá
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
  background-color: transparent;
}
.q-card {
  background-color: var(--bg-semitransparent);
}
</style>
<script lang="ts" setup>
import imgUrl from "@assets/megan-thomas-xMh_ww8HN_Q-unsplash copy.png";
import PageLayout from "@components/common/PageLayout.vue";
import { ref, computed, watch } from "vue";
import { Category } from "@api/category";
import { useUnits } from "@categories/UnitQuery";
import {
  useCategories,
  useDeleteCategoryMutation,
} from "@categories/CategoryQuery";
import CategoryForm from "@categories/CategoryForm.vue";
import { QTable, useQuasar } from "quasar";

const activeCategories = ref<Category[]>([]);
const categoryTableRef = ref<QTable>();

const $q = useQuasar();

const { data: categories, isLoading } = useCategories();
const { data: units, isLoading: unitsLoading } = useUnits();
const { mutateAsync } = useDeleteCategoryMutation();

const onSubmitted = ([data, updated]: [Category, boolean]) => {
  activeCategories.value = [data];
  if (!updated) {
    categoryTableRef.value?.lastPage();
  }
};

const onDelete = () => {
  if (activeCategories.value[0]?.id === undefined) return;
  mutateAsync(activeCategories.value[0].id).then((resp) => {
    activeCategories.value = [];
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
    name: "Jednotky",
    align: "center",
    label: "Merné jednotky",
    field: (item: Category) => {
      if (unitsLoading.value) return "";
      const unit = units.value?.find((a) => a.name === item.defaultUnit);
      switch (unit?.quantity) {
        case "mass":
          return `Váha (${unit?.symbol})`;
        case "length":
          return `Dĺžka (${unit?.symbol})`;
        case "volume":
          return `Objem (${unit?.symbol})`;
        case "temperature":
          return `Templota (${unit?.symbol})`;
        case "time":
          return `Čas (${unit?.symbol})`;
        case "count":
          return `Počet (${unit?.symbol})`;
        default:
          return `Neznáme (${unit?.symbol})`;
      }
    },
  },
]);
</script>
