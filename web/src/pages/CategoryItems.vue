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
                v-if="activeCategories.length > 0"
                round
                icon="delete"
              ></q-btn>
              <q-btn round icon="add"></q-btn>
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
        <div class="col-7">
          <q-table
            dark
            flat
            square
            card-container-class="card-container"
            :rows="categories"
            :columns="columns as QTableProps['columns']"
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
        <div class="col-5">
          <q-card v-if="activeCategories[0]" flat square>
            <q-card-section>
              <Form
                :key="activeCategories[0].id"
                @submit="onSubmit"
                :initial-values="activeCategories[0]"
                :validation-schema="schema"
              >
                <Field name="title" v-slot="{ errorMessage, value, field }">
                  <q-input
                    label="Názov"
                    :model-value="(value as string)"
                    v-bind="field"
                    type="text"
                    :error="!!errorMessage"
                    :error-message="errorMessage"
                  ></q-input>
                </Field>
                <Field
                  name="defaultUnit"
                  v-slot="{
                    errorMessage,
                    value,
                    field: { value: _, ...field },
                  }"
                >
                  <q-select
                    label="Jednotka"
                    v-model="(value as string)"
                    :options="unitOptions"
                    v-bind="field"
                    stack-label
                    input-debounce="10"
                    map-options
                    emit-value
                    use-chips
                    use-input
                    @filter="filterUnits"
                    :error="!!errorMessage"
                    :error-message="errorMessage"
                  >
                    <template v-slot:selected-item="scope">
                      <q-chip
                        v-if="scope"
                        square
                        color="white"
                        text-color="brown"
                        class="q-mb-none q-mt-sm q-ml-xs q-mr-none"
                      >
                        <q-avatar
                          color="brown"
                          text-color="white"
                          :icon="scope.opt.icon"
                        />
                        <span class="q-mx-sm">
                          {{ scope.opt.label }}
                        </span>
                      </q-chip>
                      <q-badge v-else>prázdny</q-badge>
                    </template>
                    <template v-slot:no-option>
                      <q-item>
                        <q-item-section class="text-grey">
                          Žiadne jednotky
                        </q-item-section>
                      </q-item>
                    </template>
                  </q-select>
                </Field>
              </Form>
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
import { getCategories, getUnits } from "@api";
import PageLayout from "@components/common/PageLayout.vue";
import { useQuery } from "vue-query";
import { ref, computed, watch } from "vue";
import { QSelectProps, QTableProps } from "quasar";
import { Category } from "@api/category";
import { string, object, mixed, InferType, number } from "yup";
import { useForm, Form, Field, SubmissionHandler } from "vee-validate";
import { quantities, QuantityType, Unit } from "@api/unit";

const activeCategories = ref<Category[]>([]);

const { data: categories, isLoading } = useQuery("categories", () =>
  getCategories()
);
const { data: units, isLoading: unitsLoading } = useQuery("units", () =>
  getUnits()
);

const createUnits = (units: Unit[] = []): QSelectProps["options"] => {
  return units.map((a) => {
    return {
      label: `${[a.name]} (${a.symbol})`,
      value: a.name,
      icon: ((key: QuantityType) => {
        if (key === "mass") return "scale";
        if (key === "volume") return "takeout_dining";
        if (key === "time") return "timer";
        if (key === "count") return "tag";
        if (key === "length") return "straighten";
        if (key === "temperature") return "thermostat";
        return "square_foot";
      })(a.quantity),
    };
  });
};
const unitOptions = ref(createUnits(units.value));
watch(units, () => {
  unitOptions.value = createUnits(units.value);
});
watch(activeCategories, () => {
  unitOptions.value = createUnits(units.value);
});

const filterUnits = (val: string, update: Function) => {
  if (val === "") {
    update(() => {
      unitOptions.value = createUnits(units.value);
    });
    return;
  }
  update(() => {
    unitOptions.value = createUnits(
      units.value?.filter((i) => {
        return (
          i.name.toLowerCase().indexOf(val.toLowerCase()) > -1 ||
          i.symbol.toLowerCase().indexOf(val.toLowerCase()) > -1 ||
          i.names.some((n) => n.toLowerCase().indexOf(val.toLowerCase()) > -1)
        );
      }) ?? []
    );
  });
};

const schema = object({
  id: number().optional(),
  title: string().max(250).required(),
  path: string().max(250).optional(),
  defaultUnit: string().max(50),
  quantityType: mixed<QuantityType>().oneOf(quantities),
});

type FormState = InferType<typeof schema>;

const onSubmit = ((values) => {
  console.log(values);
}) as SubmissionHandler;

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
