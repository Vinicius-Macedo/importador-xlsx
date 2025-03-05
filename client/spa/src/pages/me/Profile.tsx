import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { useQueryClient } from "@tanstack/react-query";
import { axiosInstance } from "@/services/axiosInstance";
import { useNavigate } from "react-router-dom";
import { useAuth } from "@/hooks/useAuth";
import { Container } from "@/components/ui/container";

export function Profile() {
  const { isLoading, data } = useAuth();
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  async function handleLogout() {
    await axiosInstance.post("/logout");
    queryClient.removeQueries({ queryKey: ["repoData"] });
    navigate("/");
  }

  if (isLoading) return null;

  return (
    <>
      <Container>
        <Card>
          <CardHeader>
            <CardTitle>Meu Perfil</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-col">
              <p>
                <span className="font-bold">Nome:</span> {data?.data?.name}
              </p>
              <p>
                <span className="font-bold">Nome de Usuário:</span>{" "}
                {data?.data?.username}
              </p>
              <p>
                <span className="font-bold">Email:</span> {data?.data.email}
              </p>
            </div>
          </CardContent>
          <CardFooter>
            <div className="flex justify-end">
              <Button onClick={handleLogout}>Sair</Button>
            </div>
          </CardFooter>
        </Card>
      </Container>
    </>
  );
}
