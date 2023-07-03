<template>
  <div class="row fit q-pt-md q-pr-md">
    <div class="col-12 q-mb-md flex">
      <q-space />
      <q-btn color="red-1" outline no-caps icon="add">Pridať Filter</q-btn>
      <q-separator vertical dark spaced />
      <q-btn color="red-1" outline no-caps iconRight="sym_o_view_column"
        >Stĺpce</q-btn
      >
    </div>
    <div class="col-12 q-mb-sm">
      <q-card flat square>
        <q-markup-table separator="cell" square>
          <thead>
            <tr
              v-for="headerGroup in table.getHeaderGroups()"
              :key="headerGroup.id"
            >
              <th
                v-for="header in headerGroup.headers"
                :key="header.id"
                :colSpan="header.colSpan"
                :style="{
                  position: 'relative',
                  width:
                    header.getSize() === 120.111
                      ? 'auto'
                      : `${header.getSize()}px`,
                }"
              >
                <FlexRender
                  v-if="!header.isPlaceholder"
                  :render="header.column.columnDef.header"
                  :props="header.getContext()"
                />
                <div
                  v-if="header.column.getCanResize()"
                  @mousedown="(e) => header.getResizeHandler()(e)"
                  @touchstart="(e) => header.getResizeHandler()(e)"
                  :class="{
                    resizer: true,
                    isResizing: header.column.getIsResizing(),
                  }"
                />
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="row in table.getRowModel().rows" :key="row.id">
              <td v-for="cell in row.getVisibleCells()" :key="cell.id">
                <FlexRender
                  :render="cell.column.columnDef.cell"
                  :props="cell.getContext()"
                />
              </td>
            </tr>
          </tbody>
        </q-markup-table>
      </q-card>
    </div>
    <div class="col-12 q-mb-md">
      <div class="row">
        <div class="col-12 flex q-gutter-xs">
          <q-space />
          <q-btn
            :disabled="!table.getCanPreviousPage()"
            color="grey-2"
            outline
            icon="first_page"
            @click="() => table.setPageIndex(0)"
          ></q-btn>
          <q-btn
            :disabled="!table.getCanPreviousPage()"
            color="grey-2"
            outline
            icon="navigate_before"
            @click="() => table.previousPage()"
          ></q-btn>
          <q-btn
            v-for="btnIndex of paginationBarItems"
            color="grey-2"
            :disable="pagination.pageIndex === btnIndex"
            push
            :flat="pagination.pageIndex !== btnIndex"
            :outline="pagination.pageIndex === btnIndex"
            @click="table.setPageIndex(btnIndex)"
            >{{ btnIndex + 1 }}</q-btn
          >
          <q-btn
            :disable="!table.getCanNextPage()"
            color="grey-2"
            outline
            icon="navigate_next"
            @click="() => table.nextPage()"
          ></q-btn>
          <q-btn
            :disable="!table.getCanNextPage()"
            color="grey-2"
            outline
            icon="last_page"
            @click="() => table.setPageIndex(table.getPageCount() - 1)"
          ></q-btn>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { StorageItem } from "@api/storage";
import { StoragePlace } from "@api/storagePlace";
import { Unit } from "@api/unit";
import { createStoragePlaceOptions } from "@storage/StoragaPlace";
import { computed, ref } from "vue";
import { FlexRender } from "@tanstack/vue-table";
import {
  useVueTable,
  createColumnHelper,
  getCoreRowModel,
  getPaginationRowModel,
  PaginationState,
} from "@tanstack/vue-table";

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

const pagination = ref<PaginationState>({
  pageIndex: 0,
  pageSize: 10,
});
const pageCount = computed(() => {
  const result = Math.ceil(
    (props.items?.length ?? 0) / pagination.value.pageSize
  );
  return result;
});

const paginationBarItems = computed(() => {
  const visiblePages = 3; // Number of pages to display
  const halfVisible = Math.floor(visiblePages / 2);
  let startPage, endPage;
  if (pageCount.value <= visiblePages) {
    // Display all pages if the total is less than or equal to the visible pages
    startPage = 0;
    endPage = pageCount.value - 1;
  } else if (pagination.value.pageIndex - halfVisible <= 0) {
    // Display the first visiblePages if the current page is near the beginning
    startPage = 0;
    endPage = visiblePages - 1;
  } else if (pagination.value.pageIndex + halfVisible >= pageCount.value) {
    // Display the last visiblePages if the current page is near the end
    startPage = pageCount.value - visiblePages;
    endPage = pageCount.value - 1;
  } else {
    // Display pages around the current page
    startPage = pagination.value.pageIndex - halfVisible;
    endPage = pagination.value.pageIndex + halfVisible;
  }
  let result = [];
  for (let i = startPage; i <= endPage; i++) {
    result.push(i);
  }

  return result;
});

const storagePlaceOptions = computed(() =>
  createStoragePlaceOptions(props.storagePlaces)
);

const data = computed(() =>
  (props.items ?? []).slice(
    pagination.value.pageIndex * pagination.value.pageSize,
    (pagination.value.pageIndex + 1) * pagination.value.pageSize
  )
);
const tableState = computed(() => ({
  pagination: pagination.value,
}));

const columnHelper = createColumnHelper<StorageItem>();
const table = useVueTable<StorageItem>({
  get data() {
    return data.value;
  },
  get state() {
    return tableState.value;
  },
  manualPagination: true,
  get pageCount() {
    return pageCount.value;
  },
  onPaginationChange: (updater) => {
    const newState =
      typeof updater === "function" ? updater(pagination.value) : updater;
    pagination.value.pageIndex = newState.pageIndex;
    pagination.value.pageSize = newState.pageSize;
  },
  getCoreRowModel: getCoreRowModel<StorageItem>(),
  getPaginationRowModel: getPaginationRowModel(),
  columnResizeMode: "onChange",
  defaultColumn: {},
  columns: [
    columnHelper.accessor("title", {
      header: "Názov",
      size: 250,
    }),
    columnHelper.accessor(
      (row) =>
        storagePlaceOptions.value.find((a) => a.value === row.storagePlaceId)
          ?.label ?? "",
      {
        header: "Uloženie",
        id: "location",
      }
    ),
    columnHelper.accessor(
      (row) =>
        `${row.currentAmount} ${
          props.units?.find((u) => u.name === row.unit)?.symbol
        }`,
      {
        header: "Množstvo",
        id: "amount",
      }
    ),
  ],
});

const emit = defineEmits<Emits>();
</script>

<style scoped>
.resizer {
  position: absolute;
  right: -2px;
  top: 0;
  height: 100%;
  width: 5px;
  background: rgba(0, 0, 0, 0.5);
  cursor: col-resize;
  user-select: none;
  touch-action: none;
}

.resizer.isResizing {
  background: blue;
  opacity: 1;
}

@media (hover: hover) {
  .resizer {
    opacity: 0;
  }

  *:hover > .resizer {
    opacity: 1;
  }
}
</style>
