import { Category } from "./category";
import {
  ConsumptionRequest,
  NewStorageItem,
  StorageItemListResponse,
  StorageItem,
} from "./storage";
import { CreateStoragePlace, StoragePlace } from "./storagePlace";
import { Unit } from "./unit";

const headers: HeadersInit = {
  Accept: "application/json",
  "Content-Type": "application/json",
};

export const getCategories = () => {
  return fetch("/api/categories")
    .then((resp) => {
      return resp.json();
    })
    .then((data) => data as Category[]);
};

export const getUnits = () => {
  return fetch("api/units", { headers })
    .then((resp) => resp.json())
    .then((data) => data as Unit[]);
};

export const createCategory = (category: Category) => {
  return fetch("/api/categories", {
    method: "POST",
    headers,
    body: JSON.stringify(category),
  })
    .then((resp) => resp.json())
    .then((data) => data as Category);
};

export const updateCategory = (category: Category) => {
  return fetch(`/api/categories/${category.id}`, {
    method: "POST",
    headers,
    body: JSON.stringify(category),
  })
    .then((resp) => resp.json())
    .then((data) => data as Category);
};

export const deleteCategory = (categoryId: number) => {
  return fetch(`/api/categories/${categoryId}`, {
    method: "DELETE",
    headers,
  }).then((r) => r.text());
};

export const getStoragePlaces = () => {
  return fetch("api/storage/places", { headers })
    .then((resp) => resp.json())
    .then((data) => data as StoragePlace[]);
};

export const getStoragePlace = (storagePlaceId: number) => {
  return fetch(`api/storage/places/${storagePlaceId}`, { headers })
    .then((resp) => resp.json())
    .then((data) => data as StoragePlace);
};

export const deleteStoragePlace = (storagePlaceId: number) => {
  return fetch(`/api/storage/places/${storagePlaceId}`, {
    method: "DELETE",
    headers,
  }).then((r) => r.text());
};

export const createStoragePlace = (storagePlace: CreateStoragePlace) => {
  return fetch(`/api/storage/places`, {
    method: "POST",
    headers,
    body: JSON.stringify(storagePlace),
  })
    .then((resp) => resp.json())
    .then((data) => data as StoragePlace);
};

export const updateStoragePlace = (storagePlace: StoragePlace) => {
  return fetch(`/api/storage/places/${storagePlace.storagePlaceId}`, {
    method: "POST",
    headers,
    body: JSON.stringify(storagePlace),
  })
    .then((resp) => resp.json())
    .then((data) => data as StoragePlace);
};

export const addNewStorageItem = (storageItem: NewStorageItem) => {
  return fetch(`/api/storage/items`, {
    method: "POST",
    headers,
    body: JSON.stringify(storageItem),
  })
    .then((r) => r.json())
    .then((data) => data as { data: StorageItem });
};

export const getStorageItems = () => {
  return fetch(`/api/storage/items`, {
    method: "GET",
    headers,
  })
    .then((r) => r.json())
    .then((data) => data as StorageItemListResponse);
};

export const updateStorageItemField = (
  storagePlaceId: number,
  fieldName: string,
  value: unknown
) => {
  return fetch(`/api/storage/items/${storagePlaceId}/${fieldName}`, {
    method: "POST",
    headers,
    body: JSON.stringify({ value }),
  });
};

export const consumptSstorageItemField = (
  storageItemId: number,
  request: ConsumptionRequest
) => {
  return fetch(`/api/storage/items/${storageItemId}/consumpt`, {
    method: "POST",
    headers,
    body: JSON.stringify(request),
  }).then((r) => r.json());
};
