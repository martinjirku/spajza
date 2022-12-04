import {
  createStoragePlace,
  deleteStoragePlace,
  getStoragePlaces,
  updateStoragePlace,
} from "@api";
import { useMutation, useQuery, useQueryClient } from "vue-query";
import { StoragePlace } from "@api/storagePlace";

export const useStoragePlaces = () =>
  useQuery("storage-places", () => getStoragePlaces(), {
    refetchOnMount: false,
  });

export const useStoragePlacesMutation = () => {
  const queryClient = useQueryClient();
  return useMutation(
    (data: StoragePlace) => {
      return (data.storagePlaceId ?? -1) === -1
        ? createStoragePlace(data)
        : updateStoragePlace(data);
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("storage-places");
      },
    }
  );
};

export const useDeleteStoragePlaceMutation = () => {
  const queryClient = useQueryClient();
  return useMutation(
    (id: number) => {
      return deleteStoragePlace(id);
    },
    {
      onSettled: () => {
        queryClient.invalidateQueries("storage-places");
      },
    }
  );
};
