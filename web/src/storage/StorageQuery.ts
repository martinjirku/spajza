import {
  addNewStorageItem,
  getStorageItems,
  updateStorageItemField,
} from "@api";
import { NewStorageItem } from "@api/storage";
import { useMutation, useQuery, useQueryClient } from "vue-query";

export const useStorageItems = () =>
  useQuery("storage-items", () => getStorageItems(), {
    refetchOnMount: false,
  });

export const useUpdateStorageItemTitleMutation = () => {
  const queryClient = useQueryClient();
  return useMutation(
    ({ storageItemId, title }: { storageItemId: number; title: string }) => {
      return updateStorageItemField(storageItemId, "title", title);
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("storage-items");
      },
    }
  );
};

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
