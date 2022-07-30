import { addNewStorageItem, getStorageItems } from "@api";
import { NewStorageItem, StorageItem } from "@api/storage";
import { useMutation, useQuery, useQueryClient } from "vue-query";

export const useStorageItems = () =>
  useQuery("storage-items", () => getStorageItems(), {
    refetchOnMount: false,
  });

export const useNewStorageItemMutation = () => {
  const queryClient = useQueryClient();
  return useMutation(
    (storageItem: NewStorageItem) => {
      let expirationDate: Date | undefined;
      if (storageItem.expirationDate !== undefined) {
        expirationDate = new Date(storageItem.expirationDate);
      }
      return addNewStorageItem({
        ...storageItem,
        amount: Number(storageItem.amount),
        expirationDate,
      });
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("storage-items");
      },
    }
  );
};
