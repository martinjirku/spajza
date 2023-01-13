import {
  addNewStorageItem,
  consumptSstorageItemField,
  getStorageItems,
  updateStorageItemField,
} from "@api";
import { ConsumptionRequest, NewStorageItem, StorageItem } from "@api/storage";
import { useMutation, useQuery, useQueryClient } from "vue-query";

export const useStorageItems = () =>
  useQuery("storage-items", () => getStorageItems(), {
    refetchOnMount: false,
  });

export type FieldKey = "title" | "storagePlaceId";

const getUseUpdateStorageItem =
  <T>(field: FieldKey) =>
  () => {
    const queryClient = useQueryClient();
    return useMutation(
      ({ storageItemId, value }: { storageItemId: number; value: T }) => {
        return updateStorageItemField(storageItemId, field, value);
      },
      {
        onSuccess: () => {
          queryClient.invalidateQueries("storage-items");
        },
      }
    );
  };

export const useUpdateStorageItemTitleMutation =
  getUseUpdateStorageItem<Required<StorageItem>["title"]>("title");

export const useUpdateStorageItemLocationMutation =
  getUseUpdateStorageItem<Required<StorageItem>["storagePlaceId"]>(
    "storagePlaceId"
  );

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

export const useConsumptMutation = () => {
  const queryClient = useQueryClient();
  return useMutation(
    ({
      storageItemId,
      ...rest
    }: ConsumptionRequest & { storageItemId: number }) => {
      return consumptSstorageItemField(storageItemId, rest);
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("storage-items");
      },
    }
  );
};
