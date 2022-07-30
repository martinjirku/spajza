import { addNewStorageItem, getStorageItems } from "@api";
import { StorageItem } from "@api/storage";
import { useMutation, useQuery, useQueryClient } from "vue-query";

export const useStorageItems = () =>
  useQuery("storage-items", () => getStorageItems(), {
    refetchOnMount: false,
  });

export const useNewStorageItemMutation = () => {
  const queryClient = useQueryClient();
  return useMutation(
    (storageItem: StorageItem) => addNewStorageItem(storageItem),
    {
      onSuccess: () => {
        queryClient.invalidateQueries("storage-items");
      },
    }
  );
};
