import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { useQueryClient } from "@tanstack/react-query";
import { login } from "@/services/authService";
import { validateFields } from "@/helpers/formValidatorHelper";

export function Login() {
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const [form, setForm] = useState({
    email: "",
    password: "",
  });
  const [error, setError] = useState<string>("");
  const [fieldsError, setFieldsError] = useState<{
    [key: string]: string;
  }>({});

  const mutation = useMutation({
    mutationFn: () => login(form),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["repoData"] });
      navigate("/usuario");
    },
    onError: (error: any) => {
      console.log("Login error", error);
      setError(error.response?.data?.error || "Ocorreu um erro");
    },
  });

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();

    if (Object.keys(validateFields(form)).length > 0) {
      setError("Email ou senha inválidos");
      return;
    }

    mutation.mutate();
  };

  return (
    <div className="w-full h-full flex justify-center items-center flex-auto">
      <div className="flex items-center justify-center py-12 h-full">
        <form
          onSubmit={handleSubmit}
          className="flex flex-col gap-6 mx-auto md:w-[400px]"
        >
          <div className="flex flex-col items-center gap-2 text-center">
            <h1 className="text-3xl font-bold">Entrar</h1>
            <p className="text-balance text-muted-foreground">
              {error ? (
                <span className="text-red-500">{error}</span>
              ) : (
                "Preencha os campos abaixo para acessar sua conta"
              )}
            </p>
          </div>
          <div className="grid gap-6">
            <div className="grid gap-2">
              <Label htmlFor="email">Email</Label>
              <Input
                id="email"
                type="email"
                name="email"
                placeholder="m@example.com"
                onChange={(e) =>
                  setForm({
                    ...form,
                    email: e.target.value,
                  })
                }
              />
            </div>
            <div className="grid gap-2">
              <div className="flex items-center">
                <Label htmlFor="password">Senha</Label>
                <p
                  onClick={() => navigate("/forgot-password")}
                  className="ml-auto text-sm underline-offset-4 hover:underline cursor-pointer"
                >
                  Esqueceu sua senha?
                </p>
              </div>
              <Input
                id="password"
                type="password"
                name="password"
                onChange={(e) =>
                  setForm({
                    ...form,
                    password: e.target.value,
                  })
                }
              />
            </div>
            <Button type="submit" className="w-full">
              Entrar
            </Button>
          </div>
          <div className="text-center text-sm">
            Não tem uma conta?{" "}
            <span
              onClick={() => navigate("/register")}
              className="underline underline-offset-4 cursor-pointer"
            >
              Registrar
            </span>
          </div>
        </form>
      </div>
    </div>
  );
}