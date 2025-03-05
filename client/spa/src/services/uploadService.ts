import { axiosInstance } from "./axiosInstance";

export async function uploadFile(file: File) {
  const formData = new FormData();
  formData.append("file", file);
  const response = await axiosInstance.post("/import", formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return response.data;
}