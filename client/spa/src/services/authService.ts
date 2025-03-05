import { axiosInstance } from "@/services/axiosInstance";


export const login = (form: { email: string; password: string }) => {
  return axiosInstance.post("/login", form);
};

export const logout = async () => {
  const response = await axiosInstance.post(`/logout`);
  return response.data;
};

export const fetchUser = async () => {
  const response = await axiosInstance.get(`/user`);
  return response.data;
};
