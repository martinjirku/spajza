import { getUnits } from "@api";
import { useQuery } from "@tanstack/vue-query";

export const useUnits = () =>
  useQuery(["units"], () => getUnits(), {
    refetchOnMount: false,
    refetchOnReconnect: false,
    refetchOnWindowFocus: false,
  });
