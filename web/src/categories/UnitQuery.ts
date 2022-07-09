import { getUnits } from "@api";
import { useQuery } from "vue-query";

export const useUnits = () => useQuery("units", () => getUnits());
