import { getCategories } from "@api";
import { useQuery } from "vue-query";

export const useCategories = () =>
  useQuery("categories", () => getCategories());
