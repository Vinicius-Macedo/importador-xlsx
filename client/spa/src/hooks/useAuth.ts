import { useEffect, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { axiosInstance } from "@/services/axiosInstance";

export function useAuth() {
  const [isAuthenticated, setIsAuthenticated] = useState<Boolean | null>(null);
  const { isLoading, isError, data } = useQuery({
    queryKey: ["repoData"],
    queryFn: async () => {
      const response = await axiosInstance.get("/user");
      return response; 
    },
    retry: false, 
  });

  useEffect(() => {
    if (!isLoading) {
      if (data) {
        setIsAuthenticated(true);
      } else if (isError || data === null) {
        setIsAuthenticated(false);
      }
    }
  }, [data, isLoading, isError]);

  return { isLoading, isError, data, isAuthenticated };
}