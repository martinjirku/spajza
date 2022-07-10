import {
  createCategory,
  deleteCategory,
  getCategories,
  updateCategory,
} from "@api";
import { useMutation, useQuery, useQueryClient } from "vue-query";
import { Category } from "@api/category";

export const useCategories = () =>
  useQuery("categories", () => getCategories(), {
    refetchOnMount: false,
  });

export const useCategoryMutation = () => {
  const queryClient = useQueryClient();
  return useMutation(
    (data: Category) => {
      return (data.id ?? -1) > -1 ? updateCategory(data) : createCategory(data);
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("categories");
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
        queryClient.invalidateQueries("categories");
      },
    }
  );
};
