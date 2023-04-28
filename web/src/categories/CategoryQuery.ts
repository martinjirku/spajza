import {
  createCategory,
  deleteCategory,
  getCategories,
  updateCategory,
} from "@api";
import { useMutation, useQuery, useQueryClient } from "@tanstack/vue-query";
import { Category } from "@api/category";
import { createTreeLikeCategoryOptions } from "./Category";
import { computed } from "vue";

export const useCategories = () =>
  useQuery(["categories"], () => getCategories(), {
    refetchOnMount: false,
  });

export const useCategoryOptions = () => {
  const { data, ...rest } = useQuery(["categories"], () => getCategories(), {
    refetchOnMount: false,
  });
  const options = computed(() => createTreeLikeCategoryOptions(data.value));
  return { data: options, ...rest };
};

export const useCategoryMutation = () => {
  const queryClient = useQueryClient();
  return useMutation(
    (data: Category) => {
      return (data.id ?? -1) > -1 ? updateCategory(data) : createCategory(data);
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries(["categories"]);
      },
    }
  );
};

export const useDeleteCategoryMutation = () => {
  const queryClient = useQueryClient();
  return useMutation(
    (id: number) => {
      return deleteCategory(id);
    },
    {
      onSettled: () => {
        queryClient.invalidateQueries(["categories"]);
      },
    }
  );
};
